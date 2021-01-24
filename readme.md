# Sentiment API

This is a graphql API for the twitter sentiment project

## Developer Notes and TODOs

Some development goals

- [ ] Resolvers package
	- [x] Resolve Point
	- [ ] Resolve Points
	- [ ] Resolve Tweets
	- [ ] Resolve Point's Tweets
	- [ ] Hook up Resolvers to Schema creation
- [ ] Set up http server
- [ ] Comprehensive Caching solution
	- [ ] Probably make a cache package
	- [ ] Read and write from point cache in resolvers

Deployment goals

- [ ] Makefile and Docker multistage build
- [ ] Figure out access points
- [ ] Add all yamls to include this in active cluster
- [ ] Expose

And if this ever blows up:

- [ ] Cool image-based integration tests?
- [ ] Unit tests on packages?
