package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello World!\n")
	_, _ = writer.WriteString("God bless y'all!\n")

	err := writer.Flush()
	
	if err != nil {
		panic(err)
	}
}
