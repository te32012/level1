package tasks

import (
	"errors"
	"fmt"
)

func Task8() {
	var num int64 = 127
	var index int64 = 2
	var b int64 = 0
	f := func(number int64, b int64, index int32) (int64, error) {
		tmp := int64(1) << int64(index)
		switch b {
		case 0:
			fmt.Printf("Бит номер i = %v в числе %v был установлен, новое число %b \n", index, number, (number|tmp)^tmp)
			return (number | tmp) ^ tmp, nil
		case 1:
			fmt.Printf("Бит номер i = %v в числе %v был установлен, новое число %b \n", index, number, tmp|number)
			return tmp | number, nil
		default:
			return number, errors.New("мы можем установить в качестве бита только 0 или 1")
		}
	}
	ans, err := f(num, b, int32(index))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)
}
