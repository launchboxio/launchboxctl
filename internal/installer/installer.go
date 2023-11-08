package installer

import (
	"context"
	helmclient "github.com/mittwald/go-helm-client"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const secretName = "launchboxhq-credentials"

type Installer struct {
	Config *rest.Config
}

type InstallOpts struct {
	DryRun       bool
	Namespace    string
	ClientId     string
	ClientSecret string
}

func (i *Installer) Install(ctx context.Context, opts InstallOpts) error {
	kubeClient, err := kubernetes.NewForConfig(i.Config)
	if err != nil {
		return err
	}
	helmClient, err := helmclient.NewClientFromRestConf(&helmclient.RestConfClientOptions{
		RestConfig: i.Config,
	})
	if err != nil {
		return err
	}
	// Ensure our secret is created
	_, err = kubeClient.CoreV1().
		Secrets(opts.Namespace).
		Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		_, err := kubeClient.CoreV1().
			Secrets(opts.Namespace).
			Create(ctx, &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      secretName,
					Namespace: opts.Namespace,
				},
				Data: map[string][]byte{
					"clientId":     []byte(opts.ClientId),
					"clientSecret": []byte(opts.ClientSecret),
				},
			}, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	// Install the LaunchboxHQ operator, and
	operatorChartSpec := &helmclient.ChartSpec{}
	rel, err := helmClient.InstallOrUpgradeChart(ctx, operatorChartSpec, nil)
	if err != nil {
		return err
	}

	// wait for it to be ready

	// Create the Cluster resource, and wait for agent to start

	// Poll the cluster resources, and verify the agent is connected
	return nil
}
