### ueshka

Получение и пересылка событий [УЭШКА](https://ueshka.ru/) в сторонние системы.

Поддерживаемые системы уведомлений:
- [ ] Telegram
- [ ] WhatsApp

#### Переменные окружения
|Значение|По умолчанию|Обязательно|Описания|
|---|---|---|---|
|UESHKA_API_VERSION|LK/1.8.12|_нет_|версия api|
|UESHKA_TOKEN||_да_|token api|
|UESHKA_PUPIL_ID||_да_|id ребенка|
|CHECK_INTERVAL|10|_нет_|период опроса api в секундах|
|GATE_TYPE||_да_|система пересылки [список](https://github.com/DesSolo/ueshka#gate-types)|

#### Поддерживаемые системы уведомлений

##### telegram
`GATE_TYPE=telegram`
|Variable|Default|Required|
|---|---|---|
|TELEGRAM_TOKEN||_yes_|
|TELEGRAM_CHAT_ID||_yes_|