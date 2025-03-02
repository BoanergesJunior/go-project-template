package docs_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template.git/config"
)

func CreateSwaggerYaml() {
	os.Create(config.GetFullPath("/docs/swagger.yaml"))
}
