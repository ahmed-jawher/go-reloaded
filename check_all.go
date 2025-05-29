package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
	"regexp"
)

var Regex_For_Split = regexp.MustCompile(`[A-Za-z]+'[A-Za-z]+|\w+|\([^()]+\)|[^\s\w()]+`)

func Check_All(txt string) string {

    // split the text into words and punctuation
    words := Regex_For_Split.FindAllString(txt, -1)

    var Final_txt []string

    // loop to go word by word
    for i := 0; i < len(words); i++ {

        //------------------------------------ (cap, x) -----------------------------------//
        if cap_Regex.MatchString(words[i]) {
            Final_txt = cap_x(words[i], Final_txt)
            // continue to skip (cap, x)
            continue
        }

        //------------------------------------ (up, x) ------------------------------------//
        if up_Regex.MatchString(words[i]) {
            Final_txt = up_x(words[i], Final_txt)
            // continue to skip (up, x)
            continue
        }

        //------------------------------------ (low, x) -----------------------------------//
        if low_Regex.MatchString(words[i]) {
            Final_txt = low_x(words[i], Final_txt)
            // continue to skip (low, x)
            continue
        }

        //------------------------------------ (hex) --------------------------------------//
        if words[i] == "(hex)" {
            if i == 0 {
                fmt.Println("Error: no word before (hex)")
                os.Exit(1)
            }

            // check if the previous word is a valid hex number
            num, err := strconv.ParseInt(words[i-1], 16, 64)
            if err != nil {
                fmt.Println("Error: invalid hex value: ", words[i-1])
                os.Exit(1)
            }

            // format the previous word
            Final_txt[len(Final_txt)-1] = strconv.FormatInt(num, 10)

            // skip "(hex)"
            continue
        }

        //------------------------------------ (bin) --------------------------------------//
        if words[i] == "(bin)" {
            if i == 0 {
                fmt.Println("Error: no word before (bin)")
                os.Exit(1)
            }

            // check if the previous word is a valid bin number
            num, err := strconv.ParseInt(words[i-1], 2, 64)
            if err != nil {
                fmt.Println("Error: invalid bin value: ", words[i-1])
                os.Exit(1)
            }

            // format the previous word
            Final_txt[len(Final_txt)-1] = strconv.FormatInt(num, 10)

            // skip "(bin)"
            continue
        }

        //------------------------------------ (up) ---------------------------------------//
        if words[i] == "(up)" {
            if i == 0 {
                fmt.Println("Error: no word before (up)")
                os.Exit(1)
            }
            Final_txt = up(words[i-1], Final_txt)
            // skip "(up)"
            continue
        }

        //------------------------------------ (cap) --------------------------------------//
        if words[i] == "(cap)" {
            if i == 0 {
                fmt.Println("Error: no word before (cap)")
                os.Exit(1)
            }
            Final_txt = cap(words[i-1], Final_txt)
            // skip "(cap)"
            continue
        }

        //------------------------------------ (low) --------------------------------------//
        if words[i] == "(low)" {
            if i == 0 {
                fmt.Println("Error: no word before (low)")
                os.Exit(1)
            }
            Final_txt = low(words[i-1], Final_txt)
            // skip "(low)"
            continue
        }

        //----------------------------------- (a to an) -----------------------------------//
        if strings.ToLower(words[i]) == "a" && i+1 < len(words) {
            if words[i+1][0] == 'a' || words[i+1][0] == 'A' || words[i+1][0] == 'e' || words[i+1][0] == 'E' || words[i+1][0] == 'i' || words[i+1][0] == 'I' || words[i+1][0] == 'o' || words[i+1][0] == 'O' || words[i+1][0] == 'u' || words[i+1][0] == 'U' || words[i+1][0] == 'h' || words[i+1][0] == 'H' {
                words[i] += "n"
            }
        }

        //----------------------------------- (an to a) -----------------------------------//
        if strings.ToLower(words[i]) == "an" && i+1 < len(words) {
            if !(words[i+1][0] == 'a' || words[i+1][0] == 'A' || words[i+1][0] == 'e' || words[i+1][0] == 'E' || words[i+1][0] == 'i' || words[i+1][0] == 'I' || words[i+1][0] == 'o' || words[i+1][0] == 'O' || words[i+1][0] == 'u' || words[i+1][0] == 'U' || words[i+1][0] == 'h' || words[i+1][0] == 'H') {
                // remove the last character from the word "an"
                words[i] = words[i][:len(words[i])-1]
            }
        }

        //---------------------------------------------------------------------------------//

        // if the word is not a tag, append it to Final_txt
        Final_txt = append(Final_txt, words[i])
    }

    // Merge punctuation with the previous word
    Final_txt = punctuation(Final_txt)
    return strings.Join(Final_txt, " ")
}
