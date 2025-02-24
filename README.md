# 📌 Telegram Bot – Subscription Verification
This bot checks if a user is subscribed to specified channels and sends corresponding messages, including text, buttons, and GIF animations.

---

## 🚀 Features
✅ Verifies if a user is subscribed to a channel  
✅ Sends **photos, videos, or GIFs** after subscription  
✅ Uses **inline buttons** for easy navigation  
✅ **Multi-language support** (localization)  
✅ Logs errors  

---

## ⚙️ Installation and Setup
### 1. Clone the Repository
```sh
git clone https://github.com/vaaxooo/tgaccessbot
cd tgaccessbot
```

### 2. Install Dependencies
```sh
go mod tidy
```

### 3. Create the `.env` File
```sh
cp .env.example .env
```
Fill in the `.env` file:
```
TELEGRAM_BOT_TOKEN=YOUR_TELEGRAM_BOT_TOKEN
TELEGRAM_CHANNELS=YOUR_TELEGRAM_CHANNEL_USERNAME,YOUR_TELEGRAM_CHANNEL_USERNAME
SUCCESS_REDIRECT_URL=https://t.me
```

### 4. Run the Bot
```sh
go run cmd/main.go
```

---

## 📌 Bot Commands
| Command        | Description |
|---------------|-------------|
| `/start`      | Start the bot |
| `/check`      | Check subscription |
| **Buttons**   | **Functions** |
| 🔓 **Get Access** | Checks subscription and sends a GIF or video |

---

## 🛠 Localization Settings
The bot supports **multiple languages**. Localization files are stored in `/locales/` and contain JSON files:

```
/locales/
├── en.json  # English
├── ru.json  # Russian
```
**Example `en.json`:**
```json
{
    "start_message": "Hello! I will help you gain access to exclusive content.",
    "check_button": "🔍 Check Subscription",
    "success_message": "🎉 Congratulations! You are successfully subscribed!"
}
```

---

## 📩 Contact
🔗 [GitHub](https://github.com/vaaxooo)  
📩 Email: vaaxooo@gmail.com

Now your `README.md` includes **installation, commands, project description**! 🚀
