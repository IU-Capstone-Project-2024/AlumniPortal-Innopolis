import asyncio
from telethon.sync import TelegramClient
from telethon.sessions import StringSession
from telethon.tl.functions.messages import CreateChatRequest

async def create_chat(api_id, api_hash, session_string, chat_title, user_identifiers):
    # Инициализация клиента Telegram
    client = TelegramClient(StringSession(session_string), api_id, api_hash)
    await client.connect()

    if not await client.is_user_authorized():
        raise Exception('Авторизация не выполнена')

    # Создание чата
    try:
        users = []
        for identifier in user_identifiers:
            try:
                print(f"Fetching entity for identifier: {identifier}")
                user = await client.get_entity(identifier)
                print(f"Found entity: {user}")
                users.append(user)
            except Exception as e:
                print(f"Cannot find entity for identifier {identifier}: {e}")
                raise Exception(f"Cannot find entity for identifier {identifier}: {e}")

        chat = await client(CreateChatRequest(users=users, title=chat_title))
        return f'Чат "{chat_title}" успешно создан'
    except Exception as e:
        raise e
    finally:
        await client.disconnect()
