#!/bin/bash
echo "Installing CurrencyBot..."
read -p "Enter Telegram Bot Token: " BOT_TOKEN
cat <<EOF > config.json
{"token":"$BOT_TOKEN","db_path":"./db"}
EOF
go build -o currencybot .
echo "Installation complete."
