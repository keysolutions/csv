package csv

import (
	"encoding/csv"
	"os"
)

// Reader reads records from a CSV file. It is composed of a csv.Reader from
// the standard library. See it for more information on reading the CSV file.
type Reader struct {
	*csv.Reader

	closer func() error
}

// Open opens a CSV file for reading. If there is an error, it will be of
// type *os.PathError.
func Open(name string) (*Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return &Reader{
		Reader: csv.NewReader(f),
		closer: f.Close,
	}, nil
}

// Close will close the file opened by the CSV reader. The CSV contents will
// not be readable after closing the file.
func (r *Reader) Close() error {
	return r.closer()
}
