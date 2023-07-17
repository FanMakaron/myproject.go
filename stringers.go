// https://go.dev/tour/methods/18
package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s string
	for n, cel := range ip {
		if n < 3 {
			s += fmt.Sprintf("%v.", cel)
		} else {
			s += fmt.Sprintf("%v", cel)
		}
	}

	return s

}

// from mentor
/*
func (ip IPAddr) String() string {
	result := ""        // объявляем пустую строчку как угодно, можно например так
	for i := range ip { // проходим циклом по массиву байтов
		result += "." + strconv.Itoa(int(ip[i])) // и присоединяем точку с функцией конвертации интов в строку из стандартной библиотеки :)
	}
	return result[1:] // поскольку string - это просто алиас к []byte, то к ней тоже можно обращаться адресно - возвращаем все символы, кроме первой точки
}
*/

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
