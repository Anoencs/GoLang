package main

import(
	"time"	
	"fmt"
	"math"
	"os"
)

func Duration(times string) float64{
	now := time.Now().UTC()
	date_format,_ := time.Parse("2006-01-02",times)
	return math.Ceil(now.Sub(date_format).Hours()/24)
}

func DiffTime(time1, time2 string) float64{
	date1_format,_ := time.Parse("2006-01-02",time1)
	date2_format,_ := time.Parse("2006-01-02",time2)
	return math.Ceil(date2_format.Sub(date1_format).Hours()/24)
}

func main(){

	if len(os.Args) <= 1{
		return
	}

	if len(os.Args) == 2{
		fmt.Println(Duration(os.Args[1]))
	}else if len(os.Args) == 3{
		fmt.Println(DiffTime(os.Args[1],os.Args[2]))
	}else{
		return
	}


}