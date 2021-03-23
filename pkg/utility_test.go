package utility

import (
	"errors"
	"fmt"
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

	failai := TestActionItem{Fail: true}

	util.AddStartupItem(IActionItem(&failai))

	err = util.Startup()

	if err == nil {
		t.Errorf("'%s' - err should be returned", "TestActionItem")
	}

}

type TestActionItem struct {
	*ActionItem
	Fail bool
}

//Checks - does the Checks to make sure the app is operation
func (item *TestActionItem) DoChecks(app *Utility) error {

	fmt.Println("Doing TestActionItem checks")

	if item.Fail {
		return errors.New("This should fail")
	} else {
		return nil
	}

}
