package folders

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateFolderSeeds() {
	os.Mkdir(config.GetFullPath("/seeds"), 0755)

	os.Mkdir(config.GetFullPath("/seeds/dev"), 0755)
	os.Mkdir(config.GetFullPath("/seeds/hml"), 0755)
	os.Mkdir(config.GetFullPath("/seeds/prod"), 0755)
}
