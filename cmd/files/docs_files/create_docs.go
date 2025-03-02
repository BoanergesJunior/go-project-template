package docs_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateDocs() {
	os.Create(config.GetFullPath("/docs/docs.go"))
}
