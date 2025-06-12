package product

type List struct {
	Name     string
	Price    int
	Category string
}

var Items = []*List{
	{
		Name:     "Eskrim",
		Price:    20000,
		Category: "Snack",
	},
	{
		Name:     "Es Latte",
		Price:    40000,
		Category: "Minuman",
	},
	{
		Name:     "Capuchino",
		Price:    50000,
		Category: "Minuman",
	},
	{
		Name:     "Nasi Goreng",
		Price:    30000,
		Category: "Makanan",
	},
	{
		Name:     "Mie Ayam",
		Price:    28000,
		Category: "Makanan",
	},
	{
		Name:     "Nasi Goreng",
		Price:    15000,
		Category: "Makanan",
	},
	{
		Name:     "Ayam goeyng",
		Price:    15000,
		Category: "Makanan",
	},
	{
		Name:     "Es Kentang",
		Price:    15000,
		Category: "Minuman",
	},
	{
		Name:     "Kentang",
		Price:    15000,
		Category: "Snack",
	},
}
