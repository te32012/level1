package tasks

import (
	"fmt"
)

/*

Предпочительно минимизировать в коде использование глобальных переменных.
var justString string

func someFunc() {
	Ссылка на вырезку из строки (подстроку большой строки) будет мешать сборщику мусора очистить память
	v := createHugeString(1 << 10)
	Нужно обратить внимание, что строки - неизменяемый тип данных (в случае v = v + "a" будет выделенно дополнительные (1 << 10) + 1  памяти еще до удаления старой строки)
	Перед получением среза нужно убедиться, что строка имеет длину (len > 100)
  justString = v[:100]
}

Под капотом строки используют фиксированный (неизменяемый) срез byte

func main() {
  someFunc()
}

*/

func Task15() {
	hugeString := `My mistress' eyes are nothing like the sun*
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

	slise := hugeString[:5] // Ссылка на подстроку hugeString
	fmt.Println(hugeString)
	fmt.Println(slise)

}
