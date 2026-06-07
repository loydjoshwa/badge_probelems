package main

import (
	"fmt"
	"sync"
	"time"
)

func dog(dogch,catch chan string, wg *sync.WaitGroup){
	defer wg.Done()

	for i:=0;i<=10;i+=2{
		<-dogch
		time.Sleep(time.Second)
		fmt.Println("dog")
		catch<-"cat"
	}
}

func cat(dogch,catch chan string, wg *sync.WaitGroup){
	defer wg.Done()

	for i:=1;i<=9;i+=2{
		<-catch
		time.Sleep(time.Second)
          fmt.Println("cat")

		  dogch<-"dog"
	}
}

func reverse( str string) string{

	if len(str)<=1 {
		return str
	}

  return reverse(str[1:])+string(str[0])

}


func main() {
	str:="hello"
    m:=make(map[rune]int)

	for _,v:=range str{
     m[v]++
	}

	for v,f:=range m{
		fmt.Println(string(v),f)
	}
}     