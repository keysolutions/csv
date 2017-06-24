csv: A CSV file handling package for Go
=======================================

csv provides a set of utility functions to reduce the boilerplate necessary to use the csv package from the standard library when reading and writing to files. The Reader and Writer used by the package are composed of the standard library types and maintain full compatability with their usage.

Usage
-----

Get the package.

```
go get -u github.com/keysolutions/csv
```

Opening a csv file is done by calling the Open function, much like os.Open from the standard library. Insetad of receiving an *os.File when successful, a *csv.Reader is provided instead. The *csv.Reader can be read from just like *csv.Reader in the standard library.

```
r, _ := csv.Open("/path/to/file.csv")
for {
    row, err := r.Read()
    if err == io.EOF {
        break
    }
    // Use row contents.
}
r.Close()
```

Writer a csv file is done by calling the Create function. Like Open, the same semantics as the os package in the standard library. The *csv.Writer can be written to just like *csv.Writer in the standard library.

```
w, _ := csv.Create("/path/to/file.csv")
w.Write([]string{"column 1","column 2","column 3"})
w.Close()
```
