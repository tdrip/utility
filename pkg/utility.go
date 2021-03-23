package utility

import (
	sl "github.com/tdrip/logger/pkg"
)

//Utility this class represents a simple Utility
type Utility struct {
	sl.AppLogger

	displayname string
	version     string
	conffile    string

	Configuration *Configuration

	StartupItems  map[string]IActionItem
	ShutdownItems map[string]IActionItem
}

//NewUtility Creates a new applcation
func NewUtility(displayname string, version string, conffile string) *Utility {
	app := Utility{version: version, displayname: displayname, conffile: conffile}
	app.Configuration = NewConfiguration()

	startupItems := make(map[string]IActionItem)
	app.StartupItems = startupItems

	shutdownItems := make(map[string]IActionItem)
	app.ShutdownItems = shutdownItems

	return &app
}

//GetVersion returns the version of the application
func (app *Utility) GetVersion() string {
	return app.version
}

//GetDisplayname returns the displayname for the application
func (app *Utility) GetDisplayname() string {
	return app.displayname
}

//LoadConf load utiliity configuration
func (app *Utility) LoadConf() error {

	if len(app.conffile) > 0 {
		conf, err := LoadConfig(app.conffile)
		if err != nil {
			return err
		} else {
			app.Configuration = conf
		}
	}

	return nil
}

//SaveConf save utiliity configuration
func (app *Utility) SaveConf() error {

	// check we have a file path and data otherwise error
	if len(app.conffile) > 0 && app.Configuration != nil {
		return SaveConfig(app.conffile, app.Configuration)
	}
	return nil
}

//AddStartupItem add startup item
func (app *Utility) AddStartupItem(key string, item IActionItem) {
	app.StartupItems[key] = item
}

//AddShutdownItem add shutdown item
func (app *Utility) AddShutdownItem(key string, item IActionItem) {
	app.ShutdownItems[key] = item

}

//Startup save utiliity configuration
func (app *Utility) Startup() error {

	for key, startup := range app.StartupItems {
		app.LogDebug("Startup", key)
		err := startup.DoChecks(app)
		if err != nil {
			return err
		}
	}
	return nil
}

//Shutdown do tasks on shutdown
func (app *Utility) Shutdown() error {

	for key, startup := range app.ShutdownItems {
		app.LogDebug("Shutdown", key)
		err := startup.DoChecks(app)
		if err != nil {
			return err
		}
	}
	return nil
}
