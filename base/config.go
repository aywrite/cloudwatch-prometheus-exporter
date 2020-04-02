package base

import (
	"strings"
	"time"

	"github.com/CoverGenius/cloudwatch-prometheus-exporter/helpers"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

/*
Example config

metrics:
  AWS/ELB:
   - metric: RequestCount
     help: "This is some help about the metric"
     dimensions: [AvailabilityZone, LoadBalancerName]
       resource_type_selection: "elasticloadbalancing:loadbalancer"
       resource_id_dimension: LoadBalancerName
     statistics: [Sum]
*/
type configMetric struct {
	Metric        string    `yaml:"metric"`         // The Cloudwatch metric to use
	Help          string    `yaml:"help"`           // Custom help text for the generated metric
	Dimensions    []*string `yaml:"dimensions"`     // The resource dimensions to generate individual series for (via labels)
	Statistics    []*string `yaml:"statistics"`     // List of AWS statistics to use.
	OutputName    string    `yaml:"output_name"`    // Allows override of the generate metric name
	RangeSeconds  int       `yaml:"range_seconds"`  // How far back to request data for in seconds.
	PeriodSeconds int       `yaml:"period_seconds"` // Granularity of results from cloudwatch API.
}

type metric struct {
	Data map[string][]*configMetric `yaml:",omitempty,inline"` // Map from namespace to list of metrics to scrape.
}

// Config represents the exporter configuration passed which is read at runtime from a YAML file.
type Config struct {
	Listen       string            `yaml:"listen,omitempty"`        // TCP Dial address for Prometheus HTTP API to listen on
	APIKey       string            `yaml:"api_key"`                 // AWS API Key ID
	APISecret    string            `yaml:"api_secret"`              // AWS API Secret
	Tags         []*TagDescription `yaml:"tags,omitempty"`          // Tags to filter resources by
	Period       uint8             `yaml:"period,omitempty"`        // How far back to request data for in minutes.
	Regions      []*string         `yaml:"regions"`                 // Which AWS regions to query resources and metrics for
	PollInterval uint8             `yaml:"poll_interval,omitempty"` // How often to fetch new data from the Cloudwatch API.
	LogLevel     uint8             `yaml:"log_level,omitempty"`     // Logging verbosity level
	Metrics      metric            `yaml:"metrics"`                 // Map of per metric configuration overrides
}

// LoadConfig reads the config file located at path and reads it into the Config struct
func LoadConfig(path string) (*Config, error) {
	c := Config{}
	helpers.YAMLDecode(&path, &c)
	return &c, nil
}

func (c *Config) ConstructMetrics() map[string]map[string]*MetricDescription {
	mds := make(map[string]map[string]*MetricDescription)
	for namespace, metrics := range c.Metrics.Data {
		mds[namespace] = make(map[string]*MetricDescription)
		for _, metric := range metrics {

			name := metric.OutputName
			if name == "" {
				name = helpers.ToSnakeCase(metric.Metric)
				name = strings.ToLower(strings.TrimPrefix(namespace, "AWS/")) + "_" + name
			}

			period := metric.PeriodSeconds
			if period == 0 {
				period = int(c.Period) * int(time.Minute)
			}

			rangeSeconds := metric.RangeSeconds
			if rangeSeconds == 0 {
				rangeSeconds = int(c.Period) * int(time.Minute)
			}

			// TODO one for each stat
			// TODO read defaults for namespace
			// TODO handle dimensions
			mds[namespace][metric.Metric] = &MetricDescription{
				Help:          &metric.Help,
				OutputName:    &name,
				Dimensions:    []*cloudwatch.Dimension{},
				PeriodSeconds: period,
				RangeSeconds:  rangeSeconds,
				Statistics:    metric.Statistics,

				namespace: &namespace,
				awsMetric: &metric.Metric,
			}

		}

	}
	return mds
}
