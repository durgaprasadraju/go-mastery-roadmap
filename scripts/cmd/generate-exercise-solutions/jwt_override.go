package main

import "fmt"

func jwtExerciseSpec(m module) topicSpec {
	body := `package solutions

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
	Sub string ` + "`json:\"sub\"`" + `
	Exp int64  ` + "`json:\"exp\"`" + `
}

// CreateToken builds unsigned JWT for learning — use a JWT library in production.
func CreateToken(sub string, ttl time.Duration) (string, error) {
	if sub == "" {
		return "", ErrInvalidToken
	}
	header := base64.RawURLEncoding.EncodeToString([]byte(` + "`{\"alg\":\"none\",\"typ\":\"JWT\"}`" + `))
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
}`
	test := fmt.Sprintf(`package solutions_test

import (
	"testing"
	"time"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestCreateAndParseJWT(t *testing.T) {
	tok, err := solutions.CreateToken("user-42", time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := solutions.ParseClaims(tok)
	if err != nil || claims.Sub != "user-42" {
		t.Fatalf("claims=%%+v err=%%v", claims, err)
	}
}

func TestParseBearer(t *testing.T) {
	tok, ok := solutions.ParseBearer("Bearer abc.def.ghi")
	if !ok || tok != "abc.def.ghi" {
		t.Fatalf("tok=%%q ok=%%v", tok, ok)
	}
}

func TestExpiredToken(t *testing.T) {
	tok, err := solutions.CreateToken("u", -time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := solutions.ParseClaims(tok); err == nil {
		t.Fatal("expected expired")
	}
}
`, m.slug)
	interview := `package solutions

import "strings"

// InterviewChallengeSolution validates JWT has three base64url segments.
func InterviewChallengeSolution(token string) bool {
	parts := strings.Split(token, ".")
	return len(parts) == 3 && parts[0] != "" && parts[1] != ""
}`
	interviewTest := fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution("a.b.") {
		t.Fatal("expected valid structure")
	}
}
`, m.slug)
	return topicSpec{
		summary:           "**JWT** — create, parse, and validate JSON Web Tokens.",
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Validate JWT structure and expiry",
		interviewSolution: interview,
		interviewTest:     interviewTest,
		benchCall:         `solutions.ParseBearer("Bearer x.y.z")`,
	}
}
