package add

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/go-parrot/parrot/internal/utils"

	"github.com/spf13/cobra"
)

// CmdTask represents the new command.
var CmdTask = &cobra.Command{
	Use:   "task",
	Short: "Create a task file by template",
	Long:  "Create a task file using the task template. Example: parrot g task EmailWelcome",
	Run:   run,
}

var (
	targetDir string
)

func init() {
	CmdTask.Flags().StringVarP(&targetDir, "-target-dir", "t", "internal/tasks", "generate target directory")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Printf("Please enter the %s\n", color.RedString("task name"))
		return
	}
	// eg: parrot g task EmailWelcome
	filename := args[0]

	c := &Task{
		Name:      utils.Ucfirst(filename),                                  // 首字母大写
		LcName:    utils.Lcfirst(filename),                                  // 首字母小写
		UsName:    utils.Camel2Case(filename),                               // 下划线分隔
		ColonName: strings.ReplaceAll(utils.Camel2Case(filename), "_", ":"), // 冒号分隔
		Path:      targetDir,
		ModName:   utils.ModName(),
	}
	if err := c.Generate(); err != nil {
		fmt.Println(err)
		return
	}
}
