package infrastructure

import "os"

func CreateDevDockerfile() {
	os.Create("dev.Dockerfile")
}
