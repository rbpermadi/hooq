# hooq

## Description

Hooq is a repository for Hooq Test Assignment. Its a web app using to check website health. It's using **Go** as its programming language.

## Onboarding and Development Guide

### Prequisites

* [**Go (1.9.7 or later)**](https://golang.org/doc/install)

### Setup

Please install/clone the [Prequisites](#prequisites) and make sure it works perfectly on your local machine.

After the [Prequisites](#prequisites) have been installed, please clone **Hooq** project into your local machine.

```
> cd github.com/rbpermadi
> git clone git@github.com:rbpermadi/hooq.git
```

Finally, run **Hooq** in your local machines.

```
> make run
```
You go to [localhost:7171](#http://localhost:7171) from your browser to access the homepage

To kill the server you just need to hold `Ctrl + C`

## API Endpoints

### Retrieve Site List [GET /site_check]

```
GET /site_check HTTP/1.1
Host: localhost:7171
Cache-Control: no-cache

*************************************************************
Success (Status: 200 OK)
*************************************************************

{
    "data": [
        {
            "id": 1,
            "link": "http://goggle.com",
            "status": "unhealthy"
        },
        {
            "id": 3,
            "link": "http://goggle.com",
            "status": "unhealthy"
        }
    ]
}
```

### Create Site [POST /site_check/add]

```
POST /site_check/add HTTP/1.1
Host: localhost:7171
Content-Type: application/json
Cache-Control: no-cache

{
	"link": "http://goggle.com"
}

*************************************************************
Success (Status: 201 Created)
*************************************************************

{
    "data": {
        "id": 1,
        "link": "http://google.com",
        "status": "healthy"
    }
}
```

```
POST /site_check/add HTTP/1.1
Host: localhost:7171
Content-Type: application/json
Cache-Control: no-cache

{
	"link": "google"
}

*************************************************************
Error url validation (Status: 400 Bad Request)
*************************************************************

{
    "code": 400,
    "message": "parse goggle: invalid URI for request"
}
```

### Delete Site [DELETE /site_check/delete?id={id}]

```
DELETE /site_check/delete?id=1 HTTP/1.1
Host: localhost:7171
Content-Type: application/json
Cache-Control: no-cache

*************************************************************
Success (Status: 200 Created)
*************************************************************

{
    "message": "site has successfully deleted"
}
```

```
DELETE /site_check/delete?id=2 HTTP/1.1
Host: localhost:7171
Content-Type: application/json
Cache-Control: no-cache

*************************************************************
Site Not Found (Status: 404 Not Found)
*************************************************************

{
    "code": 404,
    "message": "Site not found"
}
```

## FAQ

> Not available yet
