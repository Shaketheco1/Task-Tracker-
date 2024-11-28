package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootcmd := &cobra.Command{
		Use:   "Task-Tracker",
		Short: "Task-Tracker is a CLI task manager",
	}
	rootcmd.AddCommand(NewAddCmd())
	rootcmd.AddCommand(NewDeleteCmd())
	rootcmd.AddCommand(NewUpdateCmd())
	rootcmd.AddCommand(NewListCmd())
	rootcmd.AddCommand(NewMark_Done_Cmd())
	rootcmd.AddCommand(NewMark_In_process_Cmd())
	return rootcmd
}
