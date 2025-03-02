package utils

import (
	"github.com/BoanergesJunior/go-project-template.git/cmd/files"
	"github.com/BoanergesJunior/go-project-template.git/cmd/folders"
)

func CreateFolders() {
	folders.CreateFolderCmd()
	folders.CreateFolderDocs()
	folders.CreateFolderInternal()
	folders.CreateFolderMigrations()
	folders.CreateFolderSeeds()

	files.CreateFiles()
}
