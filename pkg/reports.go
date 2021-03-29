package utility

import (
	"encoding/csv"
	"errors"
	"os"
	"time"
)

type Report struct {
	Name   string
	File   *os.File
	Writer *csv.Writer
}

// adds a report
func CreateReport(name string) (*Report, error) {
	filename := name + "-" + time.Now().Format("2006-01-02-15-04-05") + ".csv"

	sudfile, err := os.Create(filename)

	if err != nil {
		return nil, err
	}

	sudwriter := csv.NewWriter(sudfile)

	report := Report{name, sudfile, sudwriter}

	return &report, nil
}

//NewUtility Creates a new utility applcation
func (report *Report) WriteRecord(record []string) error {
	if report.Writer != nil {
		err := report.Writer.Write(record)
		return err
	}
	return errors.New("when writing a record the writer was nil")
}
