package http_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateNew() {
	os.Create(config.GetFullPath("/internal/delivery/http/new.go"))
}
