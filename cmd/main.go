package main

import (
	"fmt"
	"os"
)

func main() {
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
		fmt.Println("init")
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
