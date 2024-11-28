package cmd

import (
	"github.com/Shaketheco1/Task-Tracker/task"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTaskCmd(args)
		},
	}
	return cmd
}

func RunListTaskCmd(args []string) error {
	list_limited := args[0]
	return task.ListTasks(list_limited)
}
