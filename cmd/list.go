/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
			services.ListAllTasks()
		} else if showDone {
			services.ListDoneTasks()
		} else if showPending {
			services.ListPendingTasks()
		}
	},
}

var (
	showDone    bool
	showPending bool
	showAll     bool
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showDone, "done", "d", false, "show only done tasks")
	listCmd.Flags().BoolVarP(&showPending, "pending", "p", false, "show only pending tasks")
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "show only all tasks")

	listCmd.MarkFlagsMutuallyExclusive("done", "pending", "all")
}
