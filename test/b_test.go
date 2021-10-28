package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

type User struct {
	Money float64
}

func TestBP(t *testing.T) {
	go func() {
		fmt.Println("hello world")
	}()

	fmt.Println("done")
}

func TestBP2(t *testing.T) {
	var lp []string // 0 1 2 3 4 .. 99
	for i:=0;i<100;i++ {

		lp = append(lp,strconv.Itoa(i))
	}

	var wg sync.WaitGroup
	for _,v := range lp {  // v 是共享地址
		wg.Add(1)
		go func() {
			wg.Done()
			fmt.Println(v)
		}()
	}

	wg.Wait()
	fmt.Println("done")

	// 请问这些代码输出的是什么？
}

func TestChannel(t *testing.T) {
	var users []*User
	var p float64

	for i := 0; i < 99999; i++ {
		rand.Seed(time.Now().UnixNano())
		px := rand.Int63n(9999999)
		p += float64(px)
		users = append(users, &User{
			Money: float64(px),
		})
	}
	var sum float64
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := range users {
		idx := i
		wg.Add(1)

			go func() {
				defer wg.Done()

				time.Sleep(time.Microsecond * 100)

				mu.Lock()
				defer mu.Unlock()

				sum += users[idx].Money
			}()
	}

	wg.Wait()
	if p!=sum {
		panic(fmt.Sprint("代码错误 sum: ", sum))
	}

	// 并发求和  要求: 同时最大并发 30
}
