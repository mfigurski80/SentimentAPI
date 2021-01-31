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

- [ ] Create JWT specification
	- [ ] Update frontend to match
- [ ] Create new analytics database. Migrate
- [ ] Mutate graphql points for writing to analytics
- [ ] Update graphapi in cluster

Caching goals

- [ ] Comprehensive Caching solution
	- [ ] Probably make a cache package
	- [ ] Read and write from point cache in resolvers

And if this ever blows up:

- [ ] Cool image-based integration tests?
- [ ] Unit tests on packages?
