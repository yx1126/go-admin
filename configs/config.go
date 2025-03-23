package config

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

func init() {
	conf, err := ini.Load("./conf.ini")
	if err != nil {
		fmt.Printf("Failed to read file %v\n", err)
		os.Exit(1)
	}
	fmt.Print("config init", conf)
}
