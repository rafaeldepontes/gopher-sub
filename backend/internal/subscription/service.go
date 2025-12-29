package subscription

type Service interface {
	Subscribe(id int64) error
}
