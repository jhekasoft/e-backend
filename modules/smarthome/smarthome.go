package smarthome

import (
	internalModels "e-backend/internal/models"
	"e-backend/modules/smarthome/models"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SmartHomeModule struct {
}

func (m *SmartHomeModule) Name() string {
	return "SmartHome"
}

func (m *SmartHomeModule) Run(c *internalModels.Core) error {
	c.DB.AutoMigrate(&models.SmartHomeSensorValue{})

	// repo := repository.NewRepository(c.DB)
	// services := service.NewService(repo)
	// h := handler.NewHandler(services)

	c.MQTT.Subscribe("smarthome/+/+/+/sensor/meteo", 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})

	return nil
}

func NewModule() internalModels.Module {
	return &SmartHomeModule{}
}
