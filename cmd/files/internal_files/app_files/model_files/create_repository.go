package model_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateRepository() {
	os.Create(config.GetFullPath("/internal/app/model/repository.go"))
}
