This is a short demonstration of how to deploy an gno asteroid to [Vercel](https://vercel.com):

After cloning this repository (with `git clone --recurse-submodule https://github.com/gnAsteroid/asteroid0.vercel`) and `npm install vercel`, go in its directory and type `vercel`.

The deployment status will appear in the console.

## Own deployments

Place your asteroid as a subdirectory of the `api` folder.

It can be named as you want, but it needs to be a directory in `api/`. 

There seems to be currently no work around 
this is because the asteroid is to be embedded with `go:embed`.

## Credits

[riccardogiorato/template-go-vercel](https://github.com/riccardogiorato/template-go-vercel)
