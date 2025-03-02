package workflows

import "os"

func CreateDeployProdEnvYaml() {
	os.Create(".github/workflows/deploy-prod-env.yaml")
}
