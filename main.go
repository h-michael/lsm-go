package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h-michael/lsm/pkgmgr"
)

func main() {
	flag.Parse()
	args := flag.Args()
	lsName := args[0]

	if err := pkgmgr.InstallViaNpm(lsName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
