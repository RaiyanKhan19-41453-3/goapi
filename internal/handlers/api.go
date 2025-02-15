package handlers

import(
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/RaiyanKhan19-41453-3/goapi/internal/middleware"
)

func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router){
		router.use(middleware.Authorization)

		router.Get("/conins", GetCoinBalance)
	})
}
