package config

var BasePath string

func GetFullPath(relativePath string) string {
	return BasePath + relativePath
}
