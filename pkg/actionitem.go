package utility

//ActionItem this structure represents a action item to be done at start up or shutdown
type IActionItem interface {
	DoStartupChecks(app Utility) error
	DoShutdownChecks(app Utility) error
}

type ActionItem struct {
	Name string
}

//DoStartupChecks does the startup checks to make sure the app is operation
func (item ActionItem) DoStartupChecks(app Utility) error {
	return nil
}

//DoShutdownChecks does the shutdown checks to make sure the app is operation
func (item ActionItem) DoShutdownChecks(app Utility) error {
	return nil
}
