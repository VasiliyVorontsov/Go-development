// package main объявляет, что этот файл является частью главного пакета,
// который можно скомпилировать и запустить как самостоятельную программу.
package main

// Импортируем необходимые пакеты из стандартной библиотеки Go.
import (
	"encoding/json" // Пакет для работы с JSON
	"fmt"           // Для работы с форматированным вводом и выводом (например, в консоль)
	"net/http"      // Это самый главный пакет для создания HTTP-серверов и клиентов
	"time"          // Для работы со временем (в нашем случае - для логирования)
)

// Объявляем структуры для JSON-ответов

// Структура для основного ответа API
type Response struct {
	Status  string      `json:"status"`  // Статус запроса (например, "success", "error")
	Message string      `json:"message"` // Сообщение для пользователя
	Data    interface{} `json:"data"`    // Данные (может быть любого типа)
}

// Структура для данных о себе
type AboutMeData struct {
	Name         string   `json:"name"`
	Role         string   `json:"role"`
	Experience   string   `json:"experience"`
	Technologies []string `json:"technologies"`
	Goal         string   `json:"goal"`
}

// Структура для данных о выборе Go
type WhyGoData struct {
	Reasons    []string `json:"reasons"`
	Expections string   `json:"expectations"`
	Note       string   `json:"note,omitempty"` // omitempty - поле будет пропущено, если пустое
}

// Функция main — это входная точка нашей программы.
func main() {
	// Регистрируем обработчик для корневого пути "/"
	http.HandleFunc("/", rootHandler)
	// Регистрируем обработчик для пути "/api/about-me"
	http.HandleFunc("/api/about-me", aboutMeHandler)
	// Регистрируем обработчик для пути "/api/why-go"
	http.HandleFunc("/api/why-go", whyGoHandler)

	// Выводим информационное сообщение в консоль
	fmt.Printf("%s Сервер запущен и слушает на порту :8080\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("-> Локальный доступ: http://localhost:8080")
	fmt.Println("-> Внешний доступ:   попросите ссылку у разработчика (Василий)")

	// Запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s\n", err)
	}
}

// Вспомогательная функция для отправки JSON-ответов
func sendJSONResponse(w http.ResponseWriter, statusCode int, status string, message string, data interface{}) {
	// Устанавливаем заголовок, указывающий, что возвращаем JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	// Создаем структуру ответа
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	// Кодируем структуру в JSON и отправляем
	json.NewEncoder(w).Encode(response)
}

// Обработчик для корневого URL "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Логируем обращение к корню
	fmt.Printf("%s Получен запрос на главную страницу от %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr)

	// Проверяем, если запрос не к корню, возвращаем 404
	if r.URL.Path != "/" {
		sendJSONResponse(w, http.StatusNotFound, "error", "Страница не найдена", nil)
		return
	}

	// Данные для корневой страницы
	rootData := map[string]interface{}{
		"endpoints": []string{
			"/api/about-me - Узнать обо мне",
			"/api/why-go - Узнать, почему я изучаю Go",
		},
		"hint": "Добавьте эти пути к URL в адресной строке",
	}

	// Если запрос к корню, выводим приветствие и инструкцию
	sendJSONResponse(w, http.StatusOK, "success", "Добро пожаловать на мой первый веб-сервер на Go!", rootData)
}

// Функция-обработчик для эндпоинта '/api/about-me'.
func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s Получен запрос на /api/about-me от %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr)

	// Создаем данные в структурированном виде
	data := AboutMeData{
		Name:       "Василий",
		Role:       "начинающий backend-разработчик",
		Experience: "не очень большой, но я уже изучил основы Python, С++, написал маленький проект на PHP",
		Technologies: []string{
			"Python",
			"C++",
			"PHP",
			"Go (изучаю)",
		},
		Goal: "сейчас я начинаю изучать Go с нуля, кстати, сейчас ты тестируешь мой простой веб-сервер",
	}

	sendJSONResponse(w, http.StatusOK, "success", "Информация о разработчике", data)
}

// Функция-обработчик для эндпоинта '/api/why-go'.
func whyGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s Получен запрос на /api/why-go от %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr)

	// Создаем данные в структурированном виде
	data := WhyGoData{
		Reasons: []string{
			"захотел выучить что-то новое",
			"сейчас данный язык популярен",
		},
		Expections: "получить структурированное понимание языка, получить базу, после чего вступить в проект",
		Note:       "Получить работу тоже было бы неплохо",
	}

	sendJSONResponse(w, http.StatusOK, "success", "Причины изучения Go", data)
}
