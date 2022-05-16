# ElifHW1

Для старту необхідно запустити базу даних 
### docker-compose up(порт 5432 має бути вільним, або змінити в конфігураціях)

для базової структури бази даних використати файл з міграцією(migrations/000001_init.up.sql
## api для користування(testHTTP.http):

POST localhost:8000/api/name_sender/name_receiver

{
  "text": "anything"
}

GET localhost:8000/api/name_receiver/name_sender
