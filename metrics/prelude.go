package metrics

import (
	"github.com/uber/jaeger-lib/metrics"
)

type (
	Factory          = metrics.Factory
	NSOptions        = metrics.NSOptions
	Options          = metrics.Options
	TimerOptions     = metrics.TimerOptions
	HistogramOptions = metrics.HistogramOptions

	Counter   = metrics.Counter
	Gauge     = metrics.Gauge
	Timer     = metrics.Timer
	Histogram = metrics.Histogram
)

var (
	NullFactory   = metrics.NullFactory
	NullCounter   = metrics.NullCounter
	NullGauge     = metrics.NullGauge
	NullTimer     = metrics.NullTimer
	NullHistogram = metrics.NullHistogram
	MustInit      = metrics.MustInit
	Init          = metrics.Init
)
