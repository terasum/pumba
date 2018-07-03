package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli"

	"github.com/alexei-led/pumba/pkg/chaos"
	// "github.com/alexei-led/pumba/pkg/chaos/kubernetes"
)

type kubeContext struct {
	context context.Context
}

// NewKubeCLICommand initialize kubernetes main command and bind it to the kubeContext
func NewKubeCLICommand(ctx context.Context) *cli.Command {
	cmdContext := &kubeContext{context: ctx}
	return &cli.Command{
		Name:    "kubernetes",
		Aliases: []string{"kube", "k8s"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "context, c",
				Usage: "the name of the kubeconfig context to use",
				Value: "default",
			},
			cli.StringFlag{
				Name:  "kubeconfig",
				Usage: "path to the kubeconfig file to use for Kubernetes API requests",
				Value: "~/.kube/config",
			},
		},
		Usage:       "chaos testing for Kubernetes",
		ArgsUsage:   fmt.Sprintf("services/pods/deployments: name/label, list of names/labels, or RE2 regex if prefixed with %q", chaos.Re2Prefix),
		Description: "emulate different failures and resource starvation for Kubernetes services, pods and containers",
		Action:      cmdContext.kube,
	}
}

// Kubernetes Command
func (cmd *kubeContext) kube(c *cli.Context) error {
	// get kubernetes context
	// kubeContext := c.String("context")
	// kubernetes config file
	// kubeConfig := c.String("kubeconfig")
	return nil
}
