package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/go-parrot/parrot/internal/model"
	"github.com/go-parrot/parrot/internal/project"
	"github.com/go-parrot/parrot/internal/proto"
	"github.com/go-parrot/parrot/internal/run"
	"github.com/go-parrot/parrot/internal/task"
	"github.com/go-parrot/parrot/internal/upgrade"

	gcache "github.com/go-parrot/parrot/internal/generate/cache"
	ghandler "github.com/go-parrot/parrot/internal/generate/handler"
	gproto "github.com/go-parrot/parrot/internal/generate/proto"
	grepo "github.com/go-parrot/parrot/internal/generate/repo"
	gservice "github.com/go-parrot/parrot/internal/generate/service"
	gtask "github.com/go-parrot/parrot/internal/generate/task"
)

var (
	// Version is the version of the compiled software.
	Version = "v0.0.1"

	rootCmd = &cobra.Command{
		Use:     "parrot",
		Short:   "parrot: A microservice toolkit for Go",
		Long:    `parrot: A microservice toolkit for Go`,
		Version: Version,
	}

	genCmd = &cobra.Command{
		Use:     "generate",
		Short:   "Generate the GENERATOR [args] [options]",
		Long:    "Generate the GENERATOR [args] [options]",
		Aliases: []string{"g"},
	}
)

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(task.CmdTask)
	rootCmd.AddCommand(model.CmdNew)
	rootCmd.AddCommand(upgrade.CmdUpgrade)

	genCmd.AddCommand(gcache.CmdCache)
	genCmd.AddCommand(ghandler.CmdHandler)
	genCmd.AddCommand(gproto.CmdAdd)
	genCmd.AddCommand(grepo.CmdRepo)
	genCmd.AddCommand(gservice.CmdAdd)
	genCmd.AddCommand(gtask.CmdTask)
	rootCmd.AddCommand(genCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
