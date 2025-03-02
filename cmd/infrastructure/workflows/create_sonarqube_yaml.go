package workflows

import "os"

func CreateSonarqubeYaml() {
	os.Create(".github/workflows/sonarqube.yaml")
}
