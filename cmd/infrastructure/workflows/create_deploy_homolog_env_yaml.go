package workflows

import "os"

func CreateDeployHomologEnvYaml() {
	os.Create(".github/workflows/deploy-homolog-env.yaml")
}
