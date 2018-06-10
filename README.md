# PZP26

PZP26 is a query language (QL).

PZP26 is a code name for this project, the name will be set on the release of V1.0.0.

Old status : [none] > current : write tests (TDD) > next : [write code | alpha | beta | release V1]

## Standalone mode

> **"I don't like Go(lang), my app is coded in my favorite language which is not Go, I can't use it !"**
>
> This library will have an standalone binary to use it outside your app (as UNIX philisophy).
>
> Deployment will be like you are installing and using MariaDB, CouchDB, ... but to mount your API.
>
> After deployment, you have to code between this API and your database (if any).
>
> We are considering to use it with an RPC bridge, and/or using files in order to talk with your code.

---

> **"I'm developping in Go(lang), it is not a problem for me !"**
>
> So, you don't need the standalone binary, you can use it as library in your Go app !

## Moving nested objects to root response

The following query :

```
(id='1') {id, name, city{id, name}}
```

Will respond :

```JSON
[
    "root[0]" {
        "id": "1",
        "name": "Smith",
        "city": {
            "id": "JKB21",
            "name": "New York"
        }
    }
]
```

(if you use JSON parser for response, but you can change it easily)

**BUT** the following query :

```
(id='1') {id, name, @city{id, name}}
```

Will responds :

```JSON
[
    "root[0]" {
        "id": "1",
        "name": "Smith",
        "city": "@JKB21"
    },
    "JKB21" {
        "id": "JKB21",
        "name": "New York"
    }
]
```

The main goal is to help you to update the data in the store on the frontend.

## Advanced selectors

The following query :

```
(NOT (id = 1 OR (id IN (4, 5, 7)))) {id}
```

Will respond :

```JSON
[
    "root[0]" [
        {"id": 2},
        {"id": 3},
        {"id": 6},
        {"id": 8},
        // (...)
    ]
]
```

## Recursive objects

```
(id='1') {id, name, @city{id, name, @nearFrom{$city_schema}}=$city_schema}
```

This query should return every city which are "linked" with `nearFrom` attribute, but on root response :

```JSON
[
    "root[0]" {
        "id": "1",
        "name": "Smith",
        "city": "@JKB21"
    },
    "JKB21" {
        "id": "JKB21",
        "name": "New York",
        "nearFrom": ["@MHB74","@PND84"]
    },
    "MHB74" {
        "id": "MHB74",
        "name": "Hempstead",
        "nearFrom": ["@PLO22", "@ZBD88"]
    },
    "PND84" {
        "id": "PND84",
        "name": "Clifton",
        "nearFrom": ["@GSI54", "@ZBD88"]
    },
    // "PLO22" {...},
    // "GSI54" {...},
    "ZBD88" {
        "id": "ZBD88",
        "name": "Stamford",
        "nearFrom": []
    },
]
```

> Usualy, the recursivity should be an security issue (buffer overflow, database overload, etc ...).
>
> But the resolver does not ask again for a `city` if it is already in other object in query (and we don't request another new field from it).
>
> In our example, look at the ID `ZBD88`, which is in `nearFrom` field of both `MHB74` and `PND84` ID.
>
> To prevent no-end recursivity, you have to specify to the library how much "recursivity levels" you want to limit it.

## No need to declare data

Just send data which is available to this API, everything should be checked on frontend.

## Middlewares

You can hack the behaviour of the QL with middleware (WIP).
