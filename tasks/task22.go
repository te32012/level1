package tasks

import (
	"fmt"
	"math/big"
)

func Task22() {
	var a *big.Int = big.NewInt(0)
	var b *big.Int = big.NewInt(0)
	var z *big.Int = big.NewInt(0)
	a.SetString("1000000000000000000", 10)
	b.SetString("2000000000000000001", 10)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println((z.Add(a, b)).String()) // Сумма
	fmt.Println((z.Sub(b, a)).String()) // Вычитание
	fmt.Println((z.Div(b, a)).String()) // Деление
	fmt.Println((z.Mul(a, b)).String()) // Умножение
}
