package folders

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateFolderGithub() {
	os.Mkdir(config.GetFullPath("/.github"), 0755)

	os.Mkdir(config.GetFullPath("/.github/workflows"), 0755)
}
