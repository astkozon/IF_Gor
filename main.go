package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("If: go run main.go 10")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Invalid number")
		return
	}

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(id int) {
			defer wg.Done()
			for msg := 1; msg <= 5; msg++ {
				fmt.Printf("[goroutine %d] message %d\n", id, msg)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}


//Если внезапно горутины начнут работать дольше например:
//к ним добавились новые операции или они читают данные по сети и т.п то время которое мы укаазали в time.Sleeep может оказатся 
//слишком маленьким
//горутины ненадежны тем что мы не знаем успеют ли они выполнится или нет