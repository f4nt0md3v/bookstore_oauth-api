package cassandra

import (
	"os"

	"github.com/gocql/gocql"
)

const (
	cassOAuthUsername = "CASSANDRA_USERNAME"
	cassOAuthPassword = "CASSANDRA_PASSWORD"
	cassOAuthHostname = "CASSANDRA_HOSTNAME"
	cassOAuthKeyspace = "CASSANDRA_KEYSPACE"
)

var (
	// TODO: bad practice to use global session but keep it for now
	session *gocql.Session

	// Cassandra parameters
	username = os.Getenv(cassOAuthUsername)
	password = os.Getenv(cassOAuthPassword)
	host     = os.Getenv(cassOAuthHostname)
	keyspace = os.Getenv(cassOAuthKeyspace)
)

func init() {
	// Connect to Cassandra cluster:
	cluster := gocql.NewCluster(host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum

	// Creating a single session to `serve them all`
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
