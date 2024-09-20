# 🤖 Telegram bot 

[![CC BY-NC-SA 4.0][cc-by-nc-sa-shield]][cc-by-nc-sa]

[cc-by-nc-sa]: http://creativecommons.org/licenses/by-nc-sa/4.0/
[cc-by-nc-sa-image]: https://licensebuttons.net/l/by-nc-sa/4.0/88x31.png
[cc-by-nc-sa-shield]: https://img.shields.io/badge/License-CC%20BY--NC--SA%204.0-lightgrey.svg

[![stable](https://img.shields.io/badge/-stable-brightgreen?style=flat-square)](https://go-faster.org/docs/projects/status#stable)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/boeing666/tgmarket_bot/.github%2Fworkflows%2Fgo.yml?style=flat-square)
[![Release](https://img.shields.io/github/release/boeing666/telegram_bot.svg?style=flat-square)](https://github.com/boeing666/tgmarket_bot/releases)
![GitHub last commit](https://img.shields.io/github/last-commit/boeing666/tgmarket_bot?style=flat-square)

## 📘 Описание
Телеграм бот, который отслеживает цены и бонусы на товары с маркетплейсов. Пользователи могут добавлять товары, устанавливать условия для уведомлений и получать актуальные сведения о снижении цен или увеличении бонусов.
## 🎯 Цель
Помочь пользователям не упустить выгодные предложения и оптимизировать расходы, автоматически информируя их о лучших ценах и акциях на интересующие товары.
## 📋 Функции
- ➕ **Добавление товаров** – Пользователь вводит ссылку на товар для отслеживания его цены и бонусов.
- ⏳ **Настройка периодичности опроса** - Пользователь выбирает, как часто бот будет проверять цены на товары.  
- 🔔 **Уведомления о снижении цен** - Бот отправляет уведомление, если цена на товар становится ниже установленной.
- 🏆 **Уведомления об увеличении бонусов** - Бот информирует, если количество бонусов за товар превышает заданную пользователем отметку.
- 📜 **Управление списком товаров** - Пользователь может просматривать и редактировать список отслеживаемых товаров.
## 🌟 Использование
1. Скачайте последнюю версию бота и разархивируйте [Releases](https://github.com/boeing666/tgmarket_bot/releases)
2. Создайте api_token [BotFather](https://t.me/BotFather)
3. Заполните конфиг данными ```configs/config.json```:
```yml
{
    "api_token": "",
    "db_settings": {
        "host": "",
        "username": "",
        "password": "",
        "database": ""
    }
}
```
4. Запустите бота
