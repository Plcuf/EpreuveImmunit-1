package controller

import "fmt"

func dechiffrage(code string) string {
	decryptage := ""
	for _, i := range code {
		letter := int(i)
		letter += 11
		letter += 9
		letter -= 712 % 77
		letter -= 3
		decryptage += string(letter)
	}
	fmt.Println(decryptage)
	return decryptage
}
