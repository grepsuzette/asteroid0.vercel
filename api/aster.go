package handler

import (
	"net/http"

	"github.com/gnolang/gno/gno.land/pkg/gnoweb"
	"github.com/grepsuzette/gnAsteroid"
)

func Aster(w http.ResponseWriter, r *http.Request) {
	styleDir := "" // "../asteroids-styles/gnosmos"
	asteroidDir := "../asteroid"
	gnAsteroid.HandleAsteroid(
		styleDir,
		asteroidDir,
		"gnAsteroid: Asteroid starting-kit",
		gnoweb.Config{},
	).ServeHTTP(w, r)
}
