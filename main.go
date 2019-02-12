package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	DATE_TIME_FORMAT = "2006/01/02 15:04:05.000"
)

func main() {

	var (
		d = flag.String("d","2006/01/02 15:04:05.000", "date")
		z = flag.String("z","UTC", "time zone")
		e = flag.String("e","","epoch")
	)
	flag.Parse()
	if len(*e) == 13 {
		i, err :=strconv.ParseInt(*e,10,64)
		if err != nil {
			panic(err)
		}
		loc2, _ := time.LoadLocation(*z)
		tu := time.Unix(0, i * 1000000).In(loc2)
		fmt.Println(tu.Format("2006/01/02 15:04:05.000 Z07:00"))
		os.Exit(0)
	}
	loc, _ := time.LoadLocation(*z)
	t,_ := time.ParseInLocation(DATE_TIME_FORMAT, *d, loc)
	fmt.Printf("Pointed time: %s\n",t)
	fmt.Println(t.UnixNano()/ 1000000)

}
