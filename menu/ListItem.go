package menu

import (
	"fmt"
	"homework/product"
)

func ListDrinkByCategory(items []*product.List, category string) []*product.List {
	fmt.Printf("\n======= Daftar %s =======\n", category)
	filtered := []*product.List{}

	for _, item := range items {
		if item.Category == category {
			filtered = append(filtered, item)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("Tidak ada item pada kategori ini.")
		return nil
	}
	var qty int
	var num int
	for i, item := range filtered {
		fmt.Printf("%d. %s - Rp %d\n", i+1, item.Name, item.Price)
	}
	fmt.Print("Pilih: ")
	fmt.Scanln(&num)
	selectedItem := filtered[num-1]
	fmt.Printf("Berapa jumlah '%s' yang ingin dibeli? ", selectedItem.Name)
	fmt.Scanln(&qty)
	product.AddCart(qty, selectedItem.Price, selectedItem.Name)

	return filtered
}
