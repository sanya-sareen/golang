package main

import (
	"fmt"
	// "math"	
)

func binarySearch(arr []int, n int)int{
	l := 0
	h := len(arr) - 1
	

	for h>= l{
		
		m := (l + h) / 2
		if(arr[m] == n){
			return n
		} 
		if(arr[m] < n){
			l = m + 1
		}else{
			 h = m - 1
		}

	}

	return -1

}

func main(){
	arr := []int {1,2,3,4,6,1}
	n := 4

	fmt.Println(binarySearch(arr, n))

}