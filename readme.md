# TODOs REST API

This server has four methods to interact with TODOs. TODOs are stored in memmory in a slice. See code comments for why I chose a slice. (A new version may include a map)...

## Methods

| Method  | Action  | Response|
| ------- | ------  |      ---| 
| POST    | Create  | 200     |
| GET     | Read    | 200     |
| PUT     | Update  | 200/400 |
| DELETE  | Delete  | 200/400 |

## Model
```{javascript}
{
  "id": int,      // Time created in ms
  "done": bool,   // Is the TODO marked as done?
  "text": string  // What is there TODO?
}
```

## Running
1. Install go
2. Get some gin: `go get github.com/gin-gonic/gin`
3. Run main: `go run main` OR `go build && ./go-todo-api`
