package createuser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/phuoc-le/go-restapi/cmd/api/models"
	"github.com/phuoc-le/go-restapi/pkg/application"
	"github.com/phuoc-le/go-restapi/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func createUser(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		user := &models.User{}
		json.NewDecoder(r.Body).Decode(user)

		if err := user.Create(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(user)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	return middleware.Chain(createUser(app), middleware.LogRequest)
}
