package shared

import (
	"io/ioutil"
	"mime/multipart"

	model "skill-test-lp/internal"

	"github.com/gocarina/gocsv"
)

func GetData(file *multipart.FileHeader) []*model.Payroll {
	in, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer in.Close()

	fileBytes, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	records := []*model.Payroll{}
	if err := gocsv.UnmarshalBytes(fileBytes, &records); err != nil {
		panic(err)
	}

	return records
}
