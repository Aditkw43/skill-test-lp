package internal

import (
	"os"

	"skill-test-lp/shared"

	"github.com/gocarina/gocsv"
)

func GetData() []*RequestCsv {
	in, err := os.Open(shared.CsvFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	records := []*RequestCsv{}

	if err := gocsv.UnmarshalFile(in, &records); err != nil {
		panic(err)
	}

	return records
}
