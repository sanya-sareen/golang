package main

import (
	"fmt"	
)

func searchInArr(arr []int, n int) bool{
	for i:= 0; i < len(arr); i++{
		if(arr[i] == n){
			return true
		}
	}
	return false
}

func main(){
	arr := []int {1,4,3}
	fmt.Println(searchInArr(arr, 3))
}