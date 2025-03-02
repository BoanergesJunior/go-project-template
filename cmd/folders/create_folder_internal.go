package folders

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateFolderInternal() {
	os.Mkdir(config.GetFullPath("/internal"), 0755)

	os.Mkdir(config.GetFullPath("/internal/app"), 0755)
	os.Mkdir(config.GetFullPath("/internal/delivery"), 0755)

	os.Mkdir(config.GetFullPath("/internal/app/model"), 0755)
	os.Mkdir(config.GetFullPath("/internal/app/repository"), 0755)
	os.Mkdir(config.GetFullPath("/internal/app/usecase"), 0755)

	os.Mkdir(config.GetFullPath("/internal/delivery/http"), 0755)
	os.Mkdir(config.GetFullPath("/internal/delivery/grpc"), 0755)

	os.Mkdir(config.GetFullPath("/internal/delivery/http/connector"), 0755)
	os.Mkdir(config.GetFullPath("/internal/delivery/http/handler"), 0755)
}
