package model_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template.git/config"
)

func CreateConnector() {
	os.Create(config.GetFullPath("/internal/app/model/connector.go"))
}
