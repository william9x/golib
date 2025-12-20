# Golib Core

> **Note:**
> The original repository is [golibs-starter/golib](https://github.com/golibs-starter/golib). Details for usage, configurations and examples can be found there.
> This repository is a mirror copy with following changes:
> - Upgraded Go and dependencies version for better compatibility.
> - Removed unused code or dependencies.
> - Use [Fiber](https://github.com/gofiber/fiber) instead of [Gin](https://github.com/gin-gonic/gin) for HTTP server. Which means fasthttp instead of net/http is used.
> - Use [Sonic](https://github.com/bytedance/sonic) instead of encoding/json for JSON serialization/deserialization.


### Setup instruction

Both `go get` and `go mod` are supported.

```shell
go get github.com/william9x/golib-core
```

### Configuration

#### 1. Added new configurations for Fiber

```yaml
app:
  fiber:
    # When set to true, this will spawn multiple Go processes listening on the same port.
    #
    # Type: bool
    # Default: false
    prefork:

    # When set to true, the router treats "/foo" and "/foo/" as different.
    # By default, this is disabled and both "/foo" and "/foo/" will execute the same handler.
    #
    # Type: bool
    # Default: false
    strictRouting:

    # When set to true, enables case-sensitive routing.
    # E.g. "/FoO" and "/foo" are treated as different routes.
    # By default, this is disabled and both "/FoO" and "/foo" will execute the same handler.
    #
    # Type: bool
    # Default: false
    caseSensitive:

    # When set to true, this relinquishes the 0-allocation promise in certain
    # cases in order to access the handler values (e.g. request bodies) in an
    # immutable fashion so that these values are available even if you return
    # from handler.
    #
    # Type: bool
    # Default: true
    immutable:

    # When set to true, converts all encoded characters in the route back
    # before setting the path for the context, so that the routing,
    # the returning of the current url from the context `ctx.Path()`
    # and the parameters `ctx.Params(%key%)` with decoded characters will work
    #
    # Default: false
    unescapePath:

    # Max body size that the server accepts.
    # -1 will decline any body size
    #
    # Type: int
    # Default: 4 * 1024 * 1024
    bodyLimit:

    # Maximum number of concurrent connections.
    #
    # Type: int
    # Default: 256 * 1024
    concurrency:

    # The amount of time allowed to read the full request including body.
    # It is reset after the request handler has returned.
    # The connection's read deadline is reset when the connection opens.
    #
    # Type: time.Duration
    # Default: unlimited
    readTimeout:

    # The maximum duration before timing out writes of the response.
    # It is reset after the request handler has returned.
    #
    # Type: time.Duration
    # Default: unlimited
    writeTimeout:

    # The maximum amount of time to wait for the next request when keep-alive is enabled.
    # If IdleTimeout is zero, the value of ReadTimeout is used.
    #
    # Type: time.Duration
    # Default: unlimited
    idleTimeout:

    # Per-connection buffer size for requests' reading.
    # This also limits the maximum header size.
    # Increase this buffer if your clients send multi-KB RequestURIs
    # and/or multi-KB headers (for example, BIG cookies).
    #
    # Type: int
    # Default: 4096
    readBufferSize:

    # Per-connection buffer size for responses' writing.
    #
    # Type: int
    # Default: 4096
    writeBufferSize:

    # CompressedFileSuffix adds suffix to the original file name and
    # tries saving the resulting compressed file under the new file name.
    #
    # Type: string
    # Default: ".fiber.gz"
    compressedFileSuffix:

    # ProxyHeader will enable c.IP() to return the value of the given header key
    # By default c.IP() will return the Remote IP from the TCP connection
    # This property can be useful if you are behind a load balancer: X-Forwarded-*
    # NOTE: headers are easily spoofed and the detected IP addresses are unreliable.
    #
    # Type: string
    # Default: ""
    proxyHeader:

    # GETOnly rejects all non-GET requests if set to true.
    # This option is useful as anti-DoS protection for servers
    # accepting only GET requests. The request size is limited
    # by ReadBufferSize if GETOnly is set.
    #
    # Type: bool
    # Default: false
    gETOnly:

    # When set to true, disables keep-alive connections.
    # The server will close incoming connections after sending the first response to client.
    #
    # Type: bool
    # Default: false
    disableKeepalive:

    # When set to true, causes the default date header to be excluded from the response.
    #
    # Type: bool
    # Default: false
    disableDefaultDate:

    # When set to true, causes the default Content-Type header to be excluded from the response.
    #
    # Type: bool
    # Default: false
    disableDefaultContentType:

    # When set to true, disables header normalization.
    # By default, all header names are normalized: conteNT-tYPE -> Content-Type.
    #
    # Type: bool
    # Default: false
    disableHeaderNormalizing:

    # When set to true, it will not print out the «Fiber» ASCII art and listening address.
    #
    # Type: bool
    # Default: false
    disableStartupMessage:

    # StreamRequestBody enables request body streaming,
    # and calls the handler sooner when given body is
    # larger than the current limit.
    #
    # Type: bool
    streamRequestBody:

    # Will not pre parse Multipart Form data if set to true.
    #
    # This option is useful for servers that desire to treat
    # multipart form data as a binary blob, or choose when to parse the data.
    # Server pre parses multipart form data by default.
    #
    # Type: bool
    disablePreParseMultipartForm:

    # Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only)
    # WARNING: When prefork is set to true, only "tcp4" and "tcp6" can be chose.
    #
    # Type: string
    # Default: NetworkTCP4
    network:

    # If you find yourself behind some sort of proxy, like a load balancer,
    # then certain header information may be sent to you using special X-Forwarded-* headers or the Forwarded header.
    # For example, the Host HTTP header is usually used to return the requested host.
    # But when you’re behind a proxy, the actual host may be stored in an X-Forwarded-Host header.
    #
    # If you are behind a proxy, you should enable TrustedProxyCheck to prevent header spoofing.
    # If you enable EnableTrustedProxyCheck and leave TrustedProxies empty Fiber will skip
    # all headers that could be spoofed.
    # If request ip in TrustedProxies whitelist then:
    #   1. c.Protocol() get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header
    #   2. c.IP() get value from ProxyHeader header.
    #   3. c.Hostname() get value from X-Forwarded-Host header
    # But if request ip NOT in Trusted Proxies whitelist then:
    #   1. c.Protocol() WON't get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header,
    #    will return https in case when tls connection is handled by the app, of http otherwise
    #   2. c.IP() WON'T get value from ProxyHeader header, will return RemoteIP() from fasthttp context
    #   3. c.Hostname() WON'T get value from X-Forwarded-Host header, fasthttp.Request.URI().Host()
    #    will be used to get the hostname.
    #
    # Type: bool
    # Default: false
    enableTrustedProxyCheck:

    # Read EnableTrustedProxyCheck doc.
    #
    # Type: []string
    # Default: []string
    # Example: ["proxy1", "proxy2", "proxy3"]
    trustedProxies:

    # If set to true, c.IP() and c.IPs() will validate IP addresses before returning them.
    # Also, c.IP() will return only the first valid IP rather than just the raw header
    # WARNING: this has a performance cost associated with it.
    #
    # Type: bool
    # Default: false
    enableIPValidation:

    # If set to true, will print all routes with their method, path and handler.
    #
    # Type: bool
    # Default: false
    enablePrintRoutes:

    # RequestMethods provides customizability for HTTP methods. You can add/remove methods as you wish.
    #
    # Type: []string
    # Optional. Default: DefaultMethods
    requestMethods:

    # EnableSplittingOnParsers splits the query/body/header parameters by comma when it's true.
    # For example, you can use it to parse multiple values from a query parameter like this:
    #   /api?foo=bar,baz == foo[]=bar&foo[]=baz
    #
    # Type: bool
    # Optional. Default: false
    enableSplittingOnParsers:

    # TLS configuration for Fiber app.
    tls:
        # Enable TLS for Fiber app.
        #
        # Type: bool
        # Optional. Default: false
        enabled:

        # Path to the TLS certificate file.
        #
        # Type: string
        # Default: ""
        certFile:

        # Path to the TLS key file.
        #
        # Type: string
        # Default: ""
        keyFile:
```
