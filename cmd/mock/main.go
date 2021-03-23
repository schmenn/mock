package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Schmenn/mock"
	"github.com/atotto/clipboard"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var nr string
	if runtime.GOOS == "windows" {
		nr = "\r\n"
	} else {
		nr = "\n"
	}

	fmt.Print("Enter text: ")
	text, _ := r.ReadString('\n')
	text = strings.Replace(text, nr, "", -1)
	mockOut := mock.Mock(text)
	fmt.Printf("Copied to clipboard: %s", mockOut)

	if err := clipboard.WriteAll(mockOut); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
