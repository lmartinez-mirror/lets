package workdir

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const dotLetsDir = ".lets"

const defaultLetsYaml = `version: "%s"
shell: bash

commands:
  hello:
	description: Say hello
	options: |
		Usage: lets hello [<name>]
		Examples:
			lets hello
			lets hello Friend
	cmd: echo Hello, "${LETSOPT_NAME:-world}"!
`

func GetDotLetsDir(workDir string) (string, error) {
	return filepath.Abs(filepath.Join(workDir, dotLetsDir))
}

// InitLetsFile creates lets.yaml int the current dir.
func InitLetsFile(workDir string, version string) error {
	configfile := filepath.Join(workDir, "lets.yaml")

	if _, err := os.Stat(configfile); err == nil {
		return fmt.Errorf("lets.yaml already exists in %s", workDir)
	}

	output := fmt.Sprintf(defaultLetsYaml, version)
	//#nosec G306
	if err := os.WriteFile(configfile, []byte(output), 0o644); err != nil {
		return err
	}

	log.Println("lets.yaml created in the current directory")

	return nil
}
