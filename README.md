Запуск проекта при разработке:

```bash
#Windows
docker-compose -f docker-compose.dev.yml up -d

#Linux
docker compose -f docker.compose.dev.yml up -d
```

Примечание: в dev окружении контейнер с фронтендом находится в режиме *hot reload* (можно на ходу верстать и видеть изменения), но:

- контейнер с БД и бекендом по прежнему работают в статичном режиме
- при запуске команды выше, образ бекенда НЕ будет пересобираться

Полный перезапуск всего проекта возможен ТОЛЬКО с помощью команды ниже:

```bash
#Windows
docker-compose -f docker-compose.dev.yml up -d --build

#Linux
docker compose -f docker.compose.dev.yml up -d --build
```

Таким образом здесь пересоберутся контейнеры и с БД и с Go-приложением