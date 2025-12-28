package tool

import "os"

// ChecksEnv verifies if the env file exists in the project if not, it will
// try to use the .env.example path in the folder root. If none exists then
// it will panic.
func ChecksEnv(env *string) {
	if _, err := os.Stat(*env); err != nil {
		*env = ".env.example"
	}
}
