package trylock

import (
	"fmt"
	"sync"
	"time"

	letterCount "github.com/Davidmuthee12/threads/mutexes/letter-count"
)

func main() {
	mutex := sync.Mutex{}
	var frequency = make([]int, 26)
	for i := 2000; i <= 2200; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc%d.txt", i)
		go letterCount.CountLetters(url, frequency, &mutex)
	}
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		if mutex.TryLock() {
			for i, c := range letterCount.AllLetters {
				fmt.Printf("%c-%d ", c, frequency[i])
			}
			mutex.Unlock()
		} else {
			fmt.Println("Mutex already being used")
		}
	}
}