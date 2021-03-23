package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

//Configuration this class represents a simple Utility Configuration
type Configuration struct {
	Data map[string][]string `json:"data,omitempty"`
}

//NewConfiguration Creates a new configuration
func NewConfiguration() *Configuration {
	conf := Configuration{}
	DataMap := make(map[string][]string)
	conf.Data = DataMap
	return &conf
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//LoadConfig Loads connections details from a configuration file
func LoadConfig(path string) (*Configuration, error) {

	// check file exists
	ok := fileExists(path)

	if ok {

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("loading %s failed with %s", path, err.Error())
		}

		// pointer to configuration
		conf := &Configuration{}

		// try to load it
		jsonerr := json.Unmarshal(bytes, conf)
		if jsonerr != nil {
			return nil, fmt.Errorf("marshal json for %s failed with %s", path, jsonerr.Error())
		}

		// sucess
		return conf, nil

	}

	// failed to find file
	return nil, fmt.Errorf("file '%s' was not found to load", path)
}

//SaveConfig saves config
func SaveConfig(path string, data *Configuration) error {

	if data == nil {
		return errors.New("data passed in was nil")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {

		// this needs to be human readable as it will be edited by humans
		bytes, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return fmt.Errorf("marshal json for %s failed with %s", path, err.Error())
		}

		err = ioutil.WriteFile(path, bytes, 0644)
		// failed to write error
		if err != nil {
			return fmt.Errorf("saving %s failed with %s", path, err.Error())
		}
	} else {
		e := os.Remove(path)
		if e == nil {
			return SaveConfig(path, data)
		}
		return fmt.Errorf("'%s' already exists so tried to delete however an error occurred: %s", path, e.Error())
	}

	return nil
}
