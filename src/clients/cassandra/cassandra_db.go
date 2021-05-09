package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

}

func GetSession() (*gocql.Session, error) {
	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()
	// fmt.Println("session created with cassandra")
	return cluster.CreateSession()
}
