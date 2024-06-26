import json
import logging
from telethon import TelegramClient, events
from telethon.tl.functions.messages import CreateChatRequest
from telethon.tl.functions.messages import AddChatUserRequest
from telethon.tl.types import InputUser

# Укажите здесь ваши параметры API
api_id = '27007150'
api_hash = '3db17b4032beef04ef4ca239734014aa'
bot_token = '6508412518:AAFkfTZCt9BvAg74a3zp1UbuwjWTvC28kV4'

# Настройка логирования
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Инициализация клиента
client = TelegramClient('bot', api_id, api_hash).start(bot_token=bot_token)

@client.on(events.NewMessage(pattern='/start'))
async def start(event):
    await event.respond("Привет! Отправьте мне JSON с названием чата и ID участников, чтобы создать чат.")

@client.on(events.NewMessage)
async def create_chat(event):
    try:
        data = json.loads(event.message.message)
        chat_title = data.get('chat_title')
        user_ids = data.get('user_ids')

        if not chat_title or not user_ids:
            await event.respond("Неверный формат JSON. Он должен содержать 'chat_title' и 'user_ids'.")
            return

        # Получаем пользователей
        users = []
        for user_id in user_ids:
            user = await client.get_entity(user_id)
            users.append(user)

        # Создание чата
        chat = await client(CreateChatRequest(users=users, title=chat_title))

        await event.respond(f"Чат '{chat_title}' успешно создан с участниками: {user_ids}")

    except Exception as e:
        logger.error(e)
        await event.respond("Произошла ошибка при создании чата.")

if __name__ == '__main__':
    client.start()
    client.run_until_disconnected()
