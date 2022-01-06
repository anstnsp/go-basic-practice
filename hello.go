package main

import (
	"fmt"
	"strings"
)

func plus(a, b int) int {
	return a + b
}

//리턴을 여러개 할수 있음.
func lenAndUpper(str string) (string, int) {
	return strings.ToUpper(str), len(str)
}

//naked => 리턴할 변수를 선언
func upperStringNaked(str string) (upperString string, length int) {
	length = len(str)
	upperString = strings.ToUpper(str)
	return
}

/**
defer - > defer함수는 function이 끝나고 실행하는 함수.
	아래는 빼기를 실행 후 print 출력 함수
*/
func minus(a int, b int) int {
	result := a - b
	defer fmt.Println(result)
	return result
}

func main() {

	fmt.Println("Hello Go")
	const constant string = "constant" //상수선언 const 변수 타입 = 값

	fmt.Println("상수 :" + constant)
	fmt.Println("plus호출:", plus(3, 4))

	fmt.Println(upperStringNaked("assdfsdf"))
	minus(5, 2)
}
