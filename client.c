#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdio.h>
// #include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#define SERVER_IP "127.0.0.1"
#define SERVER_PORT 8000
#define BUFFER_SIZE 1024

int main() {
  int sock;
  struct sockaddr_in server_addr;
  char buffer[BUFFER_SIZE];

  sock = socket(AF_INET, SOCK_STREAM, 0);

  if (sock == -1) {
    perror("Socket creation failed");
    return -1;
  };

  server_addr.sin_family = AF_INET;
  server_addr.sin_port = htons(SERVER_PORT);
  server_addr.sin_addr.s_addr = inet_addr(SERVER_IP);

  if (connect(sock, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
    perror("Connection failed");
    close(sock);
    return -1;
  };

  printf("Conected to Server %s:%d\n", SERVER_IP, SERVER_PORT);

  while (1) {
    memset(buffer, 0, BUFFER_SIZE);
    int bytes_recieved = recv(sock, buffer, BUFFER_SIZE - 1, 0);
    if (bytes_recieved <= 0) {
      printf("Connection Closed of error\n");
      break;
    };

    buffer[bytes_recieved] = '\0';
    printf("Message from server: %s\n", buffer);

    char *msg = "";

    if (strcmp(buffer, "PING") == 0) {
      printf("Received PING, sending PONG\n");
      if (send(sock, "PONG", strlen("PONG"), 0) < 0) {
        perror("Send failed");
        break;
      }
    } else if (strcmp(buffer, "BLINK ON") == 0) {
      printf("Received BLINK ON\n");
      if (send(sock, "BLINK-ON", strlen("BLINK-ON"), 0) < 0) {
        perror("Send failed");
        break;
      }
    } else if (strcmp(buffer, "BLINK OFF") == 0) {
      printf("Received BLINK OFF\n");
      if (send(sock, "BLINK-OFF", strlen("BLINK-OFF"), 0) < 0) {
        perror("Send failed");
        break;
      }
    } else {
      printf("UNKNOWN\n");
      if (send(sock, "UNKNOWN COMMAND", strlen("UNKNOWN COMMAND"), 0) < 0) {
        perror("Send failed");
        break;
      }
    }
  }

  close(sock);
  return 0;
}

// #include <ESP8266WiFi.h>

// const char* ssid = "YOUR_WIFI_SSID";
// const char* password = "YOUR_WIFI_PASSWORD";

// const char* server_ip = "192.168.1.100";  // Change to your actual server IP
// const uint16_t server_port = 8000;

// WiFiClient client;

// void setup() {
//   Serial.begin(115200);
//   delay(100);

//   // Connect to WiFi
//   Serial.println("Connecting to WiFi...");
//   WiFi.begin(ssid, password);

//   while (WiFi.status() != WL_CONNECTED) {
//     delay(500);
//     Serial.print(".");
//   }
//   Serial.println("\nWiFi connected");
//   Serial.print("IP Address: ");
//   Serial.println(WiFi.localIP());

//   // Connect to TCP server
//   Serial.print("Connecting to server ");
//   Serial.print(server_ip);
//   Serial.print(":");
//   Serial.println(server_port);

//   if (!client.connect(server_ip, server_port)) {
//     Serial.println("Connection to server failed");
//     return;
//   }
//   Serial.println("Connected to server");
// }

// void loop() {
//   if (client.connected() && client.available()) {
//     String message = client.readStringUntil('\n');
//     message.trim();  // Remove newline or carriage return
//     Serial.println("Message from server: " + message);

//     if (message == "PING") {
//       Serial.println("Received PING, sending PONG");
//       client.println("PONG");
//     } else if (message == "BLINK ON") {
//       Serial.println("Received BLINK ON");
//       client.println("BLINK-ON");
//       // You can add digitalWrite(LED_BUILTIN, LOW); here to turn on LED
//     } else if (message == "BLINK OFF") {
//       Serial.println("Received BLINK OFF");
//       client.println("BLINK-OFF");
//       // You can add digitalWrite(LED_BUILTIN, HIGH); here to turn off LED
//     } else {
//       Serial.println("UNKNOWN COMMAND");
//       client.println("UNKNOWN COMMAND");
//     }
//   }

//   delay(10);  // Small delay to avoid overwhelming the loop
// }
