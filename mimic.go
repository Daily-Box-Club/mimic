package mimic

import (
	"log"

	http "github.com/saucesteals/fhttp"

	"github.com/saucesteals/fhttp/http2"
	utls "github.com/saucesteals/utls"
)

type FnTlsSpec func() *utls.ClientHelloSpec

func NewClientSpec(version string, clientHintUA string, h2Options *HTTP2Options, getTlsSpec FnTlsSpec) *ClientSpec {
	return &ClientSpec{
		version:      version,
		clientHintUA: clientHintUA,
		h2Options:    h2Options,
		getTlsSpec:   getTlsSpec,
	}
}

type ClientSpec struct {
	version      string
	clientHintUA string
	h2Options    *HTTP2Options
	getTlsSpec   FnTlsSpec
}

type HTTP2Options struct {
	Settings          []http2.Setting
	PseudoHeaderOrder []string
	MaxHeaderListSize uint32
	InitialWindowSize uint32
	HeaderTableSize   uint32
}

// ConfigureTransport configures a net/http HTTP/1 Transport to follow the client's spec
// Returns the given Transport for convenience
func (c *ClientSpec) ConfigureTransport(t1 *http.Transport) *http.Transport {
	t1.GetTlsClientHelloSpec = c.getTlsSpec

	t2, err := http2.ConfigureTransports(t1)

	if err != nil {
		log.Printf("error enabling Transport HTTP/2 support: %v", err)
		return t1
	}

	t2.Settings = c.h2Options.Settings
	t2.MaxHeaderListSize = c.h2Options.MaxHeaderListSize
	t2.InitialWindowSize = c.h2Options.MaxHeaderListSize
	t2.HeaderTableSize = c.h2Options.MaxHeaderListSize

	return t1
}

// Version returns the version for the mimicked client..
func (c *ClientSpec) Version() string {
	return c.version
}

// ClientHintUA returns the "sec-ch-ua" header value for the mimicked client.
func (c *ClientSpec) ClientHintUA() string {
	return c.clientHintUA
}

// PseudoHeaderOrder returns the pseudo header order for the mimicked client.
func (c *ClientSpec) PseudoHeaderOrder() []string {
	return c.h2Options.PseudoHeaderOrder
}