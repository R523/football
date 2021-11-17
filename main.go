package main

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/r523/football/internal/servo"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"
)

const (
	ServoDutyNumerator   gpio.Duty = 1
	ServoDutyDenominator gpio.Duty = 5
	ServoFreq                      = 1 * physic.MilliHertz
	ServoTimeout                   = 10 * time.Second
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

	s := servo.New(rpi.P1_33, ServoDutyNumerator, ServoDutyDenominator, ServoFreq)

	if err := s.Start(); err != nil {
		pterm.Error.Printf("cannot start the servo %s", err)

		return
	}

	time.Sleep(ServoTimeout)

	_ = s.Stop()
}
