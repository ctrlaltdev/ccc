package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type CType *string
type CScope *string
type CDesc *string

func ParseConventionalCommit(msg string) (cType CType, cScope CScope, cDesc CDesc, err error) {
	reType := regexp.MustCompile(`^(\w+)(?:\(\w+\))?!?:\s`)
	resType := reType.FindStringSubmatch(msg)

	reScope := regexp.MustCompile(`\(([\w\d\s]+)\)!?:\s`)
	resScope := reScope.FindStringSubmatch(msg)

	reDesc := regexp.MustCompile(`: ([\w\d\s]+)\n`)
	resDesc := reDesc.FindStringSubmatch(msg)

	if len(resType) > 1 {
		cType = &resType[1]
	} else {
		return nil, nil, nil, errors.New("could not identify type")
	}

	if len(resScope) > 1 {
		cScope = &resScope[1]
	}

	if len(resDesc) > 1 {
		cDesc = &resDesc[1]
	} else {
		return nil, nil, nil, errors.New("could not identify title")
	}

	return cType, cScope, cDesc, nil
}

func ParseCommit(msg string) error {
	_, _, _, err := ParseConventionalCommit(msg)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	file := os.Args[1]

	data, err := ioutil.ReadFile(file)
	CheckErr(err)

	msg := string(data)
	err = ParseCommit(msg)
	if err != nil {
		log.Println(err)
		log.Fatal(errors.New("commit message should follow the conventional commit format: <type>[optional scope]: <description>"))
	}
}
