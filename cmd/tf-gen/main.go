package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/lyraproj/terraform-bridge/pkg/bridge"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm"
	"github.com/terraform-providers/terraform-provider-github/github"
	"github.com/terraform-providers/terraform-provider-google/google"
	"github.com/terraform-providers/terraform-provider-kubernetes/kubernetes"
)

func main() {

	// AWS
	fmt.Println("Generating AWS provider ...")
	bridge.Generate(aws.Provider().(*schema.Provider), "TerraformAws", "cmd/goplugin-terraformaws/generated/generated.go")

	// Azure
	fmt.Println("Generating Azure provider ...")
	bridge.Generate(azurerm.Provider().(*schema.Provider), "TerraformAzureRM", "cmd/goplugin-terraformazurerm/generated/generated.go")

	// GitHub
	fmt.Println("Generating GitHub provider ...")
	bridge.Generate(github.Provider().(*schema.Provider), "TerraformGitHub", "cmd/goplugin-terraformgithub/generated/generated.go")

	// Google
	fmt.Println("Generating Google provider ...")
	bridge.Generate(google.Provider().(*schema.Provider), "TerraformGoogle", "cmd/goplugin-terraformgoogle/generated/generated.go")

	// Kubernetes
	fmt.Println("Generating Kubernetes provider ...")
	bridge.Generate(kubernetes.Provider().(*schema.Provider), "TerraformKubernetes", "cmd/goplugin-terraformkubernetes/generated/generated.go")

}
