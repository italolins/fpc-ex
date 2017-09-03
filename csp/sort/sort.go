package main

import (
	"fmt"
)

var to_order [6]int
var UP bool
var DOWN bool

func merge(lowVal int, n int, direction bool){
	if (n > 1){
		m := n/2
		for  i := lowVal; i < lowVal + m; i++{
			compare(i, i + m, direction)
		}
		merge(lowVal, m, direction)
		merge(lowVal + m, m, direction)
	}
}

func parallelSort(lowValue int, n int, direction bool, level int){
	if(n > 1){
		m := n / 2;
            	if (level == 0) {
                	parallelSort(lowValue, m, UP, level);
                	parallelSort(lowValue + m, m, DOWN, level);
                	merge(lowValue, n, direction);
            	} 
	}
}

func compare(i int, j int, direction bool){
	if (direction == (to_order[i] > to_order[j])){
		swap(i, j);
	}
}

func swap(i int, j int){
	to_order[i],to_order[j] = to_order[j],to_order[i]
	fmt.Println("Hello, playground", to_order)
	
}

func main() {
	to_order = [6]int{2,1,9,3,4,6}
	UP := true
	//DOWN := false
	
	fmt.Println("Hello, playground", to_order)
	
	parallelSort(0,7,UP,0)
	
	fmt.Println("Hello, playground", to_order)
}
