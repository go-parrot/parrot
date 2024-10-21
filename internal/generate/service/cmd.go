package service

import (
	"fmt"

	"github.com/go-parrot/parrot/internal/utils"

	"github.com/spf13/cobra"
)

// CmdAdd represents the new command.
var CmdAdd = &cobra.Command{
	Use:   "svc",
	Short: "Create a service file by template",
	Long:  "Create a service file using the service template. Example: parrot g svc user",
	Run:   run,
}

var (
	targetDir string
)

func init() {
	CmdAdd.Flags().StringVarP(&targetDir, "-target-dir", "t", "internal/service", "generate target directory")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter the cache filename")
		return
	}
	// eg: parrot g svc User
	filename := args[0]

	c := &Service{
		Name:    utils.Ucfirst(filename), // 首字母大写
		LcName:  utils.Lcfirst(filename),
		Path:    targetDir,
		ModName: utils.ModName(),
	}
	if err := c.Generate(); err != nil {
		fmt.Println(err)
		return
	}
}
