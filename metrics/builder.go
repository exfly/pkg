package metrics

import (
	"errors"
	"expvar"
	"net/http"
	"os"
	"time"

	"github.com/exfly/pkg/log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/uber/jaeger-lib/metrics"
	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"
	jprom "github.com/uber/jaeger-lib/metrics/prometheus"
)

var errUnknownBackend = errors.New("unknown metrics backend specified")

const (
	PrometheusBackend = "prometheus"
	ExpvarBackend     = "expvar"
	NoneBackend       = ""
)

func NewBuilder(appName, jobName, pushTo string, pushInterval time.Duration, backend, httpRoute string) *Builder {
	return &Builder{
		AppName:      appName,
		JobName:      jobName,
		PushTo:       pushTo,
		PushInterval: pushInterval,
		Backend:      backend,
		HTTPRoute:    httpRoute,
	}
}

// Builder provides command line options to configure metrics backend used by Jaeger executables.
type Builder struct {
	AppName      string
	JobName      string
	PushTo       string
	PushInterval time.Duration
	Backend      string
	HTTPRoute    string // endpoint name to expose metrics, e.g. for scraping
	handler      http.Handler
}

func isDebug() bool {
	return os.Getenv("METRICS_DEBUG") == "debug"
}

// CreateMetricsFactory creates a metrics factory based on the configured type of the backend.
// If the metrics backend supports HTTP endpoint for scraping, it is stored in the builder and
// can be later added by RegisterHandler function.
func (b *Builder) CreateMetricsFactory(namespace string) (metrics.Factory, error) {
	debug := isDebug()
	if b.Backend == PrometheusBackend {
		metricsFactory := jprom.New().Namespace(metrics.NSOptions{Name: namespace, Tags: nil})
		b.handler = promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{DisableCompression: true},
		)

		go func() {
			if b.PushTo == "" {
				return
			}

			ticker := time.NewTicker(b.PushInterval)
			for range ticker.C {
				if debug {
					log.L().Debug("metrics pusher ticker")
				}
				if err := push.New(b.PushTo, b.JobName).
					Gatherer(prometheus.DefaultGatherer).Grouping("app", b.AppName).
					Push(); err != nil {
					log.L().Error("push to gateway failed", log.Error(err))
				}
			}
		}()

		return metricsFactory, nil
	}
	if b.Backend == ExpvarBackend {
		metricsFactory := jexpvar.NewFactory(10).Namespace(metrics.NSOptions{Name: namespace, Tags: nil})
		b.handler = expvar.Handler()
		return metricsFactory, nil
	}
	if b.Backend == NoneBackend || b.Backend == "none" || b.Backend == "" {
		return metrics.NullFactory, nil
	}
	return nil, errUnknownBackend
}

// Handler returns an http.Handler for the metrics endpoint.
func (b *Builder) Handler() http.Handler {
	return b.handler
}
