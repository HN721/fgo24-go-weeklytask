package product

import "fmt"

type history struct {
	nama     string
	qty      int
	subtotal int
}

var Hisorydata = []history{}

func HistoryTransaction(name string, qty int, subTotal int) {
	data := history{
		nama:     name,
		qty:      qty,
		subtotal: subTotal,
	}
	Hisorydata = append(Hisorydata, data)
	fmt.Println(Hisorydata)
}

func ShowHistory() {
	fmt.Println("======= History Transaction =========")
	if len(Hisorydata) > 0 {
		for i, item := range Hisorydata {
			fmt.Printf("%v. %v x%v - Rp.%d\n", i+1, item.nama, item.qty, item.subtotal)
		}
	} else {
		fmt.Println("History Transaction anda kosong")

	}

}
