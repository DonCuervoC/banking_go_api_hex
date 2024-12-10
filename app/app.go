package app

import (
	"log"
	"os"

	"github.com/DonCuervoC/banking_go_api_hex/domain"
	"github.com/DonCuervoC/banking_go_api_hex/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Start es la función principal que inicia el servidor y configura las rutas.
// Aquí conectamos las diferentes partes de la aplicación (adaptadores primarios y secundarios).
func Start() {

	// Cargar variables desde .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system environment variables.")
	}

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
	port := ":8000"
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
