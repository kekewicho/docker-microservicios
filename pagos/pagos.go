package main

import (
	"encoding/json" // Necesario para trabajar con JSON
	"fmt"
	"html"
	"log"
	"net/http"
	"os" // Necesario para leer archivos del sistema de ficheros

	"github.com/gorilla/mux"
)

// --- DEFINICIÓN DE ESTRUCTURAS DE DATOS ---
// Estructura para un solo registro de pago
type PaymentRecord struct {
	Amount          float64 `json:"amount"`           // monto del pago realizado
	TransactionType string  `json:"transaction_type"` // canal digital utilizado para realizar el pago
	Status          string  `json:"status"`           // estatus del pago
	CreationDate    string  `json:"creation_date"`    // fecha en la que se realizó el pago
	TransactionID   string  `json:"transaction_id"`   // identificador de la transacción
	Source          string  `json:"source"`           // origen del pago
}

// Estructura principal que contendrá la lista de PaymentRecord
// Ahora es un mapa donde la clave es el ID (string) y el valor es un array de PaymentRecord
type Payments map[string][]PaymentRecord

// --- FIN DE DEFINICIÓN DE ESTRUCTURAS DE DATOS ---

func main() {
	router := setupRouter()
	log.Fatal(http.ListenAndServe(":8003", router))
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/records", GetPayments)
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := os.ReadFile("/data/payment_records.json")
	if err != nil {
		log.Printf("Error al leer el archivo payment_records.json: %v", err)
		http.Error(w, "Error interno del servidor al leer los datos de pagos.", http.StatusInternalServerError)
		return
	}

	var paymentsData Payments // Usamos el nuevo tipo Payments (mapa)
	err = json.Unmarshal(fileBytes, &paymentsData)
	if err != nil {
		log.Printf("Error al deserializar el JSON de pagos: %v", err)
		http.Error(w, "Error interno del servidor al procesar los datos de pagos.", http.StatusInternalServerError)
		return
	}

	// Para retornar solo los registros de pago (aplanando el mapa)
	// Creamos un slice para almacenar todos los PaymentRecord de todos los IDs
	var allPaymentRecords []PaymentRecord
	for _, records := range paymentsData {
		allPaymentRecords = append(allPaymentRecords, records...)
	}

	// Serializamos el slice de PaymentRecord a JSON
	responseBytes, err := json.Marshal(allPaymentRecords)
	if err != nil {
		log.Printf("Error al serializar la respuesta JSON: %v", err)
		http.Error(w, "Error interno del servidor al serializar la respuesta de pagos.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(responseBytes) // Escribimos los bytes serializados
	if err != nil {
		log.Printf("Error al escribir la respuesta HTTP: %v", err)
	}

	log.Println("Solicitud /records procesada exitosamente.")
}