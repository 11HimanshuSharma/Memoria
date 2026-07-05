package repository
import (
	"testing"
)

func TestDiscover(t *testing.T) {
	repo, err := Discover(".")
	if err != nil {
		t.Fatal(err)
	}
	if repo.Root == "" {
		t.Fatal("repository root should not be empty")
	}
	if repo.Name == "" {
		t.Fatal("repository name should not be empty")
	}

	if repo.Module == nil {
		t.Fatal("module should not be nil")
	}
	
	if repo.Module.Name == "" {
		t.Fatal("module name should not be empty")
	}
	
	if repo.Git == nil {
		t.Fatal("git info should not be nil")
	}
}