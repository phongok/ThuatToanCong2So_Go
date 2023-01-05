package main

import (
	"fmt"
	"strconv"
)

func sum(stn1 string, stn2 string) string {
	var finalresult int

	var s1 int // Giá trị chuổi 1
	var s2 int // Giá trị chuổi 2

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

	var len1 int // Độ dài chuổi 1
	var len2 int // Độ dài chuổi 2

	len1 = len(stn1)
	len2 = len(stn2)

	var lenmax int
	lenmax = len2
	if len1 >= len2 {
		lenmax = len1
	}

	var index1 int //Chỉ số kí tự đang xét chuổi 1
	var index2 int //Chỉ số kí tự đang xét chuổi 2

	var temp1 int //Giá trị chỉ số kí tự đang xét chuổi 1
	var temp2 int //Giá trị chỉ số kí tự đang xét chuổi 2
	temp1 = s1
	temp2 = s2

	var soNho int // Số nhớ

	for i := 0; i < lenmax; i++ {
		index1 = len1 - i - 1
		index2 = len2 - i - 1

		if i == 0 {

			fmt.Printf("Bước %d: Ta lấy %c cộng với %c được %d ghi %d nhớ %d\n", i+1, stn1[index1], stn2[index2], temp1%10+temp2%10, (temp1%10+temp2%10)%10, (temp1%10+temp2%10)/10)
			soNho = (temp1%10 + temp2%10) / 10
		}
		if index1 >= 0 && index2 >= 0 && i != 0 {
			temp1 = temp1 / 10
			temp2 = temp2 / 10
			fmt.Printf("Bước %d: Ta lấy %c cộng với %c và cộng cho số nhớ là %d được %d ghi %d nhớ %d \n", i+1, stn1[index1], stn2[index2], soNho, temp1%10+temp2%10+soNho, (temp1%10+temp2%10+soNho)%10, (temp1%10+temp2%10+soNho)/10)
			soNho = (temp1%10 + temp2%10 + soNho) / 10
		}

		if index1 < 0 && index2 > 0 {
			temp2 = temp2 / 10
			fmt.Printf("Bước %d: Ta lấy 0 cộng với %c và cộng cho số nhớ là %d  được %d ghi %d nhớ %d \n", i+1, stn2[index2], soNho, temp2%10+soNho, (temp2%10+soNho)%10, (temp2%10+soNho)/10)
		}
		if index2 < 0 && index1 > 0 {
			temp1 = temp1 / 10
			fmt.Printf("Bước %d: Ta lấy %c cộng với 0 và cộng cho số nhớ là %d được %d ghi %d nhớ %d\n", i+1, stn1[index1], soNho, temp1%10+soNho, (temp1%10+soNho)%10, (temp1%10+soNho)/10)
		}

		if index1 < 0 && index2 == 0 {
			temp2 = temp2 / 10
			fmt.Printf("Bước %d: Ta lấy 0 cộng với %c và cộng cho số nhớ là %d  được %d ghi %d  \n", i+1, stn2[index2], soNho, temp2%10+soNho, (temp2%10 + soNho))
		}
		if index2 < 0 && index1 == 0 {
			temp1 = temp1 / 10
			fmt.Printf("Bước %d: Ta lấy %c cộng với 0 và cộng cho số nhớ là %d được %d ghi %d \n", i+1, stn1[index1], soNho, temp1%10+soNho, (temp1%10 + soNho))
		}

	}

	finalresult = s1 + s2

	return strconv.Itoa(finalresult)
}
func main() {
	stn1 := "6789"
	stn2 := "567"
	total := sum(stn1, stn2)
	fmt.Println("Kết quả phép toán: ", total)
}
