/*
Package config implements the ConfigContainer structure and accompanying methods.
It holds the configuration data for all utilities.
Configuration file format is the same for all.
*/
package config

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigContainer struct {
	Logfile          string                       `json:"logfile"`
	Binpaths         map[string]string            `json:"binpaths"`
	DefaultConnector string                       `json:"defaultconnector"`
	Connectors       map[string]map[string]string `json:"connectors"`
	QosMap           map[uint64]string            `json:"qosmap"`
}

func NewConfigContainer() *ConfigContainer {
	return new(ConfigContainer)
}

// Read & unmarshall configuration from 'name' file into configContainer structure
func (cc *ConfigContainer) GetConfig(name string) error {
	f, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	err = json.Unmarshal(f, cc)
	if err != nil {
		return err
	}

        if cc.Binpaths == nil {
	    cc.Binpaths = make(map[string]string)
	}

	// set default paths
        defaultpaths := map[string]string{
            "sacct": "/usr/bin/sacct",
            "sstat": "/usr/bin/sstat",
        }

        for key, path := range defaultpaths {
            if _, exists := cc.Binpaths[key]; !exists {
                cc.Binpaths[key] = path
            }
        }

	return nil
}

func (cc *ConfigContainer) DumpConfig(l *log.Logger) {
	l.Println("DUMP CONFIG:")
	l.Printf("CONFIGURATION: %#v\n", cc)
	l.Printf("CONFIGURATION logfile: %s\n", cc.Logfile)
	l.Printf("CONFIGURATION msteams.name: %s\n", cc.Connectors["msteams"]["name"])
	l.Println("--------------------------------------------------------------------------------")
}
