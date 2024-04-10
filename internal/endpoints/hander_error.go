package endpoints

import (
	"errors"
	"net/http"

	"go-mail-sender/internal/internal_errors"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (
	interface{},
	int,
	error,
)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			obj, status, err := endpointFunc(w, r)
			if err != nil {
				if errors.Is(err, internal_errors.ErrInternal) {
					render.Status(r, http.StatusInternalServerError)
				} else {
					render.Status(r, http.StatusBadRequest)
				}

				errorsReponse := map[string]string{"error": err.Error()}

				render.JSON(w, r, errorsReponse)
				return
			}
			render.Status(r, status)
			if obj != nil {
				render.JSON(w, r, obj)
			}
		},
	)
}
