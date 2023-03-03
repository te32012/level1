package tasks

import (
	"fmt"
	"regexp"
)

func isUnicle(s string) bool {

	var mp map[string]string = make(map[string]string)
	mp["A"] = "a"
	mp["B"] = "b"
	mp["C"] = "c"
	mp["D"] = "d"
	mp["E"] = "e"
	mp["F"] = "f"
	mp["G"] = "g"
	mp["H"] = "h"
	mp["I"] = "i"
	mp["J"] = "j"
	mp["K"] = "k"
	mp["L"] = "l"
	mp["M"] = "m"
	mp["N"] = "n"
	mp["O"] = "o"
	mp["P"] = "p"
	mp["R"] = "r"
	mp["S"] = "s"
	mp["T"] = "t"
	mp["U"] = "u"
	mp["V"] = "v"
	mp["W"] = "w"
	mp["X"] = "x"
	mp["Y"] = "y"
	mp["Z"] = "z"
	var m map[string]int = make(map[string]int)
	for _, r := range s {
		tmp := string(r)
		matched, _ := regexp.MatchString("^[A-Z]{1}", tmp)
		if matched {
			m[mp[tmp]] = m[mp[tmp]] + 1
		} else {
			m[tmp] = m[tmp] + 1
		}
	}
	var ans bool = true
	for _, v := range m {
		if v > 1 {
			ans = false
		}
	}
	return ans
}

func Task26() {
	fmt.Println(isUnicle("abc"))
	fmt.Println(isUnicle("aA"))
	fmt.Println(isUnicle("aAbc"))
	fmt.Println(isUnicle("acc"))
	fmt.Println(isUnicle("bcA"))
	fmt.Println(isUnicle("abc"))
}
