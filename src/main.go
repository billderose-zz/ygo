package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Business struct {
	Business_id, Name, Full_address, City, State, Type string
	Neighborhoods, Categories                          []string
	Reviwew_count                                      int
	Stars, Longitude, Latitude                         float64
	Open                                               bool
}

func (b Business) String() string {
	if len(b.Categories) != 0 {
		return fmt.Sprintf("%f\t%f\t%s\n", b.Latitude, b.Longitude, strings.Replace(b.Categories[0], ",", "", -1))
	}
	return fmt.Sprintf("%f\t%f\t%s\n", b.Latitude, b.Longitude, "N/A")

}

func main() {
	data := flag.String("source", "", "Input data")
	flag.Parse()

	file, err := os.Open(*data)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writeFile, err := os.Create("yelps.txt")
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(writeFile)
	w.WriteString("lat\tlon\tcat\n")
	r := bufio.NewReader(file)

	for {
		b := &Business{}
		if s, err := r.ReadSlice('\n'); err != nil {
			break
		} else if err := json.Unmarshal(s, &b); err != nil {
			panic(err)
		}
		w.WriteString(b.String())
	}
}
