package db

import (
	"time"
	"github.com/gocql/gocql"
)

func CreateDatabaseConnection() *gocql.Session {
	logger := createLogger("info")
	keyspace := os.Getenv("KEYSPACE")
	if keyspace == "" {
		keyspace = "poa"
	}

	cluster      := createCluster(gocql.Quorum, keyspace, "localhost")
	session, err := gocql.NewSession(*cluster)

	if err != nil {
		logger.Fatal("Unable to connect to scylla", zap.Error(err))
		return
	}

	return session
}

func createCluster(consistency gocql.Consistency, keyspace string, hosts ...string) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		Min: 		time.Second,
		Max:        10 * time.Second,
		NumRetries: 5,
	}
	cluster.Consistency = consistency
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}