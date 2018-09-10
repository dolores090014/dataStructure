package Heap

import (
	"testing"
	"time"
	"fmt"
	"math/rand"
)

func TestLHeap_Push(t *testing.T) {
	for i := 1; i <= 15; i++ {
		bench()
	}
}

func bench() {
	const MAX = 100
	heap := NewLittleHeap()
	r := rand.New(rand.NewSource(int64(time.Now().Second()))).Perm(MAX)

	for _, v := range r {
		heap.Push(v + 1)
		//for _, v := range heap.content {
		//	for _, val := range v {
		//		fmt.Print(val.Value,",")
		//	}
		//	fmt.Print("|")
		//}
		//fmt.Println()
	}
	//return
	var l = make([]int, 0)
	for {
		//for _, v := range heap.content {
		//	for _, val := range v {
		//		fmt.Print(val.Value,",")
		//	}
		//	fmt.Print("|")
		//}
		//fmt.Println()
		v := heap.Pop()
		//fmt.Println(v)
		if v == 0 {
			break
		}
		l = append(l, v)
	}
	//return
	if len(l) != MAX {
		fmt.Println("fail:repeat")
	}
	for k, v := range l {
		if k == len(l)-1 {
			fmt.Println("pass")
			break
		}
		if v > l[k+1] {
			fmt.Println("fall")
		}
	}
}
