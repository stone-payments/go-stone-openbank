package openbank

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

func (c *Client) generateToken(claims jwt.MapClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, err := t.SignedString(c.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
