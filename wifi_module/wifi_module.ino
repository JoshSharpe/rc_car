/*
 * Taken from https://learn.sparkfun.com/tutorials/esp8266-thing-hookup-guide/example-sketch-ap-web-server
 * Libraries: 
 * - https://arduino-esp8266.readthedocs.io/en/latest/esp8266wifi/readme.html
 * 
 */

#include <ESP8266WiFi.h>
#include <ESP8266WebServer.h>

#include "config.h"
#include "pages.h"


const bool debugMode = false;

/////////////////////
// Pin Definitions //
/////////////////////
//const int LED_PIN = 5; // Thing's onboard, green LED
//const int ANALOG_PIN = A0; // The only analog pin on the Thing
//const int DIGITAL_PIN = 12; // Digital pin to be read

/*
 * Request options
 */
const int INVALID_REQUEST = -1;
const int HOME_REQUEST = 0;

/*
 * Request Paths
 */
const String HOME_PATH = "/home";

ESP8266WebServer server(80);

void setup() 
{
  initHardware();
  setupWiFi();
  initServer();
}

void loop() 
{
  server.handleClient();
  // // Check if a client has connected
  // WiFiClient client = server.available();
  // if (!client) {
  //   return;
  // }

  // // Read the first line of the request
  // String req = client.readStringUntil('\r');
  // if(debugMode) {
  //   Serial.println(req);
  // }
  // client.flush();

  // String page = getNotFoundPage();
  // int currentRequest = INVALID_REQUEST;

  // if(req.indexOf(HOME_PATH) > 0) {
  //   currentRequest = HOME_REQUEST;
  //   page = getHomePage();
  // }


  // client.flush();

  // // Send the response to the client
  // client.print(page);
  // delay(1);
  // if(debugMode) {
  //   Serial.println("Client disonnected");
  // }
}

void setupWiFi()
{
  WiFi.mode(WIFI_AP);

  // Do a little work to get a unique-ish name. Append the
  // last two bytes of the MAC (HEX'd) to "Thing-":
  uint8_t mac[WL_MAC_ADDR_LENGTH];
  WiFi.softAPmacAddress(mac);
  String macID = String(mac[WL_MAC_ADDR_LENGTH - 2], HEX) +
                 String(mac[WL_MAC_ADDR_LENGTH - 1], HEX);
  macID.toUpperCase();
  // String AP_NameString = "ESP8266 Thing " + macID;
  String AP_NameString = configuration.wifiUsername;

  char AP_NameChar[AP_NameString.length() + 1];
  memset(AP_NameChar, 0, AP_NameString.length() + 1);

  for (int i=0; i<AP_NameString.length(); i++)
    AP_NameChar[i] = AP_NameString.charAt(i);

  WiFi.softAP(AP_NameChar, configuration.wifiPassword);
}

void initHardware()
{
  Serial.begin(115200);
//  pinMode(DIGITAL_PIN, INPUT_PULLUP);
//  pinMode(LED_PIN, OUTPUT);
//  digitalWrite(LED_PIN, LOW);
  // Don't need to set ANALOG_PIN as input, 
  // that's all it can be.
}

void initServer() {
  server.on("/", handleHome);
  server.on("/movements", handleMovements);
  server.onNotFound(handleNotFound);
}

void handleHome() {
      server.send(200, "text/html",  getHomePage());
}

void handleMovements() {
  if(server.method() == HTTPMethod::HTTP_POST) {
    int argCount = server.args();
    Serial.println("Retrieving args");
    for(int i=0; i < argCount; i++) {
      Serial.println(server.arg(i));
    }

    return;
  }

  if(server.method() == HTTPMethod::HTTP_GET) {
    Serial.println("GET MOVEMENTS");
    return;
  }

  Serial.println("Invalid Method.");

  // if (server.hasArg("plain")== false){ //Check if body received
  //       server.send(200, "text/plain", "Body not received");
  //       return;
  // }

  // String message = "Body received:\n";
  //         message += server.arg("plain");
  //         message += "\n";

  // server.send(200, "text/plain", message);
  // Serial.println(message);
}

void handleNotFound() {
  server.send(404, "text/plain", getNotFoundPage());
}
