package utility

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewUtility(t *testing.T) {

	displayname := "fantastic product"
	version := "0.0.0.0"
	conffile := "./dumb.json"

	util := NewUtility(displayname, version, conffile)
	util.Configuration.Data["homealone"] = []string{"i'll", "give", "ya", "till", "the", "count", "of", "ten", "1", "2", "..", "10", "haha"}
	err := util.SaveConf()
	if err != nil {
		t.Errorf("'%s' - no err should be returned", err)
	}

	failai := TestActionItem{Fail: true, CrazyLookup: "Banana"}

	util.AddStartupItem("test", IActionItem(&failai))

	err = util.Startup()

	if err == nil {
		t.Errorf("'%s' - err should be returned", "TestActionItem")
	}

	item := util.StartupItems["test"]
	tai := item.(*TestActionItem)
	if len(tai.CrazyLookup) == 0 {
		t.Errorf("'%s' - CrazyLookup should be returned", tai.CrazyLookup)
	}
}

type TestActionItem struct {
	*ActionItem
	Fail        bool
	CrazyLookup string
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
