package model_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template.git/config"
)

func CreateUsecase() {
	os.Create(config.GetFullPath("/internal/app/model/usecase.go"))
}
