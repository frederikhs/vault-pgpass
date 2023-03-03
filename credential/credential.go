package credential

import (
	"context"
	"errors"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"gopkg.in/yaml.v3"
	"os"
)

type VaultCredentials struct {
	SecretPath string `yaml:"secretPath"`

	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`

	username string
	password string

	UsernameKey string `yaml:"usernameKey"`
	PasswordKey string `yaml:"passwordKey"`

	hydrated bool
}

// NewFromFile read an array of vault credentials from a file at path
func NewFromFile(path string) ([]*VaultCredentials, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var vc []*VaultCredentials
	err = yaml.Unmarshal(f, &vc)
	if err != nil {
		return nil, err
	}

	return vc, nil
}

// New creates a new struct of vault credentials
func New(secretPath string, hostname string, port int, database string, usernameKey string, passwordKey string) *VaultCredentials {
	return &VaultCredentials{
		SecretPath:  secretPath,
		Hostname:    hostname,
		Port:        port,
		Database:    database,
		UsernameKey: usernameKey,
		PasswordKey: passwordKey,
	}
}

// Hydrate fetches credentials from vault as specified by the configuration and maps them.
func (vc *VaultCredentials) Hydrate(kv *vault.KVv2) error {
	secret, err := kv.Get(context.Background(), vc.SecretPath)
	if err != nil {
		return fmt.Errorf("unable to read secret: %v", err)
	}

	vc.mapCredentials(secret.Data)

	vc.hydrated = true

	return nil
}

// mapCredentials maps username and password using the username- and password keys
func (vc *VaultCredentials) mapCredentials(data map[string]interface{}) {
	vc.username = fmt.Sprintf("%v", data[vc.UsernameKey])
	vc.password = fmt.Sprintf("%v", data[vc.PasswordKey])
}

// WritePGPass writes a pgpass
func (vc *VaultCredentials) WritePGPass() (string, error) {
	if !vc.hydrated {
		return "", errors.New("credentials not hydrated yet")
	}
	return fmt.Sprintf("%s:%d:%s:%s:%s", vc.Hostname, vc.Port, vc.Database, vc.username, vc.password), nil
}
