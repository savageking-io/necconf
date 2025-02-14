package necconf

import (
	"fmt"
	"io/fs"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	configDirectory string
}

// Init initializes the Config struct
func (c *Config) Init(configDirectory string) error {
	log.Traceln("necconf::Init")

	if configDirectory == "" {
		return fmt.Errorf("config directory is empty")
	}

	c.configDirectory = configDirectory

	return nil
}

// ReadConfig reads the config file and unmarshals it into the conf interface
func (c *Config) ReadConfig(fsys fs.FS, filename string, conf interface{}) error {
	log.Traceln("necconf::ReadConfig")

	log.Debugf("necconf::ReadConfig: fsys: %+v", fsys)
	log.Debugf("necconf::ReadConfig: filename: %s", filename)

	if c.configDirectory == "" {
		return fmt.Errorf("config directory is empty: not initialized")
	}

	if conf == nil {
		return fmt.Errorf("conf interface is nil")
	}

	if fsys == nil {
		return fmt.Errorf("fsys is nil")
	}

	if filename == "" {
		return fmt.Errorf("config filename is empty")
	}

	content, err := fs.ReadFile(fsys, filename)
	if err != nil {
		return fmt.Errorf("failed to read config: %s", err.Error())
	}

	log.Debugf("necconf::ReadConfig: content: %s", string(content))

	err = yaml.Unmarshal(content, conf)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err.Error())
	}
	return nil
}
