package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
				  New York,150,3,2,200000
				  Paris,200,4,3,400000
				  Istanbul,500,10,5,1000000
				  Buenos Aires,1,2,3,4
				  Oslo,8,23,4452,12`

		separator = ","
	)

	var (
		locations                  []string
		sizes, beds, baths, prices []int
		//avg                                   []float32
		avgsize, avgbeds, avgbaths, avgprices float32
	)

	headerl := strings.Split(header, separator)
	datal := []string{}
	for _, i := range strings.Split(data, "\n") { // data rows = table rows
		datal = append(datal, strings.Trim(i, "\t "))
	}

	// fmt.Println("headerl:", headerl)
	// fmt.Println("datal:", datal)

	for _, i := range datal {
		row := strings.Split(i, ",")
		locations = append(locations, row[0])
		if s, err := strconv.Atoi(row[1]); err == nil {
			sizes = append(sizes, s)
			avgsize = (avgsize + float32(s)) / float32(len(sizes))
		}
		if b, err := strconv.Atoi(row[2]); err == nil {
			beds = append(beds, b)
			avgbeds = (avgbeds + float32(b)) / float32(len(beds))
		}
		if ba, err := strconv.Atoi(row[3]); err == nil {
			baths = append(baths, ba)
			avgbaths = (avgbaths + float32(ba)) / float32(len(baths))
		}
		if p, err := strconv.Atoi(row[4]); err == nil {
			prices = append(prices, p)
			avgprices = (avgprices + float32(p)) / float32(len(prices))
		}
	}

	for i, _ := range headerl {
		fmt.Printf("%-15s", headerl[i])
	}
	fmt.Println()
	fmt.Printf("%s\n", strings.Repeat("=", 75))
	for i, _ := range datal {
		fmt.Printf("%-15v%-15d%-15d%-15d%-15d\n", locations[i], sizes[i], beds[i], baths[i], prices[i])
	}
	fmt.Printf("%s\n", strings.Repeat("=", 75))
	fmt.Printf("%-15s%-15.2f%-15.2f%-15.2f%-15.2f\n", "", avgsize, avgbeds, avgbaths, avgprices)

}
