package upgrade

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/go-parrot/parrot/internal/base"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the parrot tools",
	Long:  "Upgrade the parrot tools. Example: parrot upgrade",
	Run:   Run,
}

// Run upgrade the parrot tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoInstall(
		"github.com/go-parrot/parrot/cmd/parrot",
		"github.com/go-parrot/parrot/cmd/protoc-gen-go-gin",
		"google.golang.org/protobuf/cmd/protoc-gen-go",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc",
		"github.com/envoyproxy/protoc-gen-validate",
		"github.com/google/gnostic",
		"github.com/google/gnostic/cmd/protoc-gen-openapi",
		"github.com/google/wire/cmd/wire",
	)
	if err != nil {
		fmt.Println(err)
	}
}
