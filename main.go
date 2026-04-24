package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	TaskName  string
	Completed bool
}

var tasks []Task

func returnTasks() {
	for i, task := range tasks {
		status := "Not done"
		if task.Completed {
			status = "Done"
		}

		fmt.Printf("%d. %s		-[%s]\n", i+1, task.TaskName, status)
	}
}

func addTasks(task string) {
	newTask := Task{TaskName: task, Completed: false}

	tasks = append(tasks, newTask)
}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].Completed = true
		fmt.Printf("Task has been marked as done!")
	} else {
		fmt.Println("Invalid Index Provided!")
	}
}

func editTasks(index int, stringValue string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].TaskName = stringValue
		fmt.Printf("Task has been edited succesfully!")
	} else {
		fmt.Println("Invalid Index Provided!")
	}
}

func deleteTasks(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Printf("Task has been marked as done!")
	} else {
		fmt.Println("Invalid Index Provided!")
	}
}

func main() {
	fmt.Println("Options: -->")
	fmt.Println("1. Add Task")
	fmt.Println("2.Delete Task")
	fmt.Println("3. Edit Task")
	fmt.Println("4. List Tasks")
	fmt.Println("5. Mark Done")
	fmt.Println("6. Quit")

	var indexInput int
	var editTask, newTask string

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("--Enter Choice: (1, 2, 3, 4, 5, 6)}---")
		reader.Scan()

		input := reader.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Task Name: ")
			reader.Scan()
			newTask = reader.Text()

			addTasks(newTask)
			fmt.Println(" ")
			returnTasks()

		case 2:
			fmt.Print("Enter Index: ")
			reader.Scan()

			indexInput, _ = strconv.Atoi(reader.Text())
			deleteTasks(indexInput)
			fmt.Println(" ")

			returnTasks()
		case 3:
			fmt.Print("Enter Index: ")
			reader.Scan()

			indexInput, _ = strconv.Atoi(reader.Text())

			fmt.Print("Edited Task: ")
			reader.Scan()

			editTask = reader.Text()

			editTasks(indexInput, editTask)
			fmt.Println(" ")

			returnTasks()
		case 4:
			returnTasks()

		case 5:
			fmt.Print("Enter Index: ")
			reader.Scan()

			indexInput, _ = strconv.Atoi(reader.Text())

			markCompleted(indexInput)
			fmt.Println(" ")

			returnTasks()
		case 6:
			fmt.Print("Shutting Down...")
			os.Exit(0)

		default:
			fmt.Print("Invalid Choice!")
		}
	}
}
