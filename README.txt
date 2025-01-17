(* # IoT MQTT Simulation with Clients and Mosquitto Broker
Developed an MQTT-based simulation framework in Go that generates and streams synthetic sensor data for virtual clients across distinct topics, enabling realistic testing and evaluation of OT systems.

# Overview
This project demonstrates the use of MQTT (Message Queuing Telemetry Transport) for simulating IoT (Internet of Things) clients that communicate through a Mosquitto broker. The simulation involves five IoT clients, each publishing and subscribing to specific topics in the MQTT network.

The IoT clients publish data like sensor readings and device statuses and subscribe to topics to receive updates. The Mosquitto broker facilitates message exchange between these clients. The system simulates realistic network traffic, ideal for IoT experiments, testing, and understanding MQTT's role in smart systems.

# Key Features:
1. Sensor Data Structure: The SensorData struct represents the temperature, humidity, and pressure data. It is used for both publishing and receiving sensor data.

2. MQTT Client:

The program connects to an MQTT broker (at tcp://localhost:1883) with a client ID (golang_mqtt_client).
Subscription occurs for individual sensor topics for each machine (e.g., "IIoT/machine_1/temperature") and a general "sensor/data" topic.

3. Simulated Sensor Data:

Random values are generated using mathematical functions (like sine and cosine) to simulate sensor readings.
Each machine’s sensor data has a machine-specific trend, enhanced with additional randomness to ensure variability.
The data is constrained within realistic bounds for temperature (18°C to 30°C), humidity (40% to 60%), and pressure (1005 hPa to 1025 hPa).

4. Data Publishing:

The generated sensor data is published to the corresponding MQTT topics for each machine (e.g., IIoT/machine_1/temperature).
The data is also published periodically to a general sensor/data topic to simulate aggregated sensor information.
Message Handling:

The messageHandler processes incoming MQTT messages by unmarshalling the payload into the SensorData structure. The data is then printed to the console for logging and analysis.

5. Retries and Error Handling:

The program includes retry logic for connecting to the MQTT broker and subscribing to topics. It will retry up to five times if any operation fails.
Each MQTT message payload is expected to be in JSON format, and error handling is in place to ensure smooth operation.
Graceful Shutdown:

The program listens for OS termination signals (SIGINT, SIGTERM) to shut down gracefully, ensuring that any ongoing operations are cleanly stopped, and the MQTT client is properly disconnected.


# Visual Representation
Below is a diagram illustrating how MQTT works with the 5 IoT clients and the Mosquitto broker:

                             +---------------------+
                             | Ignition SCADA /    |
                             | Mosquitto Broker    |
                             +---------------------+
                                      |
  --------------------------------------------------------------------------------
  |                  |                |                  |                 |   
+--------+      +--------+        +--------+         +--------+         +--------+
| Client 1|<--->| Client 2|<---->| Client 3|<----->| Client 4|<--->| Client 5|
+--------+      +--------+        +--------+         +--------+         +--------+
    |                |                |                 |                 |
  Subscribes to     Subscribes to     Subscribes to      Subscribes to      Subscribes to
   IIoT/machine_1/* IIoT/machine_2/* IIoT/machine_3/*  IIoT/machine_4/*    IIoT/machine_5/*
   Publishes to     Publishes to     Publishes to       Publishes to        Publishes to
   IIoT/machine_1/* IIoT/machine_2/* IIoT/machine_3/*  IIoT/machine_4/*    IIoT/machine_5/*
   (Temperature,    (Temperature,    (Temperature,      (Temperature,        (Temperature,
   Humidity, Pressure) Humidity, Pressure) Humidity, Pressure) Humidity, Pressure) Humidity, Pressure)
   Also subscribes to 'sensor/data' for collective communication with all clients.



# Explanation of the Diagram:

Mosquitto Broker / Ignition SCADA: Central hub responsible for managing message exchanges.
Clients: Represent machines that both publish and subscribe to topics. They also interact via the sensor/data topic to exchange collective information.
Each Client subscribes to and publishes data specific to their machine (e.g., IIoT/machine_1/temperature for Client 1). They also interact with all other machines by subscribing to the collective sensor/data topic and publishing their own updates to it.



# Instructions for Running the Project
Mosquitto Broker: You need to have the Mosquitto broker installed to simulate the MQTT server.

Install Mosquitto:

On Ubuntu/Debian:
sudo apt-get install mosquitto
sudo apt-get install mosquitto-clients

On macOS (using Homebrew):
brew install mosquitto

Running the Mosquitto Broker
Once you have installed Mosquitto, you can start the broker with the following command (ensure port 1883 is open):

On Ubuntu/Debian:
mosquitto -v 
The -v option enables verbose output so you can see all incoming and outgoing messages in the terminal.

on macOS:
brew services start mosquitto
to stop: brew services stop mosquitto
 *)
