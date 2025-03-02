package infrastructure

import "os"

func CreateWorkflows() {
	os.Create(".github/workflows")
}
