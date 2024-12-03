package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var todoList []string

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Println("To-Do List CLI")
// 		fmt.Println("1. Add Task")
// 		fmt.Println("2. List Tasks")
// 		fmt.Println("3. Remove Task")
// 		fmt.Println("4. Exit")
// 		fmt.Print("Enter your choice: ")

// 		scanner.Scan()
// 		choice := scanner.Text()

// 		switch choice {
// 		case "1":
// 			addTask(scanner)
// 		case "2":
// 			listTasks()
// 		case "3":
// 			removeTask(scanner)
// 		case "4":
// 			fmt.Println("Exiting...")
// 			return
// 		default:
// 			fmt.Println("Invalid choice, please try again.")
// 		}
// 	}
// }

// func addTask(scanner *bufio.Scanner) {
// 	fmt.Print("Enter task: ")
// 	scanner.Scan()
// 	task := scanner.Text()
// 	todoList = append(todoList, task)
// 	fmt.Println("Task added.")
// }

// func listTasks() {
// 	fmt.Println("To-Do List:")
// 	for i, task := range todoList {
// 		fmt.Printf("%d. %s\n", i+1, task)
// 	}
// }

// func removeTask(scanner *bufio.Scanner) {
// 	fmt.Print("Enter task number to remove: ")
// 	scanner.Scan()
// 	var taskNumber int
// 	fmt.Sscanf(scanner.Text(), "%d", &taskNumber)
// 	if taskNumber > 0 && taskNumber <= len(todoList) {
// 		todoList = append(todoList[:taskNumber-1], todoList[taskNumber:]...)
// 		fmt.Println("Task removed.")
// 	} else {
// 		fmt.Println("Invalid task number.")
// 	}
// }
