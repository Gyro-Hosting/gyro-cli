package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestRoundTripAndPrecedence(t *testing.T) {
	p := filepath.Join(t.TempDir(), "sub", "config.toml")
	t.Setenv("GYRO_CONFIG_FILE", p)
	t.Setenv("GYRO_PROFILE", "")
	f, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if err := f.SetToken("work", "wck_1", false); err != nil {
		t.Fatal(err)
	}
	f.Get("work").Output = "yaml"
	f.DefaultProfile = "work"
	if err := f.Save(); err != nil {
		t.Fatal(err)
	}
	st, _ := os.Stat(p)
	if runtime.GOOS != "windows" && st.Mode().Perm() != 0o600 { // Windows has no POSIX modes
		t.Fatalf("perm %v", st.Mode())
	}
	g, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if g.ActiveName("") != "work" || g.ActiveName("x") != "x" {
		t.Error("active name")
	}
	t.Setenv("GYRO_PROFILE", "env")
	if g.ActiveName("") != "env" {
		t.Error("env profile")
	}
	tok, src, _ := g.TokenFor("work")
	if tok != "wck_1" || src != "config" || g.Profiles["work"].Output != "yaml" {
		t.Error("token/output")
	}
	g.Delete("work")
	if len(g.Profiles) != 0 || g.DefaultProfile != "" {
		t.Error("delete")
	}
}
