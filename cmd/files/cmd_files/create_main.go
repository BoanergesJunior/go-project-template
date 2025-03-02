package cmd_files

import (
	"os"

	"github.com/BoanergesJunior/go-project-template/config"
)

func CreateMain() {
	os.Create(config.GetFullPath("/cmd/application/main.go"))
}
