/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/optizephyr/todo-cli/services"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean the done todo item",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := services.CleanTasks()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("成功清理以下任务：")
			for _, task := range tasks {
				fmt.Printf("任务%d:%s 完成状态:%v 创建于:%s\n", task.ID, task.Description, task.Done, task.CreatedAt.Format("2006-01-02 15:04:05"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
