package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

// upture the number from (cap, x)
var cap_Regex = regexp.MustCompile(`\(cap,\s*(-?\d+)\)`)

//------------------------------------ (cap, x) -----------------------------------//
func cap_x(token string, Final_txt []string) []string {

    // fmt.Println(" - ")
    // here we check if the word is "(cap, x)"
    matches := cap_Regex.FindStringSubmatch(token)
    if len(matches) != 2 {
        fmt.Println("Error: invalid ub format: ", token)
        os.Exit(1)
    }

    // save the value of x
    // and convert it from string to int
    x, err := strconv.Atoi(matches[1])
        if err != nil || x <= 0 {
            fmt.Println("Error: invalid or non-positive number of x , (cap, x)")
            os.Exit(1)
        }

        // check if there are enough words before (cap, x)
        if len(Final_txt) < x {
            fmt.Println("Error: not enough words before (cap, x)")
            os.Exit(1)
        }

        // here we add the words to Final_txt
        for j := len(Final_txt) - x; j < len(Final_txt); j++ {
            if !isAlpha(Final_txt[j]) {
                fmt.Println("Error: word before (cap, x) is not alphabetic: ", Final_txt[j])
                os.Exit(1)
            }
            Final_txt[j] = strings.ToUpper(string(Final_txt[j][0])) + strings.ToLower(Final_txt[j][1:])
        }

    return Final_txt
}
