package cmd

import (
	"fmt"
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"log"
)

var updateKubeconfigCmd = &cobra.Command{
	Use:   "update-kubeconfig",
	Short: "Get kubeconfig credentials for a project",
	Run:   updateKubeconfig,
}

func init() {
	updateKubeconfigCmd.Flags().Int("project-id", 0, "ID of the project")
	_ = updateKubeconfigCmd.MarkFlagRequired("project-id")

	updateKubeconfigCmd.Flags().String("oidc-issuer-url", "https://launchboxhq.io", "OIDC Endpoint for Launchbox")
	updateKubeconfigCmd.Flags().String("client-id", "fYfghZAGwU5r2H_UXK_VtNSEOxzC0A0nGni_QJKLNfE", "Client ID for cluster authentication")

	rootCmd.AddCommand(updateKubeconfigCmd)
}

func updateKubeconfig(cmd *cobra.Command, args []string) {
	projectId, _ := cmd.Flags().GetInt("project-id")
	oidcIssuer, _ := cmd.Flags().GetString("oidc-issuer-url")
	clientId, _ := cmd.Flags().GetString("client-id")
	projectSdk := project.New(conf)
	result, err := projectSdk.Get(&project.GetProjectInput{
		ProjectId: projectId,
	})
	if err != nil {
		log.Fatal(err)
	}
	//host := result.Project.Host
	caCert := result.Project.CaCertificate
	slug := result.Project.Slug

	currentConfig, err := clientcmd.LoadFromFile("/Users/rwittman/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	currentConfig.Clusters[slug] = &clientcmdapi.Cluster{
		Server:                   fmt.Sprintf("https://api.%s.launchboxhq.io", slug),
		CertificateAuthorityData: []byte(caCert),
	}

	currentConfig.AuthInfos[slug] = &clientcmdapi.AuthInfo{
		Exec: &clientcmdapi.ExecConfig{
			APIVersion: "client.authentication.k8s.io/v1beta1",
			Args: []string{
				"oidc-login",
				"get-token",
				"--oidc-issuer-url=" + oidcIssuer,
				"--oidc-client-id=" + clientId,
				"--oidc-extra-scope=email",
			},
			Command: "kubectl",
		},
	}
	currentConfig.Contexts[slug] = &clientcmdapi.Context{
		AuthInfo: slug,
		Cluster:  slug,
	}
	currentConfig.CurrentContext = slug

	err = clientcmd.WriteToFile(*currentConfig, "/Users/rwittman/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current context set to " + slug)
}
