package tasks

import (
	"strings"
)

//Tasks 2
/*
Given words first and second, consider occurrences in some text of the form "first second third", where second comes immediately after first, and third comes immediately after second.
For each such occurrence, add "third" to the answer, and return the answer.

Example 1:
Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
Output: ["girl","student"]
Example 2:
Input: text = "we will we will rock you", first = "we", second = "will"
Output: ["we","rock"]
*/

var itemTextExmp1 string = "alice is a good girl she is a good student"

var itemTextExmp2 string = "we will we will rock you"

func findlerOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	n := len(words)
	res := make([]string, 0, n)
	for i := 0; i+2 < n; i++ {
		if words[i] == first &&
			words[i+1] == second {
			res = append(res, words[i+2])
		}
	}
	return res
}