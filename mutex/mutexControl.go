package main

import (
	"fmt"
	"sync"
	"time"
)

var balanceNew int = 50
var mx sync.Mutex

//sync.Mutex进行控制并发安全
func main()  {

	go func() {
		Deposit(100)
		fmt.Printf("first %v\n",Balance())
	}()

	go func() {
		Deposit(200)
		fmt.Printf("second %v\n",Balance())

	}()

	time.Sleep(10)
	fmt.Printf("总金额：%v\n",Balance())

}

func Deposit1(amount int) {
	mx.Lock()
	balanceNew = balanceNew + amount
	mx.Unlock()
}

func Balance1() int {
	mx.Lock()
	b := balanceNew
	mx.Unlock()
	return b
}
