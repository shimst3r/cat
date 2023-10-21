package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
)

// regex is the simplest regex to identify proper words, using two word boundaries.
var regex = regexp.MustCompile(`\b(\w+)\b`)

func main() {
	cmd := exec.Command("cat", os.Args[1:]...)

	// If there's only one argument, expect input to come from stdin instead of a file.
	// In that case, the standard input must be redirected to the command's stdin using
	// io.Copy().
	if len(os.Args) == 1 {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			io.Copy(stdin, os.Stdin)
		}()
	}

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Finally, make the program work by replacing stuff properly. 🐈
	newOut := regex.ReplaceAllString(string(out), "Meow")
	fmt.Printf("%s\n", newOut)
}
