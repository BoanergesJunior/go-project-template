package usecase_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateNew() {
	os.Create(config.GetFullPath("/internal/app/usecase/new.go"))
}
