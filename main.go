package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// size -- name
	store := make(map[int64]string)
	needDeleteArray := make(map[string]bool)
	fmt.Println("please input dir name, just like ./temp:")
	dirName := "./temp"
	_, err := fmt.Scan(&dirName)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Printf("open error:%v\n", err)
	}

	fmt.Println("find file------------------")
	for _, fir := range dir {
		stat, err := os.Stat(dirName + "/" + fir.Name())
		if err != nil {
			fmt.Printf("open error:%v\n", err)
		}
		size := stat.Size()
		name := stat.Name()
		fmt.Printf("name:%v, size:%v\n", name, size)
		s := store[size]
		if s == "" {
			store[size] = name
		} else {
			needDeleteArray[name] = true
		}
	}
	fmt.Println("find file------------------")

	// delete
	fmt.Println("--------------------------start delete--------------------------------")
	fmt.Printf("will delete %d file\n", len(needDeleteArray))
	for k, _ := range needDeleteArray {
		fmt.Printf("will delete name(%v)\n", k)
	}
	fmt.Print("please input yes(y) or no(n):")
	var choose string
	_, err = fmt.Scan(&choose)
	if err != nil {
		fmt.Println("input not support")
		return
	}
	if choose == "y" {
		for k, _ := range needDeleteArray {
			err := os.Remove(dirName + "/" + k)
			if err != nil {
				fmt.Printf("delete error:%v\n", err)
			}
		}
		fmt.Println("----------------delete done---------------------")
	} else if choose == "n" {
		fmt.Println("you cancel it")
		return
	} else {
		fmt.Printf("input(%v) not correct, it should be y or n\n", choose)
		return
	}

}
