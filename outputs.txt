# Here's what is outputted if your mosquitto broker is not running in the background (attempts connection 5 times with a 2 second break in between each time):

[Running] go run "(your personal path to )/main.go"
2025/01/16 23:44:34 Connection attempt 1 failed: network Error : dial tcp [::1]:1883: connect: connection refused
2025/01/16 23:44:36 Connection attempt 2 failed: network Error : dial tcp [::1]:1883: connect: connection refused
2025/01/16 23:44:38 Connection attempt 3 failed: network Error : dial tcp [::1]:1883: connect: connection refused
2025/01/16 23:44:40 Connection attempt 4 failed: network Error : dial tcp [::1]:1883: connect: connection refused
2025/01/16 23:44:42 Connection attempt 5 failed: network Error : dial tcp [::1]:1883: connect: connection refused
2025/01/16 23:44:44 failed to connect to the broker after 5 attempts
exit status 1


# You have to start mosquitto (I use a macOS so I have to use homebrew):

(base) ____@MacBook-Pro ~ % brew services start mosquitto
==> Successfully started `mosquitto` (label: homebrew.mxcl.mosquitto)


# Then, you can run the program (ensure you are in the correct directory & run: go run main.go):

[Running] go run "(your personal path to )/main.go"
Connected to MQTT broker.
Subscribed to topic 'IIoT/machine_1/temperature'
Subscribed to topic 'IIoT/machine_1/humidity'
Subscribed to topic 'IIoT/machine_1/pressure'
Subscribed to topic 'IIoT/machine_2/temperature'
Subscribed to topic 'IIoT/machine_2/humidity'
Subscribed to topic 'IIoT/machine_2/pressure'
Subscribed to topic 'IIoT/machine_3/temperature'
Subscribed to topic 'IIoT/machine_3/humidity'
Subscribed to topic 'IIoT/machine_3/pressure'
Subscribed to topic 'IIoT/machine_4/temperature'
Subscribed to topic 'IIoT/machine_4/humidity'
Subscribed to topic 'IIoT/machine_4/pressure'
Subscribed to topic 'IIoT/machine_5/temperature'
Subscribed to topic 'IIoT/machine_5/humidity'
Subscribed to topic 'IIoT/machine_5/pressure'
Published to topic 'IIoT/machine_1/temperature': {"temperature":27.389,"humidity":60,"pressure":1017.125}
Published to topic 'IIoT/machine_1/humidity': {"temperature":27.389,"humidity":60,"pressure":1017.125}
Published to topic 'IIoT/machine_1/pressure': {"temperature":27.389,"humidity":60,"pressure":1017.125}
Published to topic 'IIoT/machine_2/temperature': {"temperature":26.412,"humidity":60,"pressure":1017.286}
Published to topic 'IIoT/machine_2/humidity': {"temperature":26.412,"humidity":60,"pressure":1017.286}
Published to topic 'IIoT/machine_2/pressure': {"temperature":26.412,"humidity":60,"pressure":1017.286}
Published to topic 'IIoT/machine_3/temperature': {"temperature":28.205,"humidity":58.881,"pressure":1018.435}
Published to topic 'IIoT/machine_3/humidity': {"temperature":28.205,"humidity":58.881,"pressure":1018.435}
Published to topic 'IIoT/machine_3/pressure': {"temperature":28.205,"humidity":58.881,"pressure":1018.435}
Published to topic 'IIoT/machine_4/temperature': {"temperature":28.9,"humidity":57.597,"pressure":1020.658}
Published to topic 'IIoT/machine_4/humidity': {"temperature":28.9,"humidity":57.597,"pressure":1020.658}
Published to topic 'IIoT/machine_4/pressure': {"temperature":28.9,"humidity":57.597,"pressure":1020.658}
Published to topic 'IIoT/machine_5/temperature': {"temperature":29.3,"humidity":57.549,"pressure":1020.761}
Published to topic 'IIoT/machine_5/humidity': {"temperature":29.3,"humidity":57.549,"pressure":1020.761}
Published to topic 'IIoT/machine_5/pressure': {"temperature":29.3,"humidity":57.549,"pressure":1020.761}
Published to topic 'sensor/data': {"temperature":28.138,"humidity":59.793,"pressure":1016.754}
Received on topic 'IIoT/machine_1/temperature': Temperature: 27.39°C, Humidity: 60.00%, Pressure: 1017.12 hPa
Received on topic 'IIoT/machine_1/humidity': Temperature: 27.39°C, Humidity: 60.00%, Pressure: 1017.12 hPa
Received on topic 'IIoT/machine_1/pressure': Temperature: 27.39°C, Humidity: 60.00%, Pressure: 1017.12 hPa
Received on topic 'IIoT/machine_2/temperature': Temperature: 26.41°C, Humidity: 60.00%, Pressure: 1017.29 hPa
Received on topic 'IIoT/machine_2/humidity': Temperature: 26.41°C, Humidity: 60.00%, Pressure: 1017.29 hPa
Received on topic 'IIoT/machine_2/pressure': Temperature: 26.41°C, Humidity: 60.00%, Pressure: 1017.29 hPa
Received on topic 'IIoT/machine_3/temperature': Temperature: 28.20°C, Humidity: 58.88%, Pressure: 1018.43 hPa
Received on topic 'IIoT/machine_3/humidity': Temperature: 28.20°C, Humidity: 58.88%, Pressure: 1018.43 hPa
Received on topic 'IIoT/machine_3/pressure': Temperature: 28.20°C, Humidity: 58.88%, Pressure: 1018.43 hPa
Received on topic 'IIoT/machine_4/temperature': Temperature: 28.90°C, Humidity: 57.60%, Pressure: 1020.66 hPa
Received on topic 'IIoT/machine_4/humidity': Temperature: 28.90°C, Humidity: 57.60%, Pressure: 1020.66 hPa
Received on topic 'IIoT/machine_4/pressure': Temperature: 28.90°C, Humidity: 57.60%, Pressure: 1020.66 hPa
Received on topic 'IIoT/machine_5/temperature': Temperature: 29.30°C, Humidity: 57.55%, Pressure: 1020.76 hPa
Received on topic 'IIoT/machine_5/humidity': Temperature: 29.30°C, Humidity: 57.55%, Pressure: 1020.76 hPa
Received on topic 'IIoT/machine_5/pressure': Temperature: 29.30°C, Humidity: 57.55%, Pressure: 1020.76 hPa
Published to topic 'IIoT/machine_4/humidity': {"temperature":28.87,"humidity":59.11,"pressure":1018.124}
Received on topic 'IIoT/machine_2/temperature': Temperature: 28.39°C, Humidity: 59.19%, Pressure: 1018.95 hPa
Received on topic 'IIoT/machine_2/humidity': Temperature: 28.39°C, Humidity: 59.19%, Pressure: 1018.95 hPa
Published to topic 'IIoT/machine_4/pressure': {"temperature":28.87,"humidity":59.11,"pressure":1018.124}
Received on topic 'IIoT/machine_2/pressure': Temperature: 28.39°C, Humidity: 59.19%, Pressure: 1018.95 hPa
Received on topic 'IIoT/machine_3/temperature': Temperature: 28.97°C, Humidity: 60.00%, Pressure: 1017.42 hPa
Published to topic 'IIoT/machine_5/temperature': {"temperature":29.17,"humidity":59.228,"pressure":1018.784}
Received on topic 'IIoT/machine_3/humidity': Temperature: 28.97°C, Humidity: 60.00%, Pressure: 1017.42 hPa
Published to topic 'IIoT/machine_5/humidity': {"temperature":29.17,"humidity":59.228,"pressure":1018.784}
Published to topic 'IIoT/machine_5/pressure': {"temperature":29.17,"humidity":59.228,"pressure":1018.784}
Received on topic 'IIoT/machine_3/pressure': Temperature: 28.97°C, Humidity: 60.00%, Pressure: 1017.42 hPa
Received on topic 'IIoT/machine_4/temperature': Temperature: 28.87°C, Humidity: 59.11%, Pressure: 1018.12 hPa
Received on topic 'IIoT/machine_4/humidity': Temperature: 28.87°C, Humidity: 59.11%, Pressure: 1018.12 hPa
Received on topic 'IIoT/machine_4/pressure': Temperature: 28.87°C, Humidity: 59.11%, Pressure: 1018.12 hPa
Received on topic 'IIoT/machine_5/temperature': Temperature: 29.17°C, Humidity: 59.23%, Pressure: 1018.78 hPa
Published to topic 'sensor/data': {"temperature":28.968,"humidity":60,"pressure":1017.417}
Received on topic 'IIoT/machine_5/humidity': Temperature: 29.17°C, Humidity: 59.23%, Pressure: 1018.78 hPa
Received on topic 'IIoT/machine_5/pressure': Temperature: 29.17°C, Humidity: 59.23%, Pressure: 1018.78 hPa
Published to topic 'IIoT/machine_1/temperature': {"temperature":26.438,"humidity":58.834,"pressure":1016.73}
Published to topic 'IIoT/machine_1/humidity': {"temperature":26.438,"humidity":58.834,"pressure":1016.73}
Published to topic 'IIoT/machine_1/pressure': {"temperature":26.438,"humidity":58.834,"pressure":1016.73}
Published to topic 'IIoT/machine_2/temperature': {"temperature":28.156,"humidity":59.126,"pressure":1017.288}
Published to topic 'IIoT/machine_2/humidity': {"temperature":28.156,"humidity":59.126,"pressure":1017.288}
Published to topic 'IIoT/machine_2/pressure': {"temperature":28.156,"humidity":59.126,"pressure":1017.288}
Published to topic 'IIoT/machine_3/temperature': {"temperature":28.653,"humidity":57.774,"pressure":1019.361}
Published to topic 'IIoT/machine_3/humidity': {"temperature":28.653,"humidity":57.774,"pressure":1019.361}
Published to topic 'IIoT/machine_3/pressure': {"temperature":28.653,"humidity":57.774,"pressure":1019.361}
Received on topic 'IIoT/machine_1/temperature': Temperature: 26.44°C, Humidity: 58.83%, Pressure: 1016.73 hPa
Received on topic 'IIoT/machine_1/humidity': Temperature: 26.44°C, Humidity: 58.83%, Pressure: 1016.73 hPa
Received on topic 'IIoT/machine_1/pressure': Temperature: 26.44°C, Humidity: 58.83%, Pressure: 1016.73 hPa
Received on topic 'IIoT/machine_2/temperature': Temperature: 28.16°C, Humidity: 59.13%, Pressure: 1017.29 hPa
Published to topic 'IIoT/machine_4/temperature': {"temperature":29.831,"humidity":57.918,"pressure":1018.646}
Received on topic 'IIoT/machine_2/humidity': Temperature: 28.16°C, Humidity: 59.13%, Pressure: 1017.29 hPa
Published to topic 'IIoT/machine_4/humidity': {"temperature":29.831,"humidity":57.918,"pressure":1018.646}
Received on topic 'IIoT/machine_2/pressure': Temperature: 28.16°C, Humidity: 59.13%, Pressure: 1017.29 hPa
Received on topic 'IIoT/machine_3/temperature': Temperature: 28.65°C, Humidity: 57.77%, Pressure: 1019.36 hPa
Published to topic 'IIoT/machine_4/pressure': {"temperature":29.831,"humidity":57.918,"pressure":1018.646}
Received on topic 'IIoT/machine_3/humidity': Temperature: 28.65°C, Humidity: 57.77%, Pressure: 1019.36 hPa
Published to topic 'IIoT/machine_5/temperature': {"temperature":29.234,"humidity":59.674,"pressure":1018.732}
Received on topic 'IIoT/machine_3/pressure': Temperature: 28.65°C, Humidity: 57.77%, Pressure: 1019.36 hPa
Received on topic 'IIoT/machine_4/temperature': Temperature: 29.83°C, Humidity: 57.92%, Pressure: 1018.65 hPa
Published to topic 'IIoT/machine_5/humidity': {"temperature":29.234,"humidity":59.674,"pressure":1018.732}
Received on topic 'IIoT/machine_4/humidity': Temperature: 29.83°C, Humidity: 57.92%, Pressure: 1018.65 hPa
Published to topic 'IIoT/machine_5/pressure': {"temperature":29.234,"humidity":59.674,"pressure":1018.732}
Received on topic 'IIoT/machine_4/pressure': Temperature: 29.83°C, Humidity: 57.92%, Pressure: 1018.65 hPa
Published to topic 'sensor/data': {"temperature":29.831,"humidity":57.918,"pressure":1018.646}
Received on topic 'IIoT/machine_5/temperature': Temperature: 29.23°C, Humidity: 59.67%, Pressure: 1018.73 hPa
Received on topic 'IIoT/machine_5/humidity': Temperature: 29.23°C, Humidity: 59.67%, Pressure: 1018.73 hPa
Received on topic 'IIoT/machine_5/pressure': Temperature: 29.23°C, Humidity: 59.67%, Pressure: 1018.73 hPa
Published to topic 'IIoT/machine_1/temperature': {"temperature":26.099,"humidity":56.784,"pressure":1014.954}
Published to topic 'IIoT/machine_1/humidity': {"temperature":26.099,"humidity":56.784,"pressure":1014.954}
Published to topic 'IIoT/machine_1/pressure': {"temperature":26.099,"humidity":56.784,"pressure":1014.954}
Published to topic 'IIoT/machine_2/temperature': {"temperature":28.057,"humidity":60,"pressure":1018.321}
Published to topic 'IIoT/machine_2/humidity': {"temperature":28.057,"humidity":60,"pressure":1018.321}
Published to topic 'IIoT/machine_2/pressure': {"temperature":28.057,"humidity":60,"pressure":1018.321}
Received on topic 'IIoT/machine_1/temperature': Temperature: 26.10°C, Humidity: 56.78%, Pressure: 1014.95 hPa
Published to topic 'IIoT/machine_3/temperature': {"temperature":28.864,"humidity":60,"pressure":1019.302}
Published to topic 'IIoT/machine_3/humidity': {"temperature":28.864,"humidity":60,"pressure":1019.302}
Published to topic 'IIoT/machine_3/pressure': {"temperature":28.864,"humidity":60,"pressure":1019.302}
Published to topic 'IIoT/machine_4/temperature': {"temperature":29.357,"humidity":58.731,"pressure":1018.208}
Published to topic 'IIoT/machine_4/humidity': {"temperature":29.357,"humidity":58.731,"pressure":1018.208}
Received on topic 'IIoT/machine_1/humidity': Temperature: 26.10°C, Humidity: 56.78%, Pressure: 1014.95 hPa
Published to topic 'IIoT/machine_4/pressure': {"temperature":29.357,"humidity":58.731,"pressure":1018.208}
Received on topic 'IIoT/machine_1/pressure': Temperature: 26.10°C, Humidity: 56.78%, Pressure: 1014.95 hPa
Published to topic 'IIoT/machine_5/temperature': {"temperature":29.495,"humidity":55.19,"pressure":1019.793}
Published to topic 'IIoT/machine_5/humidity': {"temperature":29.495,"humidity":55.19,"pressure":1019.793}
Published to topic 'IIoT/machine_5/pressure': {"temperature":29.495,"humidity":55.19,"pressure":1019.793}
Published to topic 'sensor/data': {"temperature":26.099,"humidity":56.784,"pressure":1014.954}
Received on topic 'IIoT/machine_2/temperature': Temperature: 28.06°C, Humidity: 60.00%, Pressure: 1018.32 hPa
Received on topic 'IIoT/machine_2/humidity': Temperature: 28.06°C, Humidity: 60.00%, Pressure: 1018.32 hPa
Received on topic 'IIoT/machine_2/pressure': Temperature: 28.06°C, Humidity: 60.00%, Pressure: 1018.32 hPa
Received on topic 'IIoT/machine_3/temperature': Temperature: 28.86°C, Humidity: 60.00%, Pressure: 1019.30 hPa
Received on topic 'IIoT/machine_3/humidity': Temperature: 28.86°C, Humidity: 60.00%, Pressure: 1019.30 hPa
Received on topic 'IIoT/machine_3/pressure': Temperature: 28.86°C, Humidity: 60.00%, Pressure: 1019.30 hPa
Received on topic 'IIoT/machine_4/temperature': Temperature: 29.36°C, Humidity: 58.73%, Pressure: 1018.21 hPa
Received on topic 'IIoT/machine_4/humidity': Temperature: 29.36°C, Humidity: 58.73%, Pressure: 1018.21 hPa
Received on topic 'IIoT/machine_4/pressure': Temperature: 29.36°C, Humidity: 58.73%, Pressure: 1018.21 hPa
Received on topic 'IIoT/machine_5/temperature': Temperature: 29.50°C, Humidity: 55.19%, Pressure: 1019.79 hPa
Received on topic 'IIoT/machine_5/humidity': Temperature: 29.50°C, Humidity: 55.19%, Pressure: 1019.79 hPa
Received on topic 'IIoT/machine_5/pressure': Temperature: 29.50°C, Humidity: 55.19%, Pressure: 1019.79 hPa

Shutting down gracefully...
MQTT client disconnected. Goodbye!

[Done] exited with code=null in 6.209 seconds


# Don't forget to stop your mosquitto broker from running in the background afterward!

(base) ____@MacBook-Pro ~ % brew services stop mosquitto
Stopping `mosquitto`... (might take a while)
==> Successfully stopped `mosquitto` (label: homebrew.mxcl.mosquitto)
