# go-zero-config

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/go-zero-config)](https://goreportcard.com/report/github.com/go-packagist/go-zero-config)
[![tests](https://github.com/go-packagist/go-zero-config/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/go-zero-config/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/go-zero-config)](https://pkg.go.dev/github.com/go-packagist/go-zero-config)
[![codecov](https://codecov.io/gh/go-packagist/go-zero-config/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/go-zero-config)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Introduction

**go-zero-config** is a configuration package for go-zero and [go-packagist framework](https://github.com/go-packagist/go-zero-config).

## Installation

```bash
go get github.com/go-packagist/go-zero-config
```

## Usage

```go
package main

import (
    "github.com/go-packagist/framework/foundation"
    "github.com/go-packagist/go-zero-config"
    "github.com/stretchr/testify/assert"
    "testing"
)

func main() {
    app := foundation.NewApplication()

    app.Register(go_zero_config.NewProvider(app.Container, ".test/config.yaml"))

    assert.Equal(t, "test.config", app.MustMake("zero.config").(*go_zero_config.Config).Name)
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
