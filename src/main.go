package main

import (
    "github.com/gofiber/fiber/v2"

    "fmt"
    "net/http"
    "github.com/gorilla/mux"

    "log"
    "github.com/joho/godotenv"
)

func loadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }

    username, exists := os.LookupEnv("GITHUB_USERNAME")

    if exists {
        fmt.Println(username)
    }
}

func runFiber() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello Motto")
    })

    app.Get("/books/:title/page/:page", func(c *fiber.Ctx) error {
        return c.SendString("Fiber Requested: book " + c.Params("title") + " on page " + c.Params("page"))
    })

    app.Listen(":3000")
}

// needs both net/http and gorilla/mux
// basically the same as net/http
func runMux() {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)

        fmt.Fprintf(w, "Mux Requested: book %s on page %s \n", vars["title"], vars["page"])
    })

    http.ListenAndServe(":3000", r)
}

func main() {
    runFiber()
    // runMux()
}

