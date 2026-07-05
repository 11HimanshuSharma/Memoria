package repository
import (
	"testing"
)

func TestDiscover(t *testing.T) {
	ctx, err := Discover(".")
	if err != nil {
		t.Fatal(err)
	}
	if ctx.Root == "" {
		t.Fatal("repository root should not be empty")
	}
	if ctx.Name == "" {
		t.Fatal("repository name should not be empty")
	}

	if ctx.Language != Go {
		t.Fatalf("expected language %q got %q", Go, ctx.Language)
	}

	if ctx.BuildSystem != GoModules {
		t.Fatalf("expected build system %q got %q", GoModules, ctx.BuildSystem)
	}

	if ctx.Module == "" {
		t.Fatal("module should not be empty")
	}
}