package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

var sum = 0

type sender = chan<- int
type reviser = <-chan int
type waitGroup = sync.WaitGroup
type mutex = sync.Mutex

func main() {
	fmt.Println("开始")
	errorTest(20)
	fmt.Println("结束")

}
func errorTest(num int) {
	arr := [3]int{1, 1, 2}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(arr[num])
}

func syncSum(num *int, wg *waitGroup) {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Microsecond)
		*num += 1
		println(*num)
	}
	wg.Done()
}

func syncSumLock(lock *mutex, num *int, wg *waitGroup) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		time.Sleep(time.Microsecond)
		*num += 1
		println(*num)
		lock.Unlock()
	}
	wg.Done()
}

func syncX(s int, wg *sync.WaitGroup) {
	fmt.Printf("%d号线程已完成\n", s)
	wg.Done()
}

func ExampleReadAtLeast() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Println("error:", err)
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err)
	}

	// Output:
	// some io.Reader
	// error: short buffer
	// error: unexpected EOF
}

func add2(num int, c chan int) {
	sum += num
	for i := 0; i < 2; i++ {
		time.Sleep(1000)
		fmt.Println(1)
		c <- i + 'a'
	}
}
func add3(num int, c chan int) {
	sum += num
	c <- sum
	a, ok := <-c
	b, ok2 := <-c
	if ok {
		fmt.Println(a)
	}
	if ok2 {
		fmt.Println(b)
	}
}

func add(num int) {
	sum += num
}

type Animal interface {
	eat()
}

type Dog struct {
}

func (dog Dog) eat() string {
	return "dod"
}
func eat2(dog Dog) string {
	return "dod"
}

func fn1() (int, int) {
	return 1, 2
}

type Student struct {
	age  int
	name string
	no   int
}
