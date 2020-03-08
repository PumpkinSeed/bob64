package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

const (
	welcome = "bob64 - Interactive base64 encoder\n---"
	exitStr = "exit"
	bye = "Bye :)"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(welcome)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		fmt.Println(base64Encode(text))
	}
}

func base64Encode(text string) string {
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	if text == exitStr {
		fmt.Println(bye)
		os.Exit(0)
	}
	return base64.StdEncoding.EncodeToString([]byte(text))
}

