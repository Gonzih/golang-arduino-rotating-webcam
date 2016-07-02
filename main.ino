#include <Servo.h>

Servo myservo;

int oldPos = 0;
int pos = 0;
int tmp = 0;
int delta = 0;

void setup() {
  myservo.attach(9);
  myservo.write(0);
  Serial.begin(9600);
}

void loop() {
  if (Serial.available() > 0) {

    tmp = Serial.parseInt();

    if (tmp >= 0 && tmp <= 180) {
      Serial.print("Using new value: ");
      Serial.println(tmp);

      oldPos = pos;
      pos = tmp;

      if (oldPos < pos) {
        delta = 1;
      } else {
        delta = -1;
      }

      while (oldPos != pos) {
        oldPos += delta;

        myservo.write(oldPos);
        Serial.print("Wrote value: ");
        Serial.println(oldPos);
        delay(25);
      }
    }

    delay(500);
  }
}
