package workflows

import "os"

func CreateDeployDevYaml() {
	os.Create(".github/workflows/deploy-dev.yaml")
}
