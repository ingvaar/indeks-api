# Indeks API

> A simple API to store links to articles with tags on top.

## Purpose

This API aims to only do one thing but do it right: store links, anonymously in a first time, with tags related to them and serve them back if needed.

This is also a Go playground for me to experiment and improve my skills.

The app is built around MongoDB for the database and Go-Chi for the routing part.

I want this project to be idiomatic and straightforward, at least for now.

## Philosophy

With this repository, I also aim to promote:

- the use of [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/), [semantic versioning](https://semver.org/) and [semantic releasing](https://github.com/cycjimmy/semantic-release-action) which automates the whole package release workflow including: determining the next version number, generating the release notes, and publishing the artifacts (project tarball, docker images, etc.)
- unit testing
- linting via [golangci-lint](https://github.com/golangci/golangci-lint)
- a uniformed way of building the project for several platforms via a Makefile using a docker image

> Makefile could possibly be replaced by Magefile in a near futur.
