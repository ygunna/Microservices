package wrappa

import (
	"github.com/tedsuo/rata"
)

type APIWrappa struct {
}

func NewAPIWrappa(
) *APIWrappa {
	return &APIWrappa{
	}
}

func (wrappa *APIWrappa) Wrap(handlers rata.Handlers) rata.Handlers {
	wrapped := rata.Handlers{}

	for name, handler := range handlers {
		//wrapped[name] = CorsHandler{Handler: handler}
		wrapped[name] = handler
	}

	return wrapped
}
