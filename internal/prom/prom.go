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
			Name: "create_project_count", // metric name
			Help: "Number of successful created projects.",
		},
		[]string{"status"}, // labels
	)
	getUpdateProjectCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "update_project_count", // metric name
			Help: "Number of successful updated projects.",
		},
		[]string{"status"}, // labels
	)
	getRemoveProjectCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "remove_project_count", // metric name
			Help: "Number of successful removed projects.",
		},
		[]string{"status"}, // labels
	)

	// must register counter on init
	prometheus.MustRegister(getCreateProjectCounter, getUpdateProjectCounter, getRemoveProjectCounter)
}

func CreateProjectCounterInc(status string) {
	if getCreateProjectCounter != nil {
		getCreateProjectCounter.WithLabelValues(status).Inc()
	}
}

func UpdateProjectCounterInc(status string) {
	if getUpdateProjectCounter != nil {
		getUpdateProjectCounter.WithLabelValues(status).Inc()
	}
}

func RemoveProjectCounterInc(status string) {
	if getRemoveProjectCounter != nil {
		getRemoveProjectCounter.WithLabelValues(status).Inc()
	}
}
