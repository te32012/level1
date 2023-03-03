package tasks

/*
Интерфейс car с методом drive реализует налагаемые требования на новый класс
Класс simplecar - простая машина которая может ездить только по дорогам
Класс customizedСar - сперциальная машина, которая может ездить не толлько по дорогам
Метод customizedСar() создает из простой машины специальную машину (адаптирует машину к бездорожью)
*/

type cars interface {
	drive() string
}

type simpleСar struct {
}

func (sc *simpleСar) drive() string {
	return "drive on road"
}

type customizedСar struct {
	*simpleСar
}

func (cc *customizedСar) drivenottoroad() string {
	return "not " + cc.drive()
}

func (cc *customizedСar) customizedСar(c *simpleСar) cars {
	return &customizedСar{c}
}
