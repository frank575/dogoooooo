// !!沒什麼內容的正則範例
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
	fmt.Println(r.FindAllString("hello hello", -1)) // [llo ]
}
