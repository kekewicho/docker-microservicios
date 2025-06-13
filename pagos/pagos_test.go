package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPayments(t *testing.T) {
	// Crear archivo de prueba temporal
	// Crear request y recorder
	req, err := http.NewRequest("GET", "/records", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(rr, req)

	// Verificar código de estado
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar que el cuerpo tenga parte del JSON esperado
	assert.Contains(t, rr.Body.String(), `"transaction_id"`)
	assert.Contains(t, rr.Body.String(), `"amount"`)
}
