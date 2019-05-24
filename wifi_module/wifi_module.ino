/*
   Taken from https://learn.sparkfun.com/tutorials/esp8266-thing-hookup-guide/example-sketch-ap-web-server
   Libraries:
   - https://arduino-esp8266.readthedocs.io/en/latest/esp8266wifi/readme.html

*/

#include <ESP8266WiFi.h>
#include <ESP8266WebServer.h>

#include "config.h"
#include "pages.h"


const bool debugMode = false;

const IPAddress IP = (192, 168, 4, 1);
const int PORT = 80;

/////////////////////
// Pin Definitions //
/////////////////////
//const int LED_PIN = 5; // Thing's onboard, green LED
//const int ANALOG_PIN = A0; // The only analog pin on the Thing
//const int DIGITAL_PIN = 12; // Digital pin to be read

/*
   Request options
*/
const int INVALID_REQUEST = -1;
const int HOME_REQUEST = 0;

/*
   Request Paths
*/
const String HOME_PATH = "/home";

ESP8266WebServer server(PORT);

void setup()
{
  initHardware();
  setupWiFi();
  initServer();
}

void loop()
{
  server.handleClient();
}

void setupWiFi()
{
  Serial.println("Setting up Wifi.");
  Serial.print("configuration.wifiUsername: ");
  Serial.println(configuration.wifiUsername);
  Serial.print("configuration.wifiPassword: ");
  Serial.println(configuration.wifiPassword);
  WiFi.mode(WIFI_STA);
  WiFi.begin(configuration.wifiUsername, configuration.wifiPassword);

  Serial.println("");
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.printf("Web server started, open %s in a web browser\n", WiFi.localIP().toString().c_str());
  Serial.println("Finished setting up Wifi.");
  Serial.printf("Connection status: %d\n", WiFi.status());
}

void initHardware()
{
  Serial.begin(115200);
}

void initServer() {
  server.on("/", handleHome);
  server.on("/movements", handleMovements);
  server.onNotFound(handleNotFound);
  server.begin();
}

void handleHome() {
  server.send(200, "text/html",  getHomePage());
}

void handleMovements() {
  if (server.method() == HTTPMethod::HTTP_POST) {


    return;
  }

  if (server.method() == HTTPMethod::HTTP_GET) {
    Serial.println("GET MOVEMENTS");
    int argCount = server.args();
    Serial.println("Retrieving args");
    for (int i = 0; i < argCount; i++) {
      Serial.println(server.arg(i));
    }
    server.send(200, "text/plain", "Found!");
    return;
  }

  Serial.println("Invalid Method.");
}

void handleNotFound() {
  server.send(404, "text/plain", getNotFoundPage());
}
