package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// Создаем новый GET запрос с count больше фактического
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	// Создаем новый экземпляр ResponseRecorder для записи ответа
	responseRecorder := httptest.NewRecorder()
	//Преобразуем функцию mainHandle в объект, который можно использовать, как обработчик HTTP-запросов
	handler := http.HandlerFunc(mainHandle)
	//Вызываем метод ServeHTTP для обработчика handler, передаем в него responseRecorder для записи ответа и сам запрос
	handler.ServeHTTP(responseRecorder, req)
	//Преобразуем тело ответа в строку и разделяем строку на подстроки, получая слайс строк
	cafeList := strings.Split(responseRecorder.Body.String(), ",")

	//Сравниваем ожидаемый статус ответа с фактическим
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected HTTP status OK")

	//Сравниваем длину списка из тела ответа с ожидаемой длиной списка
	assert.Equal(t, totalCount, len(cafeList), "The list of cafes is not complete")

}

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK, "Expected HTTP status OK")
	assert.NotEmpty(t, responseRecorder.Body, "The response body contains no information")
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=NOCITY", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected HTTP status Bad Request")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Wrong city value error expected")
}
