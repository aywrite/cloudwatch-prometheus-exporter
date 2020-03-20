package network

import (
	b "github.com/CoverGenius/cloudwatch-prometheus-exporter/base"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var metrics = map[string]*b.MetricDescription{
	"ActiveConnectionCount": {
		Help:       aws.String("The total number of concurrent active TCP connections through the NAT gateway"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_active_connection_count"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"BytesInFromDestination": {
		Help:       aws.String("The number of bytes received by the NAT gateway from the destination"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_bytes_in_from_destination"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"BytesInFromSource": {
		Help:       aws.String("The number of bytes received by the NAT gateway from clients in your VPC"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_bytes_in_from_source"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"BytesOutToDestination": {
		Help:       aws.String("The number of bytes sent out through the NAT gateway to the destination"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_bytes_out_to_destination"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"BytesOutToSource": {
		Help:       aws.String("The number of bytes sent through the NAT gateway to the clients in your VPC"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_bytes_out_to_source"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"ConnectionAttemptCount": {
		Help:       aws.String("The number of connection attempts made through the NAT gateway"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_connection_attempt_count"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"ConnectionEstablishedCount": {
		Help:       aws.String("The number of connections established through the NAT gateway"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_connection_established_count"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"ErrorPortAllocation": {
		Help:       aws.String("The number of times the NAT gateway could not allocate a source port"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_error_port_allocation"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"IdleTimeoutCount": {
		Help:       aws.String("The number of connections that transitioned from the active state to the idle state. An active connection transitions to idle if it was not closed gracefully and there was no activity for the last 350 seconds"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_idle_timeout_count"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"PacketsDropCount": {
		Help:       aws.String("The number of packets dropped by the NAT gateway"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_packets_drop_count"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"PacketsInFromDestination": {
		Help:       aws.String("The number of packets received by the NAT gateway from the destination"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_packets_in_from_destination"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"PacketsInFromSource": {
		Help:       aws.String("The number of packets received by the NAT gateway from clients in your VPC"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_packets_in_from_source"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"PacketsOutToDestination": {
		Help:       aws.String("The number of packets sent out through the NAT gateway to the destination"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_packets_out_to_destination"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"PacketsOutToSource": {
		Help:       aws.String("The number of packets sent through the NAT gateway to the clients in your VPC"),
		Type:       aws.String("counter"),
		OutputName: aws.String("nat_gateway_packets_out_to_source"),
		Data:       map[string][]*string{},
		Statistic:  aws.String("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
}

func GetMetrics() map[string]*b.MetricDescription {
	return metrics
}