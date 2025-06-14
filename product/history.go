package product

import "fmt"

type history struct {
	nama string
	qty  int
}

var Hisorydata = []history{}

func HistoryTransaction() {
	data := history{
		nama: "sss",
		qty:  2,
	}
	Hisorydata = append(Hisorydata, data)
	fmt.Println(Hisorydata)
}
