package product

type List struct {
	Name     string
	Price    int
	Category string
}

var Items = []*List{
	// Makanan
	{Name: "Nasi Goreng Spesial", Price: 30000, Category: "Makanan"},
	{Name: "Mie Ayam Bakso", Price: 28000, Category: "Makanan"},
	{Name: "Ayam Goreng Krispi", Price: 25000, Category: "Makanan"},
	{Name: "Nasi Ayam Geprek", Price: 27000, Category: "Makanan"},
	{Name: "Sate Ayam", Price: 35000, Category: "Makanan"},
	{Name: "Bakso Urat", Price: 32000, Category: "Makanan"},
	{Name: "Soto Ayam", Price: 29000, Category: "Makanan"},
	{Name: "Rendang Sapi", Price: 40000, Category: "Makanan"},

	// Minuman
	{Name: "Es Teh Manis", Price: 10000, Category: "Minuman"},
	{Name: "Es Jeruk", Price: 12000, Category: "Minuman"},
	{Name: "Es Latte", Price: 40000, Category: "Minuman"},
	{Name: "Capuccino", Price: 50000, Category: "Minuman"},
	{Name: "Es Milo", Price: 15000, Category: "Minuman"},
	{Name: "Air Mineral", Price: 8000, Category: "Minuman"},
	{Name: "Kopi Hitam", Price: 15000, Category: "Minuman"},
	{Name: "Thai Tea", Price: 18000, Category: "Minuman"},

	// Snack
	{Name: "Eskrim Cokelat", Price: 20000, Category: "Snack"},
	{Name: "Kentang Goreng", Price: 15000, Category: "Snack"},
	{Name: "Cireng Isi", Price: 12000, Category: "Snack"},
	{Name: "Tahu Krispi", Price: 13000, Category: "Snack"},
	{Name: "Dimsum", Price: 17000, Category: "Snack"},
	{Name: "Pisang Coklat", Price: 16000, Category: "Snack"},
	{Name: "Siomay Mini", Price: 14000, Category: "Snack"},
}
