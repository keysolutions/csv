package csv

import (
	"encoding/csv"
	"os"
)

// Writer writes records to a CSV file. It is composed of a csv.Writer from
// the standard library. See it for more information on writing the CSV file.
type Writer struct {
	*csv.Writer

	closer func() error
}

// Create creates a new CSV file, truncating the file if it already exists.
// If there is an error, it will be of type *PathError.
func Create(name string) (*Writer, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return &Writer{
		Writer: csv.NewWriter(f),
		closer: f.Close,
	}, nil
}

// Close will close the file created by the CSV writer. Flush will be called to
// flush any outstanding buffered data. The CSV contents will  not be writable after
// closing the file.
func (w *Writer) Close() error {
	w.Flush()
	return w.closer()
}
