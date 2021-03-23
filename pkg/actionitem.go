package utility

//ActionItem this class represents a action item to valid the app at startup
type ActionItem struct {
}

//DoCheck - does tshe Checks to make sure the app is operation
func (item *ActionItem) DoCheck(app *Utility) error {
	return nil
}
