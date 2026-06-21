package solutions

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var ErrInvalidToken = errors.New("invalid jwt token")

// Claims represents simplified JWT payload.
type Claims struct {
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

// CreateToken builds unsigned JWT for learning — use a JWT library in production.
func CreateToken(sub string, ttl time.Duration) (string, error) {
	if sub == "" {
		return "", ErrInvalidToken
	}
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"none","typ":"JWT"}`))
	payload, err := json.Marshal(Claims{Sub: sub, Exp: time.Now().Add(ttl).Unix()})
	if err != nil {
		return "", err
	}
	body := base64.RawURLEncoding.EncodeToString(payload)
	return header + "." + body + ".", nil
}

// ParseClaims decodes payload from unsigned JWT.
func ParseClaims(token string) (Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return Claims{}, ErrInvalidToken
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return Claims{}, err
	}
	var c Claims
	if err := json.Unmarshal(raw, &c); err != nil {
		return Claims{}, err
	}
	if c.Exp < time.Now().Unix() {
		return Claims{}, ErrInvalidToken
	}
	return c, nil
}

// ParseBearer extracts token from Authorization header.
func ParseBearer(header string) (string, bool) {
	const p = "Bearer "
	if !strings.HasPrefix(header, p) {
		return "", false
	}
	tok := strings.TrimSpace(header[len(p):])
	return tok, tok != ""
}