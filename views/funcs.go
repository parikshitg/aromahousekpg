package views

import (
	"fmt"
	"html/template"
)

// HTMLString renders HTML from string
func HTMLString(s string) interface{} {
	return template.HTML(s)
}

// IsEven returns true if nubmer is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd returns true if nubmer is odd
func IsOdd(n int) bool {
	return n%2 == 1
}

// FirstChar returns 1st character from the stricng
func FirstChar(s string) string {
	chars := []rune(s)
	if len(chars) > 0 {
		return string(chars[0])
	}
	return ""
}

// Image returns placeholder image if image is not found
func Image(url string, width, height int) string {
	if url != "" {
		return url
	}
	return fmt.Sprintf("/placeholder/%dx%d.png", width, height)
}
