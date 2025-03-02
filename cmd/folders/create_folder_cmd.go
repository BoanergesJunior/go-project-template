package folders

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateFolderCmd() {
	os.Mkdir(config.GetFullPath("/cmd"), 0755)

	os.Mkdir(config.GetFullPath("/cmd/application"), 0755)
}
