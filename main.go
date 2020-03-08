package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	welcomeEncoder = "bob64 - Interactive base64 encoder"
	welcomeDecoder = "bob64 - Interactive base64 decoder"
	exitStr = "exit"
	bye = "Bye :)"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var decode bool
	if len(os.Args) > 1 {
		decode = os.Args[1] == "-d"
	}
	if decode {
		fmt.Println(welcomeDecoder)
	} else {
		fmt.Println(welcomeEncoder)
	}
	fmt.Printf("To leave the interactive mode type 'exit'\n---\n")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		if decode {
			fmt.Println(base64Decode(text))
		} else {
			fmt.Println(base64Encode(text))
		}
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

func base64Decode(text string) string {
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	if text == exitStr {
		fmt.Println(bye)
		os.Exit(0)
	}
	result, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

