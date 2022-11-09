package internal

import (
	"io/ioutil"
	"mime/multipart"

	"github.com/gocarina/gocsv"
)

func GetData(file *multipart.FileHeader) []*Payroll {
	in, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer in.Close()

	fileBytes, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	records := []*Payroll{}
	if err := gocsv.UnmarshalBytes(fileBytes, &records); err != nil {
		panic(err)
	}

	return records
}
