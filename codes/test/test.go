package main

import "fmt"

func main() {
	//s := []int{1, 2, 3, 4}
	//s2 := s[2:]
	//fmt.Printf("s1 %v,len=%v,cap=%v \n", s, len(s), cap(s))
	//fmt.Printf("s2 %v,len=%v,cap=%v \n", s2, len(s2), cap(s2))
	//
	//s2[0] = 99
	//fmt.Printf("s1 %v,len=%v,cap=%v \n", s, len(s), cap(s))
	//fmt.Printf("s2 %v,len=%v,cap=%v \n", s2, len(s2), cap(s2))
	//
	//s2 = append(s2, 199)
	//fmt.Printf("s1 %v,len=%v,cap=%v \n", s, len(s), cap(s))
	//fmt.Printf("s2 %v,len=%v,cap=%v  \n", s2, len(s2), cap(s2))
	//
	//s2[1] = 1999
	//fmt.Printf("s1 %v,len=%v,cap=%v \n", s, len(s), cap(s))
	//fmt.Printf("s2 %v,len=%v,cap=%v \n ", s2, len(s2), cap(s2))
	arr := make(map[string]int, 10)
	fmt.Printf("len=%v,cap=%v", len(arr))
}
