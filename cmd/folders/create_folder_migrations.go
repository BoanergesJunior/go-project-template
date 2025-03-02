package folders

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateFolderMigrations() {
	os.Mkdir(config.GetFullPath("/migrations"), 0755)

	os.Mkdir(config.GetFullPath("/migrations/migrations/up"), 0755)
	os.Mkdir(config.GetFullPath("/migrations/migrations/down"), 0755)
}
