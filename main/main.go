package main

import (
	"Rot"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cipherText string
	scanner := bufio.NewScanner(os.Stdin)
	var RotType int
	flag := "y"
	var plainText string

	service := Rot.Welcome()

	for flag == "y" {

		if service == "e" {
			RotType = Rot.RotType()
			fmt.Println("Insert the text you want to encrypt")
			_, plainText = scanner.Scan(), scanner.Text()
			fmt.Println("The cipherText is :\n", Rot.Ciphering(plainText, RotType))
			flag,service = Rot.End()

		} else if service == "d" {
			RotType = Rot.RotType()
			fmt.Println("Insert the text you want to decrypt")
			_, cipherText = scanner.Scan(), scanner.Text()
			fmt.Println("The PlainText is :\n", Rot.Deciphering(cipherText, RotType))
			flag,service = Rot.End()

		} else {
			fmt.Println("Invalid Input")
			flag = "n"
		}
	}
}