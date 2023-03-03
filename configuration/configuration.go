package configuration

import (
	"errors"
	"github.com/frederikhs/vault-pgpass/credential"
	"gopkg.in/yaml.v3"
	"os"
)

type Configuration struct {
	Address string                         `yaml:"address"`
	Hosts   []*credential.VaultCredentials `yaml:"hosts"`
}

// LoadFromFile loads a configuration file from a file specified at path
func LoadFromFile(path string) (*Configuration, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Configuration
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		return nil, err
	}

	if len(c.Address) == 0 {
		return nil, errors.New("no vault address specified")
	}

	if len(c.Hosts) == 0 {
		return nil, errors.New("no hosts specified")
	}

	return &c, nil
}

// GenerateExample generates an example configuration file
func GenerateExample() ([]byte, error) {
	vc := []*credential.VaultCredentials{
		{
			SecretPath:  "production/postgres/X/X",
			Hostname:    "db.example.com",
			Port:        5432,
			Database:    "postgres",
			UsernameKey: "username-key",
			PasswordKey: "password-key",
		},
		{
			SecretPath:  "production/postgres/Y/Y",
			Hostname:    "db2.example.com",
			Port:        5432,
			Database:    "postgres",
			UsernameKey: "username-key",
			PasswordKey: "password-key",
		},
	}

	c := Configuration{
		Address: "https://127.0.0.1:8200",
		Hosts:   vc,
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		return nil, err
	}

	return b, nil
}
