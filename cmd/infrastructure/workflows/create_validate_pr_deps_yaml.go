package workflows

import "os"

func CreateValidatePrDepsYaml() {
	os.Create(".github/workflows/validate-pr-deps.yaml")
}
