# IoT MQTT Simulation with Clients and Mosquitto Broker
Developed an MQTT-based simulation framework in Go that generates and streams synthetic sensor data for virtual clients across three distinct topics (Temperature, Humidity, and Pressure), enabling realistic testing and evaluation of OT systems.

# Overview
This project demonstrates the use of MQTT (Message Queuing Telemetry Transport) for simulating IoT (Internet of Things) clients that communicate through a Mosquitto broker. The simulation involves five IoT clients, each publishing and subscribing to specific topics in the MQTT network.

The IoT clients publish data like sensor readings and device statuses and subscribe to topics to receive updates. The Mosquitto broker facilitates message exchange between these clients. The system simulates realistic network traffic, ideal for IoT experiments, testing, and understanding MQTT's role in smart systems.

# Visual Representation
Below is a diagram illustrating how MQTT works with the 5 IoT clients, the Mosquitto broker, and their respective topics:



# Explanation of the Diagram:

Mosquitto Broker: Acts as the central hub for MQTT communication, receiving and dispatching messages to clients.
IoT Clients: Five clients are connected to the Mosquitto broker, each subscribing to relevant topics and publishing data.
Topics: Each client subscribes to different topics (e.g., temperature, humidity, status) and publishes data to other topics. Topics allow selective messaging.
Client 1 subscribes to temperature/+/status and publishes to temperature/+/data.
Client 2 subscribes to humidity/+/data and publishes to humidity/+/status.
Client 3 subscribes to status/+/data and publishes to status/+/command.
Client 4 subscribes to +/status/data and publishes to +/command.
Client 5 subscribes to +/command/data and publishes status updates to +/status.



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




# Running the IoT Clients (optional):
Each client simulates an IoT device that connects to the broker and interacts with different topics. You can run the clients one at a time or in parallel.

Client 1 (Publisher and Subscriber):

Publishes data to temperature/+/data.
Subscribes to temperature/+/status.
Client 2 (Publisher and Subscriber):

Publishes data to humidity/+/data.
Subscribes to humidity/+/status.
Client 3 (Publisher and Subscriber):

Publishes data to status/+/command.
Subscribes to status/+/data.
Client 4 (Publisher and Subscriber):

Publishes data to +/command.
Subscribes to +/status/data.
Client 5 (Publisher and Subscriber):

Publishes data to +/status.
Subscribes to +/command/data.
Run the clients as individual scripts or in parallel, ensuring each client runs on its own terminal window or in the background. You can start each client with the following command:

bash
Copy code
python client_1.py
python client_2.py
python client_3.py
python client_4.py
python client_5.py

