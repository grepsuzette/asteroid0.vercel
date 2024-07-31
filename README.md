This template shows a deployment of [grepsuzette/gnosmos](https://grepsuzette/gnosmos) to [Vercel][https://vercel.com].

## Prerequisite

* `npm install vercel`

## How to deploy this example

* `git clone --recurse-submodules https://github.com/grepsuzette/gnosmos.vercel`
* go in the repo
* type `vercel`
* the deployment status will appear in the console.

## How to create your own deployments

Your asteroid need to be in `api/gnosmos`.

You may change this last part, but it needs 
to be a directory in `api/`. This is because 
we need the asteroid to be `go:embed`-ed.

go ahead and replace `api/gnosmos` by your asteroid,
it can be a submodule if you want or a local symlink.

<!--TODO-------------------------
For example:
rm -fr gnosmos
git submodule add git@github.com:grepsuzette/gnosmos.git
----TODO-------------------------->

Also change `const asteroidName = "gnAsteroid"`
in `api/index.go` to the name you want to use.

**Note**: If someone knows a better solution (ideally, 
to not need to embed the asteroid on Vercel, but 
also to simplify the process), please let me know by
opening an issue.

## Credits

Originally learnt from [riccardogiorato/template-go-vercel](https://github.com/riccardogiorato/template-go-vercel)
