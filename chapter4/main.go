package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"bufio"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	arr3 := [5]int{1: 23, 4: 25}
	fmt.Println(arr3)

	//pointer arrays

	arrayP := [5]*int{0: new(int), 4: new(int)}
	fmt.Println(arrayP)

	*arrayP[0] = 23
	*arrayP[4] = 25
	for _, val := range arrayP {
		if val != nil {
			fmt.Println(*val)
		}

	}
	multiArray := [4][2]int{1: {2, 3}, 2: {4, 5}, 3: {6, 7}, 0: {2, 5}}
	for outerIndex, mValue := range multiArray {
		fmt.Printf("Outer Index %d \n", outerIndex)
		for i, value := range mValue {
			fmt.Printf("Index %d \n", i)
			fmt.Printf("Value %d \n ", value)
		}
	}

	var bigArr [1e6]int
	foo(&bigArr)


	d1 := []byte("hello how are you??")
	file, err := os.Create("/tmp/d1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Write(d1)
	defer file.Close()

	d2 := []byte("sfdsdfds/sfsdfds/ sdsfdsfd/fdsfdsf")
	err = ioutil.WriteFile("/tmp/d1.txt", d2, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}

	writer := bufio.NewWriter(file)
	bytesWrote, err := writer.WriteString("Riya Dennis")
	if err != nil {
		log.Fatal(err.Error())
	}

	file.Sync()

	fmt.Printf("Bytes wrote %d \n", bytesWrote)
	fileReader, err := os.Open("/tmp/d1.txt")
	reader := bufio.NewReader(fileReader)
	fi, err := file.Stat()
	fz := fi.Size()
	b4, err := reader.Peek(int(fz))

	fmt.Printf("%d bytes read: %s\n", len(b4), string(b4))

}
func foo(arr *[1e6]int) {
	var i int;
	for i < 10 {
		arr[i] = 34 + i
		fmt.Printf("%d \n", arr[i])
		i++
	}
}
