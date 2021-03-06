// Copyright © 2017-2019 The OpenEBS Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pool

import (
	"strconv"

	zpool "github.com/openebs/maya/pkg/zpool/v1alpha1"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	size                prometheus.Gauge
	usedCapacity        prometheus.Gauge
	freeCapacity        prometheus.Gauge
	usedCapacityPercent prometheus.Gauge
	status              *prometheus.GaugeVec

	zpoolCommandErrorCounter     prometheus.Gauge
	zpoolRejectRequestCounter    prometheus.Gauge
	zpoolListparseErrorCounter   prometheus.Gauge
	noPoolAvailableErrorCounter  prometheus.Gauge
	incompleteOutputErrorCounter prometheus.Gauge
}

type statsFloat64 struct {
	status              float64
	size                float64
	used                float64
	free                float64
	usedCapacityPercent float64
}

// List returns list of type float64 of various stats
// NOTE: Please donot change the order, add the new stats
// at the end of the list.
func (s *statsFloat64) List() []float64 {
	return []float64{
		s.size,
		s.used,
		s.free,
		s.usedCapacityPercent,
	}
}

func parseFloat64(e string, m *metrics) float64 {
	num, err := strconv.ParseFloat(e, 64)
	if err != nil {
		m.zpoolListparseErrorCounter.Inc()
	}
	return num
}

func (s *statsFloat64) parse(stats zpool.Stats, p *pool) {
	s.size = parseFloat64(stats.Size, &p.metrics)
	s.used = parseFloat64(stats.Used, &p.metrics)
	s.free = parseFloat64(stats.Free, &p.metrics)
	s.status = zpool.Status[stats.Status]
	s.usedCapacityPercent = parseFloat64(stats.UsedCapacityPercent, &p.metrics)
}

// newMetrics initializes fields of the metrics and returns its instance
func newMetrics() metrics {
	return metrics{
		size: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "pool_size",
				Help:      "Size of pool",
			},
		),

		status: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "pool_status",
				Help:      `Status of pool (0, 1, 2, 3, 4, 5, 6)= {"Offline", "Online", "Degraded", "Faulted", "Removed", "Unavail", "NoPoolsAvailable"}`,
			},
			[]string{"pool"},
		),

		usedCapacity: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "used_pool_capacity",
				Help:      "Capacity used by pool",
			},
		),

		freeCapacity: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "free_pool_capacity",
				Help:      "Free capacity in pool",
			},
		),

		usedCapacityPercent: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "used_pool_capacity_percent",
				Help:      "Capacity used by pool in percent",
			},
		),

		zpoolListparseErrorCounter: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "zpool_list_parse_error_count",
				Help:      "Total no of parsing errors",
			},
		),

		zpoolRejectRequestCounter: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "zpool_reject_request_count",
				Help:      "Total no of rejected requests of zpool command",
			},
		),

		zpoolCommandErrorCounter: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "zpool_command_error",
				Help:      "Total no of zpool command errors",
			},
		),

		noPoolAvailableErrorCounter: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "no_pool_available_error",
				Help:      "Total no of no pool available errors",
			},
		),

		incompleteOutputErrorCounter: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "openebs",
				Name:      "zpool_list_incomplete_stdout_error",
				Help:      "Total no of incomplete stdout of zpool list command errors",
			},
		),
	}
}
