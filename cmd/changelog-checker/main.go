package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	command := exec.Command("git", "tag")
	output, err := command.CombinedOutput()
	if err != nil {
		panic(err)
	}
	var versions []string
	versions = splitLines(string(output))
    if len(versions) < 1 {
        panic("There are no version tags.")
    } 
    latestLocalVersion := versions[0]
	fmt.Println(latestLocalVersion)
}

func splitLines(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
