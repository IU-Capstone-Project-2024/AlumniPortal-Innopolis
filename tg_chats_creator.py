import logging
import os
import time
from getpass import getpass

import requests
from dotenv import load_dotenv
from telethon import TelegramClient, errors
from telethon.tl.functions.messages import CreateChatRequest

load_dotenv()
# Укажите в .env ваши параметры API и данные пользователя
api_id = os.getenv("API_ID")
api_hash = os.getenv("API_HASH")
phone_number = os.getenv("PHONE")
data_url = os.getenv("DATA_URL")

# Настройка логирования
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Инициализация клиента
client = TelegramClient('user', api_id, api_hash)

async def main():
    await client.connect()

    if not await client.is_user_authorized():
        await client.send_code_request(phone_number)
        try:
            await client.sign_in(phone_number, input('Введите код: '))
        except errors.SessionPasswordNeededError:
            password = getpass('Введите пароль: ')
            await client.sign_in(password=password)

    while True:
        try:
            logger.info("Получение данных с веб-сервиса.")
            response = requests.get(data_url)
            logger.info(f"Ответ от веб-сервиса: {response.status_code}, {response.text}")

            if response.status_code == 200:
                data = response.json()
                chat_title = data.get('chat_title')
                user_ids = data.get('user_ids')

                if not chat_title or not user_ids:
                    logger.error("Неверный формат данных. Они должны содержать 'chat_title' и 'user_ids'.")
                    continue

                # Преобразуем user_ids в целые числа
                user_ids = [int(user_id) for user_id in user_ids]

                # Получаем пользователей
                users = []
                for user_id in user_ids:
                    user = await client.get_entity(user_id)
                    users.append(user)

                # Создание чата
                chat = await client(CreateChatRequest(users=users, title=chat_title))
                logger.info(f"Чат '{chat_title}' успешно создан с участниками: {user_ids}")

            else:
                logger.error("Ошибка при получении данных с веб-сервиса.")

            # Задержка перед следующим запросом
            time.sleep(60)  # Пауза 60 секунд

        except Exception as e:
            logger.error(f"Произошла ошибка: {e}")
            time.sleep(60)  # Пауза перед повторной попыткой в случае ошибки

if __name__ == '__main__':
    with client:
        client.loop.run_until_complete(main())
