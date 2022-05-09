package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

func (c *Counter) String() string {
	return fmt.Sprintf("{count: %d}", c.count)
}

// var mapChan = make(chan map[string]Counter, 1)
// 指针类型才能修改通道发送前的值
var mapChan = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Microsecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}
