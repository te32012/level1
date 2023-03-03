package tasks

import (
	"fmt"
	"unicode/utf8"
)

func Task19() {
	var text string = `My mistress' eyes are nothing like the sun*
	Coral is far more red than her lips' red
	If snow be white, why then her breasts are dun*
	If hairs be wires, black wires grow on her head 
	I have seen roses damasked, red and white
	But no such roses see I in her cheeks
	And in some perfumes is there more delight
	Than in the breath that from my mistress reeks
	I love to hear her speak, yet well I know
	That music hath a far more pleasing sound*
	I grant I never saw a goddess go
	My mistress when she walks treads on the ground 
	And yet, by heaven, I think my love as rare
	As any she belied with false compare`
	fmt.Println(reverseString(text))
}

func reverseString(s string) string {
	l := utf8.RuneCountInString(s)
	ans := make([]rune, l)
	index := 1
	for _, r := range s {
		ans[l-index] = r
		index++
	}
	return string(ans)
}
