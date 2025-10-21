package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LeituraSensor struct {
	NomeSensor string
	Valor      int
}

func iniciarSensor(nome string, canal chan<- LeituraSensor) {
	for {
		leituraAleatoria := rand.Intn(100)

		leitura := LeituraSensor{
			NomeSensor: nome,
			Valor:      leituraAleatoria,
		}

		canal <- leitura

		intervalo := time.Duration(rand.Intn(1000)+500) * time.Millisecond
		time.Sleep(intervalo)
	}
}

func main() {
	canalComum := make(chan LeituraSensor)

	go iniciarSensor("Temperatura", canalComum)
	go iniciarSensor("Pressão", canalComum)
	go iniciarSensor("Umidade", canalComum)

	fmt.Println("Monitor de sensores iniciado. Aguardando leituras no canal comum...")

	for {
		select {
		case leitura := <-canalComum:
			fmt.Printf("[SENSOR %s]: %d\n", leitura.NomeSensor, leitura.Valor)
		}
	}
}