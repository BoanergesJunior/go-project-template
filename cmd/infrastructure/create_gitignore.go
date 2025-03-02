package infrastructure

import "os"

func CreateGitIgnore() {
	os.Create(".gitignore")
}
