package main

import "strings"
import "fmt"

func main() {
	message := "this is my secret message"
	fmt.Println(cipher(message, 13))

	camel := "saveChangesInTheEditor"
	fmt.Println(camelCount(camel))
}

func camelCount(text string) int {
	count := 1
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			count++
		}
	}
	return count
}

func cipher(text string, key int) string {
	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	text = strings.ToUpper(text)
	var num int
	var ret strings.Builder
	for _, char := range text {
		for _, val := range alphabets {
			if char == val {
				num = strings.IndexRune(alphabets, char)
				num = num + key
				if num >= len(alphabets) {
					num = num - len(alphabets)
				} else if num < 0 {
					num = num + len(alphabets)
				}
			}
		}
		if char == ' ' {
			ret.WriteRune(' ')
		}
		ret.WriteByte(alphabets[num])
	}
	return ret.String()
}
