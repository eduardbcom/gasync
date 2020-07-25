# gasync 
[![Build Status](https://travis-ci.org/eduardbcom/gasync.svg?branch=master)](https://travis-ci.org/eduardbcom/gasync)
[![Coverage Status](https://coveralls.io/repos/github/eduardbcom/gasync/badge.svg?branch=master)](https://coveralls.io/github/eduardbcom/gasync?branch=master)

Set of utilities to handle concurrency tasks in Go.

Inspired by `yarn add async`. :speak_no_evil: :fire: :ok_hand:

# control flow
- [parallel](#parallel)
- [series](#series)
- [tryEach](#tryEach)
- [times](#times)
- [retry](#retry)

## parallel:
```go
import (
    "github.com/eduardbcom/gasync/parallel"
)

ctx := context.TODO()

res, err := parallel.Do(
    // both functions run in parallel
    func() (interface{}, error) { return identity(ctx, "task1") },
    func() (interface{}, error) { return identity(ctx, "task2") },
)

// res == []interface{}{"task1", "task2"}
// err == nil
```

```go
import (
    "github.com/eduardbcom/gasync/parallel"
)

ctx := context.TODO()

res, err := parallel.Do(
    // both functions run in parallel
    func() (interface{}, error) { return identity(ctx, "task1") },
    func() (interface{}, error) { return nil, errors.New("some error here") },
)

// res == nil
// err == errors.New("some error here")
```

```go
import (
    "github.com/eduardbcom/gasync/parallel"
)

ctx := context.TODO()

res, err := parallel.DoWithLimit(
    2,
    // both functions run in parallel
    func() (interface{}, error) { return identity(ctx, "task1") },
    func() (interface{}, error) { return identity(ctx, "task2") },
)

// res == []interface{}{"task1", "task2"}
// err == nil
```

## series:
```go
import (
    "github.com/eduardbcom/gasync/series"
)

ctx := context.TODO()

res, err := series.Do(
    // functions run sequentially
    func() (interface{}, error) { return identity(ctx, "task1") },
    func() (interface{}, error) { return identity(ctx, "task2") },
)

// res == []interface{}{"task1", "task2"}
// err == nil
```

```go
import (
    "github.com/eduardbcom/gasync/series"
)

ctx := context.TODO()

res, err := series.Do(
    // functions run sequentially
    func() (interface{}, error) { return identity(ctx, "task1") },
    func() (interface{}, error) { return nil, errors.New("some error here") },
)

// res == nil
// err == errors.New("some error here")

```

## tryEach:
```go
import (
    "github.com/eduardbcom/gasync/tryEach"
)

ctx := context.TODO()

res, err := tryEach.Do(
    // functions run sequentially
    func() (interface{}, error) { return identity(ctx, "task1") },
    func() (interface{}, error) { return identity(ctx, "task2") },
)

// res == "task1"
// err == nil
```

```go
import (
    "github.com/eduardbcom/gasync/tryEach"
)

ctx := context.TODO()

res, err := tryEach.Do(
    // functions run sequentially
    func() (interface{}, error) { return nil, errors.New("some error here") },
    func() (interface{}, error) { return identity(ctx, "task2") },
)

// res == "task2"
// err == nil
```

```go
import (
    "github.com/eduardbcom/gasync/tryEach"
)

ctx := context.TODO()

res, err := tryEach.Do(
    // functions run sequentially
    func() (interface{}, error) { return nil, errors.New("some error here") },
    func() (interface{}, error) { return nil, errors.New("some second error here") },
)

// res == nil
// err == errors.New("some second error here")
```

## times:
```go
import (
    "github.com/eduardbcom/gasync/times"
)

ctx := context.TODO()

res, err := times.Do(
    2,
    // both calls run in parallel
    func() (interface{}, error) { return identity(ctx, "task") },
)

// res == []interface{}{"task", "task"}
// err == nil
```

```go
import (
    "github.com/eduardbcom/gasync/times"
)

ctx := context.TODO()

res, err := times.DoWithLimit(
    2,
    5,
    // 2 calls run in parallel
    // 2 calls run in parallel
    // 1 call
    func() (interface{}, error) { return identity(ctx, "task") },
)

// res == []interface{}{"task", "task", "task", "task", "task"}
// err == nil
```

```go
import (
    "github.com/eduardbcom/gasync/times"
)

ctx := context.TODO()

res, err := times.DoSeries(
    2,
    // functions run in sequential way
    func() (interface{}, error) { return identity(ctx, "task") }
)

// res == []interface{}{"task", "task"}
// err == nil
```

## retry:
```go
import (
    "github.com/eduardbcom/gasync/retry"
)

ctx := context.TODO()

res, err := retry.Do(
    2,
    func() (interface{}, error) { return identity(ctx, "task") },
)

// res == "task1"
// err == nil
```

```go
import (
    "github.com/eduardbcom/gasync/retry"
)

ctx := context.TODO()

res, err := retry.Do(
    2,
    func() (interface{}, error) { return nil, errors.New("some error here") },
)

// res == nil
// err == errors.New("some error here")
```

```go
import (
    "github.com/eduardbcom/gasync/retry"
)

ctx := context.TODO()

intervalMs := 1000
res, err := retry.DoWithInterval(
    2,
    intervalMs,
    func() (interface{}, error) { return nil, errors.New("some error here") }
)

// res == nil
// err == errors.New("some error here")
```

## Motivation:
At some point of my journey with Go I realized that I want to hide all concurrency mess under pleasurable interface.
As long as I worked with JS for a while, the first idea was "I need something like the 'async' module. But in Go.".

## TODO:
- retry
- fix TODOs
