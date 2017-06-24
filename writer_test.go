package csv

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteError(t *testing.T) {
	const path = "/path/to/a/nonexistant/directory/to/cause/an/error"
	r, err := Create(path)
	if err == nil {
		r.Close()
		os.Remove(path)
	}
	if _, ok := err.(*os.PathError); !ok {
		t.Fatalf("expected *os.PathError; got: %v", err)
	}
}

func TestWrite(t *testing.T) {
	const path = "writer_test.csv"
	r, err := Create(path)
	if err != nil {
		t.Fatalf("unexpected open error: %v", err)
	}

	for _, row := range exampleData {
		if err := r.Write(row); err != nil {
			t.Errorf("unexpected write error: %v", err)
		}
	}
	if err := r.Close(); err != nil {
		t.Errorf("unexpected close error: %v", err)
	}

	f1, _ := os.Open(path)
	f2, _ := os.Open("testdata/example.csv")
	b1, _ := ioutil.ReadAll(f1)
	b2, _ := ioutil.ReadAll(f2)
	if bytes.Compare(b1, b2) != 0 {
		t.Errorf("expected testdata/example.csv to be the same as %s", path)
	}

	f1.Close()
	f2.Close()
	os.Remove(path)
}
