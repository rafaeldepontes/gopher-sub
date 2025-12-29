package migration

import (
	"os"
	"os/exec"

	"github.com/rafaeldepontes/gopher-sub/internal/logger"
)

func Init() error {
	log := logger.GetLogger()

	log.Infoln("Initializing migrations...")
	cmd := exec.Command("flyway", "migrate", "-configFiles=flyway.conf")
	cmd.Env = os.Environ()

	resp, err := cmd.CombinedOutput()
	log.Infoln("Migration response:", string(resp))
	if err != nil {
		return err
	}

	log.Infoln("Migrations finished successfully!")
	return nil
}
