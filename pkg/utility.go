package utility

import (
	sl "github.com/tdrip/logger/pkg"
)

//Utility this class represents a simple Utility
type Utility struct {
	sl.AppLogger

	displayname string
	version     string
	productcode string
	conffile    string

	Configuration *Configuration

	StartupItems  []*ActionItem
	ShutdownItems []*ActionItem
}

//NewUtility Creates a new applcation
func NewUtility(productcode string, displayname string, version string, conffile string) *Utility {
	app := Utility{productcode: productcode, version: version, displayname: displayname, conffile: conffile}
	app.Configuration = NewConfiguration()
	app.StartupItems = []*ActionItem{}
	app.ShutdownItems = []*ActionItem{}
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

//GetProductCode gets the product code
func (app *Utility) GetProductCode() string {
	return app.productcode
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
func (app *Utility) AddStartupItem(item *ActionItem) {
	items := app.StartupItems
	items = append(items, item)
	app.StartupItems = items
}

//AddShutdownItem add shutdown item
func (app *Utility) AddShutdownItem(item *ActionItem) {
	items := app.ShutdownItems
	items = append(items, item)
	app.ShutdownItems = items
}

//Startup save utiliity configuration
func (app *Utility) Startup() error {

	for _, startup := range app.StartupItems {
		err := startup.DoChecks(app)
		if err != nil {
			return err
		}
	}
	return nil
}

//Shutdown do tasks on shutdown
func (app *Utility) Shutdown() error {

	for _, startup := range app.ShutdownItems {
		err := startup.DoChecks(app)
		if err != nil {
			return err
		}
	}
	return nil
}
