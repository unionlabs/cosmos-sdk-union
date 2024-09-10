// Code generated by metricsgen. DO NOT EDIT.

package state

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		BlockProcessingTime: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_processing_time",
			Help:      "Time spent processing FinalizeBlock",

			Buckets: stdprometheus.LinearBuckets(1, 10, 10),
		}, labels).With(labelsAndValues...),
		ConsensusParamUpdates: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "consensus_param_updates",
			Help:      "Number of consensus parameter updates returned by the application since process start.",
		}, labels).With(labelsAndValues...),
		ValidatorSetUpdates: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_set_updates",
			Help:      "Number of validator set updates returned by the application since process start.",
		}, labels).With(labelsAndValues...),
		PruningServiceBlockRetainHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "pruning_service_block_retain_height",
			Help:      "PruningServiceBlockRetainHeight is the accepted block retain height set by the data companion",
		}, labels).With(labelsAndValues...),
		PruningServiceBlockResultsRetainHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "pruning_service_block_results_retain_height",
			Help:      "PruningServiceBlockResultsRetainHeight is the accepted block results retain height set by the data companion",
		}, labels).With(labelsAndValues...),
		PruningServiceTxIndexerRetainHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "pruning_service_tx_indexer_retain_height",
			Help:      "PruningServiceTxIndexerRetainHeight is the accepted transactions indices retain height set by the data companion",
		}, labels).With(labelsAndValues...),
		PruningServiceBlockIndexerRetainHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "pruning_service_block_indexer_retain_height",
			Help:      "PruningServiceBlockIndexerRetainHeight is the accepted blocks indices retain height set by the data companion",
		}, labels).With(labelsAndValues...),
		ApplicationBlockRetainHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "application_block_retain_height",
			Help:      "ApplicationBlockRetainHeight is the accepted block retain height set by the application",
		}, labels).With(labelsAndValues...),
		BlockStoreBaseHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_store_base_height",
			Help:      "BlockStoreBaseHeight shows the first height at which a block is available",
		}, labels).With(labelsAndValues...),
		ABCIResultsBaseHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "abciresults_base_height",
			Help:      "ABCIResultsBaseHeight shows the first height at which abci results are available",
		}, labels).With(labelsAndValues...),
		TxIndexerBaseHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "tx_indexer_base_height",
			Help:      "TxIndexerBaseHeight shows the first height at which tx indices are available",
		}, labels).With(labelsAndValues...),
		BlockIndexerBaseHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_indexer_base_height",
			Help:      "BlockIndexerBaseHeight shows the first height at which block indices are available",
		}, labels).With(labelsAndValues...),
		StoreAccessDurationSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "store_access_duration_seconds",
			Help:      "The duration of accesses to the state store labeled by which method was called on the store.",

			Buckets: stdprometheus.ExponentialBuckets(0.0002, 10, 5),
		}, append(labels, "method")).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		BlockProcessingTime:                    discard.NewHistogram(),
		ConsensusParamUpdates:                  discard.NewCounter(),
		ValidatorSetUpdates:                    discard.NewCounter(),
		PruningServiceBlockRetainHeight:        discard.NewGauge(),
		PruningServiceBlockResultsRetainHeight: discard.NewGauge(),
		PruningServiceTxIndexerRetainHeight:    discard.NewGauge(),
		PruningServiceBlockIndexerRetainHeight: discard.NewGauge(),
		ApplicationBlockRetainHeight:           discard.NewGauge(),
		BlockStoreBaseHeight:                   discard.NewGauge(),
		ABCIResultsBaseHeight:                  discard.NewGauge(),
		TxIndexerBaseHeight:                    discard.NewGauge(),
		BlockIndexerBaseHeight:                 discard.NewGauge(),
		StoreAccessDurationSeconds:             discard.NewHistogram(),
	}
}
