// Package metrics provides Prometheus metrics for the CertWatch Agent.
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Certificate metrics
	CertDaysUntilExpiry = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certwatch",
			Subsystem: "certificate",
			Name:      "days_until_expiry",
			Help:      "Days until certificate expires",
		},
		[]string{"hostname", "port"},
	)

	CertValid = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certwatch",
			Subsystem: "certificate",
			Name:      "valid",
			Help:      "Certificate validity (1=valid, 0=invalid)",
		},
		[]string{"hostname", "port"},
	)

	CertChainValid = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certwatch",
			Subsystem: "certificate",
			Name:      "chain_valid",
			Help:      "Certificate chain validity (1=valid, 0=invalid)",
		},
		[]string{"hostname", "port"},
	)

	CertExpiryTimestamp = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certwatch",
			Subsystem: "certificate",
			Name:      "expiry_timestamp_seconds",
			Help:      "Certificate expiry as Unix timestamp",
		},
		[]string{"hostname", "port"},
	)

	// Scan metrics
	ScanTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "certwatch",
			Subsystem: "scan",
			Name:      "total",
			Help:      "Total number of certificate scans",
		},
		[]string{"status"}, // "success" or "failure"
	)

	ScanDurationSeconds = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "certwatch",
			Subsystem: "scan",
			Name:      "duration_seconds",
			Help:      "Certificate scan duration in seconds",
			Buckets:   prometheus.ExponentialBuckets(0.1, 2, 10),
		},
		[]string{"hostname"},
	)

	// Sync metrics
	SyncTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "certwatch",
			Subsystem: "sync",
			Name:      "total",
			Help:      "Total number of sync operations",
		},
		[]string{"status"}, // "success" or "failure"
	)

	SyncDurationSeconds = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "certwatch",
			Subsystem: "sync",
			Name:      "duration_seconds",
			Help:      "Sync operation duration in seconds",
			Buckets:   prometheus.DefBuckets,
		},
	)

	SyncCertsCreated = promauto.NewCounter(
		prometheus.CounterOpts{
			Namespace: "certwatch",
			Subsystem: "sync",
			Name:      "certificates_created_total",
			Help:      "Total number of certificates created via sync",
		},
	)

	SyncCertsUpdated = promauto.NewCounter(
		prometheus.CounterOpts{
			Namespace: "certwatch",
			Subsystem: "sync",
			Name:      "certificates_updated_total",
			Help:      "Total number of certificates updated via sync",
		},
	)

	SyncCertsOrphaned = promauto.NewCounter(
		prometheus.CounterOpts{
			Namespace: "certwatch",
			Subsystem: "sync",
			Name:      "certificates_orphaned_total",
			Help:      "Total number of certificates orphaned via sync",
		},
	)

	// Heartbeat metrics
	HeartbeatTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "certwatch",
			Subsystem: "heartbeat",
			Name:      "total",
			Help:      "Total number of heartbeat operations",
		},
		[]string{"status"}, // "success" or "failure"
	)

	HeartbeatDurationSeconds = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "certwatch",
			Subsystem: "heartbeat",
			Name:      "duration_seconds",
			Help:      "Heartbeat operation duration in seconds",
			Buckets:   prometheus.DefBuckets,
		},
	)

	// Agent info
	AgentInfo = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certwatch",
			Subsystem: "agent",
			Name:      "info",
			Help:      "Agent information (always 1)",
		},
		[]string{"version", "name", "agent_id"},
	)

	CertificatesConfigured = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "certwatch",
			Subsystem: "agent",
			Name:      "certificates_configured",
			Help:      "Number of certificates configured for monitoring",
		},
	)

	AgentUptime = promauto.NewCounter(
		prometheus.CounterOpts{
			Namespace: "certwatch",
			Subsystem: "agent",
			Name:      "uptime_seconds_total",
			Help:      "Total uptime of the agent in seconds",
		},
	)
)

// RecordCertificateMetrics updates all certificate-related metrics for a single certificate.
func RecordCertificateMetrics(hostname, port string, daysUntilExpiry, expiryTimestamp float64, valid, chainValid bool) {
	CertDaysUntilExpiry.WithLabelValues(hostname, port).Set(daysUntilExpiry)
	CertExpiryTimestamp.WithLabelValues(hostname, port).Set(expiryTimestamp)

	if valid {
		CertValid.WithLabelValues(hostname, port).Set(1)
	} else {
		CertValid.WithLabelValues(hostname, port).Set(0)
	}

	if chainValid {
		CertChainValid.WithLabelValues(hostname, port).Set(1)
	} else {
		CertChainValid.WithLabelValues(hostname, port).Set(0)
	}
}

// RecordScanSuccess records a successful scan operation.
func RecordScanSuccess(hostname string, duration float64) {
	ScanTotal.WithLabelValues("success").Inc()
	ScanDurationSeconds.WithLabelValues(hostname).Observe(duration)
}

// RecordScanFailure records a failed scan operation.
func RecordScanFailure(hostname string, duration float64) {
	ScanTotal.WithLabelValues("failure").Inc()
	ScanDurationSeconds.WithLabelValues(hostname).Observe(duration)
}

// RecordSyncSuccess records a successful sync operation.
func RecordSyncSuccess(duration float64, created, updated, orphaned int) {
	SyncTotal.WithLabelValues("success").Inc()
	SyncDurationSeconds.Observe(duration)
	SyncCertsCreated.Add(float64(created))
	SyncCertsUpdated.Add(float64(updated))
	SyncCertsOrphaned.Add(float64(orphaned))
}

// RecordSyncFailure records a failed sync operation.
func RecordSyncFailure(duration float64) {
	SyncTotal.WithLabelValues("failure").Inc()
	SyncDurationSeconds.Observe(duration)
}

// RecordHeartbeatSuccess records a successful heartbeat operation.
func RecordHeartbeatSuccess(duration float64) {
	HeartbeatTotal.WithLabelValues("success").Inc()
	HeartbeatDurationSeconds.Observe(duration)
}

// RecordHeartbeatFailure records a failed heartbeat operation.
func RecordHeartbeatFailure(duration float64) {
	HeartbeatTotal.WithLabelValues("failure").Inc()
	HeartbeatDurationSeconds.Observe(duration)
}

// SetAgentInfo sets the agent info metric.
func SetAgentInfo(version, name, agentID string) {
	AgentInfo.WithLabelValues(version, name, agentID).Set(1)
}

// SetCertificatesConfigured sets the number of configured certificates.
func SetCertificatesConfigured(count int) {
	CertificatesConfigured.Set(float64(count))
}
