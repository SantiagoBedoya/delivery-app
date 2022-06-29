package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SantiagoBedoya/delivery-app-customers/accounts"
	"github.com/SantiagoBedoya/delivery-app-customers/repositories/postgres"
	"github.com/SantiagoBedoya/delivery-app-customers/utils/ratelimit"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var (
	dbHost         string
	dbPort         string
	dbUser         string
	dbPass         string
	dbName         string
	privateKeyPath string
	publicKeyPath  string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file")
	}
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
	privateKeyPath = os.Getenv("PRIVATE_KEY_PATH")
	publicKeyPath = os.Getenv("PUBLIC_KEY_PATH")
}

func main() {
	privateKey, err := readCert(privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	publicKey, err := readCert(publicKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	dbsource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("pgx", dbsource)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app := gin.Default()
	app.Use(helmet.Default())
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(cors.Default())
	{
		repository := postgres.NewPostgresRepository(db)
		service := accounts.NewService(repository, privateKey, publicKey)
		handler := accounts.NewHandler(service)
		app.POST("/api/v1/auth/sign-in", ratelimit.GetRateLimit("5-S"), handler.SignIn)
		app.POST("/api/v1/auth/sign-up", ratelimit.GetRateLimit("5-H"), handler.SignUp)
	}

	errs := make(chan error, 2)
	go func() {
		port := handlePort()
		log.Printf("service is running on port %s\n", port)
		server := http.Server{
			Addr:         port,
			WriteTimeout: 60 * time.Second,
			ReadTimeout:  60 * time.Second,
			Handler:      app,
		}
		errs <- server.ListenAndServe()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	log.Printf("terminated %s\n", <-errs)
}

func handlePort() string {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func readCert(path string) ([]byte, error) {
	cert, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return cert, nil
}
