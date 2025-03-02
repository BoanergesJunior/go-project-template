package infrastructure

import "os"

func CreateDockerfile() {
	os.Create("Dockerfile")
}
