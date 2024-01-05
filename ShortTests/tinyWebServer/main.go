package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const prodName string = "WebServer info"

// const prodVersion string = "v1.0.214214205"
// const prodVersion string = "v2.0.214214102"
const prodVersion string = "v3.0.214107102"

var productInfo string
var hostName string

func main() {

	productInfo = fmt.Sprintf("%s (%s)\n", prodName, prodVersion)
	hostName, _ = os.Hostname()
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr).IP

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		// ServerHeader:     productInfo,
		// AppName:          prodName,
		// DisableKeepalive: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {

		// currentTime := time.Now().Format("2006/01/02 15:04:05.000000")

		return c.Render("index", fiber.Map{
			"Title":     productInfo,
			"product":   prodName,
			"version":   prodVersion,
			"hostName":  hostName,
			"ipAddress": ipAddress,
			"dateTime":  time.Now().Format("2006/01/02 15:04:05.000000"),
		})
	})

	log.Fatal(app.Listen(":8080"))

}
