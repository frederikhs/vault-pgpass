package main

import (
	"flag"
	"fmt"
	"github.com/frederikhs/vault-pgpass/configuration"
	vault "github.com/hashicorp/vault/api"
	"os"
)

func main() {
	g := flag.Bool("g", false, "generate an example configuration file")
	f := flag.String("f", "", "specify location for your configuration file")
	t := flag.String("t", "", "vault token to use")
	o := flag.String("o", "", "output file or omit for stdout")
	flag.Parse()

	if *g && *f != "" {
		flag.Usage()
		os.Exit(1)
	}

	if *g {
		b, err := configuration.GenerateExample()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *o != "" {
			if err = os.WriteFile(*o, b, 0644); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println(string(b))
		}

		return
	}

	if *t == "" {
		flag.Usage()
		os.Exit(1)
	}

	c, err := configuration.LoadFromFile(*f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config := vault.DefaultConfig()
	config.Address = c.Address

	vaultClient, err := vault.NewClient(config)
	if err != nil {
		fmt.Printf("unable to initialize Vault vaultClient: %v\n", err)
		os.Exit(1)
	}

	vaultClient.SetToken(*t)
	kv := vaultClient.KVv2("secret")

	var output []byte
	for _, host := range c.Hosts {
		err = host.Hydrate(kv)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		s, err := host.WritePGPass()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *o != "" {
			output = append(output, []byte(s+"\n")...)
		} else {
			fmt.Println(s)
		}
	}

	if *o != "" {
		err = os.WriteFile(*o, output, 0600)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
