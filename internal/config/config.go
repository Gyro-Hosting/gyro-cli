// Package config stores named profiles in ~/.config/gyro/config.toml (0600), aws-style.
//
//	default_profile = "default"
//
//	[profiles.default]
//	api_url = "https://www.gyrohosting.com/api/v1/client"
//	token = "wck_..."            # or token_storage = "keyring" (token lives in the OS keychain)
//	output = "table"
//	lang = "tr"
//
// Precedence: command-line flag > environment (GYRO_TOKEN, GYRO_PROFILE, GYRO_API_URL, GYRO_OUTPUT,
// GYRO_LANG) > profile > built-in default.
package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/BurntSushi/toml"
	"github.com/zalando/go-keyring"
)

const keyringService = "gyro-cli"

// Profile is one named set of credentials and preferences.
type Profile struct {
	APIURL       string `toml:"api_url,omitempty"`
	Token        string `toml:"token,omitempty"`
	TokenStorage string `toml:"token_storage,omitempty"` // "" / "file" | "keyring"
	Output       string `toml:"output,omitempty"`
	Lang         string `toml:"lang,omitempty"`
	SSHUser      string `toml:"ssh_user,omitempty"`
	SSHIdentity  string `toml:"ssh_identity,omitempty"`
}

// File is the whole config file.
type File struct {
	DefaultProfile string              `toml:"default_profile,omitempty"`
	Profiles       map[string]*Profile `toml:"profiles"`
	path           string
}

// Path returns the config file location (GYRO_CONFIG_FILE overrides).
func Path() string {
	if p := os.Getenv("GYRO_CONFIG_FILE"); p != "" {
		return p
	}
	if x := os.Getenv("XDG_CONFIG_HOME"); x != "" {
		return filepath.Join(x, "gyro", "config.toml")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, ".config", "gyro", "config.toml")
}

// Load reads the config file; a missing file yields an empty config.
func Load() (*File, error) {
	f := &File{Profiles: map[string]*Profile{}, path: Path()}
	b, err := os.ReadFile(f.path)
	if errors.Is(err, os.ErrNotExist) {
		return f, nil
	}
	if err != nil {
		return nil, err
	}
	if _, err := toml.Decode(string(b), f); err != nil {
		return nil, fmt.Errorf("%s: %w", f.path, err)
	}
	if f.Profiles == nil {
		f.Profiles = map[string]*Profile{}
	}
	return f, nil
}

// FilePath is where Save writes.
func (f *File) FilePath() string { return f.path }

// InsecurePerms reports a config file readable by group/others (unix only).
func (f *File) InsecurePerms() bool {
	if runtime.GOOS == "windows" {
		return false
	}
	st, err := os.Stat(f.path)
	return err == nil && st.Mode().Perm()&0o077 != 0
}

// Save writes the file atomically with mode 0600 (directory 0700).
func (f *File) Save() error {
	if err := os.MkdirAll(filepath.Dir(f.path), 0o700); err != nil {
		return err
	}
	var buf bytes.Buffer
	buf.WriteString("# gyro CLI configuration. Keep this file private (mode 0600).\n")
	if err := toml.NewEncoder(&buf).Encode(f); err != nil {
		return err
	}
	tmp := f.path + ".tmp"
	if err := os.WriteFile(tmp, buf.Bytes(), 0o600); err != nil {
		return err
	}
	_ = os.Chmod(tmp, 0o600)
	return os.Rename(tmp, f.path)
}

// ProfileNames returns the sorted profile names.
func (f *File) ProfileNames() []string {
	var out []string
	for k := range f.Profiles {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// Get returns a profile, creating an empty one in memory when missing.
func (f *File) Get(name string) *Profile {
	p, ok := f.Profiles[name]
	if !ok {
		p = &Profile{}
		f.Profiles[name] = p
	}
	return p
}

// ActiveName resolves the profile name: flag > GYRO_PROFILE > default_profile > "default".
func (f *File) ActiveName(flag string) string {
	switch {
	case flag != "":
		return flag
	case os.Getenv("GYRO_PROFILE") != "":
		return os.Getenv("GYRO_PROFILE")
	case f.DefaultProfile != "":
		return f.DefaultProfile
	}
	return "default"
}

// TokenFor returns the stored token of a profile (file or keychain) and where it came from.
func (f *File) TokenFor(name string) (token, source string, err error) {
	p, ok := f.Profiles[name]
	if !ok {
		return "", "", nil
	}
	if p.TokenStorage == "keyring" {
		t, err := keyring.Get(keyringService, name)
		if err != nil {
			return "", "keyring", fmt.Errorf("keychain: %w", err)
		}
		return t, "keyring", nil
	}
	if p.Token != "" {
		return p.Token, "config", nil
	}
	return "", "", nil
}

// SetToken stores a token in the file or in the OS keychain.
func (f *File) SetToken(name, token string, useKeyring bool) error {
	p := f.Get(name)
	if useKeyring {
		if err := keyring.Set(keyringService, name, token); err != nil {
			return fmt.Errorf("keychain: %w", err)
		}
		p.Token, p.TokenStorage = "", "keyring"
		return nil
	}
	if p.TokenStorage == "keyring" {
		_ = keyring.Delete(keyringService, name)
	}
	p.Token, p.TokenStorage = token, ""
	return nil
}

// Delete removes a profile (and its keychain entry).
func (f *File) Delete(name string) {
	if p, ok := f.Profiles[name]; ok && p.TokenStorage == "keyring" {
		_ = keyring.Delete(keyringService, name)
	}
	delete(f.Profiles, name)
	if f.DefaultProfile == name {
		f.DefaultProfile = ""
	}
}
