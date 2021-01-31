# Sentiment API

This is a graphql API for the twitter sentiment project

## Developer Notes and TODOs

Some development goals

- [x] Resolvers package
	- [x] Resolve Point
	- [x] Resolve Points
	- [x] Resolve Tweets
	- [x] Resolve Point's Tweets
	- [x] Hook up Resolvers to Schema creation
- [x] Set up http server

Deployment goals

- [x] Makefile and Docker multistage build
- [x] Figure out access points
- [x] Add all yamls to include this in active cluster
- [x] Expose

Analytics goals

- [x] Add identity token to parameters
- [x] Add sentiment database in mysql. Transition to new structure
- [ ] Add request identity log
- [ ] Update graphapi in cluster

Caching goals:

- [ ] Comprehensive Caching solution
	- [ ] Put it in client
	- [ ] Cache points
	- [ ] Cache tweets
	- [ ] Potentially cache by sql query?

And if this ever blows up:

- [ ] Cool image-based integration tests?
- [ ] Unit tests on packages?

