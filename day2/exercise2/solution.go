package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func takeRating(student int,ratings chan int,wg *sync.WaitGroup) {
	n := rand.Intn(10)
	time.Sleep(time.Duration(n)*time.Second)

	ratings<-rand.Intn(6) //Rating from 0-5
	wg.Done()
}
func main() {
	noOfStudents:=200
	ratings:=make(chan int,noOfStudents)
	var wg sync.WaitGroup
	for i:=0;i<noOfStudents;i++ {
		wg.Add(1)
		go takeRating(i,ratings,&wg)
	}
	wg.Wait()
	close(ratings)
	sum:=0
	for i:=0;i<noOfStudents;i++ {
		sum += <-ratings
	}
	fmt.Println(float32(sum)/float32(noOfStudents))
}
