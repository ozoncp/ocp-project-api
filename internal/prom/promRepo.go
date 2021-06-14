package prom

import "github.com/prometheus/client_golang/prometheus"

var (
	getCreateRepoCounter *prometheus.CounterVec
	getUpdateRepoCounter *prometheus.CounterVec
	getRemoveRepoCounter *prometheus.CounterVec
)

func RegisterRepoMetrics() {
	// create a new counter vector
	getCreateRepoCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "create_repo_count", // metric name
			Help: "Number of successful created repos.",
		},
		[]string{"status"}, // labels
	)
	getUpdateRepoCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "update_repo_count", // metric name
			Help: "Number of successful updated repos.",
		},
		[]string{"status"}, // labels
	)
	getRemoveRepoCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "remove_repo_count", // metric name
			Help: "Number of successful removed repos.",
		},
		[]string{"status"}, // labels
	)

	// must register counter on init
	prometheus.MustRegister(getCreateRepoCounter, getUpdateRepoCounter, getRemoveRepoCounter)
}

func CreateRepoCounterInc(status string) {
	if getCreateRepoCounter != nil {
		getCreateRepoCounter.WithLabelValues(status).Inc()
	}
}

func UpdateRepoCounterInc(status string) {
	if getUpdateRepoCounter != nil {
		getUpdateRepoCounter.WithLabelValues(status).Inc()
	}
}

func RemoveRepoCounterInc(status string) {
	if getRemoveRepoCounter != nil {
		getRemoveRepoCounter.WithLabelValues(status).Inc()
	}
}
