// toc#reflect:::基本反射#toc
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 詳盡教學 https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/
// 知乎教學(本範例也皆源於此) https://zhuanlan.zhihu.com/p/55038960
func main() {
	user := User{
		"frank", 25,
	}

	// 取得值類型
	userType := reflect.TypeOf(user)
	firstField := userType.Field(0)
	fmt.Println(1, userType.Name()) // User
	fmt.Println(2, userType.Kind()) // struct
	fmt.Println(3, firstField.Name) // Name
	fmt.Println(4, firstField.Type) // string
	fmt.Println(5, firstField.Tag)  // json:"name"

	helloWorldType := reflect.TypeOf("hello world")
	fmt.Println(6, helloWorldType.Name()) // string
	fmt.Println(7, helloWorldType.Kind()) // string

	fmt.Println("")

	// 改變既有值
	userValue := reflect.ValueOf(&user).Elem()
	userValue.FieldByName("Name").SetString("rose")
	fmt.Println(8, user) // {rose 25}

	// 創建新值
	userType2 := reflect.TypeOf(user)
	user2 := reflect.New(userType2).Elem()
	user2.FieldByName("Name").SetString("jeff")
	user2.FieldByName("Age").SetInt(30)
	fmt.Println(9, user)   // {rose 25}
	fmt.Println(10, user2) // {jeff 30}

	fmt.Println("")

	// 返回新值
	iUser := reflect.ValueOf(user).Interface()
	if originalUser, ok := iUser.(User); ok {
		fmt.Println(11, originalUser.Name, originalUser.Age) // rose 25
	}

	// make 實例創建
	intSlice := make([]int, 0)
	intMap := make(map[string]int)

	sliceType := reflect.TypeOf(intSlice)
	mapType := reflect.TypeOf(intMap)

	reflectSlice := reflect.MakeSlice(sliceType, 0, 0)
	reflectMap := reflect.MakeMap(mapType)

	v := 10
	rv := reflect.ValueOf(v)
	reflectSlice = reflect.Append(reflectSlice, rv)
	intSlice2 := reflectSlice.Interface().([]int)
	fmt.Println(12, intSlice, intSlice2) // [] [10]

	k := "hello"
	rk := reflect.ValueOf(k)
	reflectMap.SetMapIndex(rk, rv)
	intMap2 := reflectMap.Interface().(map[string]int)
	fmt.Println(13, intMap, intMap2) // map[] map[hello:10]

	// make 創建函數
	//funType := reflect.TypeOf()
}

func count(len int) (sum int) {
	for i := 0; i < len; i++ {
		sum++
	}
	return
}
