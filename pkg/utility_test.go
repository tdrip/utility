package utility

import (
	"testing"
)

func TestNewUtility(t *testing.T) {

	productcode := "dumbprod123"
	displayname := "fantastic product"
	version := "0.0.0.0"
	conffile := "./dumb.json"

	util := NewUtility(productcode, displayname, version, conffile)
	util.Configuration.Data["homealone"] = []string{"i'll", "give", "ya", "till", "the", "count", "of", "ten", "1", "2", "..", "10", "haha"}
	err := util.SaveConf()
	if err != nil {
		t.Errorf("'%s' - no err should be returned", err)
	}
}
