package infrastructure

import "os"

func CreateAirDebugToml() {
	os.Create(".air.debug.toml")
}
