package files

import (
	"github.com/BoanergesJunior/go-project-template/cmd/files/cmd_files"
	"github.com/BoanergesJunior/go-project-template/cmd/files/docs_files"
	"github.com/BoanergesJunior/go-project-template/cmd/files/internal_files/app_files/model_files"
	"github.com/BoanergesJunior/go-project-template/cmd/files/internal_files/app_files/repository_files"
	"github.com/BoanergesJunior/go-project-template/cmd/files/internal_files/app_files/usecase_files"
	"github.com/BoanergesJunior/go-project-template/cmd/files/internal_files/delivery_files/connector_files"
	"github.com/BoanergesJunior/go-project-template/cmd/files/internal_files/delivery_files/http_files"
)

func CreateFiles() {
	cmd_files.CreateMain()

	docs_files.CreateDocs()
	docs_files.CreateSwaggerJson()
	docs_files.CreateSwaggerYaml()

	model_files.CreateConnector()
	model_files.CreateRepository()
	model_files.CreateUsecase()

	repository_files.CreateNew()

	usecase_files.CreateNew()

	connector_files.CreateNew()

	http_files.CreateNew()
}
