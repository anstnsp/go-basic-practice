package main

import (
	"fmt"
)

/**
해당 소스들의 참고 및 Go 자료
1)http://golang.site/go/article/16-Go-%EA%B5%AC%EC%A1%B0%EC%B2%B4-struct
2)https://edu.goorm.io/lecture/2010/%25ED%2595%259C-%25EB%2588%2588%25EC%2597%2590-%25EB%2581%259D%25EB%2582%25B4%25EB%258A%2594-%25EA%25B3%25A0%25EB%259E%25AD-%25EA%25B8%25B0%25EC%25B4%2588
3)https://programmers.co.kr/learn/courses/13
4)https://go.dev/learn/
*/

//① 변수를 하나 선언
var num1 int

//② 같은 타입을 가지는 변수를 여러 개 선언
var num2, num3 int

//③ 여러 변수에 한 번에 값을 초기화 : 선언과 동시에 값을 초기화하면 타입을 명시할 필요가 없습니다.
var num4, num5, str1 = 4, 5, "example"

//④ 함수 밖에서는 :=를 쓸 수 없습니다.
//errorvar := str1

//⑤ 다른 타입을 가지는 변수를 여러 개 선언
var (
	i int
	b bool
	s string
)

func main() {

	fmt.Println("①", num1)
	fmt.Println("②", num2, num3)
	fmt.Println("③", num4, num5, str1)

	//④ 함수 안에서는 :=를 쓰면 var과 타입을 지정하지 않고 변수를 선언과 동시에 초기화할 수 있습니다.
	num6 := 6
	fmt.Println("④", num6)

	fmt.Println("⑤", i, b, s)

	var f, t int = 10, 20
	fmt.Println(f, t)

	i, j, k := 1, 2, 3
	fmt.Println(i, j, k)

	var str1, str2 string = "hello", "goorm"
	fmt.Println(str1, str2)

	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names { //range는 배열의 인덱스와 값을 리턴함.
		println(index, name)
	}

	상수선언과초기화()

	증감연산자()

	포인터연산자()

	콘솔입력함수()

	문자열표현()

	자료형의변환()

	함수()

	가변인자함수("a", "b", "c", "d")

	클로저()

	맵사용()

	패키지설명()
}

/**
괄호로 같이 묶여있는 상수들은 다른 형으로 초기화될 수 있습니다.
괄호 시작(과 괄호 마지막)의 위치는 상관 없지만 각 상수들은 개행하여 초기화해야 합니다. 개행하지 않고 초기화하면 에러가 발생합니다.
각 상수들 사이에 콤마(,)를 입력하면 안 됩니다. 입력하면 에러가 발생합니다.
묶어서 선언된 상수들 중에서 첫번째 값은 꼭 선언되어야 합니다. 선언되지 않은 값은 바로 전 상수의 값을 가집니다.
iota라는 식별자를 값으로 초기화하면 그 후에 초기화하지 않고 이어지는 상수들은 순서(index)가 값으로 저장됩니다.
*/
const (
	c1 = 10 //첫 번째 값은 무조건 초기화해야 함
	c2
	c3
	c4 = "goorm" //다른 형 선언 가능
	c5
	c6 = iota //c8까지 index값 저장
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func 상수선언과초기화() {
	fmt.Println()
	fmt.Println("## 상수선언과초기화 ##")
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)
}

func 증감연산자() {
	/**
	증감 연산자는 값을 1만큼 증가시키거나 감소시키는 연산자입니다.Go언어에서 증감 연산자를 사용할 때 주의할 점이 두 가지가 있습니다.
	첫 번째는 증감 연산자를 사용하고 동시에 대입할 수 없습니다.
	두 번째는 전위 연산을 할 수 없습니다. 아래 예시가 있습니다.
	num := count++ (X)
	++count (X)
	*/
	fmt.Println()
	fmt.Println("## 증감연산자 ##")
	count1, count2 := 1, 10.4

	count1++
	count2--

	fmt.Println("count1++ :", count1)
	fmt.Println("count2-- :", count2)
}

func 포인터연산자() {
	/**
	&	변수의 메모리 주소를 참조한다.
	*	포인터 변수에 저장된 메모리에 접근하여 값을 참조한다.
	*/
	fmt.Println()
	fmt.Println("## 포인터 연산자 ##")
	var num int = 5
	var pnum = &num

	fmt.Println("num : ", num)   //num 값  =5
	fmt.Println("pnum :", pnum)  //num의 메모리 주소 =  0xc0000b0028
	fmt.Println("pnum :", *pnum) //num의 주소로 메모리에 할당돼있는 값 접근 = 5

	*pnum++
	fmt.Println("num : ", num)   // 6
	fmt.Println("pnum :", *pnum) // 6
}

func 콘솔입력함수() {
	/**
	Scanln은 여러 값을 동시에 입력받을 수 있습니다.
	빈칸(스페이스바)으로 값을 구분하고 엔터(개행)를 입력하면 입력이 종료됩니다.
	입력받는 변수에 '&' 연산자를 붙여 입력받습니다. 물론 입력받는 변수는 미리 선언되어야 합니다.
	*/
	fmt.Println()
	fmt.Println("## 콘솔입력함수 ##")
	var num1, num2, num3 int

	fmt.Println("정수 3개를 입력하세요 :")
	fmt.Scanln(&num1, &num2, &num3)
	fmt.Println("입력하신 값 :", num1, num2, num3)
}

func 문자열표현() {
	/**
	  	첫 번째는 Back Quote(`문자열`)을 이용한 방법입니다.
	  	모양이 인용부호(' ')와 비슷해서 혼동할 수 있지만 다른 기호입니다.
	  	Back Quote로 둘러 싸인 문자열은 Raw String Literal이라고 부릅니다. 쉽게 말하자면,
	  	이 안에 있는 문자열은 어느 기호든 문자열 자체로 인식되는 Raw String 값이라는 것입니다.
	  	예를 들어, 개발자라면 익숙하게 알고 있는 이스케이프 시퀀스가 특별한 의미로 인식되지 않는다는 것입니다.
	  	개행의 의미를 가지고 있는 \n이 Back Quote에 있으면 개행 기능이 되지 않고 문자열 자체로 출력됩니다.
		두 번째는 이중인용부호("문자열")를 이용한 방법입니다.
	  	이중인용부호로 둘러싸인 문자열은 Interpreterd String literal이라고 부릅니다.
	  	쉽게 말해, 안에 있는 문자열에 이스케이프 시퀀스같은 문자열들은 특별한 의미로 해석돼 그 기능을 수행한다는 것입니다.
	  	그리고 이중인용부호를 사용할 시에 복수라인에 걸쳐 쓸 수 없습니다.
	  	코딩을할때 아무리 엔터를 치고 길게 쳐도 한 줄에 표현된다는 뜻입니다.
	  	하지만 \n과 같은 기호가 있으면 개행의 의미로 해석돼 여러 라인에 걸쳐 쓸 수 있습니다.
	*/
	fmt.Println()
	fmt.Println("## 문자열표현 ##")
	var rawLiteral string = `바로 실행해보면서 배우는 \n Golang`

	var interLiteral string = "바로 실행해보면서 배우는 \nGolang"

	plusString := "구름 " + "EDU\n" + "Golang"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	fmt.Println()
	fmt.Println(plusString)
}

func 자료형의변환() {
	/**
	Go언어에서는 형 변환을 할 때 변환을 명시적으로 지정해주어야합니다.
	예를 들어 float32에서 uint로 변환할 때, 암묵적 변환은 일어나지 않으므로 "uint(변수이름)"과 같이 반드시 변환을 지정해줘야합니다.
	만약 명시적인 지정이 없다면 런타임 에러가 발생합니다.
	*/
	fmt.Println()
	fmt.Println("## 자료형의변환 ##")
	var num int = 10
	var changef float32 = float32(num) //int형을 float32형으로 변환
	changei := int8(num)               //int형을 int8형으로 변환

	var str string = "goorm"
	changestr := []byte(str)  //바이트 배열
	str2 := string(changestr) //바이트 배열을 다시 문자열로 변환

	fmt.Println(num)
	fmt.Println(changef, changei)

	fmt.Println(str)
	fmt.Println(changestr)
	fmt.Println(str2)
}

/**
[Pass By Value]
위의 [1. 함수]의 예제에서는 msg의 값 "Hello" 문자열이 복사되어 함수 say()에 전달된다.
즉, 만약 파라미터 msg의 값이 say() 함수 내에서 변경된다하더라도 호출함수 main()에서의 msg 변수는 변함이 없다.
[Pass By Reference]
아래의 예제에서처럼 msg 변수앞에 & 부호를 붙이면 msg 변수의 주소를 표시하게 된다.
흔히 포인터라 불리우는 이 용법을 사용하면 함수에 msg 변수의 값을 복사하지 않고 msg 변수의 주소를 전달하게 된다.
피호출 함수 say()에서는 *string 과 같이 파라미터가 포인터임을 표시하고 이때 say 함수의 msg는
문자열이 아니라 문자열을 갖는 메모리 영역의 주소를 갖게 된다. msg 주소에 데이타를 쓰기 위해서는 *msg = "" 과 같이 앞에 *를 붙이는데 이를 흔히 Dereferencing 이라 한다.
*/
func 함수() {
	fmt.Println()
	fmt.Println("## 함수 ##")
	msg := "hello"
	say(&msg)
	println(msg) //변경된 메시지 출력
}

func say(msg *string) {
	println(*msg)
	*msg = "changed" //메시지 변경
}

//가변 파라미터를 갖는 함수를 호출할 때는 0개, 1개, 2개, 혹은 n개의 동일타입 파라미터를 전달할 수 있다.
func 가변인자함수(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func 클로저() {
	fmt.Println()
	fmt.Println("## 클로저 ##")
	next := nextValue()
	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func 맵사용() {
	fmt.Println()
	fmt.Println("## 맵사용 ##")
	var idMap map[int]string
	idMap = make(map[int]string)
	//추가 혹은 갱신
	idMap[901] = "Apple"
	idMap[134] = "Grape"
	idMap[777] = "Tomato"

	// 키에 대한 값 읽기
	str := idMap[134]
	println(str)

	noData := idMap[999] // 값이 없으면 nil 혹은 zero 리턴
	println(noData)

	// 삭제
	delete(idMap, 777)

	/**
	  map을 사용하는 경우 종종 map안에 특정 키가 존재하는지를 체크할 필요가 있다.
	  이를 위해 Go에선 "map변수[키]" 읽기를 수행할 때 2개의 리턴값을 리턴한다.
	  첫번째는 키에 상응하는 값이고, 두번째는 그 키가 존재하는지 아닌지를 나타내는 bool 값이다.
	  즉, 아래 예제에서 val, exists := tickers["MSFT"] 의 val 에는 Microsoft라는 값이 리턴되고, 변수 exists에는 키가 존재하므로 true가 리턴된다.
	*/
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	// map 키 체크
	val, exists := tickers["MSFT"]
	if !exists {
		println(val)
		println("No MSFT ticker")
	}

	//Map이 갖고 있는 모든 요소들을 출력하기 위해, for range 루프를 사용할 수 있다. Map 컬렉션에 for range를 사용하면, Map 키와 Map 값 2개의 데이타를 리턴한다.
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}

func 패키지설명() {
	/**
	1. Main 패키지
	일반적으로 패키지는 라이브러리로서 사용되지만, "main" 이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다.
	패키지명이 main 인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다.
	그리고 이 main 패키지 안의 main() 함수가 프로그램의 시작점 즉 Entry Point가 된다.
	패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안된다.

	2. 패키지 Import
	패키지를 import 할 때, Go 컴파일러는 GOROOT 혹은 GOPATH 환경변수를 검색하는데,
	표준 패키지는 GOROOT/pkg 에서 그리고 사용자 패키지나 3rd Party 패키지의 경우 GOPATH/pkg 에서 패키지를 찾게 된다.
	GOROOT 환경변수는 Go 설치시 자동으로 시스템에 설정되지만, GOPATH는 사용자가 지정해 주어야 한다.
	GOPATH 환경변수는 3rd Party 패키지를 갖는 라이브러리 디렉토리나 사용자 패키지가 있는 작업 디렉토리를 지정하게 되는데,
	복수 개일 경우 세미콜론(윈도우즈의 경우)을 사용하여 연결한다.

	3. 패키지 Scope
	패키지 내에는 함수, 구조체, 인터페이스, 메서드 등이 존재하는데, 이들의 이름(Identifier)이 첫문자를 대문자로 시작하면 이는 public 으로 사용할 수 있다.
	즉, 패키지 외부에서 이들을 호출하거나 사용할 수 있게 된다. 반면, 이름이 소문자로 시작하면 이는 non-public 으로 패키지 내부에서만 사용될 수 있다.

	4. 패키지 init 함수와 alias
	개발자가 패키지를 작성할 때, 패키지 실행시 처음으로 호출되는 init() 함수를 작성할 수 있다.
	즉, init 함수는 패키지가 로드되면서 실행되는 함수로 별도의 호출 없이 자동으로 호출된다.
	package testlib

	var pop map[string]string

	func init() {   // 패키지 로드시 map 초기화
	    pop = make(map[string]string)
	}
	경우에 따라 패키지를 import 하면서 단지 그 패키지 안의 init() 함수만을 호출하고자 하는 케이스가 있다.
	이런 경우는 패키지 import 시 _ 라는 alias 를 지정한다. 아래는 other/xlib 패키지를 호출하면서 _ alias를 지정한 예이다.

	package main
	import _ "other/xlib"

	만약 패키지 이름이 동일하지만, 서로 다른 버젼 혹은 서로 다른 위치에서 로딩하고자 할 때는, 패키지 alias를 사용해서 구분할 수 있다.

	import (
	    mongo "other/mongo/db"
	    mysql "other/mysql/db"
	)
	func main() {
	    mondb := mongo.Get()
	    mydb := mysql.Get()
	    //...
	}
	6. 사용자 정의 패키지 생성
	개발자는 사용자 정의 패키지를 만들어 재사용 가능한 컴포넌트를 만들어 사용할 수 있다.
	사용자 정의 라이브러리 패키지는 일반적으로 폴더를 하나 만들고 그 폴더 안에 .go 파일들을 만들어 구성한다.
	하나의 서브 폴더안에 있는 .go 파일들은 동일한 패키지명을 가지며, 패키지명은 해당 폴더의 이름과 같게 한다.
	즉, 해당 폴더에 있는 여러 *.go 파일들은 하나의 패키지로 묶인다.
	*/
}

func Go구조체() {
	fmt.Println()
	fmt.Println("## Go구조체 ##")
	/**
	1. Struct (구조체)
	Go에서 struct는 Custom Data Type을 표현하는데 사용되는데, Go의 struct는 필드들의 집합체이며 필드들의 컨테이너이다.
	Go에서 struct는 필드 데이타만을 가지며, (행위를 표현하는) 메서드를 갖지 않는다.
	Go 언어는 객체지향 프로그래밍(Object Oriented Programming, OOP)을 고유의 방식으로 지원한다.
	즉, Go에는 전통적인 OOP 언어가 가지는 클래스, 객체, 상속 개념이 없다.
	전통적인 OOP의 클래스(class)는 Go 언어에서 Custom 타입을 정의하는 struct로 표현되는데,
	전통적인 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go 언어의 struct는 필드만을 가지며, 메서드는 별도로 분리하여 정의한다 (Go Method 에서 설명).
	*/

	// struct 정의
	type person struct { //person 구조체를 패키지외부에서사용할 수 있게 하려면 struct명을 Person 으로 변경.
		name string
		age  int
	}
	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)
	/**
	struct 객체를 생성할 때, 초기값을 함께 할당하는 방법도 있다.
	즉, 아래 첫번째 예처럼, struct 필드값을 순서적으로 { } 괄호안에 넣을 수 있으며,
	두번째 예처럼 순서에 상관없이 필드명을 지정하고(named field) 그 값을 넣을 수 도 있다.
	특히 두번째 예처럼 필드명을 지정하는 경우, 만약 일부 필드가 생략될 경우 생략된 필드들은 Zero value (정수인 경우 0, float인 경우 0.0, string인 경우 "", 포인터인 경우 nil 등)를 갖는다.
	*/
	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	println(p1)
	println(p2)
	/**
	또 다른 객체 생성 방법으로 Go 내장함수 new()를 쓸 수 있다.
	new()를 사용하면 모든 필드를 Zero value로 초기화하고 person 객체의 포인터(*person)를 리턴한다.
	객체 포인터인 경우에도 필드 엑세스 시 . (dot)을 사용하는데, 이 때 포인터는 자동으로 Dereference 된다 (이는 C에서 포인터의 경우 -> 을 사용하는 문법과 다르다).
	Go에서 struct는 기본적으로 mutable 개체로서 필드값이 변화할 경우 (별도로 새 개체를 만들지 않고) 해당 개체 메모리에서 직접 변경된다.
	하지만, struct 개체를 다른 함수의 파라미터로 넘긴다면, Pass by Value에 따라 객체를 복사해서 전달하게 된다.
	그래서 만약 Pass by Reference로 struct를 전달하고자 한다면, struct의 포인터를 전달해야 한다.
	*/
	p3 := new(person)
	p3.name = "Lee"
	/**
	4. 생성자(constructor) 함수
	때로 구조체(struct)의 필드가 사용 전에 초기화되어야 하는 경우가 있다.
	예를 들어, struct 의 필드가 map 타입인 경우 map을 사전에 미리 초기화해 놓으면, 외부 struct 사용자가 매번 map을 초기화 해야 된다는 것을 기억할 필요가 없다.
	이러한 목적을 위해 생성자 함수를 사용할 수 있다. 생성자 함수는 struct를 리턴하는 함수로서 그 함수 본문에서 필요한 필드를 초기화한다.
	아래 예제에서 생성자 함수 newDict()는 dict라는 struct의 map 필드를 초기화한 후 그 struct 포인터를 리턴하고 있다.
	이어 main() 함수에서 struct 개체를 만들 때 dict 를 직접 생성하지 않고 대신 newDict() 함수를 호출하여 이미 초기화된 data 맵 필드를 사용하고 있다.
	*/

	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
}

type dict struct {
	data map[int]string
}

//생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}
