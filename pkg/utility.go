package utility

import (
	sl "github.com/tdrip/logger/pkg
)

//Utility this class represents a simple Utility
type Utility struct {
	sl.AppLogger

	displayname string
	version     string
	productcode string
}

//NewUtility Creates a new applcation
func NewUtility(productcode string, displayname string, version string) *Utility {
	app := Utility{productcode: productcode, version: version, displayname: displayname}
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
