package folders

import (
	"os"

	"github.com/BoanergesJunior/go-project-template.git/config"
)

func CreateFolderGithub() {
	os.Mkdir(config.GetFullPath("/.github"), 0755)

	os.Mkdir(config.GetFullPath("/.github/workflows"), 0755)
}
