package main

import (
	"fmt"
)

func list(){
	tasks,err := load()
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(tasks)
	}
}