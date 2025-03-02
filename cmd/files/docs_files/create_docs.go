package docs_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template.git/config"
)

func CreateDocs() {
	os.Create(config.GetFullPath("/docs/docs.go"))
}
