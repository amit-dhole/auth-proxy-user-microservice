package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"errors"

	"gopkg.in/go-playground/validator.v9"
)

var (
	// Config is a package variable, which is populated during init() execution and shared to whole application
	Config Configuration
	// APIVersions stores slice of supported versions
	APIVersions = []string{"v1"}

	// filePath defines a path to JSON-config file
	filePath = "config.json"
)

// Configuration options
// nolint: lll
type Configuration struct {
	ListenURL          string `json:"listenURL"`
	URLPrefix          string `json:"urlPrefix"`
	HTTPProtocol       string `json:"httpProtocol"`
	User               string `json:"user"`
	Password           string `json:"password"`
	CertificateFile    string `json:"certificateFile"`
	CertificateKeyFile string `json:"certificateKeyFile"`
	Authentication     bool   `json:"authentication"`
}

// SetFilePath set configuration file path
func SetFilePath(path string) {
	filePath = path
}

// Load reads and loads configuration to Config variable
func Load() (err error) {

	if len(filePath) != 0 {
		err = readConfigFromJSON(filePath)

		if err != nil {
			return errors.New("config initialization failed - %v " + err.Error())
		}

		if err = validate(&Config); err != nil {
			return errors.New("config validation failed" + err.Error())
		}
	}
	return
}

// readConfigFromJSON reads config data from JSON-file
func readConfigFromJSON(configFilePath string) error {
	log.Printf("Looking for JSON config file (%s)", configFilePath)

	contents, err := ioutil.ReadFile(configFilePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&Config)
	}

	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %s\n", configFilePath, err)
		return errors.New("Reading configuration from JSON failed: " + err.Error())
	}
	log.Printf("Configuration has been read from JSON (%s) successfully\n", configFilePath)
	return nil
}

// validate validates a struct and nested fields
func validate(c *Configuration) error {
	v := validator.New()
	return v.Struct(c)
}
