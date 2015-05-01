// main
package main

import (
	"fmt"
	"reflect"
)

func main() {
	str:=reflect.ValueOf(112).Int()
	fmt.Println("Hello World!"+str)
}
