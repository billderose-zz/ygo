package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type business struct {
	Business_id, Name, Full_address, City, State, Type string
	Neighborhoods, Categories                          []string
	Reviwew_count                                      int
	Stars, Longitude, Latitude                         float64
	Open                                               bool
}



func (b business) String() string {
	return fmt.Sprintf("%s\n %s\n", b.Name, b.Full_address)
}

var data = flag.String("source", "/Users/Bill/Downloads/yelp/yelp_academic_dataset_business.json", "Data to read")

func main() {
	file, err := os.Open(*data)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		b := &business{}
		if s, err := r.ReadSlice('\n'); err != nil {
			break
		} else if err := json.Unmarshal(s, &b); err != nil {
			panic(err)
		}
		fmt.Println(b)
	}
}
