package main

import (
	"fmt"
	"sync"
	"time"
)

/**
1. Go루틴
Go루틴(goroutine)은 Go 런타임이 관리하는 Lightweight 논리적 (혹은 가상적) 쓰레드(주1)이다.
Go에서 "go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행한다.
goroutine은 비동기적으로(asynchronously) 함수루틴을 실행하므로, 여러 코드를 동시에(Concurrently) 실행하는데 사용된다.
(주1)goroutine은 OS 쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위하여 만든 것으로, 기본적으로 Go 런타임이 자체 관리한다.
Go 런타임 상에서 관리되는 작업단위인 여러 goroutine들은 종종 하나의 OS 쓰레드 1개로도 실행되곤 한다.
즉, Go루틴들은 OS 쓰레드와 1 대 1로 대응되지 않고, Multiplexing으로 훨씬 적은 OS 쓰레드를 사용한다.
메모리 측면에서도 OS 쓰레드가 1 메가바이트의 스택을 갖는 반면, goroutine은 이보다 훨씬 작은 몇 킬로바이트의 스택을 갖는다(필요시 동적으로 증가).
Go 런타임은 Go루틴을 관리하면서 Go 채널을 통해 Go루틴 간의 통신을 쉽게 할 수 있도록 하였다.

아래 예제에서 main 함수를 보면, 먼저 say()라는 함수를 동기적으로 호출하고, 다음으로 동일한 say() 함수를 비동기적으로 3번 호출하고 있다.
첫번째 동기적 호출은 say() 함수가 완전히 끝났을 때 다음 문장으로 이동하고, 다음 3개의 go say() 비동기 호출은 별도의 Go루틴들에서 동작하면서,
메인루틴은 계속 다음 문장(여기서는 time.Sleep)을 실행한다. 여기서 goroutine들은 그 실행순서가 일정하지 않으므로 프로그램 실행시 마다 다른 출력 결과를 나타낼 수 있다.
*/
func main() {
	//함수를 동기적으로 실행
	say2("Sync")

	//함수를 비동기적으로 실행
	go say2("Async1")
	go say2("Async2")
	go say2("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)

	/**
	2. 익명함수 Go루틴
	Go루틴(goroutine)은 익명함수(anonymous function)에 대해 사용할 수도 있다.
	즉, go 키워드 뒤에 익명함수를 바로 정의하는 것으로, 이는 익명함수를 비동기로 실행하게 된다.
	아래 예제에서 첫번째 익명함수는 간단히 Hello 문자열을 출력하는데, 이를 goroutine으로 실행하면 비동기적으로 그 익명함수를 실행하게 된다.
	두번째 익명함수는 파라미터를 전달하는 예제로 익명함수에 파라미터가 있는 경우, go 익명함수 바로 뒤에 파라미터(Hi)를 함께 전달하게 된다.

	여기서 sync.WaitGroup을 사용하고 있는데, 이는 기본적으로 여러 Go루틴들이 끝날 때까지 기다리는 역활을 한다.
	WaitGroup을 사용하기 위해서는 먼저 Add() 메소드에 몇 개의 Go루틴을 기다릴 것인지 지정하고, 각 Go루틴에서 Done() 메서드를 호출한다 (여기서는 defer 를 사용하였다).
	그리고 메인루틴에서는 Wait() 메서드를 호출하여, Go루틴들이 모두 끝나기를 기다린다.
	*/

	// WaitGroup 생성. 2개의 go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	//익명함수를 사용한 고루틴
	go func() {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	//익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() //고루틴모두 끝날때까지 대기

	/**
	3. 다중 CPU 처리
	Go는 디폴트로 1개의 CPU를 사용한다. 즉, 여러 개의 Go 루틴을 만들더라도, 1개의 CPU에서 작업을 시분할하여 처리한다 (Concurrent 처리).
	만약 머신이 복수개의 CPU를 가진 경우, Go 프로그램을 다중 CPU에서 병렬처리 (Parallel 처리)하게 할 수 있는데,
	병렬처리를 위해서는 아래와 같이 runtime.GOMAXPROCS(CPU수) 함수를 호출하여야 한다 (여기서 CPU 수는 Logical CPU 수를 가리킨다).
	*/
	//4개의 CPU 사용
	//runtime.GOMAXPROCS(4)
}

func say2(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}
