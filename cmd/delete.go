package cmd

import (
	"errors"
	"strconv"

	"github.com/Shaketheco1/Task-Tracker/task"
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}
	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("task id is required")
	}
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return errors.New("invalid task id")
	}
	return task.DeleteTask(id)
}
