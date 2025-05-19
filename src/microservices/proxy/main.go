package main

import (
    "io"
    "log"
    "math/rand"
    "net/http"
    "os"
    "strconv"
    "time"
)

var (
    monolithURL string
    moviesURL string
    eventsURL string
    usersURL string
    migrationPercent int
    gradualMigration bool
)

func init() {
    rand.Seed(time.Now().UnixNano())
    monolithURL = os.Getenv("MONOLITH_URL")
    moviesURL = os.Getenv("MOVIES_SERVICE_URL")
    eventsURL = os.Getenv("EVENTS_SERVICE_URL")
    usersURL = os.Getenv("USERS_SERVICE_URL") // Используем monolith как fallback
    migrationPercent, _ = strconv.Atoi(os.Getenv("MOVIES_MIGRATION_PERCENT"))
    gradualMigration, _ = strconv.ParseBool(os.Getenv("GRADUAL_MIGRATION"))
    log.Printf("Migration percent: %d", migrationPercent)
    log.Printf("Users URL: %s", usersURL)
}

func main() {
    // Добавляем новые маршруты
    http.HandleFunc("/health", handleHealthCheck)
    http.HandleFunc("/api/users", handleUsersRequest)
    http.HandleFunc("/api/movies", handleMoviesRequest)
    
    port := os.Getenv("PORT")
    log.Printf("Starting proxy service on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleMoviesRequest(w http.ResponseWriter, r *http.Request) {
    if gradualMigration {
        if rand.Intn(100) < migrationPercent {
            forwardRequest(w, r, moviesURL)
        } else {
            forwardRequest(w, r, monolithURL)
        }
    } else {
        forwardRequest(w, r, moviesURL)
    }
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

// Обработчик для пользователей
func handleUsersRequest(w http.ResponseWriter, r *http.Request) {
    if gradualMigration {
        if rand.Intn(100) < migrationPercent {
            forwardRequest(w, r, usersURL)
        } else {
            forwardRequest(w, r, monolithURL)
        }
    } else {
        forwardRequest(w, r, usersURL)
    }
}

func forwardRequest(w http.ResponseWriter, r *http.Request, targetURL string) {
    client := &http.Client{}
    req, err := http.NewRequest(r.Method, targetURL+r.URL.Path, r.Body)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    
    // правильное копирование заголовков
    for k, v := range r.Header {
        for _, value := range v {
            req.Header.Add(k, value)
        }
    }
    
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
        return
    }
    
    defer resp.Body.Close()
    
    // копируем ответ
    w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
    w.WriteHeader(resp.StatusCode)
    
    _, err = io.Copy(w, resp.Body)
    if err != nil {
        log.Printf("Ошибка при копировании ответа: %v", err)
    }
}
