package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/05-packages-modules/exercises/solutions"
)

func TestParseSemVer(t *testing.T) {
	v, err := solutions.ParseSemVer("1.2.3")
	if err != nil || v.Major != 1 || v.Minor != 2 || v.Patch != 3 {
		t.Fatalf("v=%+v err=%v", v, err)
	}
}
