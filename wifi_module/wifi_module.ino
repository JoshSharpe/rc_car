/*
 * Taken from https://learn.sparkfun.com/tutorials/esp8266-thing-hookup-guide/example-sketch-ap-web-server
 * Libraries: 
 * - https://arduino-esp8266.readthedocs.io/en/latest/esp8266wifi/readme.html
 * 
 */

#include <ESP8266WiFi.h>

#include "config.h"
#include "pages.h"


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
const string HOME_PATH = "/home"

WiFiServer server(80);

void setup() 
{
  initHardware();
  setupWiFi();
  server.begin();
}

void loop() 
{
  // Check if a client has connected
  WiFiClient client = server.available();
  if (!client) {
    return;
  }

  // Read the first line of the request
  String req = client.readStringUntil('\r');
  Serial.println(req);
  client.flush();

  string page = getNotFoundPage();
  int currentRequest = INVALID_REQUEST;

  if(req.indexOf(HOME_PATH) > 0) {
    currentRequest = HOME_REQUEST;
    page = getHomePage();
  }


  client.flush();

  // Send the response to the client
  client.print(page);
  delay(1);
  Serial.println("Client disonnected");
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
