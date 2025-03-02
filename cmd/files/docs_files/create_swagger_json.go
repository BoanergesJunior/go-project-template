package docs_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateSwaggerJson() {
	os.Create(config.GetFullPath("/docs/swagger.json"))
}
