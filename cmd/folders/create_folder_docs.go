package folders

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateFolderDocs() {
	os.Mkdir(config.GetFullPath("/docs"), 0755)
}
