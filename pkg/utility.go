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

	Reports map[string]*Report

	Items map[string]IActionItem
}

//NewUtility Creates a new utility applcation
func NewUtility(displayname string, version string, conffile string) *Utility {
	util := Utility{version: version, displayname: displayname, conffile: conffile}
	util.Configuration = NewConfiguration()

	items := make(map[string]IActionItem)
	util.Items = items

	reports := make(map[string]*Report)
	util.Reports = reports

	return &util
}

//GetVersion returns the version of the application
func (util *Utility) GetVersion() string {
	return util.version
}

//GetDisplayname returns the displayname for the application
func (util *Utility) GetDisplayname() string {
	return util.displayname
}

//LoadConf load utiliity configuration
func (util *Utility) LoadConf() error {

	if len(util.conffile) > 0 {
		conf, err := LoadConfig(util.conffile)
		if err != nil {
			return err
		} else {
			util.Configuration = conf
		}
	}

	return nil
}

//SaveConf save utiliity configuration
func (util *Utility) SaveConf() error {

	// check we have a file path and data otherwise error
	if len(util.conffile) > 0 && util.Configuration != nil {
		return SaveConfig(util.conffile, util.Configuration)
	}
	return nil
}

//AddItem add startup item
func (util *Utility) AddItem(key string, item IActionItem) {
	data := util.Items
	data[key] = item
	util.Items = data
}

//Startup save utiliity configuration
func (util *Utility) Startup() error {

	for key, startup := range util.Items {
		util.LogDebug("Startup", key)
		err := startup.DoStartupChecks(util)
		if err != nil {
			util.LogErrorE("Startup", err)
			return err
		}
	}
	return nil
}

//Shutdown do tasks on shutdown
func (util *Utility) Shutdown() error {

	for key, startup := range util.Items {
		util.LogDebug("Shutdown", key)
		err := startup.DoShutdownChecks(util)
		if err != nil {
			util.LogErrorE("Shutdown", err)
			return err
		}
	}

	// close reports
	for _, sud := range util.Reports {
		if sud.File != nil {
			sud.File.Close()
		}
	}

	return nil
}

func (util *Utility) FlushRecords() {

	for _, sud := range util.Reports {
		if sud.Writer != nil {
			sud.Writer.Flush()
		}
	}
}

//AddItem utility and key, item
func AddUtilityItem(util *Utility, key string, item IActionItem) *Utility {
	util.AddItem(key, item)
	return util
}

// WriteRecord to report
func (util *Utility) WriteRecord(name string, record []string) error {
	report := util.Reports[name]
	return report.WriteRecord(record)
}

// Add a custom report to the tool
func (util *Utility) AddCustomReport(name string) error {

	rep, err := CreateReport(name)
	if err != nil {
		return err
	}

	util.Reports[name] = rep

	return nil
}
