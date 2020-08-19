package errors

import "fmt"

type ArgError struct {
	System           string `json:"system"`
	Status           int    `json:"status"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

var (
	InvalidPhoneNumber = &ArgError{"Alfa", 422, "Неправильный номер телефона", ""}
	//если во время валидации номера телефона будет ошибка
	InvalidCharacter    = &ArgError{"Alfa", 400, "Неправильные входные данные. Неправильный JSON", ""}
	ElasticConnectError = &ArgError{"Alfa", 503, "Сервис недоступен, недоступен движок поиска", ""}
	NoFound             = &ArgError{"Alfa", 404, "Ресурс не найден", ""}
	AccessDenied        = &ArgError{"Alfa", 403, "Доступ к ресурсу запрещен",""}
)

func (e *ArgError) Error() string {
	return fmt.Sprintf("%s %s", e.Status, e.Message)
}
