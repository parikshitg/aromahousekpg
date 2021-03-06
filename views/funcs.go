package views

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

// HTMLString renders HTML from string
func HTMLString(s string) interface{} {
	str := RemoveDiv(s)
	return template.HTML(str)
}

// RemoveDiv renders HTML from string
func RemoveDiv(s string) string {
	str := strings.TrimSpace(s)
	if strings.HasPrefix(str, "<div>") && strings.HasSuffix(str, "</div>") {
		str = str[5 : len(str)-6]
	}
	return str
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

func Day(tm time.Time) int {
	return tm.Day()
}

func Month(tm time.Time) string {
	return fmt.Sprint(tm.Format("January"))
}

func Year(tm time.Time) int {
	return tm.Year()
}

func Numbers(n1, n2 int) []int {
	var numbers []int
	for i := n1; i <= n2-n1+1; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}
