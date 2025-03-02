package http_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template.git/config"
)

func CreateNew() {
	os.Create(config.GetFullPath("/internal/delivery/http/new.go"))
}
