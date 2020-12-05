// toc#regexp:::基本正則#toc
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 台灣手機或家電格式
	telR, _ := regexp.Compile("^((\\+8869|09)|(0([2-8]|37|49|89|82|826|836))-?)\\d{8}$")
	tel := "09-25151515"
	fmt.Println(telR.MatchString(tel), tel) // false
	tel = "0925151515"
	fmt.Println(telR.MatchString(tel), tel) // true
	tel = "0836-25151515"
	fmt.Println(telR.MatchString(tel), tel) // true
	tel = "+886981689689"
	fmt.Println(telR.MatchString(tel), tel) // true
	tel = "+886-981689689"
	fmt.Println(telR.MatchString(tel), tel) // false

	r, _ := regexp.Compile("[A-z]{3}\\s+")
	// n -> 取幾筆，-1 為全部
	fmt.Println("r:", r.FindAllString("hello hello", -1)) // [llo ]
	fmt.Println("r:", r.FindString("he"))                 // ""

	r2, _ := regexp.Compile("am")
	r2Word := "i\n" +
		"      am\n" +
		"apple!" +
		"i\n" +
		"am\n" +
		"    guava!"
	fmt.Println("r2:", r2.FindAllIndex([]byte(r2Word), -1))
}
