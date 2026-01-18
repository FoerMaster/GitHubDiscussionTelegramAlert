# Simple service for sending notifications from GitHub WebHook Discussions to Telegram Bot

<img width="390" height="321" alt="image" src="https://github.com/user-attachments/assets/874a8b6f-f4b2-48aa-a083-4d32b50ca194" />

## Build

```sh
docker build -t tg-notify-service .
```

## Run

```sh
docker run \
  -e TELEGRAM_BOT_TOKEN=telegram_bot_token \
  -e TELEGRAM_USERID=your_telegram_id \
  -e GITHUB_WEBHOOK_SECRET=github_webhook_secret \
  -p 8080:8080 \
  tg-notify-service
```

After starting, the service will be available on port **8080**.

## Endpoints

- `/webhook/github` - Main endpoint where GitHub should send webhook requests
- `/health` - Health check endpoint for container status monitoring

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `TELEGRAM_BOT_TOKEN` | Your Telegram bot token from [@BotFather](https://t.me/BotFather) | ✅ |
| `TELEGRAM_USERID` | Your Telegram user ID | ✅ |
| `GITHUB_WEBHOOK_SECRET` | Secret key for webhook signature validation | ✅ |
| `PORT` | Server port (default: 8080) | ❌ |

## GitHub Webhook Setup

1. Go to your repository **Settings** → **Webhooks** → **Add webhook**
2. Set **Payload URL**: `https://your-domain.com/webhook/github`
3. Set **Content type**: `application/json`
4. Set **Secret**: Same value as `GITHUB_WEBHOOK_SECRET`
5. Select events: **Discussions** and **Discussion comments**

## Supported Events

✅ Discussion created
✅ Discussion edited
✅ Discussion deleted
✅ Discussion comment created
✅ Discussion comment edited
✅ Discussion comment deleted

---

*Created for monitoring [giscus](https://giscus.app/) comments on a website.*
