package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/72-tcp/exercises/solutions"
)

func TestNetworking(t *testing.T) {
	h, p, ok := solutions.ParseHostPort("localhost:8080")
	if !ok || h != "localhost" || p != "8080" {
		t.Fatalf("host=%s port=%s ok=%v", h, p, ok)
	}
	if solutions.BuildURL("api.local", "/v1") != "http://api.local/v1" {
		t.Fatal("build url failed")
	}
}
