/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/optizephyr/todo-cli/services"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all todo items",
	Aliases: []string{"ls", "show"},
	Long:    ``,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if showAll {
			tasks, err := services.GetAllTasks()
			if err != nil {
				fmt.Println(err)
			} else {
				if len(tasks) == 0 {
					fmt.Println("任务列表为空")
				} else {
					fmt.Println("任务列表:")
					for _, task := range tasks {
						fmt.Printf("任务%d:%s 完成状态:%v 创建于:%s\n", task.ID, task.Description, task.Done, task.CreatedAt.Format("2006-01-02 15:04:05"))
					}
				}
			}

		} else if showDone {
			tasks, err := services.GetDoneTasks()
			if err != nil {
				fmt.Println(err)
			} else {
				if len(tasks) == 0 {
					fmt.Println("已完成任务列表为空")
				} else {
					fmt.Println("以下任务已完成:")
					for _, task := range tasks {
						fmt.Printf("任务%d:%s 完成状态:%v 创建于:%s\n", task.ID, task.Description, task.Done, task.CreatedAt.Format("2006-01-02 15:04:05"))
					}
				}
			}
		} else {
			tasks, err := services.GetPendingTasks()
			if err != nil {
				fmt.Println(err)
			} else {
				if len(tasks) == 0 {
					fmt.Println("未完成任务列表为空")
				} else {
					fmt.Println("以下任务未完成:")
					for _, task := range tasks {
						fmt.Printf("任务%d:%s 完成状态:%v 创建于:%s\n", task.ID, task.Description, task.Done, task.CreatedAt.Format("2006-01-02 15:04:05"))
					}
				}
			}
		}
	},
}

var (
	showDone bool
	showAll  bool
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showDone, "done", "d", false, "show only done tasks")
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "show only all tasks")

	listCmd.MarkFlagsMutuallyExclusive("done", "all")
}
