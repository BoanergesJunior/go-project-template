package repository_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateNew() {
	os.Create(config.GetFullPath("/internal/app/repository/new.go"))
}
