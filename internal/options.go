package internal

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/uber/jaeger-client-go"
	//clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
)

// ResolvedOpts is a set of knobs controllable by Options.
type ResolvedOpts struct {
	Hostname     string
	PortHTTP     uint
	PortAdmin    uint
	PortGRPC     uint
	PortChannelz uint

	AdminServer                   chi.Router
	UnaryInterceptor              []grpc.UnaryServerInterceptor
	StreamInterceptor             []grpc.StreamServerInterceptor
	ServiceDescriptionInterceptor []grpc.UnaryServerInterceptor
	StatsHandler                  []stats.Handler
	GrpcServerOptions             []grpc.ServerOption
	// TODO
	//EtcdClient                    *clientv3.Client

	PublicTLSConfig *tls.Config

	PublicMiddleware []func(http.Handler) http.Handler
	AdminMiddleware  []func(http.Handler) http.Handler

	// TODO
	CorsOptions                             cors.Options
	CorsCallbacks                           []func(*cors.Options)
	CorsRequireOriginInDefaultAndCustomList bool // флаг, обяжет Origin присутсвовать и в кастомном списке(который задается через опцию) и в дефолтном списке, все остальные запросы будут отброшены
	CorsDisableOriginInDefaultList          bool // флаг, обяжет Origin присутствовать в кастомном списке(который задается через опцию), присутствие лишь в дефолтном списке не даст разрешения

	OperationNameFunc func(*http.Request) string

	// TODO
	SpanObserver jaeger.ContribObserver

	// TODO
	ServeMuxOpts []runtime.ServeMuxOption

	PublicURLTagFunc func(url *url.URL) string

	DisabledChannelz         bool
	SwaggerInt64Faked        bool
	DisabledServeMux         bool
	DisabledLogLevelWatching bool
	DisabledHTTPRecoverMw    bool
	DisabledHTTPAgent        bool
	DisabledHTTPRatelimiter  bool
	DisabledCors             bool
	DisabledPublicHTTP       bool
	DisabledTracer           bool
	DisabledGzip             bool

	BindAddress string
}

// Options holds configuration for an `App`.
type Options struct {
	ResolvedOpts
	DisableErrorLogInterceptor     bool
	DisableContextErrorInterceptor bool
	ConfigClient                   interface{} // образуется циклическая зависимость при указании типа.
	SecretClient                   interface{} // образуется циклическая зависимость при указании типа.
}
