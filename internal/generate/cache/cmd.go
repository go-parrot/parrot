package add

import (
	"fmt"
	"strings"

	"github.com/go-parrot/parrot/internal/utils"

	"github.com/spf13/cobra"
)

// CmdCache represents the new command.
var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Create a cache file by template",
	Long:  "Create a cache file using the cache template. Example: parrot g cache UserCache",
	Run:   run,
}

var (
	targetDir string
)

func init() {
	CmdCache.Flags().StringVarP(&targetDir, "-target-dir", "t", "internal/cache", "generate target directory")
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter the cache filename")
		return
	}
	// eg: parrot g cache UserCache
	filename := args[0]

	c := &Cache{
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
