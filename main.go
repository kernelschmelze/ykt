package main

import (
	"fmt"
	"os"

	"github.com/yawn/ykoath"

	"github.com/pkg/errors"
)

func main() {
	var action string

	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	oath, err := ykoath.New()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer oath.Close()

	_, err = oath.Select()

	if err != nil {
		fmt.Println(errors.Wrapf(err, "failed to select"))
		os.Exit(1)
	}

	names, err := oath.List()

	if err != nil {
		fmt.Println(errors.Wrapf(err, "failed to list"))
		os.Exit(1)
	}

	switch action {

	case "list":

		for _, name := range names {
			fmt.Println(name)
		}

	case "get":

		if len(os.Args) < 3 {
			fmt.Println("name missing")
			os.Exit(1)
		}
		name := os.Args[2]
		calc, err := oath.Calculate(name, func(name string) error {
			fmt.Printf("*** PLEASE TOUCH YOUR YUBIKEY TO UNLOCK %q ***\n", name)
			return nil
		})
		if err != nil {
			fmt.Println(errors.Wrapf(err, "failed to calculate name for %q", name))
			os.Exit(1)
		}

		fmt.Println(calc)

	case "set":

		if len(os.Args) < 3 {
			fmt.Println("name missing")
			os.Exit(1)
		}
		name := os.Args[2]

		if len(os.Args) < 4 {
			fmt.Println("secret missing")
			os.Exit(1)
		}
		secret := os.Args[3]
		if err := oath.Put(name, ykoath.HmacSha1, ykoath.Totp, 6, []byte(secret), false); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("OK")

	case "del":
		if len(os.Args) < 3 {
			fmt.Println("name missing")
			os.Exit(1)
		}
		name := os.Args[2]
		if err := oath.Delete(name); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("OK")

	default:
		fmt.Println("list")
		fmt.Println("get name")
		fmt.Println("del name")
		fmt.Println("set name secret")
		os.Exit(1)
	}

	os.Exit(0)

}
