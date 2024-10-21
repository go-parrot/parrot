package proto

import (
	"github.com/spf13/cobra"

	"github.com/go-parrot/parrot/internal/proto/client"
	"github.com/go-parrot/parrot/internal/proto/server"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
	Run:   run,
}

func init() {
	CmdProto.AddCommand(client.CmdClient)
	CmdProto.AddCommand(server.CmdServer)
}

func run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
