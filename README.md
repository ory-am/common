# ory/common

[![Build Status](https://travis-ci.org/ory/common.svg)](https://travis-ci.org/ory/common)
[![Coverage Status](https://coveralls.io/repos/ory/common/badge.svg?branch=master&service=github)](https://coveralls.io/github/ory/common?branch=master)

A library for common tasks:

* [`env`](env/README.md)  adds defaults to `os.GetEnv()` and saves you 3 lines of code
* [`rand`](rand/README.md)  is a library based on crypto/rand to create random sequences, which are cryptographically strong.
* [`compiler`](compiler/README.md) enables you to compile regex expressions from templates like `foo{.*}bar`. Useful for URL pattern matching.


You'll find READMEs in each package directory.

This library also includes packages called `pkg` and `context`. Both are subject to frequent changes. *Don't use them.*