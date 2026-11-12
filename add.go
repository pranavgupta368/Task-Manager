package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func genNextID() int {
	maxId := 0

	tasks,err := load()
	if err != nil{
		return 1
	}

	for i:=0;i<len(tasks);i++{
		num,err := strconv.Atoi(tasks[i].TaskId)
		if err != nil{
			fmt.Println(err)
		}

		if num > maxId{
			maxId = num
		}
	}
	return maxId+1
}

func add(){
	if len(os.Args) < 3 {
            fmt.Println("Please provide a task name.")
            os.Exit(1)
        }

		task := os.Args[2:]
		counter := genNextID()
		todo := Task{
			TaskId: strconv.Itoa(counter),
			Name: strings.Join(task," "),
			Status: false,
			CreatedAt: time.Now(),
		}
		err := save(todo)
		if err != nil{
			fmt.Println(err)
		} else{
			fmt.Println("The task is saved")
		}
}