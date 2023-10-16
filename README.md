# STC - statistical intelligence
Coding assignment for Future memories

I decided to build the solution like a little lib, instead of just a straight-up script.

There is probably some gains to be made in performance by threading the fetching of group classes. 
But I decided not to. I could have built a pool of go routines that could do the fetching, which would dramatically increase performance while not relentlessly hammering the API.

I could probably refactor the Service a bit more to have less code repetition. The fetching and summarizing methods are a bit repetitive.

`go test ./...` to run tests

`go run main.go` to run the application
