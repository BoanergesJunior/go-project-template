package infrastructure

import "os"

func CreateAirToml() {
	os.Create(".air.toml")
}
