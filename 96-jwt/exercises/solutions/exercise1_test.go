package solutions_test

import (
	"testing"
	"time"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/96-jwt/exercises/solutions"
)

func TestCreateAndParseJWT(t *testing.T) {
	tok, err := solutions.CreateToken("user-42", time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := solutions.ParseClaims(tok)
	if err != nil || claims.Sub != "user-42" {
		t.Fatalf("claims=%+v err=%v", claims, err)
	}
}

func TestParseBearer(t *testing.T) {
	tok, ok := solutions.ParseBearer("Bearer abc.def.ghi")
	if !ok || tok != "abc.def.ghi" {
		t.Fatalf("tok=%q ok=%v", tok, ok)
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
