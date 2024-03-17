func isValid(s string) bool {
    parentheses := []rune{}
    for _, c := range s {
        parentheses = append(parentheses, c)
        if (len(parentheses) < 2) {
            continue
        }
        if (c == ')') {
            if (parentheses[len(parentheses)-2] == '(') {
                parentheses = parentheses[:len(parentheses)-2]
            } else {
                break
            }
        } else if (c == ']') {
            if (parentheses[len(parentheses)-2] == '[') {
                parentheses = parentheses[:len(parentheses)-2]
            } else {
                break
            }
        } else if (c == '}') {
            if (parentheses[len(parentheses)-2] == '{') {
                parentheses = parentheses[:len(parentheses)-2]
            } else {
                break
            }
        }
    }
    return len(parentheses) == 0
}