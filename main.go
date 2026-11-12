package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct{
	TaskId string `json:"TaskID"`
	Name string `json:"Task"`
	Status bool `json:"Status"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func load() ([]Task,error){
	filecontent,err := os.ReadFile("tasks.json")

	var tasks []Task
	if err != nil{
		if os.IsNotExist(err){
			return []Task{},nil
		}
		return nil,fmt.Errorf("Failed to load the file %w",err)
	}

	err1 := json.Unmarshal(filecontent,&tasks)
	if err1 != nil{
		return nil,fmt.Errorf("Error in converting the test %w",err1)
	}
	return tasks,nil
}

func save(task Task) error{
	filecontent,err := os.ReadFile("tasks.json")

	var tasks []Task
	if err != nil{
		if os.IsNotExist(err){
			var newList []Task
			newList = append(newList, task)
			newTask,err1 := json.Marshal(newList)
			if err1 != nil{
				return fmt.Errorf("Error in converting the text %w",err1)
			}
			return os.WriteFile("tasks.json",newTask,0644)
		}
		return fmt.Errorf("Error in opening the file %w",err)
	}

	err1 := json.Unmarshal(filecontent,&tasks)
	if err1 != nil{
		return fmt.Errorf("Error in converting the test %w",err1)
	}
	
	tasks = append(tasks,task)
	updatedTasks,err2 := json.Marshal(tasks)
	if err2 != nil{
		return fmt.Errorf("Failed to convert the list into json %w",err2)
	}

	return os.WriteFile("tasks.json",updatedTasks,0644)
}

func main(){
	if len(os.Args)<2{
		fmt.Println("Please first enter a command")
		fmt.Println("go run . add [Task]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command{
	case "add":
			add()
	case "list":
			list()
	case "done":
			done()
	case "delete":
			delete()
		}
}