package main

import (
	"fmt"
	"os"
)

var (
	// Version is set the release script
	Version = "dev"
)

func main() {
	cmd, args := getCommand()

	switch cmd {
	default:
		fmt.Fprintf(os.Stderr, "command %q not implemented yet\n", cmd)
		os.Exit(1)
	case testCmdName:
		testCmd(args)
	case importCmdName:
		importCmd(args)
	case listCmdName:
		listCmd(args)
	case mvCmdName:
		mvCmd(args)
	case cpCmdName:
		cpCmd(args)
	case versionCmdName:
		versionCmd()
	}
}

func versionCmd() {
	fmt.Println(Version)
}
