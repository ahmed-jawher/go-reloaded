package pkg

import "regexp"

// Check if it is ( punctuation ) or ( special char )
var punctRe = regexp.MustCompile(`^[^\w\s()]+$`)

func IsPunctuation(word string) bool {
	return punctRe.MatchString(word)
}

// ====================================================================================

func IsAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}
