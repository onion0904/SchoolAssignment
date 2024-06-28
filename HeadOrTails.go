import (
	"fmt"
	"rand"
)

func main() {
	fmt.Println("Tossing a coin...")
	var count1, count2 int
	for i := 0; i < 3; i++ {
		R :=rand.Intn(100)
		if R % 2 == 0 {
			fmt.Println("Round %d:Heads",i+1)
			count1++
		}else if R%2 == 1 {
			fmt.Println("Round %d:Tails",i+1)
			count2++
		}
	}
	fmt.Println("Heads: %d, Tails: %d",count1,count2)
	return
}