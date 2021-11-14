package main

import (
	"cryptolab6/pkg/esign"
	"fmt"
)

func main() {
	fmt.Println("Введите текст для шифрования:")
	var txt string
	fmt.Scan(&txt)
	if esign.Sign(txt) {
		fmt.Println("СОВПАДАЕТ")
	} else {
		fmt.Println("НЕ СОВПАДАЕТ")
	}
}
