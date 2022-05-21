package main
  
import (
    "fmt"
    "os"
	"sort"
	"strconv"
)
  
func main() {
	var numbers []float64
	if len(os.Args) <= 1{
		return
	}
	for _, args := range os.Args[1:] {
        if n, err := strconv.ParseFloat(args, 64); err == nil {
            numbers = append(numbers, n)
        }
	}
	sort.Float64s(numbers)
	min := numbers[0]
	max := numbers[len(numbers)-1]
	var median float64
	if len(numbers) % 2 == 1{
		median = float64(numbers[len(numbers)/2])
	}else{
		median = (float64(numbers[len(numbers)/2 - 1]) + float64(numbers[len(numbers)/2]))/2.0
	}

	var mean float64 = 0
	for i := 0; i < len(numbers); i++{
		mean = mean + numbers[i]
	}
	mean = mean/ float64(len(numbers))
	fmt.Printf("Min: %f\n",min)
	fmt.Printf("Max: %f\n",max)
	fmt.Printf("Median: %f\n",median)
	fmt.Printf("Mean: %f\n",mean)
}