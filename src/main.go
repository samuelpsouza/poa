package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"fmt"
	"strconv"
	"time"
	"github.com/gocql/gocql"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)



type Authentication struct {
	Id int
	Key string
}

type Request struct {
	Id int
	Body string
}

type Response struct {
	Id int
	Body string
	StatusCode int
}

var auths = []Authentication{
	{Id:1, Key: "Basic"},
}

var answers = []Response{
	{Id:1, Body: "hello ans", StatusCode: 200},
}

func init() {
	godotenv.Load()
}

func main() {
	createDatabaseConnection()
	initWebServer()
}

func greeting(c *gin.Context) {
	authId, authErr := strconv.Atoi(c.Query("authId"))
	answerId, ansErr := strconv.Atoi(c.Query("answerId"))

	if authErr != nil || ansErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad request")
	}

	fmt.Println(authId)

	for _, res := range answers {
        if res.Id == answerId {
            c.IndentedJSON(http.StatusOK, res)
            return
        }
    }

	c.IndentedJSON(http.StatusOK, gin.H{"message": "answer not found"})
}

func greetingPost(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "post message"})
}

func createDatabaseConnection() *gocql.Session {
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

func initWebServer() {
	port     := os.Getenv("PORT")	

	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	initRoutes(router)
	router.Run("localhost:" + port)
}

func initRoutes(router *gin.Engine) {
	router.GET("/echo", greeting)
	router.POST("/echo", greetingPost)
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

func createLogger(level string) *zap.Logger {
	lvl := zap.NewAtomicLevel()
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		lvl.SetLevel(zap.InfoLevel)
	}
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		lvl,
	))
	return logger
}