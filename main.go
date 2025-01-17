package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// SensorData represents the structure of the sensor data.
type SensorData struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Pressure    float64 `json:"pressure"`
}

// Global variable
var data SensorData

// MessageHandler processes incoming MQTT messages.
var messageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	if err := json.Unmarshal(msg.Payload(), &data); err != nil {
		log.Printf("Error parsing JSON from topic '%s': %v\n", msg.Topic(), err)
		return
	}
	fmt.Printf("Received on topic '%s': Temperature: %.2f°C, Humidity: %.2f%%, Pressure: %.2f hPa\n",
		msg.Topic(), data.Temperature, data.Humidity, data.Pressure)
}

// Function to connect to the MQTT broker with retries.
func connectToBroker(client MQTT.Client) error {
	for i := 0; i < 5; i++ { // Retry up to 5 times
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Printf("Connection attempt %d failed: %v\n", i+1, token.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		return nil // Connection successful
	}
	return fmt.Errorf("failed to connect to the broker after 5 attempts")
}

// Function to subscribe with retries.
func subscribeToTopic(client MQTT.Client, topic string) error {
	for i := 0; i < 5; i++ {
		if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
			log.Printf("Subscription attempt %d to topic '%s' failed: %v\n", i+1, topic, token.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		return nil // Subscription successful
	}
	return fmt.Errorf("failed to subscribe to topic '%s' after 5 attempts", topic)
}

// Round to the nearest thousandth.
func roundToThousandth(value float64) float64 {
	return math.Round(value*1000) / 1000
}

func main() {
	// MQTT broker configuration.
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("golang_mqtt_client")
	opts.SetDefaultPublishHandler(messageHandler)
	opts.SetConnectTimeout(10 * time.Second)
	opts.SetKeepAlive(60 * time.Second)

	client := MQTT.NewClient(opts)

	// Connect to the broker.
	err := connectToBroker(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MQTT broker.")

	// Topics for machines and their sensors, including "sensor/data".
	topics := []string{
		"IIoT/machine_1/temperature", "IIoT/machine_1/humidity", "IIoT/machine_1/pressure",
		"IIoT/machine_2/temperature", "IIoT/machine_2/humidity", "IIoT/machine_2/pressure",
		"IIoT/machine_3/temperature", "IIoT/machine_3/humidity", "IIoT/machine_3/pressure",
		"IIoT/machine_4/temperature", "IIoT/machine_4/humidity", "IIoT/machine_4/pressure",
		"IIoT/machine_5/temperature", "IIoT/machine_5/humidity", "IIoT/machine_5/pressure",
		"sensor/data", // Topic for publishing data.
	}

	// Subscribe to each topic except "sensor/data" (for publishing).
	for _, topic := range topics[:len(topics)-1] {
		if err := subscribeToTopic(client, topic); err != nil {
			log.Fatalf("Failed to subscribe to topic '%s': %v", topic, err)
		}
		fmt.Printf("Subscribed to topic '%s'\n", topic)
	}

	// Random number generator for sensor data simulation.
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Simulate publishing data for each machine and "sensor/data".
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				for i := 1; i <= 5; i++ {
					// Machine-specific trend for sensor data
					temperature := 25.0 + 5.0*math.Sin(float64(i)*math.Pi/12) + random.Float64()*2.0 - 1.0
					humidity := 50.0 + 10.0*math.Cos(float64(i)*math.Pi/24) + random.Float64()*5.0 - 2.5
					pressure := 1015.0 + 5.0*math.Sin(float64(i)*math.Pi/12) + random.Float64()*3.0 - 1.5
	
					// More randomness!!! More variability!!!
					temperature += random.Float64()*1.0 - 0.5
					humidity += random.Float64()*2.0 - 1.0
					pressure += random.Float64()*1.0 - 0.5
	
					// Ensure values stay within realistic bounds:
					temperature = math.Min(math.Max(temperature, 18.0), 30.0)
					humidity = math.Min(math.Max(humidity, 40.0), 60.0)
					pressure = math.Min(math.Max(pressure, 1005.0), 1025.0)

					// Publish to respective machine topics.
					data := SensorData{
						Temperature: roundToThousandth(temperature),
						Humidity:    roundToThousandth(humidity),
						Pressure:    roundToThousandth(pressure),
					}

					topics := []string{
						fmt.Sprintf("IIoT/machine_%d/temperature", i),
						fmt.Sprintf("IIoT/machine_%d/humidity", i),
						fmt.Sprintf("IIoT/machine_%d/pressure", i),
					}

					for _, topic := range topics {
						payload, err := json.Marshal(data)
						if err != nil {
							log.Printf("Error encoding JSON: %v\n", err)
							continue
						}

						token := client.Publish(topic, 0, false, payload)
						token.Wait()

						fmt.Printf("Published to topic '%s': %s\n", topic, payload)
					}
				}

				// Publish data to the "sensor/data" topic.
				payload, err := json.Marshal(data)
				if err != nil {
					log.Printf("Error encoding JSON: %v\n", err)
					continue
				}
				token := client.Publish("sensor/data", 0, false, payload)
				token.Wait()

				fmt.Printf("Published to topic 'sensor/data': %s\n", payload)

				// Wait for 1 second.
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Graceful shutdown.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	fmt.Println("\nShutting down gracefully...")
	close(stop)
	client.Disconnect(250)
	fmt.Println("MQTT client disconnected. Goodbye!") 
}
