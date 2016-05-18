# Thot
## Description
Thot is a web service that allows you to connect many services through a single gateway, doing something like an API Gateway. `Thot` basically works through subscriptions. Subscribers provide certain data (see [Section Subscribe]) which serve to `Thot` to know where to redirect traffic and the incoming requests.

## Subscriber
The information provided by subscribers should be as follows:
```json
{
  "endpoint": "service endpoint",
  "method": "request method",
  "name": "name of service",
  "url": "url, dns, etc ..."
}
```
And it should be posted to `/POST/subscribe`, After signing up you can start making request to `Thot` as follows.

### Example
Suppose we subscribe a service called `Example` with an endpoint `hello` who receives information via `POST` to `http://localhost:8000`:

The information that we send has the following structure:
```json
{
  "endpoint": "hello",
  "method": "POST",
  "name": "Example",
  "url": "http://localhost:8000"
}
```

```sh
curl -X POST -H "Content-Type: application/json" \
  -d '{"endpoint": "hello","method": "POST","name": "Example","url": "http://localhost:8000"}' "http://thot-domain/subscribe"
```

This makes `Thot` internally create the following endpoint:

`POST /example/hello`

And with that, for a request to the entry point Example, `hello`, just enough to make a request to` `POST http://thot-domain/example/hello`.

eg:
```sh
curl -X POST -d {'foo': 'var'} http://thot-domain/example/hello
```

This will cause all traffic sent to `http://thot-domain/example/hello` be redirected to `http://localhost:8000/hello`


Too see a working example execute:
```sh
$ go run main.go
```

And in another terminal window run:
```sh
$ go run server_example/server.go
```

Then, you can send POST to Thot as follow:

First, subscribe:
```sh
curl -X POST -H "Content-Type: application/json" \
  -d '{"endpoint": "hello","method": "POST","name": "Example","url": "http://localhost:8000"}' "http://localhost:8080/subscribe"
```

Then, send POST:
```sh
curl -X POST -d {'name': 'rod'} http://localhost:8080/example/hello
```

and this should be response with:
```sh
$ Hello rod.
```
Rf.
