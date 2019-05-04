# [ggm] - Graph of Go Modules

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/ggm.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/ggm)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/ggm/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/ggm.svg)](https://github.com/spiegel-im-spiegel/ggm/releases/latest)

## Download and Install

```
$ go get github.com/spiegel-im-spiegel/ggm@latest
```

## Usage

```
$ ggm -h
Usage:
  ggm [flags] [input file]

Flags:
  -c, --config string   Configuration file
      --debug           Debug flag
  -h, --help            help for ggm
  -v, --version         Output version of ggm

$ go mod graph | ggm  -c config/sample.toml | dot -Tpng -o ggm.png
```

![ggm](./ggm.png)

[ggm]: https://github.com/spiegel-im-spiegel/ggm "spiegel-im-spiegel/ggm: Graph of Go Modules"
