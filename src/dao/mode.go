package dao

var Mode = map[string]string{
	"10": " order by goods_sold_amount ",
	"11": " order by goods_sold_amount desc ",
	"20": " order by goods_price ",
	"21": " order by goods_price desc ",
	"30": " order by goods_score ",
	"31": " order by goods_score desc ",
}
