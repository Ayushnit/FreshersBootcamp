package main

import "fmt"

type FreqMap map[string]int

func freq(s string,results chan FreqMap) {
	m:=FreqMap{}
	for _,r :=range s {
		m[string(r)]++
	}
	results <-m
}
func findFreq(strings []string ) FreqMap{
	output:=FreqMap{}
	count:=len(strings)
	results:=make(chan FreqMap,count)

	for _,s:=range strings {
		go freq(s,results)
	}
	//close(results)

	for i:=0;i<count;i++ {
		for elem, freq := range <-results {
			output[elem] += freq
		}
	}
	return output
}

func main() {
	var arr []string=[]string{"quick","brown","fox","lazy","dog"}
	mp:=findFreq(arr)

	for elem,freq := range mp {
		fmt.Println(elem,freq)
	}
}

