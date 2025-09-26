/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/optizephyr/todo-cli/services"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo item",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task, err := services.AddTask(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("成功添加任务%d:%s\n", task.ID, task.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
