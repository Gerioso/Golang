package main

import "fmt"

func main() {
	str := "someString"

	fmt.Println("first val:", str)

	dontCahngeStr(str)

	fmt.Println("after dontCahngeStr val:", str)

	fmt.Println("val addr in main:", &str)

	cahngeStr(&str)

	fmt.Println("after cahngeStr val:", str)
	fmt.Println("val addr in main:", &str)
}

func dontCahngeStr(str string) {
	str = "nextStr"
}

func cahngeStr(str *string) {
	fmt.Println("addr addr in cahngeStr:", &str)

	*str = "nextStr"
}
