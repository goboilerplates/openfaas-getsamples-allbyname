# Boilerplate for GetSamplesAllByName OpenFaas Function in Go
[![Build Status](https://travis-ci.org/goboilerplates/openfaas-getsamples-allbyname.svg?branch=master)](https://travis-ci.org/goboilerplates/openfaas-getsamples-allbyname)
[![codecov](https://codecov.io/gh/goboilerplates/openfaas-getsamples-allbyname/branch/master/graph/badge.svg)](https://codecov.io/gh/goboilerplates/openfaas-getsamples-allbyname)
[![Go Report Card](https://goreportcard.com/badge/github.com/goboilerplates/openfaas-getsamples-allbyname)](https://goreportcard.com/report/github.com/goboilerplates/openfaas-getsamples-allbyname)
[![GoDoc](https://godoc.org/github.com/goboilerplates/openfaas-getsamples-allbyname/function?status.svg)](https://godoc.org/github.com/goboilerplates/openfaas-getsamples-allbyname/function)
[![](https://images.microbadger.com/badges/image/goboilerplates/openfaas-getsamples-allbyname.svg)](https://microbadger.com/images/goboilerplates/openfaas-getsamples-allbyname)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/goboilerplates/openfaas-getsamples-allbyname/blob/master/LICENSE)

## Features
- OpenFaas Function
- CI with Travis
- Docker Build
- Docker Compose

## Installation

Get the openfaas-getsamples-allbyname repository

```
go get github.com/goboilerplates/openfaas-getsamples-allbyname

cd echo $GOPATH/src/github.com/goboilerplates/openfaas-getsamples-allbyname
```
## Running the tests

Run all tests

```
go test ./test/...
```

Or run all tests with coverage

```
bash script/coverage.sh
```

## Docker support 

Deploy OpenFaas with Docker Swarm

```
docker swarm init

bash script/deploy_stack.sh
```

Build function docker image

```
bash script/Dockerbuild.sh
```

Deploy function to OpenFaas

```
bash script/deploy.sh
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/goboilerplates/openfaas-getsamples-allbyname/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details