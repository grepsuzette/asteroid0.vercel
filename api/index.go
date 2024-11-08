package handler

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gnAsteroid/gno/gno.land/pkg/gnoweb"
	"github.com/gnAsteroid/gnAsteroid"
)

// Below embed name should be the same as the subdirectory
//go:embed asteroid
var embedFS embed.FS
var app http.Handler

const embedRelName = "asteroid" // this must be the same as the name in go:embed above
const asteroidName = "gnAsteroid" // change this to your asteroid name

func init() {
	if asteroid, e := fs.Sub(embedFS, embedRelName); e != nil {
		panic(e.Error())
	} else {
		app = gnAsteroid.HandleAsteroid(
			asteroid,
			gnAsteroid.DefaultStyle(),
			asteroidName,
			gnoweb.Config{
				RemoteAddr:  "gno.land:26657",
				HelpChainID: "portal-loop",
				HelpRemote:  "gno.land:26657",
			},
		)

	}
}

// func needs to be exported name otherwise does not matter
func Abcdef(w http.ResponseWriter, r *http.Request) { app.ServeHTTP(w, r) }
