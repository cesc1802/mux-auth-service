package cmd

import (
	"auth-service/component/appctx"
	"auth-service/handlers"
	"auth-service/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "start the back order service gateway",
	Run:   runServeHTTPCmd,
}

func init() {
	serveCmd.AddCommand(httpCmd)
}

type DBConfiguration struct {
	Username  string
	Password  string
	Database  string
	Host      string
	Port      string
	Loc       string
	Charset   string
	ParseTime string
}

func NewDBConfiguration(username string, password string, database string, host string, port string, loc string, charset string) *DBConfiguration {
	d := &DBConfiguration{Username: username, Password: password, Database: database, Host: host, Port: port, Loc: loc, Charset: charset}
	d.ParseTime = "True"
	return d
}

func initDB() *DBConfiguration {
	u := viper.GetString(MySQLUserName)
	p := viper.GetString(MySQLPassword)
	d := viper.GetString(MySQLDatabase)
	h := viper.GetString(MySQLHost)
	po := viper.GetString(MySQLPort)
	char := viper.GetString(MySQLCharset)
	l := viper.GetString(MySQLLoc)
	return NewDBConfiguration(u, p, d, h, po, l, char)
}

// ToDSN returns the mysql data source name based on configuration.
func (d *DBConfiguration) ToDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", d.Username, d.Password, d.Host, d.Port, d.Database, d.Charset, d.ParseTime, d.Loc)
}

func runServeHTTPCmd(cmd *cobra.Command, args []string) {
	logger := log.Default()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	d := initDB()
	mysqlDsn := d.ToDSN()
	dbConnection, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	maxOpenConnections := viper.GetInt(MySQLMaxOpenConnections)
	maxIdleConnections := viper.GetInt(MySQLMaxIdleConnections)
	sqlConn, err := dbConnection.DB()

	if err != nil {
		panic(err)
	}
	sqlConn.SetMaxOpenConns(maxOpenConnections)
	sqlConn.SetMaxIdleConns(maxIdleConnections)
	sqlConn.SetConnMaxLifetime(200 * time.Minute)
	ac := appctx.NewAppContext(dbConnection, "", "dev-secret")

	go func() {

		router := mux.NewRouter()
		router.Use(middleware.AllowCors)
		router.Use(middleware.Recover(ac))

		v1 := router.PathPrefix("/v1").Subrouter()

		handlers.PrivateRoute(ac, v1)

		srv := &http.Server{
			Addr:         "localhost:8080",
			Handler:      router,
			IdleTimeout:  60 * time.Second,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		}
		logger.Print("server started")
		err = srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	<-c

	logger.Print("server graceful shutdown")
}
