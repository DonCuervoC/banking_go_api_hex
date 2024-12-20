package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DonCuervoC/banking_go_api_hex/domain"
	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// checkEnvVars valida que todas las variables de entorno requeridas estén definidas.
func checkEnvVars(vars ...string) error {
	for _, v := range vars {
		if os.Getenv(v) == "" {
			return fmt.Errorf("environment variable %s is not defined", v)
		}
	}
	return nil
}

// loadEnv carga el archivo .env (si existe) y valida las variables requeridas.
func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to system environment variables.")
	}
	// Variables requeridas
	requiredVars := []string{"SERVER_ADDRESS", "SERVER_PORT"}
	return checkEnvVars(requiredVars...)
}

// Start es la función principal que inicia el servidor y configura las rutas.
// Aquí conectamos las diferentes partes de la aplicación (adaptadores primarios y secundarios).
func Start() {
	// Cargar y validar variables de entorno
	if err := loadEnv(); err != nil {
		log.Fatalf("Sanity check failed: %v", err)
	}

	// Continuar con la ejecución del programa
	log.Println("Sanity check passed. Starting the application.")

	// Leer el modo de ejecución de la variable de entorno
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode // Por defecto, usa modo debug
	}
	gin.SetMode(mode)

	// Gin es un framework que nos permite manejar solicitudes HTTP (Adaptador Primario).
	router := gin.Default()
	if mode == gin.ReleaseMode {
		router.SetTrustedProxies([]string{"192.168.0.1", "192.168.0.2"})
	} else {
		// En desarrollo, no confiamos en proxies externos
		router.SetTrustedProxies(nil)
	}

	dbClient := getDbClient()

	authRepository := domain.NewAuthRepository(dbClient)
	auH := AuthHandler{service.NewLoginService(authRepository, domain.GetRolePermissions())}

	log.Printf("Getting router ready ...")
	// Definimos una ruta HTTP GET para obtener todos los clientes.
	log.Printf("Route login ...")
	router.POST("/auth/login", auH.Login)
	//router.POST("/auth/register", auH.Register)
	log.Printf("Route Verify ...")
	router.GET("/auth/verify", auH.Verify)
	log.Printf("Route Refresh token ...")
	router.POST("/auth/refresh", auH.Refresh)

	// Ejecutar el servidor
	port := os.Getenv("SERVER_PORT")
	log.Printf("Starting server in %s mode on port %s...", gin.Mode(), port)
	if err := router.Run(port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func getDbClient() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Configuración de la conexión con PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error while connecting to the database: ", err)
		// logger.Error("Error while connecting to the database: " + err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
