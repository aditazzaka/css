package main

import "fmt"

type arr [5]int

func printArray(ar arr, n int) int {
	var size int = n
	panjang := len(ar)
	if size == panjang {
		return 0
	}
	fmt.Printf("%d ", ar[size])
	return printArray(ar, size+1)

}

func main() {
	var arr arr
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	printArray(arr, 0)
}


package main

import "fmt"

func banyakDigit(angka int, banyak int) int {
	var temp int = banyak
	if angka > -1 && angka < 10 {
		temp++
		return temp
	}
	var num int = angka
	num = (num - (num % 10)) / 10
	temp++
	return banyakDigit(num, temp)
}

func main() {
	var angka int = 12345
	hasil := banyakDigit(angka, 0)
	fmt.Println(hasil)
}


package main

import "fmt"

func konversi_desimalTobiner(angka int) {
	if angka == 1 {
		fmt.Print("1")
		return
	} else {
		konversi_desimalTobiner(angka / 2)
		fmt.Print(angka % 2)
	}
}

func main() {
	konversi_desimalTobiner(13)
}
