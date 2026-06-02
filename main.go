package main

import "fmt"

func countdown(n int) int {
	if n==1{
		return 1
	}
	return n*countdown(n-1)
}

func main() {
	n:=10
	a:=0
	b:=1

	for i:=0;i<n;i++{
		fmt.Println(a," ")
        
		
		next:=a+b
		a=b
		b=next

	}
}