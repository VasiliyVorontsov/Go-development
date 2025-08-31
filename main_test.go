// package main объявляет, что этот файл является частью главного пакета,
// который можно скомпилировать и запустить как самостоятельную программу.
package main

// Импортируем необходимые пакеты из стандартной библиотеки Go.
import (
	"fmt"      // Для работы с форматированным вводом и выводом (например, в консоль)
	"net/http" // Это самый главный пакет для создания HTTP-серверов и клиентов
	"time"     // Для работы со временем (в нашем случае - для логирования)
)

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

// Обработчик для корневого URL "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Логируем обращение к корню
	fmt.Printf("%s Получен запрос на главную страницу от %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr)

	// Устанавливаем заголовок, указывающий, что возвращаем простой текст
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Проверяем, если запрос не к корню, возвращаем 404
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound) // Status Code 404
		w.Write([]byte("404 - Страница не найдена\n"))
		return
	}

	// Если запрос к корню, выводим приветствие и инструкцию
	w.Write([]byte("Добро пожаловать на мой первый веб-сервер на Go!\n\n"))
	w.Write([]byte("Для тестирования перейдите по одному из доступных эндпоинтов:\n\n"))
	w.Write([]byte("• /api/about-me    - Узнать обо мне\n"))
	w.Write([]byte("• /api/why-go     - Узнать, почему я изучаю Go\n\n"))
	w.Write([]byte("Подсказка: добавьте эти пути к URL в адресной строке выше."))
}

// Функция-обработчик для эндпоинта '/api/about-me'.
func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s Получен запрос на /api/about-me от %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Привет, меня зовут Василий. Я начинающий backend-разработчик.\n"))
	w.Write([]byte("Мой опыт в программировании не очень большой, но я уже изучил основы Python, С++, написал маленький проект на PHP.\n"))
	w.Write([]byte("Сейчас я начинаю изучать Go с нуля, кстати, сейчас ты тестируешь мой простой веб-сервер.\n"))
}

// Функция-обработчик для эндпоинта '/api/why-go'.
func whyGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s Получен запрос на /api/why-go от %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Я изучаю Go, потому что захотел выучить что-то новое. Сейчас данный язык популярен, поэтому выбор пал на него.\n"))
	w.Write([]byte("От факультатива я ожидаю получить структурированное понимание языка, получить базу, после чего вступить в проект!\n"))
	w.Write([]byte("P.S. Получить работу тоже было бы неплохо.."))
}
