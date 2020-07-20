# gasync
Set of utilities to handle concurrency tasks in Go. Inspired by 'yarn add async'.

## Parallel:
```go
import (
    "github.com/eduardbcom/gasync/parallel"
)

ctx := context.TODO()

res, err := parallel.Do(
    // both functions runs in parallel
    func() (interface{}, error) { return request(ctx, "task1") },
    func() (interface{}, error) { return request(ctx, "task2") },
)
if err != nil {
    return nil, err
}

res == ["task1", "task2"]
```

## Motivation:
At some point of my journey with Go I realized that I want to hide all concurrency mess under pleasurable interface.
As long as I worked with JS for a while, the first idea was "I need something like the 'async' module. But in Go.".
