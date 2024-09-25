This template shows how easily an asteroid can be deployed to [Vercel](https://vercel.com):

```
npm install vercel
git clone --recurse-submodules https://github.com/gnAsteroid/gnosmos.vercel
cd gnosmos.vercel
vercel
```

The deployment status will appear in the console.

## Own deployments

Place your asteroid as a subdirectory of the `api` folder.

It can be named as you want, but it needs 
to be a directory in `api/`. There seems currently no work around 
this is because the asteroid is to be embedded with `go:embed`.

## Credits

[riccardogiorato/template-go-vercel](https://github.com/riccardogiorato/template-go-vercel)
