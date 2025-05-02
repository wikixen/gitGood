package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	if len(os.Args) < 2 {
		fmt.Println("usage: git-good <command> [<args>]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		fmt.Println("add")
	case "cat-file":
		fmt.Println("cat-file")
	case "check-ignore":
		fmt.Println("check-ignore")
	case "checkout":
		fmt.Println("checkout")
	case "commit":
		fmt.Println("commit")
	case "hash-object":
		fmt.Println("hash-object")
	case "init":
		initFn(initCmd)
	case "log":
		fmt.Println("log")
	case "ls-files":
		fmt.Println("ls-files")
	case "ls-tree":
		fmt.Println("ls-tree")
	case "rev-parse":
		fmt.Println("rev-parse")
	case "rm":
		fmt.Println("rm")
	case "show-ref":
		fmt.Println("show-ref")
	case "status":
		fmt.Println("status")
	case "tag":
		fmt.Println("tag")
	default:
		fmt.Println("unknown command")
	}
}

func initFn(initCmd *flag.FlagSet) {
	initCmd.Usage = func() {
		fmt.Println("Initialize a new, empty repository.")
	}
	if len(os.Args) < 3 {
		err := createRepo(".")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if err := initCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal(err)
		}
		err := createRepo(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}
}
