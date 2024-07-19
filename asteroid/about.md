**GnAsteroid** is a small project based on gno.land allowing anyone to create their own asteroid(s).

Technically, an asteroid is an alternative frontend to https://gno.land.

An asteroid: 

* can be a blog, a personal or official page, a website,
* uses markdown (the markdown can embed html for things like `<form>` and `<map>`),
* can directly render Realms (GNO smart contracts) from https://gno.land,
* can have routes to/from other asteroids.

<!--------------------------------------
# wait, why not just r/demo/microblog

A blog on a blockchain, rendered by a Realm, is different from a set of editable markdown documents which can render Realms.

Since Asteroids can render Realms, they can always embed blogs like [r/demo/microblog](https://gno.land/r/demo/microblog), but the opposite is not true at the moment.
----------------------------------------->

# interdependent independance

Asteroids are independent: 

* **editorially** :: they're a bunch of markdown files, 
* **stylistically** :: each Asteroid can use whichever style it wants,
* **physically** :: not served by the webserver at gno.land, implying they have their own workload.

Asteroids all rely on gno.land for the Realms. They can also have routes to other asteroids (TODO link to doc).

# inspiration

The world wide web, even pre-HTML, really started connecting documents of computer scientists, researchers or students through hyperlinks. It was creating lines between islands, too remote to be discerned.  A bit later, personal and corporate webpages became a thing, surviving until now as blogs.

# dependence/independence 

If GNO is to emerge from the busy noise of the coinverse, its editorial content (everything non-code related) needs to find its **balance** on the dependence/independence continuum. 

Here are 2 extremes of this continuum:

1. All the content is on gno.land (i.e. everything on-chain):
   * same colorscheme,
   * same header/footer,
   * inability to work with offchain documents,
   * => dampened creativity.
2. dApps, SPA (single-page application) with connection to gno.land:
   * lack of simplicity (lot of work),
   * lack of standards,
   * => high barrier of entry
   * => missed opportunity of content creation.

There remains in the beginning some unaddressed space in the middle. One particular point of balance can initially be asteroids, which may give shape to loose communities of:

* enthusiasts, 
* programmers, 
* bloggers, 
* influencers,
* teams. 

While Asteroids don't amount to much technologically, they're easy and fast to work with.
