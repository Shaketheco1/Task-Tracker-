package cmd

import (
	"errors"
	"strconv"

	"github.com/Shaketheco1/Task-Tracker/task"
	"github.com/spf13/cobra"
)

func NewMark_In_process_Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-process",
		Short: "Mark a task as in-process",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunMarkTaskCmd(args, task.TASK_STATUS_INPROGRESS)
		},
	}
	return cmd
}

func NewMark_Done_Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-completed",
		Short: "Mark a task as completed",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunMarkTaskCmd(args, task.TASK_STATUS_DONE)
		},
	}
	return cmd
}

func RunMarkTaskCmd(args []string, status task.TaskStatus) error {
	if len(args) == 0 {
		return errors.New("task id is required")
	}
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return errors.New("invalid task id")
	}
	return task.MarkTask(id, status)
}
