package subscription

import "net/http"

type Controller interface {
	Subscriber(w http.ResponseWriter, r *http.Request)
}
