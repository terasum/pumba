package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"

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
	kubeconfig := c.String("kubeconfig")
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// TEMP: invoke clientset
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	return nil
}
