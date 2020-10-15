### ueshka

Получение и пересылка событий [УЭШКА](https://ueshka.ru/) в сторонние системы.

Поддерживаемые системы уведомлений:
- [x] Telegram
- [ ] WhatsApp

#### Переменные окружения
|Значение|По умолчанию|Обязательно|Описания|
|---|---|---|---|
|UESHKA_API_VERSION|LK/1.8.12|_нет_|версия api|
|UESHKA_TOKEN||_да_|token api|
|UESHKA_PUPIL_ID||_да_|id ребенка|
|CHECK_INTERVAL|10|_нет_|период опроса api в секундах|
|GATE_TYPE||_да_|система пересылки [список](https://github.com/DesSolo/ueshka#%D0%BF%D0%BE%D0%B4%D0%B4%D0%B5%D1%80%D0%B6%D0%B8%D0%B2%D0%B0%D0%B5%D0%BC%D1%8B%D0%B5-%D1%81%D0%B8%D1%81%D1%82%D0%B5%D0%BC%D1%8B-%D1%83%D0%B2%D0%B5%D0%B4%D0%BE%D0%BC%D0%BB%D0%B5%D0%BD%D0%B8%D0%B9)|

#### Поддерживаемые системы уведомлений

##### telegram
`GATE_TYPE=telegram`
|Значение|По умолчанию|Обязательно|
|---|---|---|
|TELEGRAM_TOKEN||_yes_|
|TELEGRAM_CHAT_ID||_yes_|