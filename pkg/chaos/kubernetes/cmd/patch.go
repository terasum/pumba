package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/kubernetes"
)

// NewPatchCLICommand initialize patch command and bind it to the kubeContext
func NewPatchCLICommand(ctx context.Context) *cli.Command {
	cmdContext := &kubeContext{context: ctx}
	return &cli.Command{
		Name: "patch",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "context, c",
				Usage: "the name of the kubeconfig context to use",
				Value: "default",
			},
		},
		Usage:       "chaos testing for Kubernetes",
		ArgsUsage:   fmt.Sprintf("services/pods/deployments: name/label, list of names/labels, or RE2 regex if prefixed with %q", chaos.Re2Prefix),
		Description: "emulate different failures and resource starvation for Kubernetes services, pods and containers",
		Action:      cmdContext.patch,
	}
}

// Patch Command
func (cmd *kubeContext) patch(c *cli.Context) error {
	// get random
	random := c.GlobalBool("random")
	// get dry-run mode
	dryRun := c.GlobalBool("dry-run")
	// get interval
	interval := c.GlobalString("interval")
	// get names or pattern
	// names, pattern := chaos.GetNamesOrPattern(c)
	// init patch command
	patchCommand, err := kubernetes.NewPatchCommand(dryRun)
	if err != nil {
		return err
	}
	// run chaos command
	return chaos.RunChaosCommand(cmd.context, patchCommand, interval, random)
	return nil
}
