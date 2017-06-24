package csv

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestOpenError(t *testing.T) {
	r, err := Open("testdata/notfound.csv")
	if err == nil {
		r.Close()
	}
	if _, ok := err.(*os.PathError); !ok {
		t.Fatalf("expected *os.PathError; got: %v", err)
	}
}

func TestRead(t *testing.T) {
	r, err := Open("testdata/example.csv")
	if err != nil {
		t.Fatalf("unexpected open error: %v", err)
	}
	defer r.Close()

	for i := 0; ; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("unexpected read error: %v", err)
		}
		if !reflect.DeepEqual(row, exampleData[i]) {
			t.Errorf("expected %v; got: %v", exampleData[i], row)
		}
	}
}
