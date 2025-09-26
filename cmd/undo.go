/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/optizephyr/todo-cli/services"
	"github.com/spf13/cobra"
)

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo a todo task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				panic(fmt.Sprintf("%v转换失败", arg))
			}
			ids = append(ids, id)
		}
		tasks, err := services.UndoTasks(ids)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Undo以下任务:")
			for _, task := range tasks {
				fmt.Printf("任务%d:%s 完成状态:%v 创建于:%s\n", task.ID, task.Description, task.Done, task.CreatedAt.Format("2006-01-02 15:04:05"))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(undoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
