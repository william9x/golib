package properties

import (
	"time"

	"github.com/william9x/golib/config"
)

func NewFiberProperties(loader config.Loader) (*FiberProperties, error) {
	props := FiberProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type FiberProperties struct {
	// When set to true, this will spawn multiple Go processes listening on the same port.
	//
	// Default: false
	Prefork bool

	// When set to true, the router treats "/foo" and "/foo/" as different.
	// By default, this is disabled and both "/foo" and "/foo/" will execute the same handler.
	//
	// Default: false
	StrictRouting bool

	// When set to true, enables case-sensitive routing.
	// E.g. "/FoO" and "/foo" are treated as different routes.
	// By default, this is disabled and both "/FoO" and "/foo" will execute the same handler.
	//
	// Default: false
	CaseSensitive bool

	// When set to true, this relinquishes the 0-allocation promise in certain
	// cases in order to access the handler values (e.g. request bodies) in an
	// immutable fashion so that these values are available even if you return
	// from handler.
	//
	// Default: true
	Immutable bool `default:"true"`

	// When set to true, converts all encoded characters in the route back
	// before setting the path for the context, so that the routing,
	// the returning of the current url from the context `ctx.Path()`
	// and the parameters `ctx.Params(%key%)` with decoded characters will work
	//
	// Default: false
	UnescapePath bool

	// Max body size that the server accepts.
	// -1 will decline any body size
	//
	// Default: 4 * 1024 * 1024
	BodyLimit int

	// Maximum number of concurrent connections.
	//
	// Default: 256 * 1024
	Concurrency int

	// The amount of time allowed to read the full request including body.
	// It is reset after the request handler has returned.
	// The connection's read deadline is reset when the connection opens.
	//
	// Default: unlimited
	ReadTimeout time.Duration

	// The maximum duration before timing out writes of the response.
	// It is reset after the request handler has returned.
	//
	// Default: unlimited
	WriteTimeout time.Duration

	// The maximum amount of time to wait for the next request when keep-alive is enabled.
	// If IdleTimeout is zero, the value of ReadTimeout is used.
	//
	// Default: unlimited
	IdleTimeout time.Duration

	// Per-connection buffer size for requests' reading.
	// This also limits the maximum header size.
	// Increase this buffer if your clients send multi-KB RequestURIs
	// and/or multi-KB headers (for example, BIG cookies).
	//
	// Default: 4096
	ReadBufferSize int

	// Per-connection buffer size for responses' writing.
	//
	// Default: 4096
	WriteBufferSize int

	// CompressedFileSuffix adds suffix to the original file name and
	// tries saving the resulting compressed file under the new file name.
	//
	// Default: ".fiber.gz"
	CompressedFileSuffix string

	// ProxyHeader will enable c.IP() to return the value of the given header key
	// By default c.IP() will return the Remote IP from the TCP connection
	// This property can be useful if you are behind a load balancer: X-Forwarded-*
	// NOTE: headers are easily spoofed and the detected IP addresses are unreliable.
	//
	// Default: ""
	ProxyHeader string

	// GETOnly rejects all non-GET requests if set to true.
	// This option is useful as anti-DoS protection for servers
	// accepting only GET requests. The request size is limited
	// by ReadBufferSize if GETOnly is set.
	//
	// Default: false
	GETOnly bool

	// When set to true, disables keep-alive connections.
	// The server will close incoming connections after sending the first response to client.
	//
	// Default: false
	DisableKeepalive bool

	// When set to true, causes the default date header to be excluded from the response.
	//
	// Default: false
	DisableDefaultDate bool

	// When set to true, causes the default Content-Type header to be excluded from the response.
	//
	// Default: false
	DisableDefaultContentType bool

	// When set to true, disables header normalization.
	// By default, all header names are normalized: conteNT-tYPE -> Content-Type.
	//
	// Default: false
	DisableHeaderNormalizing bool

	// When set to true, it will not print out the «Fiber» ASCII art and listening address.
	//
	// Default: false
	DisableStartupMessage bool

	// StreamRequestBody enables request body streaming,
	// and calls the handler sooner when given body is
	// larger than the current limit.
	StreamRequestBody bool

	// Will not pre parse Multipart Form data if set to true.
	//
	// This option is useful for servers that desire to treat
	// multipart form data as a binary blob, or choose when to parse the data.
	//
	// Server pre parses multipart form data by default.
	DisablePreParseMultipartForm bool

	// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only)
	// WARNING: When prefork is set to true, only "tcp4" and "tcp6" can be chose.
	//
	// Default: NetworkTCP4
	Network string

	// If you find yourself behind some sort of proxy, like a load balancer,
	// then certain header information may be sent to you using special X-Forwarded-* headers or the Forwarded header.
	// For example, the Host HTTP header is usually used to return the requested host.
	// But when you’re behind a proxy, the actual host may be stored in an X-Forwarded-Host header.
	//
	// If you are behind a proxy, you should enable TrustedProxyCheck to prevent header spoofing.
	// If you enable EnableTrustedProxyCheck and leave TrustedProxies empty Fiber will skip
	// all headers that could be spoofed.
	// If request ip in TrustedProxies whitelist then:
	//   1. c.Protocol() get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header
	//   2. c.IP() get value from ProxyHeader header.
	//   3. c.Hostname() get value from X-Forwarded-Host header
	// But if request ip NOT in Trusted Proxies whitelist then:
	//   1. c.Protocol() WON't get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header,
	//    will return https in case when tls connection is handled by the app, of http otherwise
	//   2. c.IP() WON'T get value from ProxyHeader header, will return RemoteIP() from fasthttp context
	//   3. c.Hostname() WON'T get value from X-Forwarded-Host header, fasthttp.Request.URI().Host()
	//    will be used to get the hostname.
	//
	// Default: false
	EnableTrustedProxyCheck bool

	// Read EnableTrustedProxyCheck doc.
	//
	// Default: []string
	TrustedProxies []string

	// If set to true, c.IP() and c.IPs() will validate IP addresses before returning them.
	// Also, c.IP() will return only the first valid IP rather than just the raw header
	// WARNING: this has a performance cost associated with it.
	//
	// Default: false
	EnableIPValidation bool

	// If set to true, will print all routes with their method, path and handler.
	// Default: false
	EnablePrintRoutes bool

	// RequestMethods provides customizability for HTTP methods. You can add/remove methods as you wish.
	//
	// Optional. Default: DefaultMethods
	RequestMethods []string

	// EnableSplittingOnParsers splits the query/body/header parameters by comma when it's true.
	// For example, you can use it to parse multiple values from a query parameter like this:
	//   /api?foo=bar,baz == foo[]=bar&foo[]=baz
	//
	// Optional. Default: false
	EnableSplittingOnParsers bool

	Tls TLSProperties
}

type TLSProperties struct {
	Enabled  bool `default:"false"`
	CertFile string
	KeyFile  string
}

func (t FiberProperties) Prefix() string {
	return "app.fiber"
}
