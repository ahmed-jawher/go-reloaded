package main

func punctuation(words []string) []string {
    var result []string
    is_open := true

    // loop to go word by word again
    for i := 0; i < len(words); i++ {

        //------------------------------- ( single quote ) -------------------------------//

        if words[i] == "'" {
            if is_open {
                if i+1 < len(words) {
                    words[i+1] = "'" + words[i+1]
                }
            } else {
                if len(result) > 0 {
                    result[len(result)-1] += "'"
                }
            }
            is_open = !is_open
            continue
        }

        //---------------------- ( Special case for (...) or (!?) ) ----------------------//

        // if it is "." "." "." , we merge them to "..."
        if i+2 < len(words) && words[i] == "." && words[i+1] == "." && words[i+2] == "." {
            if len(result) > 0 {
                result[len(result)-1] += "..."
            } else {
                result = append(result, "...")
            }
            // skip two dots
            i += 2
            continue
        }

        // if it is "!" "?", we merge them to "!?"
        if i+1 < len(words) && words[i] == "!" && words[i+1] == "?" {
            if len(result) > 0 {
                result[len(result)-1] += "!?"
            } else {
                result = append(result, "!?")
            }
            // skip "?" and "!"
            i += 1
            continue
        }

        //------------------------- ( if we fund a punctuation ) -----------------------//

        if isPunctuation(words[i]) {
            if len(result) > 0 {
                // here we go to the last word in result and join it with the punctuation
                result[len(result)-1] += words[i]
            } else {
                result = append(result, words[i])
            }
            continue
        }

        // if it is a normal word, add it
        result = append(result, words[i])
    }

    return result
}
