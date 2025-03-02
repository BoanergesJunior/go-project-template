package docs_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template.git/config"
)

func CreateSwaggerJson() {
	os.Create(config.GetFullPath("/docs/swagger.json"))
}
