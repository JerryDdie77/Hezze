# Запуск проекта при разработке:

```bash
#Начало
git clone https://github.com/JerryDdie77/Hezze
cd Hezze
cp .env.example .env # или просто создать .env файл, скопировав туда всё из .env.example
``` 

```bash
#Windows
docker-compose -f docker-compose.dev.yml up -d

#Linux
docker compose -f docker-compose.dev.yml up -d
```

Примечание: в dev окружении контейнер с фронтендом находится в режиме *hot reload* (можно на ходу верстать и видеть изменения), но:

- контейнер с БД и бекендом по прежнему работают в статичном режиме
- при запуске команды выше, образ бекенда НЕ будет пересобираться

Полный перезапуск всего проекта возможен ТОЛЬКО с помощью команды ниже:

```bash
#Windows
docker-compose -f docker-compose.dev.yml up -d --build

#Linux
docker compose -f docker-compose.dev.yml up -d --build
```

# Полная очистка проекта

```bash
#Windows

#Остановить всё и удалить тома (данные из БД пропадут)
docker-compose -f docker-compose.dev.yml down -v

#Остановить всё и удалить тома + удалить образы
docker-compose -f docker-compose.dev.yml down -v --rmi all

#Linux
docker compose -f docker-compose.dev.yml down -v

docker compose -f docker-compose.dev.yml down -v --rmi all
```