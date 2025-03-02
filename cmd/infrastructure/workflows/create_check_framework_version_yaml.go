package workflows

import "os"

func CreateCheckFrameworkVersionYaml() {
	os.Create(".github/workflows/check-framework-version.yaml")
}
