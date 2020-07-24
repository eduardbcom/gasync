# gasync 
[![Build Status](https://travis-ci.org/eduardbcom/gasync.svg?branch=master)](https://travis-ci.org/eduardbcom/gasync)
[![Coverage Status](https://coveralls.io/repos/github/eduardbcom/gasync/badge.svg?branch=master)](https://coveralls.io/github/eduardbcom/gasync?branch=master)

Set of utilities to handle concurrency tasks in Go.

Inspired by `yarn add async`. :speak_no_evil: :fire: :ok_hand:

# control flow
- [parallel](#parallel)
- [series](#series)
- [tryEach](#tryEach)

## parallel:
```go
import (
    "github.com/eduardbcom/gasync/parallel"
)

ctx := context.TODO()

res, err := parallel.Do(
    // both functions run in parallel
    func() (interface{}, error) { return request(ctx, "task1") },
    func() (interface{}, error) { return request(ctx, "task2") },
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
    func() (interface{}, error) { return request(ctx, "task1") },
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
    func() (interface{}, error) { return request(ctx, "task1") },
    func() (interface{}, error) { return request(ctx, "task2") },
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
    func() (interface{}, error) { return request(ctx, "task1") },
    func() (interface{}, error) { return request(ctx, "task2") },
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
    func() (interface{}, error) { return request(ctx, "task1") },
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
    func() (interface{}, error) { return request(ctx, "task1") },
    func() (interface{}, error) { return request(ctx, "task2") },
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
    func() (interface{}, error) { return request(ctx, "task2") },
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

## Motivation:
At some point of my journey with Go I realized that I want to hide all concurrency mess under pleasurable interface.
As long as I worked with JS for a while, the first idea was "I need something like the 'async' module. But in Go.".

## TODO:
- times
- timesLimit
- timesSeries
- retry
- fix TODOs
