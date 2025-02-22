package Rot

import (
	"fmt"
	"os"
)

func Ciphering(c string, rot int) string {
	var cipherText string
	var num int
	for _, plainText := range c {
		if plainText >= 'a' && plainText <= 'z' {
			plainText = plainText - 'a'
			num = int(plainText) + rot
			if num >= 26 {
				num -= 26
			}
			cipherText += string(rune(num + 'a'))
		} else if plainText >= 'A' && plainText <= 'Z' {
			plainText = plainText - 'A'
			num = int(plainText) + rot
			if num >= 26 {
				num -= 26
			}
			cipherText += string(rune(num + 'A'))

		} else {
			cipherText += string(plainText)
		}
	}
	return cipherText
}

func Deciphering(c string, rot int) string {
	var Text string
	var num int
	for _, cipherText := range c {
		if cipherText >= 'a' && cipherText <= 'z' {
			cipherText = cipherText - 'a'
			num = int(cipherText) - rot
			if num < 0 {
				num += 26
			}
			Text += string(rune(num + 'a'))
		} else if cipherText >= 'A' && cipherText <= 'Z' {
			cipherText = cipherText - 'A'
			num = int(cipherText) - rot
			if num < 0 {
				num += 26
			}
			Text += string(rune(num + 'A'))

		} else {
			Text += string(cipherText)
		}
	}
	return Text
}

func Welcome() string {
	var service string
	fmt.Println("Welcome to cipher Rot")
	fmt.Println("This tool help you to encrypt and decrypt any type of Rot")
	fmt.Println("\nPlease chose what you want encryption or decryption (Insert e or d )")
	fmt.Scan(&service)
	return service
}

func RotType() int {
	var rotType int
	fmt.Println("What type of Rot you want (1 - 25) ?")
	fmt.Println("Please insert the number only like '14'")
	fmt.Scan(&rotType)
	for( rotType > 25 || rotType < 1){
		fmt.Println("Invalid Number\nPlease insert the numbers between (1-25)")
		fmt.Scan(&rotType)
	}
	return rotType
}

func End() (string, string) {
	var flag string
	var service string
	fmt.Println("Do you want to encrypt or decrypt more things ?")
	fmt.Println("Input ( y or n )")
	fmt.Scan(&flag)
	if flag == "y" {
		fmt.Println("\nPlease chose what you want encryption or decryption (Insert e or d )")
		fmt.Scan(&service)
	}else{
		os.Exit(1);
	}
	return flag, service
}
