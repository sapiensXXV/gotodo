package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Manager interface {
	handleCommand(args []string)
}

type TaskManager struct {
	todoList []Task
}

const taskFile = "tasks.json"

//func init() {
//	file, err1 := os.Open(taskFile)
//	if err1 != nil {
//		log.Fatalf("cannot open file: ", err1)
//		return
//	}
//	defer file.Close()
//
//	decoder := json.NewDecoder(file)
//	//if err2 := decoder.Decode(&manager.todoList); err2 != nil {
//	//	log.Fatalf("cannot parsing data: ", err2)
//	//}
//	if err2 := decoder.Decode(&manager.todoList); err2 != nil {
//		log.Fatalf("cannot parsing data: ", err2)
//	}
//}

func NewTaskManager() Manager {
	manager := &TaskManager{}

	file, err1 := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0644)
	if err1 != nil {
		log.Fatalf("cannot open file: ", err1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&manager.todoList); err != nil {
		log.Fatalf("error: cannot decode file - ", err)
	}
	return manager
}

func (tm *TaskManager) handleCommand(args []string) {

}

func (tm *TaskManager) addTask(task Task) {
	tasks := append(tm.todoList, task)
	saveAtFile(tasks)

	fmt.Println("add task complete.")
}

func (tm *TaskManager) listTask() {
	// task number, title, category, status, created_at, completed_at
	fmt.Printf("%-5s %-25s %-15s %-15s %-22s %-22s\n", "#", "Title", "Category", "Status", "Created At", "Completed At")
	fmt.Println(strings.Repeat("=", 110))

	num := 1
	for _, task := range tm.todoList {
		created := task.createdAt().Format("2006-01-02 15:04:05")
		completed := ""
		if !task.completedAt().IsZero() {
			completed = task.completedAt().Format("2006-01-02 15:04:05")
		}

		fmt.Printf("%-5d %-25s %-15s %-15s %-22s %-22s\n", num, task.title(), task.category(), task.status(), created, completed)
		num++
	}
}

func (tm *TaskManager) deleteTask(num int) {
	if num <= 0 || num > len(tm.todoList) {
		log.Fatalf("error: wrong task number (input: %d, number of task: %d)\n", num, len(tm.todoList))
		return
	}

	deleteTask := tm.todoList[num-1]
	tm.todoList = append(tm.todoList[:num-1], tm.todoList[num:]...)

	// 새로운 todoList를 기반으로 JSON을 인코딩하고 파일에 쓴다.
	saveAtFile(tm.todoList)

	fmt.Println("task remove complete")
	fmt.Println("Title: %s", deleteTask.Title)
}

func (tm *TaskManager) changeTaskStatus(number int, status string) {
	taskStatus, err := ParseTaskStatus(status)
	if err != nil {
		log.Fatal(err)
	}

	if number <= 0 || number > len(tm.todoList) {
		log.Fatalf("task number [%d] does not exist. Current number of tasks: %d", number, len(tm.todoList))
		return
	}

	tm.todoList[number].Status = taskStatus
	saveAtFile(tm.todoList)
}

func saveAtFile(tasks []Task) {
	file, err := os.OpenFile(taskFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("error: cannot open or create file")
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "   ")
	if err2 := encoder.Encode(tasks); err2 != nil {
		log.Fatalf("There was a problem saving the file.")
	}
}
