package prom

import "github.com/prometheus/client_golang/prometheus"

var (
	getCreateProjectCounter *prometheus.CounterVec
	getUpdateProjectCounter *prometheus.CounterVec
	getRemoveProjectCounter *prometheus.CounterVec
)

func RegisterProjectMetrics() {
	// create a new counter vector
	getCreateProjectCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "success_create_project_count", // metric name
			Help: "Number of successful created projects.",
		},
		[]string{"status"}, // labels
	)
	getUpdateProjectCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "success_update_project_count", // metric name
			Help: "Number of successful updated projects.",
		},
		[]string{"status"}, // labels
	)
	getRemoveProjectCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "success_remove_project_count", // metric name
			Help: "Number of successful removed projects.",
		},
		[]string{"status"}, // labels
	)

	// must register counter on init
	prometheus.MustRegister(getCreateProjectCounter, getUpdateProjectCounter, getRemoveProjectCounter)
}

func GetCreateProjectCounter() *prometheus.CounterVec {
	return getCreateProjectCounter
}

func GetUpdateProjectCounter() *prometheus.CounterVec {
	return getUpdateProjectCounter
}

func GetRemoveProjectCounter() *prometheus.CounterVec {
	return getRemoveProjectCounter
}
