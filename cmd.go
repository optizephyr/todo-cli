package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/optizephyr/todo-cli/models"
	"github.com/optizephyr/todo-cli/storage"
)

var tasks []models.Task

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'list' subcommands")
		os.Exit(1)
	}

	tasks, _ = storage.LoadTasks()

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:]) // 解析 'add' 子命令后的参数
		args := addCmd.Args()
		if len(args) < 1 {
			fmt.Println("Please provide a task description")
			os.Exit(1)
		}
		AddTask(args[0])

	case "list":
		listCmd.Parse(os.Args[2:]) // 解析 'list' 子命令后的参数
		ListTasks()

	case "remove":
		removeCmd.Parse(os.Args[2:])
		args := removeCmd.Args()
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Please input number")
				return
			}
			RemoveTask(id)
		}

	case "done":
		doneCmd.Parse(os.Args[2:])
		args := doneCmd.Args()
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Please input number")
				return
			}
			DoneTask(id)
		}
	default:
		fmt.Printf("Unknown subcommand '%s'\n", os.Args[1])
		os.Exit(1)
	}
	storage.SaveTasks(tasks)
}
