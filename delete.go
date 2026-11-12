package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

func delete(){
	if len(os.Args) < 3 {
            fmt.Println("Please provide a task id.")
            os.Exit(1)
        }

		tasks,err := load()
		flag := 0
		taskID := os.Args[2]
		if err != nil{
			fmt.Println(err)
		}else{
			for i:=0;i<len(tasks);i++{
				if(tasks[i].TaskId == taskID){
					tasks = slices.Delete(tasks,i,i+1)
					newList,err := json.Marshal(tasks)
					
					if err != nil{
						fmt.Println(err)
					} else{
						err1 := os.WriteFile("tasks.json",newList,0644)
						if err1 != nil{
							fmt.Println("The storage has been occupuied")
							os.Exit(1)
						}
						fmt.Println("The task has been deleted")
					}
					flag = 1
					break
				}
			}
			if flag == 0{
				fmt.Println("The task is not present in the list")
			}
		}
}