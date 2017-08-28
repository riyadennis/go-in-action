package main
import(
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func printer(c chan int){
	for i := range c{
		fmt.Printf("Recieved %d \n", i)
	}
	wg.Done()
}
func main(){
	c := make(chan int)
	go printer(c)

	for i:=1;i < 10; i++ {
		c <- i
	}

	close(c)

	wg.Wait()
}
