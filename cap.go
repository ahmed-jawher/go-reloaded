package main

import (
    "fmt"
    "os"
    "strings"
)

//------------------------------------ (cap) --------------------------------------//
func cap(token string, Final_txt []string) []string {

    // check if the previous word is a valid word
    if !isAlpha(token) {
        fmt.Println("Error: invalid word before (cap): ", token)
        os.Exit(1)
    }

    // format the previous word
    Final_txt[len(Final_txt)-1] = strings.ToUpper(token[0:1]) + strings.ToLower(token[1:])

    return Final_txt
}
