# Record Hash

Securing Data from Direct Manipulation

## Motivation

When we store data to database, commonly DBA can manipulate the data directly without application. Using this simple library we againt DBA to do that.

## Install

```bash
go get github.com/ad3n/RecordHash
```

## Usage

```go
type User struct {
    ID       int    `hash:"id"`
    Username string `hash:"username"`
    Hash     string
}

user := User{
    ID: 1,
    Username: "admin"   
}

hasher := recordhash.New(recordhash.MD5, user)
user.Hash = hasher.Hash()

user.Hash = recordhash.Hash(user)
```
