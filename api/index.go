package handler

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gnolang/gno/gno.land/pkg/gnoweb"
	"github.com/grepsuzette/gnAsteroid"
)

//go:embed gnosmos
var embedFS embed.FS
var app http.Handler

const embedRelName = "gnosmos" // to "chroot". Whatever name was after your go:embed directive.
const asteroidName = "gnAsteroid"

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
