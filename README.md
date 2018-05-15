# Phony - JSON Server

Phony is a JSON server that allows you to serve perdefined JSON responses to routes declared in a config file (`phonyconfig.json`)

This may be particularly useful when developing front-end functionality where no backend implementation exists yet. Using Phony, one can specify routes and the expected JSON response which, when requested, can be consumed by front-end code.

## Usage

For example, given the following config file:
```json
[
  {
    "path": "/phony",
    "method": "GET",
    "status": 200,
    "data": {"version": "x.x.x", "repository": "https://github.com/cgons/phony"}
  }
]
```

Run the Phony CLI:
```bash
./phony

2018/05/15 14:00:33 Loaded config file: ./phonyconfig.json
2018/05/15 14:00:33 Phony JSON Server running on: 0.0.0.0:9191
2018/05/15 14:00:33 -------------------------------------------------
```
Hit the URL defined in the config file to receive a response from Phony.
```bash
http GET localhost:9191/phony

HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Length: 66
Content-Type: application/json
Date: Tue, 15 May 2018 18:00:59 GMT

{
    "repository": "https://github.com/cgons/phony",
    "version": "0.1.0"
}
```

Phony will also log all requests made to it...
```bash
2018/05/15 14:00:37 Request: GET -> /test-route
2018/05/15 14:00:49 Request: GET -> /phony
```

When Phony receives a request for a route not defined in `phonyconfig.json`, a 404 response is returned.

## Upcoming Features
- [ ] Ability to apped a `?delay=3` query param to a request URL to mimic a slow response
- [ ] Ability to add a `delay` parameter to config to mimic a slow response
- [ ] A `/files/<pdf|xls|xlxs|txt|doc|docx>` route to serve empty files of PDF, Excel (xls, xlsx), Text and Word
- [ ] Unit tests
