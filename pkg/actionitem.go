package utility

//ActionItem this class represents a action item to valid the app at startup
type ActionItem struct {
}

//DoChecks does the Checks to make sure the app is operation
func (item *ActionItem) DoChecks(app *Utility) error {
	return nil
}
