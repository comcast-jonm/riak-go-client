你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Riak Go Client

The **Riak Go Client** is a client which makes it easy to communicate with [Riak](http://basho.com/riak/), an open source, distributed database that focuses on high availability, horizontal scalability, and *predictable* latency. Both Riak and this code is maintained by [Basho](http://www.basho.com/).

The [latest version](https://github.com/basho/riak-go-client/releases/latest) of the client supports both Riak KV 2.0+, and Riak TS 1.0+. 

## Build Status

[![Build Status](https://travis-ci.org/basho/riak-go-client.svg?branch=master)](https://travis-ci.org/basho/riak-go-client)

# Installation

`go get github.com/basho/riak-go-client`

# Documentation

* [API documentation on Godoc](https://godoc.org/github.com/basho/riak-go-client)
* [Wiki](https://github.com/basho/riak-go-client/wiki)
* [Release Notes](https://github.com/basho/riak-go-client/blob/master/RELNOTES.md). 

# Testing / Contributing

This repository's maintainers are engineers at Basho and we welcome your contribution to the project! Review the details in [CONTRIBUTING.md](CONTRIBUTING.md) in order to give back to this project.

*Note:* Please clone this repository in such a manner that submodules are also cloned:

```
git clone --recursive https://github.com/basho/riak-go-client
```

OR:

```
git clone https://github.com/basho/riak-go-client
git submodule init --update
```

## Unit Tests

```sh
make unit-test
```

## Integration Tests

You have two options to run Riak locally - either build from source, or use a pre-installed Riak package.

### Source

To setup the default test configuration, build a Riak node from a clone of `github.com/basho/riak`:

```sh
# check out latest release tag
git checkout riak-2.1.4
make locked-deps
make rel
```

[Source build documentation](http://docs.basho.com/riak/kv/latest/setup/installing/source/).

When building from source, the protocol buffers port will be `8087` and HTTP will be `8098`.

### Package

Install using your platform's package manager ([docs](http://docs.basho.com/riak/kv/latest/setup/installing/))

When installing from a package, the protocol buffers port will be `8087` and HTTP will be `8098`.

### Running Integration Tests

* Ensure you've initialized this repo's submodules:

```sh
git submodule update --init
```

* Run the following:

```sh
./tools/setup-riak
make integration-test
```

This repository's maintainers are engineers at Basho and we welcome your contribution to the project! Review the details in [CONTRIBUTING.md](CONTRIBUTING.md) in order to give back to this project.

### An honest disclaimer

Due to our obsession with stability and our rich ecosystem of users, community updates on this repo may take a little longer to review. 

The most helpful way to contribute is by reporting your experience through issues. Issues may not be updated while we review internally, but they're still incredibly appreciated.

Thank you for being part of the community! We love you for it. 

## Roadmap

* 1.0.0 - Full Riak 2 support with command queuing and retries.

## License

The **Riak Go Client** is Open Source software released under the Apache 2.0
License. Please see the [LICENSE](LICENSE) file for full license details.

These excellent community projects inspired this client and parts of their code
are in `riak-go-client` as well:

* [`goriakpbc`](https://github.com/tpjg/goriakpbc)
* [`riaken-core`](https://github.com/riaken/riaken-core)
* [`backoff`](https://github.com/jpillora/backoff)

## Authors

* [Luke Bakken](https://github.com/lukebakken)
* [Christopher Mancini](https://github.com/christophermancini)

## Contributors

Thank you to all of our contributors!

* [Ian Lozinski](https://github.com/i)
* [Sergio C. Arteaga](https://github.com/tegioz)
* [Andrew Zeneski](https://github.com/andrewzeneski)
* [Кирилл Александрович Журавлев](https://github.com/kazhuravlev)
* [Paul Guelpa](https://github.com/pguelpa)
* [Xabier Larrakoetxea Gallego](https://github.com/slok)
* [Paul Maseberg](https://github.com/pmaseberg)
* [Weerasak Chongnguluam](https://github.com/iporsut)
* [Hawk Newton](https://github.com/hawknewton)
* [Beau Brewer](https://github.com/beaubrewer)
