package main

import "time"
import "fmt";

func main(){
	
	start := time.Now()

	for i := 0; i < 100000000; i++ {
		if i% 1000000000 == 0{
			fmt.Println(i)
		}
	}

	end := time.Now()

	fmt.Println(end.Sub(start))
	
}
