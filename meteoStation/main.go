package main

import (
	"errors"
	"fmt"
)

const numStations = 10

type WeatherStation struct {
	temperature float64
}
type City struct {
	stations [numStations]WeatherStation
}

func (c *City) updateTemperature(stationIndex int, temperature float64) error {
	if stationIndex < 0 || stationIndex >= numStations {
		return errors.New("invalid station index")
	}

	c.stations[stationIndex].temperature = temperature
	return nil
}
func (c *City) calculateAverageTemperature() float64 {
	sum := 0.0
	for _, station := range c.stations {
		sum = station.temperature
	}
	return sum / float64(numStations)
}

func main() {
	city := City{}

	for {
		var stationIndex int
		var temperature float64

		fmt.Print("Введите номер метеостанции (0-9): ")
		_, err := fmt.Scanln(&stationIndex)
		if err != nil {
			fmt.Println("Ошибка ввода номера метеостанции: ", err)
			continue
		}

		fmt.Print("Введите температуру: ")
		_, err = fmt.Scanln(&temperature)
		if err != nil {
			fmt.Println("Ошибка ввода температуры: ", err)
			continue
		}

		err = city.updateTemperature(stationIndex, temperature)
		if err != nil {
			fmt.Println("Ошибка обновления температуры: ", err)
			continue
		}

		averageTemperature := city.calculateAverageTemperature()
		fmt.Printf("Средняя температура в городе: %.2f\n", averageTemperature)
	}
}
