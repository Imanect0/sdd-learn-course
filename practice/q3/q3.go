package q3

import "fmt"

type Person struct {
	Name string
	Age int
}

func GetPersonInfo(p Person) string {
	return fmt.Sprint(p)
}

// 以上のコードを変更してはいけません。
// GetPersonInfo関数を実行した時に、
// "Name is <name>, Age is <age>"という文字列が出力されるようにしてください。
// <name>と<age>にはそれぞれPersonのNameとAgeが入ります。
// 例：GetPersonInfo(Person{"Taro", 20})の戻り値は"Name is Taro, Age is 20"となります。

// 以下にコードを記述して下さい。
