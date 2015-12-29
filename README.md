Todo Backend
----

- [Overview](#overview)
- [Routes](#routes)
- [Usage](#usage)
  - [Starting the Server](#starting-the-server)
  - [Formatting](#formatting)
  - [Fetching Data](#fetching-data)
  - [Sending Data](#sending-data)
- [Persistence](#persistence)

## Overview

This is a very simple todo backen written in Go using [jsonapi][jsonapi] and
[httprouter][httprouter].

Todos are represented in code as follows:

```go
type Todo struct {                                                                 
  Id    int
  Title string
  Body  string
  Done  bool
}
```

## Routes
The app supports a single route at the moment, `/todos`. One can either `GET`
from this endpoint or `POST` to it.

## Usage

### Starting the server
Once compiled, you need only have the `todo` binary on the system. To start the
server, navigate the the directory where the compiled binary lives and run the
following commands.

```bash
$ cd /directory/where/todo/lives
$ ./todo -port 3000
```

The server will now be listening on port 3000 until you hit CTRL+C to terminate
terminate it. You may specify any port number you wish. Conventionally, live web
services run on port 80. The default is 8080.

### Formatting
JSON data should be formatted according the the [JSON API Specification][jsonspec].

### Fetching Data
A `GET` request will return all existing todos. The data will look like this:

```json
{
  "data": [
    {
      "type": "todos",
      "id": "0",
      "attributes": {
        "body": "This is my first todo",
        "done": false,
        "title": "A Title"
      }
    },
    {
      "type": "todos",
      "id": "1",
      "attributes": {
        "body": "Some Text",
        "done": false,
        "title": "Another Title"
      }
    },
    {
      "type": "todos",
      "id": "2",
      "attributes": {
        "body": "This todo is already done",
        "done": true,
        "title": "A Completed Todo"
      }
    }
  ]
}
```

### Sending Data
A `POST` request should have its data formatted according to the JSON API spec
for [creating][jsoncreation] a new resource. That may look something like this:

```json
{
  "data": {
    "type": "todos",
    "attributes": {
        "title": "A Title",
        "body": "Some Text",
        "done": true
    }
  }
}
```

Notice that no `id` is specified as this will be assigned and then returned by
the server.

## Persistence
Todos are persisted using [BoltDB][boltdb]. When the app is started, a file
named `todo.db` will be created. To erase all todos and start from scratch,
simply remove this file with `rm todo.db`.

[jsonapi]: https://github.com/shwoodard/jsonapi
[httprouter]: https://github.com/julienschmidt/httprouter
[jsonspec]: http://jsonapi.org/
[jsoncreation]: http://jsonapi.org/format/#crud-creating
[boltdb]: https://github.com/boltdb/bolt
