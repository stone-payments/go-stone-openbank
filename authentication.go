package openbank

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

func (c *Client) Authenticate() error {
	if c.validToken() {
		return nil
	}

	claims := c.authClaims()
	tokenString, err := c.generateToken(claims)
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Set("client_assertion", tokenString)
	data.Set("client_id", c.ClientID)
	data.Set("grant_type", "client_credentials")

	u, err := c.AccountURL.Parse("/auth/realms/stone_bank/protocol/openid-connect/token")
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("user-agent", c.UserAgent)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	var token oauth2.Token
	_, err = c.Do(req, &token)
	if err != nil {
		return err
	}
	ctx := context.Background()
	config := &oauth2.Config{}
	ts := config.TokenSource(ctx, &token)

	c.m.Lock()
	defer c.m.Unlock()
	c.client = oauth2.NewClient(ctx, ts)
	c.token = token

	return nil
}

func (c *Client) authClaims() jwt.MapClaims {
	now := time.Now()
	u, _ := c.AccountURL.Parse("/auth/realms/stone_bank")
	claims := jwt.MapClaims{
		"aud":       u.String(),
		"client_id": c.ClientID,
		"exp":       now.Add(time.Hour * time.Duration(2)).Unix(),
		"iat":       now.Unix(),
		"jti":       uuid.New().String(),
		"iss":       c.ClientID,
		"nbf":       now.Unix(),
		"realm":     "stone_bank",
		"sub":       c.ClientID,
	}
	return claims
}

func (c *Client) validToken() bool {
	if !c.token.Valid() {
		return false
	}

	src := strings.Split(c.token.AccessToken, ".")
	if len(src) != 3 {
		return false
	}

	if l := len(src[1]) % 4; l > 0 {
		src[1] += strings.Repeat("=", 4-l)
	}

	decoded, err := base64.URLEncoding.DecodeString(src[1])
	if err != nil {
		c.log.Error(fmt.Errorf("decoding base64 error %s", err))
		return false
	}

	var output tokenData
	err = json.Unmarshal(decoded, &output)
	if err != nil {
		c.log.Error(fmt.Errorf("decoding json error %s", err))
		return false
	}

	tm := time.Unix(int64(output.Exp), 0)
	remainder := tm.Sub(time.Now())
	if remainder < 30 {
		return false
	}

	return true
}

type tokenData struct {
	Exp int `json:"exp"`
}
