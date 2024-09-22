# Svetozar

_Светозар — это имя, которое имеет глубокий исторический корень. В древнеславянской мифологии, Светозар был солнечным богом, который воплощал красоту и свет. Он был богатырем и воином, который защищал свой народ, но также являлся богом любви и миролюбия._

## 📚 О проекте

Svetozar умеет обнаруживать нежелательную активность в цифровых социальных сервисах (Telegram).

Svetozar устроен просто и надежно. Фоновый процесс "прослушивает" чаты авторизованного пользователя, используя библиотеку _TDLib (Telegram Database Library)_.

_TDLib_ - is a cross-platform, fully functional Telegram client. We designed it to help third-party developers create their own custom apps using the Telegram platform.

Для классификации сообщений используется нейросеть _ruBERT-toxic_ https://huggingface.co/sismetanin/rubert-toxic-pikabu-2ch.

## ⚡️ Запуск

1. Установка переменных окружения  
   `source .env.test`

2. Запуск приложения  
   `go run cmd/app/main.go serve`
