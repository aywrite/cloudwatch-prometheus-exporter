package rds

import (
	b "github.com/CoverGenius/cloudwatch-prometheus-exporter/base"
	h "github.com/CoverGenius/cloudwatch-prometheus-exporter/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var metrics = map[string]*b.MetricDescription{
	"BinLogDiskUsage": {
		Help:       aws.String("The amount of disk space occupied by binary logs on the master. Applies to MySQL read replicas"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_bin_log_disk_usage"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"BurstBalance": {
		Help:       aws.String("The percent of General Purpose SSD (gp2) burst-bucket I/O credits available"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_burst_balance"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"CPUCreditBalance": {
		Help:       aws.String("The number of earned CPU credits that an instance has accrued since it was launched or started"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_cpu_credit_balance"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"CPUCreditUsage": {
		Help:       aws.String("The number of CPU credits spent by the instance for CPU utilization. One CPU credit equals one vCPU running at 100 percent utilization for one minute or an equivalent combination of vCPUs, utilization, and time"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_cpu_credit_usage"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"CPUSurplusCreditBalance": {
		Help:       aws.String("The number of surplus credits that have been spent by an unlimited instance when its CPUCreditBalance value is zero"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_cpu_surplus_credit_balance"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"CPUSurplusCreditsCharged": {
		Help:       aws.String("The number of spent surplus credits that are not paid down by earned CPU credits, and which thus incur an additional charge"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_cpu_surplus_credits_charged"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"CPUUtilization": {
		Help:       aws.String("The percentage of CPU utilization"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_cpu_utilization"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"DatabaseConnections": {
		Help:       aws.String("The number of database connections in use"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_database_connections"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"DBLoad": {
		Help:       aws.String("The number of active sessions for the DB engine. Typically, you want the data for the average number of active sessions"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_db_load"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"DBLoadCPU": {
		Help:       aws.String("The number of active sessions where the wait event type is CPU"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_db_load_cpu"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"DBLoadNonCPU": {
		Help:       aws.String("The number of active sessions where the wait event type is not CPU"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_db_load_non_cpu"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"DiskQueueDepth": {
		Help:       aws.String("The number of outstanding IOs (read/write requests) waiting to access the disk"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_disk_queue_depth"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"FreeableMemory": {
		Help:       aws.String("The amount of available random access memory"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_freeable_memory"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"FreeStorageSpace": {
		Help:       aws.String("The amount of available storage space"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_free_storage_space"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"MaximumUsedTransactionIDs": {
		Help:       aws.String("The maximum transaction ID that has been used. Applies to PostgreSQL"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_maximum_used_transaction_ids"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"NetworkReceiveThroughput": {
		Help:       aws.String("The incoming (Receive) network traffic on the DB instance, including both customer database traffic and Amazon RDS traffic used for monitoring and replication"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_network_receive_throughput"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"NetworkTransmitThroughput": {
		Help:       aws.String("The outgoing (Transmit) network traffic on the DB instance, including both customer database traffic and Amazon RDS traffic used for monitoring and replication"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_network_transmit_throughput"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"OldestReplicationSlotLag": {
		Help:       aws.String("The lagging size of the replica lagging the most in terms of WAL data received. Applies to PostgreSQL"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_oldest_replication_slot_lag"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"ReadIOPS": {
		Help:       aws.String("The average number of disk read I/O operations per second"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_read_iops"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"ReadLatency": {
		Help:       aws.String("The average amount of time taken per disk I/O operation"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_read_latency"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"ReadThroughput": {
		Help:       aws.String("The average number of bytes read from disk per second"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_read_throughput"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"ReplicaLag": {
		Help:       aws.String("The amount of time a Read Replica DB instance lags behind the source DB instance. Applies to MySQL, MariaDB, and PostgreSQL Read Replicas"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_replica_lag"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"ReplicationSlotDiskUsage": {
		Help:       aws.String("The disk space used by replication slot files. Applies to PostgreSQL"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_replication_slot_disk_usage"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"SwapUsage": {
		Help:       aws.String("The amount of swap space used on the DB instance. This metric is not available for SQL Server"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_swap_usage"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"TransactionLogsDiskUsage": {
		Help:       aws.String("The disk space used by transaction logs. Applies to PostgreSQL"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_transaction_logs_disk_usage"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"TransactionLogsGeneration": {
		Help:       aws.String("The size of transaction logs generated per second. Applies to PostgreSQL"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_transaction_logs_generation"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"WriteIOPS": {
		Help:       aws.String("The average number of disk write I/O operations per second"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_write_iops"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"WriteLatency": {
		Help:       aws.String("The average amount of time taken per disk I/O operation"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_write_latency"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
	"WriteThroughput": {
		Help:       aws.String("The average number of bytes written to disk per second"),
		Type:       aws.String("counter"),
		OutputName: aws.String("rds_write_throughput"),
		Data:       map[string][]*string{},
		Statistic:  h.StringPointers("Average"),
		Period:     5,
		Dimensions: []*cloudwatch.Dimension{},
	},
}

func GetMetrics() map[string]*b.MetricDescription {
	return metrics
}
