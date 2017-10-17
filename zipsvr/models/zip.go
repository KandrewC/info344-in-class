package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Zip struct {
	/* These have to be upper case because we have to export the file
	We'll make these lower case in the json later
	*/
	Code  string
	City  string
	State string
}

type ZipSlice []*Zip

type ZipIndex map[string]ZipSlice

func LoadZips(fileName string) (ZipSlice, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}

	reader := csv.NewReader(f)
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("Error reading header row: %v", err)
	}
	// We can preallocate the underlying array so golang doesn't take time
	// allocating memory
	zips := make(ZipSlice, 0, 43000)
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("Error reading record: %v", err)
		}
		z := &Zip{ //why is the & there
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z)
	}
}
