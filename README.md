# Простой сервис для отправки уведомлений из GitHub WebHook - Discussion в Telegram Bot

<img width="390" height="321" alt="image" src="https://github.com/user-attachments/assets/874a8b6f-f4b2-48aa-a083-4d32b50ca194" />

### Сборка:
```sh
docker build -t tg-notify-service .
```

### Запуск:
```sh
docker run \                                                                                                             
  -e TELEGRAM_BOT_TOKEN=telegram_bot_token \
  -e TELEGRAM_USERID=your_telegram_id \
  -e GITHUB_WEBHOOK_SECRET=github_webhook_secret \
  tg-notify-service
```

После запуска сервис будет доступен на порту **8080**

### Эндпоинты:
`/webhook/github` - Основной url, на него github должен слать запросы
`/health` - Url для проверки статуса контейнера

---

Создавался для мониторинга giscus на сайте.
