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

	Items map[string]IActionItem
}

//NewUtility Creates a new utility applcation
func NewUtility(displayname string, version string, conffile string) *Utility {
	app := Utility{version: version, displayname: displayname, conffile: conffile}
	app.Configuration = NewConfiguration()

	items := make(map[string]IActionItem)
	app.Items = items

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

//AddItem add startup item
func (app *Utility) AddItem(key string, item IActionItem) {
	data := app.Items
	data[key] = item
	app.Items = data
}

//Startup save utiliity configuration
func (app *Utility) Startup() error {

	for key, startup := range app.Items {
		app.LogDebug("Startup", key)
		err := startup.DoStartupChecks(app)
		if err != nil {
			app.LogErrorE("Startup", err)
			return err
		}
	}
	return nil
}

//Shutdown do tasks on shutdown
func (app *Utility) Shutdown() error {

	for key, startup := range app.Items {
		app.LogDebug("Shutdown", key)
		err := startup.DoShutdownChecks(app)
		if err != nil {
			app.LogErrorE("Shutdown", err)
			return err
		}
	}
	return nil
}
