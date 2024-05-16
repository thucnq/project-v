package paginator

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

func flip(order Order) Order {
	if order == ASC {
		return DESC
	}
	return ASC
}
