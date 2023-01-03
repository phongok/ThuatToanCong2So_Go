package main

import (
	"fmt"
	"strconv"
)

func sum(stn1 string, stn2 string) string {
	var finalresult int
	var s1 int
	var s2 int

	var len1 int // Độ dài chuổi 1
	var len2 int // Độ dài chuổi 2

	len1 = len(stn1)
	len2 = len(stn2)

	var index1 int //Chỉ số kí tự đang xét chuổi 2
	var index2 int //Chỉ số kí tự đang xét chuổi 2

	var lenmax int
	lenmax = len2
	if len1 >= len2 {
		lenmax = len1
	}

	for i := 0; i < lenmax; i++ {
		index1 = len1 - i - 1
		index2 = len2 - i - 1
		var d1 int //Chỉ số kí tự đang xét chuổi 2
		var d2 int //Chỉ số kí tự đang xét chuổi 2
		d1, err1 := strconv.Atoi(stn1)
		if err1 != nil {
			// ... handle error
			panic(err1)
		}
		d2, err2 := strconv.Atoi(stn2)
		if err2 != nil {
			// ... handle error
			panic(err2)
		}
		if i == 0 {

			fmt.Printf("Bước %d: Ta lấy %c cộng với %c được %d\n", i+1, stn1[index1], stn2[index2], d1+d2)
		}
		if index1 >= 0 && index2 >= 0 && i != 0 {
			fmt.Printf("Bước %d: Ta lấy %c cộng với %c được kkk \n", i+1, stn1[index1], stn2[index2])
		}

		if index1 < 0 {
			fmt.Printf("Bước %d: Ta lấy 0 cộng với %c được \n", i+1, stn2[index2])
		}
		if index2 < 0 {
			fmt.Printf("Bước %d: Ta lấy %c cộng với 0 được \n", i+1, stn1[index1])
		}

	}

	s1, err1 := strconv.Atoi(stn1)
	if err1 != nil {
		// ... handle error
		panic(err1)
	}
	s2, err2 := strconv.Atoi(stn2)
	if err2 != nil {
		// ... handle error
		panic(err2)
	}

	finalresult = s1 + s2

	return strconv.Itoa(finalresult)
}
func main() {
	stn1 := "1234"
	stn2 := "123"
	total := sum(stn1, stn2)
	fmt.Println("Result: ", total)
}
