package base

import (
	"context"
	"testing"
)

func TestRepo(t *testing.T) {
	r := NewRepo("https://github.com/go-parrot/parrot-layout.git", "main")
	if err := r.Clone(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := r.CopyTo(context.Background(), "/tmp/test_parrot_repo", "github.com/go-parrot/parrot-layout", nil); err != nil {
		t.Fatal(err)
	}
}
