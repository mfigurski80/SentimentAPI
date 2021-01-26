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

- [x] Dockerfile for supporting analytics mysql database
	- [ ] Decide on database type (nosql?)
	- [ ] Figure out how to restrict permissions (is there a way to do write only?)
- [ ] Mutate graphql points for writing to analytics
- [ ] Deploy database
- [ ] Update graphapi in cluster

And if this ever blows up:

- [ ] Cool image-based integration tests?
- [ ] Unit tests on packages?
- [ ] Comprehensive Caching solution
	- [ ] Probably make a cache package
	- [ ] Read and write from point cache in resolvers
