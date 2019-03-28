# Right now this will just be in the api folder but the parts will be getting distributed to the correct sections

# Look First Drive Later API

## Auth

The API uses authentication on some of the routes to protect deleting and modifying events.

We use jwt bear tokens. To login use the login route with user name and a password as the parameters this will return a token that will be passed into the headers of authenticated requests. These tokens last for 12hr but standard practice states that every half that time we should request a refresh of the token. This can be done by using the refresh_token route. If the token is invalid then authentication will fail and the route will be blocked.

## Routes

### Login route

The query 

```POST https://{{server}}/api/login?username=admin&password=admin```

Would produce on a successful request

``` { "code": 200,  "expire": "2019-03-26T17:26:29Z", "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTM2MjExODksImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU1MzYxNzU4OX0.DWw8IOkenoVnEUOigte64oiirCjkI9lbxkDx2d1py3E"} ```

Would produce on a failed request 

``` { "code": 401, "message": "incorrect Username or Password" } ```

### Refresh_token 

The query

``` GET https://{{server}}/api/auth/refresh_token```

Would produce on a successful request

``` { "code": 200, "expire": "2019-03-26T17:35:54Z", "token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJleHAiOjE1NTM2MjE3NTQsImlkIjoiYWRtW4iLCJvcmlnX2lhdCI6MTU1MzYxODE1NH0EWb2p6jjgYWCZBQ8CoBz0kbmcLeMCz5pkFUq80dgoCs"} ```

On a failed request 

``` { "code": 401, "message": "invalid character '\\b' in string literal"} ```

### Hello

The query

``` GET https://{{server}}/api/auth/hello ```

Would produce on a successful request

``` { "text": "Hello World.", "userID": "admin", "userName": "admin"} ```

on a failed requests 

``` { "code": 401, "message": "invalid character '\\b' in string literal"} ```

### Ping

The query ``` GET https://{{server}}/api/ping ``` 

Would produce on a successful request 

```"message": "pong"```