package base

import "testing"

func TestModuleVersion(t *testing.T) {
	v, err := ModuleVersion("golang.org/x/mod")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestParrotMod(t *testing.T) {
	out := ParrotMod()
	t.Log(out)
}
