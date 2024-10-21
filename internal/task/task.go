package task

import (
	"github.com/go-parrot/parrot/internal/task/list"
	"github.com/spf13/cobra"
)

// CmdProto represents the proto command.
var CmdTask = &cobra.Command{
	Use:   "task",
	Short: "Generate the task file",
	Long:  "Generate the task file.",
	Run:   run,
}

func init() {
	CmdTask.AddCommand(list.CmdList)
}

func run(cmd *cobra.Command, args []string) {
}
