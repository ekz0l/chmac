package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ekz0l/chmac/pkg/ifconfig"
	"log"
)

func main() {
	if err := changeMac(); err != nil {
		log.Fatal(err)
	}
}

var iFace = flag.String("i", "", "interface name")

func changeMac() error {
	if *iFace == "" {
		return errors.New("required -i flag")
	}

	exists, err := ifconfig.ExistIF(*iFace)
	if err != nil {
		return err
	}

	fmt.Println(exists)

	return nil
}
