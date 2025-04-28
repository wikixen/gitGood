package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)

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
		initCmd.Usage = func() {
			fmt.Println("Initialize a new, empty repository.")
		}
		if err := initCmd.Parse(os.Args[2:]); err != nil {
			err := createRepo(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
		}
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
