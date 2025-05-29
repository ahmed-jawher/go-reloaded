package pkg

import (
	"fmt"
	"os"
	"strings"
)

// ------------------------------------ (cap) --------------------------------------//
func Cap(token string, Final_txt []string) []string {

	// check if the previous word is a valid word
	if !IsAlpha(token) {
		fmt.Println("Error: invalid word before (cap): ", token)
		os.Exit(1)
	}

	// format the previous word
	Final_txt[len(Final_txt)-1] = strings.ToUpper(token[0:1]) + strings.ToLower(token[1:])

	return Final_txt
}
