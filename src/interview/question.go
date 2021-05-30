package interview

import (
	"context"
	"fmt"
	"time"
)

func cacl(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func DeferAction() {
	a := 1
	b := 2

	defer cacl("1", a, cacl("10", a, b))

	b = 0
	defer cacl("2", a, cacl("20", a, b))
	b = 3

	panic("exit")

	// output
	// 10 1 2 3
	// 20 1 0 1
	// 2 1 1 2
	// 1 1 3 4
	// panic: exit
}

func CloseChan() {
	type A struct {
		val int
	}

	c := make(chan *A, 10)

	c <- &A{val: 1}
	c <- &A{val: 2}
	close(c)

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)

	// output

	//&{1}
	//&{2}
	//<nil>
	//<nil>
}

func MapInit() {
	type A struct {
		val int
	}

	// 以下用法错误 map的值是无法取址的，也就无法对结构体里的field进行操作; 因为map是无序的，并且随时存在扩容缩容，其地址也就不固定
	// m := map[string]A{"x": {1}}
	// m["x"].val = 2

	// 解决方案
	// m := map[string]&A
}

func ConText() {
	ctx, _ := context.WithCancel(context.Background())
	//defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	resp := make(chan int, 3)
	err := make(chan interface{}, 3)

	operation := func(pid int, ctx context.Context, dst int, resp chan int, err chan interface{}) {
		n := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println(pid, "canceled", ctx.Err())
				err <- ctx.Err()
				return
			case <-time.After(time.Second):
				fmt.Println(pid, "1 second pass", n)
				n++
				if n == dst {
					resp <- pid
					return
				}
			}
		}
	}

	go operation(1, ctx, 10, resp, err)
	go operation(2, ctx, 10, resp, err)
	go operation(3, ctx, 7, resp, err)

	select {
	case pid := <-resp:
		fmt.Println(pid, "find dst and cancel other goroutines")
	case e := <-err:
		fmt.Println(e)
	}

	cancel()
	time.Sleep(1 * time.Second) // wait for goroutines return
}
