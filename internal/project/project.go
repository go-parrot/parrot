package project

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// CmdNew represents the new command.
var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create a project template",
	Long:  "Create a project project using the repository template. Example: parrot new helloworld",
	Run:   run,
}

var (
	repoURL    string
	branch     string
	timeout    string
	moduleName string
)

func init() {
	if repoURL = os.Getenv("PARROT_REPO"); repoURL == "" {
		repoURL = "https://github.com/go-parrot/parrot-starter.git"
	}

	// default http, only include http server
	branch = "master"
	// default timeout
	timeout = "60s"

	CmdNew.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")
	CmdNew.Flags().StringVarP(&branch, "branch", "b", branch, "default is http server, empty is http and gRPC")
	CmdNew.Flags().StringVarP(&timeout, "timeout", "t", timeout, "request timeout time")
	CmdNew.Flags().StringVarP(&moduleName, "module", "m", "", "module name")
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t, err := time.ParseDuration(timeout)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	// get project name
	name := ""
	if len(args) == 0 {
		prompt := &survey.Input{
			Message: "What is project name ?",
			Help:    "Created project name.",
		}
		err = survey.AskOne(prompt, &name)
		if name == "" || err != nil {
			return
		}
	} else {
		name = args[0]
	}

	if moduleName == "" {
		moduleName = name
	}

	p := &Project{Name: name}
	done := make(chan error, 1)
	go func() {
		done <- p.New(ctx, wd, repoURL, branch, moduleName)
	}()

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Fprint(os.Stderr, "\033[31mERROR: project creation timed out\033[m\n")
		} else {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: failed to create project(%s)\033[m\n", ctx.Err().Error())
		}
	case err = <-done:
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: Failed to create project(%s)\033[m\n", err.Error())
		}
	}
}
