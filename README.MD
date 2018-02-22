# Manparse

> converts manpages to Program Interface Model

[![GoDoc](https://godoc.org/github.com/cmdse/manparse?status.svg)](https://godoc.org/github.com/cmdse/manparse)
[![Build Status](https://travis-ci.org/cmdse/manparse.svg?branch=master)](https://travis-ci.org/cmdse/manparse)
[![Go Report Card](https://goreportcard.com/badge/github.com/cmdse/manparse)](https://goreportcard.com/report/github.com/cmdse/manparse)
[![Maintainability](https://api.codeclimate.com/v1/badges/b07e1f08a9b118679e17/maintainability)](https://codeclimate.com/github/cmdse/manparse/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b07e1f08a9b118679e17/test_coverage)](https://codeclimate.com/github/cmdse/manparse/test_coverage)

`manparse` is a tool to extract a program interface model from man-pages. It uses doclifter from Eric S. Raymond which converts man pages to DocBook xml files. Those files already hold a good level of semantics, and `manparse` will do its best to grab as much information as possible. The program interface model should be serialized to be consumed by other tools. Serialization format has not yet been chosen but it will likely be JSON + JSON Schema.

See documentation in [cmdse.github.io](https://cmdse.github.io/pages/manparse.html).

Also, visit [edge cases list reference](/EDGE_CASES.MD) to see how `manparse` handle ambiguous scenario.
