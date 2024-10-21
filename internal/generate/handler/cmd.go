package handler

import (
	"fmt"

	"github.com/go-parrot/parrot/internal/utils"

	"github.com/spf13/cobra"
)

// CmdHandler represents the new command.
var CmdHandler = &cobra.Command{
	Use:   "handler",
	Short: "Create a handler file by template",
	Long:  "Create a handler file using the handler template. Example: parrot handler add demo",
	Run:   run,
}

var (
	targetDir string
	version   string
	method    string
)

func init() {
	CmdHandler.Flags().StringVarP(&targetDir, "target-dir", "t", "internal/handler", "generate target directory")
	CmdHandler.Flags().StringVarP(&version, "version", "v", "v1", "handler version")
	CmdHandler.Flags().StringVarP(&method, "method", "m", "get", "http method")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter the handler filename")
		return
	}
	// eg: parrot handler add demo
	filename := args[0]

	c := &Handler{
		Name:    utils.Ucfirst(filename),    // 首字母大写
		LcName:  utils.Lcfirst(filename),    // 首字母小写
		UsName:  utils.Camel2Case(filename), // 下划线分隔
		Path:    targetDir,
		Version: version,
		Method:  method,
	}
	if err := c.Generate(); err != nil {
		fmt.Println(err)
		return
	}
}
