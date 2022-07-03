package main

import (
	"fmt"
	"github.com/leekchan/accounting"
	"go-module/c4f"
)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 3}
	fmt.Println(ac.FormatMoney(123456789.213123))
	c4f.Println("We have red")
}
