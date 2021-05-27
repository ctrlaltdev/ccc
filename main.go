package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	version              = "v1.0.0-alpha.1"
	commitMsgFile        = flag.String("f", "", "Path to the commit message file")
	shouldInit           = flag.Bool("init", false, "Install git hook in the current repository")
	shouldDisplayVersion = flag.Bool("version", false, "Display version of the binary")
	hookCmd              = "ccc -f $@"
)

func DisplayVersion() {
	fmt.Printf("Conventional Commit Checker\nCCC - %s\n", version)
	os.Exit(0)
}

func InitHook() error {
	cwd, err := os.Getwd()
	CheckErr(err)

	gitFolder, err := FindFolderInParent(cwd, ".git")
	CheckErr(err)

	hooksFolder := filepath.Join(gitFolder, "hooks")

	err = CreateFolderIfNotExists(hooksFolder, 0755)
	CheckErr(err)

	hookFile := filepath.Join(hooksFolder, "commit-msg")

	data := ReadFile(hookFile)

	if !strings.Contains(data, hookCmd) {
		WriteFile(hookFile, fmt.Sprintf("%s\n%s\n", data, hookCmd))
		fmt.Println("Conventional Commit Checker Git Hook Installed")
	}

	fmt.Println("Conventional Commit Checker Git Hook Already Installed")

	return nil
}

func main() {
	flag.Parse()

	if *shouldDisplayVersion {
		DisplayVersion()
	}

	if *shouldInit {
		err := InitHook()
		CheckErr(err)
		os.Exit(0)
	}

	if *commitMsgFile != "" {
		msg := ReadFile(*commitMsgFile)
		err := ParseCommit(msg)
		if err != nil {
			log.Println(err)
			log.Fatal(errors.New("commit message should follow the conventional commit format: <type>[optional scope]: <description>"))
		}
		os.Exit(0)
	}
}
