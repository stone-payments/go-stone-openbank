package openbank

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/square/go-jose.v2"
)

const (
	stonePublicKeysEndpoint = `/api/v1/discovery/keys`
)

func (c *Client) DecryptAndValidateWebhook(encryptedJWE string) ([]byte, error) {
	decryptedData, err := c.DecryptJWE(encryptedJWE)
	if err != nil {
		return nil, fmt.Errorf(`failed at decrypting webhook: %w`, err)
	}

	jwe, err := jose.ParseSigned(string(decryptedData))
	if err != nil {
		return nil, fmt.Errorf(`err parsing webhook data: %w`, err)
	}

	signatureKey, err := c.getSignatureKey(jwe.Signatures)
	if err != nil {
		return nil, fmt.Errorf(`error geting signature key: %w`)
	}

	if _, err := jwe.Verify(signatureKey); err != nil {
		return nil, fmt.Errorf(`err verifying webhook signature: %w`, err)
	}

	payload, err := getPayload(jwe)
	if err != nil {
		return nil, fmt.Errorf(`failed serializing jwe payload: %w`, err)
	}

	return payload, nil
}

// this is a terrible workaround over JSONWebSignature to get it's private payload content
func getPayload(jwe *jose.JSONWebSignature) ([]byte, error) {
	jweSerialized := jwe.FullSerialize()
	type payload struct {
		Payload string `json:"payload"`
	}
	var p payload
	if err := json.Unmarshal([]byte(jweSerialized), &p); err != nil {
		return nil, err
	}

	b, err := base64.RawURLEncoding.DecodeString(p.Payload)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) getSignatureKey(signatures []jose.Signature) (*jose.JSONWebKey, error) {
	if len(signatures) != 1 {
		return nil, fmt.Errorf(`multi signature not supported`)
	}

	signature := signatures[0]
	jwk, err := c.getStonePublicKey(signature.Header.KeyID)
	if err != nil {
		return nil, fmt.Errorf(`failure refreshing public keys: %w`, err)
	}

	return jwk, nil
}

func (c *Client) getStonePublicKey(id string) (*jose.JSONWebKey, error) {
	if key := c.StonePublicKeys.Get(id); key != nil {
		return key, nil
	}

	if err := c.refreshPublicKeys(); err != nil {
		return nil, fmt.Errorf(`failure refreshing stone public keys: %w`, err)
	}

	return c.StonePublicKeys.Get(id), nil
}

func (c *Client) refreshPublicKeys() error {
	keysURL, err := c.ApiBaseURL.Parse(stonePublicKeysEndpoint)
	if err != nil {
		return fmt.Errorf(`failure parsing endpoint: %w`, err)
	}

	response, err := c.client.Get(keysURL.String())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	type responseBody struct {
		Keys []jose.JSONWebKey `json:"keys"`
	}
	var r responseBody
	if err = json.Unmarshal(body, &r); err != nil {
		return err
	}

	for i := range r.Keys {
		c.StonePublicKeys[r.Keys[i].KeyID] = &r.Keys[i]
	}

	return nil
}

func (c *Client) DecryptJWE(encryptedBody string) ([]byte, error) {
	jwe, err := jose.ParseEncrypted(encryptedBody)
	if err != nil {
		return nil, err
	}

	data, err := jwe.Decrypt(c.privateKey)
	if err != nil {
		return nil, err
	}

	return data, nil
}
