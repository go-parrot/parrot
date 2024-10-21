package repo

import (
	"fmt"

	"github.com/go-parrot/parrot/internal/utils"

	"github.com/spf13/cobra"
)

// CmdRepo represents the new command.
var CmdRepo = &cobra.Command{
	Use:   "repo",
	Short: "Create a repo file by template",
	Long:  "Create a repo file using the repo template. Example: parrot g repo user",
	Run:   run,
}

var (
	targetDir string
	withCache bool
)

func init() {
	CmdRepo.Flags().StringVarP(&targetDir, "-target-dir", "t", "internal/repository", "generate target directory")
	CmdRepo.Flags().BoolVarP(&withCache, "-with-cache", "c", true, "add cache operate")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter the repo filename")
		return
	}
	// eg: parrot g repo user
	filename := args[0]

	c := &Repo{
		Name:      utils.Ucfirst(filename),    // 首字母大写
		LcName:    utils.Lcfirst(filename),    // 首字母小写
		UsName:    utils.Camel2Case(filename), // 下划线分隔
		Path:      targetDir,
		ModName:   utils.ModName(),
		WithCache: withCache,
	}
	if err := c.Generate(); err != nil {
		fmt.Println(err)
		return
	}
}
