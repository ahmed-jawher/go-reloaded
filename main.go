package main

import (
    "fmt"
    "os"
    "strings"
)

// ----------------------------------------
//  main
// ----------------------------------------

func main() {
    // x:= "it (cap) was the best of times ..."
    // fmt.Println(Check_All(x))

    if len(os.Args) != 3 {
        fmt.Println("Error: should be 2 arguments, Ex: sample.txt result.txt")
        os.Exit(1)
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    // HasSuffix to check if the file last 4 char is ".txt"
    if !strings.HasSuffix(outputFile, ".txt") {
        fmt.Println("Error: output file must be .txt")
        os.Exit(1)
    }
    if !strings.HasSuffix(inputFile, ".txt") {
        fmt.Println("Error: input file must be .txt")
        os.Exit(1)
    }

    // Read sample.txt and catch any errors
    content, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Println("Error: reading sample.txt: ", err)
        os.Exit(1)
    }

    // save the content of sample.txt
    x := strings.TrimSpace(string(content))

    // Check if the file is empty
    if x == "" {
        fmt.Println("Error: sample.txt is empty")
        return
    }

    // Check for invalid characters
    for _, y := range x {
        if y == '\n' || y == '\r' {
            continue
        }
        if y < 32 || y > 126 {
            fmt.Printf("Error: there is invalid character in the file")
            os.Exit(1)
        }
    }

    // Print the output in "result.txt"
    err = os.WriteFile(outputFile, []byte(Check_All(x)), 0644)
    if err != nil {
        fmt.Println("Error: writing to result.txt:", err)
        os.Exit(1)
    }
}
