package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pterm/pterm"
	"github.com/r523/football/internal/http/handler"
	"github.com/r523/football/internal/servo"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
)

const (
	ServoDutyNumerator   gpio.Duty = 1
	ServoDutyDenominator gpio.Duty = 20
	ServoFreq                      = 1 * physic.Hertz

	Offset = 10
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Foot", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("ball", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	if _, err := host.Init(); err != nil {
		pterm.Error.Printf("host initiation failed %s\n", err)

		return
	}

	ch := make(chan int)

	app := fiber.New()

	handler.Static(app)

	d := handler.Rotate{
		Channel: ch,
	}
	d.Register(app.Group("/api"))

	app.Use(logger.New())

	go func() {
		if err := app.Listen(":1378"); err != nil {
			pterm.Error.Printf("listen on port 1378 failed %s\n", err)
		}
	}()

	s := servo.New(rpi.P1_33, ServoDutyNumerator, ServoDutyDenominator, ServoFreq)

	go func(ch <-chan int) {
		for {
			angle := <-ch

			pterm.Info.Printf("change angle to %d\n", angle)

			if err := s.Start(); err != nil {
				pterm.Error.Printf("cannot start the servo %s", err)

				return
			}

			time.Sleep(time.Duration(Offset+angle/10) * time.Second)

			_ = s.Stop()
		}
	}(ch)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	pterm.Info.Printf("Bye!\n")
}
