package openbank

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (c *Client) ConsentLink(sessionID string) (string, error) {
	claims := c.consentClaims(sessionID)
	tokenString, err := c.generateToken(claims)
	if err != nil {
		return "", err
	}

	pathURL := fmt.Sprintf("/#/consent?type=consent&client_id=%s&jwt=%s", c.ClientID, tokenString)
	u, err := c.AccountURL.Parse(pathURL)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func (c *Client) consentClaims(sessionID string) jwt.MapClaims {
	now := time.Now()
	if sessionID == "" {
		sessionID = uuid.New().String()
	}
	claims := jwt.MapClaims{
		"aud":              "accounts-hubid@openbank.stone.com.br",
		"client_id":        c.ClientID,
		"exp":              now.Add(time.Hour * time.Duration(2)).Unix(),
		"iat":              now.Unix(),
		"iss":              c.ClientID,
		"jti":              sessionID,
		"nbf":              now.Unix(),
		"redirect_uri":     c.ConsentRedirectURL,
		"session_metadata": map[string]string{"client_session": sessionID},
		"type":             "consent",
	}
	return claims
}
