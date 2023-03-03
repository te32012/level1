package tasks

import (
	"fmt"
	"strings"
)

func Task20() {
	var text string = "snow dog sun"
	var words []string = strings.Split(text, " ")
	var sb strings.Builder
	for x := len(words) - 1; x >= 0; x-- {
		sb.WriteString(words[x])
		sb.WriteString(" ")
	}
	fmt.Println(sb.String())
}
