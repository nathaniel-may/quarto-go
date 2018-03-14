# quarto-go

[Quarto](https://en.wikipedia.org/wiki/Quarto_(board_game)) is a game with simple rules yet strategic game play. This code will implement a playable version online for personal, non-commercial use.

Developement is currently paused until the mongo driver implements the 3.6 wire protocols [GODRIVER-263](https://jira.mongodb.org/browse/GODRIVER-263)

## Stack

GUI - TBD  
AWS API Gateway  
AWS Lambda Functions - Go  
MongoDB Go Driver (alpha)  
MongoDB Atlas v3.6

## Testing

```
cd $GOPATH/src/quarto-go/
go test ./...
```