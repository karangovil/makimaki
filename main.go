package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/karangovil/makimaki/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello %s! This is the MakiMaki programming language!\n",
        user.Username)
    fmt.Printf("Feel free to type in commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
