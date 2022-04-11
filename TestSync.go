package main

import(
	"fmt"
	"sync"
	"math/rand"
	"time"
)

func main(){
	const N = 5
	var timers [N]int32
	
	var wg sync.WaitGroup
	
	wg.Add(N)
	
	rand.Seed(time.Now().UnixNano())
		

	
	for i := 1; i <= N; i++ {
		i := i		//very important... otherwise panic will occur..
		go func(){
			defer func (){
				fmt.Println("i quit:", i)
			}()
			timers[i - 1] = 10 + rand.Int31n(10)
			time.Sleep(time.Duration(timers[i - 1]) * time.Second)
			wg.Done()
		}()
	}
	
	wg.	Wait()
	
	fmt.Println("timers:", timers)
}