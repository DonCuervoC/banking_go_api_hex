package app

import (
	"fmt"
	"log"
	"os"

	"github.com/DonCuervoC/banking_go_api_hex/domain"
	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	// Configurar proxies confiables (nil para confiar en todos en desarrollo)
	// Ejemplo para producción, ajusta a tus necesidades:
	if mode == gin.ReleaseMode {
		// Definir una lista de proxies confiables, si es necesario
		router.SetTrustedProxies([]string{"192.168.0.1", "192.168.0.2"})
	} else {
		// En desarrollo, no confiamos en proxies externos
		router.SetTrustedProxies(nil)
	}

	//4.
	// Inyección de dependencias:
	// Creamos un "repositorio" (CustomerRepositoryStub), que es un adaptador secundario, y lo inyectamos en el servicio.
	// Luego pasamos ese servicio al controlador (CustomerHandlers), que maneja las solicitudes HTTP.
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Definimos una ruta HTTP GET para obtener todos los clientes.
	// Esta ruta usa la función `getAllCustomer` del controlador `CustomerHandlers`.
	router.GET("/customers", ch.getAllCustomer)
	router.GET("/customer/:customer_id", ch.getCustomer)

	// Ejecutar el servidor
	port := os.Getenv("SERVER_PORT")
	log.Printf("Starting server in %s mode on port %s...", gin.Mode(), port)
	if err := router.Run(port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func Start01() {

	// Gin es un framework que nos permite manejar solicitudes HTTP (Adaptador Primario).
	router := gin.Default()
	// Configurar proxies confiables (nil para confiar en todos en desarrollo)
	//router.SetTrustedProxies(nil)

	//4.
	// Inyección de dependencias:
	// Creamos un "repositorio" (CustomerRepositoryStub), que es un adaptador secundario, y lo inyectamos en el servicio.
	// Luego pasamos ese servicio al controlador (CustomerHandlers), que maneja las solicitudes HTTP.
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Definimos una ruta HTTP GET para obtener todos los clientes.
	// Esta ruta usa la función `getAllCustomer` del controlador `CustomerHandlers`.
	router.GET("/customers", ch.getAllCustomer)
	router.GET("/customer/:customer_id", ch.getCustomer)

	// Ejecutar el servidor
	port := ":8000"
	log.Printf("Starting server in %s mode on port %s...", gin.Mode(), port)
	if err := router.Run(port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
