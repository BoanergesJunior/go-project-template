package infrastructure

import "os"

func CreateEnvExample() {
	os.Create(".env.example")
}
