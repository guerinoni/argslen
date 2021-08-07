# argslen linter

[![CI](https://github.com/guerinoni/argslen/actions/workflows/CI.yml/badge.svg)](https://github.com/guerinoni/argslen/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/guerinoni/argslen/branch/main/graph/badge.svg?token=N5M67KW7KV)](https://codecov.io/gh/guerinoni/argslen)

Argslen is a linter that checks for long list of argument in functions.

The default limit is 5 (`maxArguments`) and skip the test files (`skipTests`), but you can configure these.

## Usage

```zsh
argslen ./...                 (scan all pkg)
argslen .                     (scan current pkg)
argslen -maxArguments=2 .     (scan currente with max args 2)
argslen -skipTests=true ./... (scan all pkg skipping tests files)
```