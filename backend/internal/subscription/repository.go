package subscription

type Repository interface {
	SubscribeUser(id int64) error
}
