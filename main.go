package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h-michael/lsm/installer"
)

func main() {
	flag.Parse()
	args := flag.Args()
	cmd := args[0]
	lsName := args[1]

	switch cmd {
	case "install":
		if err := installer.InstallViaNpm(lsName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "uninstall":
		if err := installer.UninstallViaNpm(lsName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
