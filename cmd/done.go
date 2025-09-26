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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "set the todo item to done",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				panic(fmt.Sprintf("输入%s，需要数字", arg))
			}
			ids = append(ids, id)
		}
		tasks, err := services.DoneTasks(ids)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("完成以下任务:")
			for _, task := range tasks {
				fmt.Printf("任务%d:%s 完成状态:%v 创建于:%s\n", task.ID, task.Description, task.Done, task.CreatedAt.Format("2006-01-02 15:04:05"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
