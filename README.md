Todo Toy
--------

- [Overview](#overview)
- [Routes](#routes)
- [Usage](#usage)
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
As of this writing, the app does not persist resources for longer than the
lifetime of the server. Once the server is shut down, any added data will be
lost. This should change in the near future.

[jsonapi]: https://github.com/shwoodard/jsonapi
[httprouter]: https://github.com/julienschmidt/httprouter
[jsonspec]: http://jsonapi.org/
[jsoncreation]: http://jsonapi.org/format/#crud-creating
