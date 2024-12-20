package cmd

import (
	"errors"
	"strconv"

	"github.com/Shaketheco1/Task-Tracker/task"
	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(args)
		},
	}
	return cmd
}

func RunUpdateTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("task id is required")
	}

	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return errors.New("invalid task id")
	}
	description := args[1]
	return task.UpdateTask(id, description)
}
