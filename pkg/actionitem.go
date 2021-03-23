package utility

//ActionItem this class represents a action item to be done at start up or shutdown
type IActionItem interface {
	DoChecks(app *Utility) error
}

type ActionItem struct {
	Name string
}

//DoChecks does the Checks to make sure the app is operation
func (item *ActionItem) DoChecks(app *Utility) error {
	return nil
}
