# CurrencyBot

A high‑performance Telegram bot that scrapes Bonbast & Navasan for currency, gold, and coin prices, and posts updates to Telegram groups. Optionally uses official APIs when an API key is provided (e.g., Navasan). Supports fully configurable posting logic, templates, downtime scheduling, admin controls, multi‑language UI, self‑healing, auto‑failover monitoring, and more.

## Features

✔ Sell/Buy/Both price modes  
✔ Multi‑language (Persian & English) bot UI  
✔ Trigger‑based price posting  
✔ Downtime scheduling & aligned interval posting  
✔ Database cache clearing & resource stats  
✔ Multiple templates per channel  
✔ Admin alerting on errors & scraper failures  
✔ Optional Navasan API fallback  
✔ Docker & systemd support  
✔ Auto restart on crash

---

## Installation

**Upload to GitHub and deploy** via installer:

```bash
bash <(curl -s https://raw.githubusercontent.com/Armin-kho/currencybot/main/install/install.sh)
