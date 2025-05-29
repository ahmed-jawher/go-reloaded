package pkg

import (
	"fmt"
	"os"
	"strings"
)

// ------------------------------------ (up) ---------------------------------------//
func Up(token string, Final_txt []string) []string {

	// check if the previous word is a valid word
	if !IsAlpha(token) {
		fmt.Println("Error: invalid word before (up): ", token)
		os.Exit(1)
	}

	// format the previous word
	Final_txt[len(Final_txt)-1] = strings.ToUpper(token)

	return Final_txt
}
