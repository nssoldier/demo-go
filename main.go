package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func crawlData(length int, name string) (n []int) {
	for count := 0; count < length; count++ {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		var elem = rand.Intn(100)
		n = append(n, elem)
		fmt.Println(fmt.Sprintf(
			"added %d to %s: %d/%d(%.2f%s)",
			elem,
			name,
			count+1,
			length,
			(float32(count+1)/float32(length))*100,
			"%",
		),
		)
	}

	return n
}

func main() {
	var waitGroup sync.WaitGroup
	var result []int

	fmt.Println("dcm start")
	waitGroup.Add(3)

	for count := 0; count < 3; count++ {
		count := count
		go func() {
			var arrLength = rand.Intn(20)
			var name = fmt.Sprintf("Array %d", count+1)

			fmt.Println(fmt.Sprintf("Starting add elems to %s with length is %d", name, arrLength))
			var response = crawlData(arrLength, name)
			fmt.Print(fmt.Sprintf("%s done: ", name))
			fmt.Println(response)
			result = append(result, response...)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println(result)

	fmt.Println("Done!")
}
