# Telegram File Stream Bot

A Telegram bot that generates **direct streamable links** for your Telegram files. Uses MTProto (not Bot API), supporting files up to 2GB (4GB with Telegram Premium).

---

## Prerequisites

| Item | Where to get it |
|---|---|
| `API_ID` & `API_HASH` | https://my.telegram.org/apps — create an app |
| `BOT_TOKEN` | [@BotFather](https://t.me/BotFather) on Telegram |
| `LOG_CHANNEL` | Create a **private channel**, add your bot as admin, then get the channel ID (use [@userinfobot](https://t.me/userinfobot)) |

---

## Deploy on Render (Git-based)

1. Push this repo to GitHub
2. Go to https://dashboard.render.com → **New Web Service**
3. Connect your GitHub repo
4. Set **Language** to **Go** (not Docker)
5. Set **Start Command** to `fsb run` (auto-read from Procfile)
6. Set **Health Check Path** to `/` (optional)
7. Add environment variables:

| Variable | Required | Description |
|---|---|---|
| `API_ID` | Yes | From my.telegram.org |
| `API_HASH` | Yes | From my.telegram.org |
| `BOT_TOKEN` | Yes | From @BotFather |
| `LOG_CHANNEL` | Yes | Channel ID (e.g. `-100123456789`) |
| `HOST` | Yes | `https://your-app.onrender.com` |
| `HASH_LENGTH` | No | Hash length in URLs (default: `10`, min: `10`, max: `64`) |
| `ALLOWED_USERS` | No | Comma-separated Telegram user IDs (restricts bot access) |
| `USER_SESSION` | No | Pyrogram session string for auto-adding bots to channel |
| `STREAM_CONCURRENCY` | No | Parallel chunk downloads (default: `4`) |
| `STREAM_BUFFER_COUNT` | No | Chunks to prefetch in memory (default: `8`) |
| `STREAM_TIMEOUT_SEC` | No | Per-chunk timeout (default: `30`) |
| `STREAM_MAX_RETRIES` | No | Retries per failed chunk (default: `3`) |

8. Click **Deploy Web Service**

---

## Environment Variables (.env)

Copy `fsb.sample.env` to `fsb.env` and fill in:

```env
API_ID=
API_HASH=
BOT_TOKEN=
LOG_CHANNEL=

HOST=https://your-app.onrender.com
HASH_LENGTH=10
ALLOWED_USERS=
```

| Variable | Default | Description |
|---|---|---|
| `API_ID` | — | Telegram API ID (required) |
| `API_HASH` | — | Telegram API Hash (required) |
| `BOT_TOKEN` | — | Bot token from @BotFather (required) |
| `LOG_CHANNEL` | — | Channel ID for storing forwarded files (required) |
| `HOST` | auto-detect | Public URL of your instance |
| `HASH_LENGTH` | `10` | Length of hash in stream URLs |
| `ALLOWED_USERS` | empty | Restrict bot to specific Telegram user IDs (comma-separated) |
| `DEV` | `false` | Enable debug mode |
| `PORT` | `8080` | Server port |
| `USE_SESSION_FILE` | `true` | Persist bot sessions for faster startup |
| `USER_SESSION` | empty | Pyrogram session string for userbot features |
| `STREAM_CONCURRENCY` | `4` | Parallel chunk downloads per request |
| `STREAM_BUFFER_COUNT` | `8` | Chunks to keep in memory |
| `STREAM_TIMEOUT_SEC` | `30` | Per-chunk timeout in seconds |
| `STREAM_MAX_RETRIES` | `3` | Max retries per failed chunk |
| `MULTI_TOKEN1-4` | empty | Additional bot tokens for load balancing |
| `USE_PUBLIC_IP` | `false` | Auto-detect public IP (not recommended) |

---

## Usage

1. Start the bot on Telegram
2. Send any file/photo/video to the bot
3. Bot replies with a **streamable link**
4. Open the link in browser to stream or download

Add `?d=true` to the link to force download instead of streaming.

---

## Security Notes

- Hash uses **SHA-256** (min 10 characters)
- Rate limiting: 10 requests/second per IP
- Security headers: X-Content-Type-Options, X-Frame-Options, X-XSS-Protection, Referrer-Policy
- Error messages are sanitized (no internal info leaked)
- Filenames sanitized to prevent header injection
- `ALLOWED_USERS` restricts who can generate links

---

## Credits

- [@celestix](https://github.com/celestix) for [gotgproto](https://github.com/celestix/gotgproto)
- [@divyam234](https://github.com/divyam234/teldrive) for [Teldrive](https://github.com/divyam234/teldrive)
- [@krau](https://github.com/krau) for image support

## License

Copyright (C) 2026 under [GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.en.html).
