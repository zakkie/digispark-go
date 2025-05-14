#include <Arduino.h>
#include <DigiUSB.h>

void setup() {
  pinMode(LED_BUILTIN, OUTPUT);
  DigiUSB.begin();
}

typedef enum {
  LED_OFF = 0,
  LED_ON,
  LED_BLINK,
} led_mode;

led_mode mode = LED_OFF;
bool blink_status = false;

void loop() {
  if (DigiUSB.available()) {
    char c = DigiUSB.read();

    switch (c) {
      case '0':  // off
        digitalWrite(LED_BUILTIN, LOW);
        mode = LED_OFF;
        break;
      case '1':  // on
        digitalWrite(LED_BUILTIN, HIGH);
        mode = LED_ON;
        break;
      case 'b':  // blink
        mode = LED_BLINK;
        break;
      default:
        break;
    }
  }
  DigiUSB.refresh(); // Don't remove this

  if (mode == LED_BLINK) {
    digitalWrite(LED_BUILTIN, blink_status ? HIGH : LOW);
    blink_status = !blink_status;
  }
  delay(100);
}
