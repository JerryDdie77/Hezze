FROM node:22-alpine

WORKDIR /app

# Копируем package.json и устанавливаем зависимости
COPY package*.json ./
RUN npm install

# Копируем исходный код
COPY . .


# Запускаем dev сервер
CMD ["npm", "run", "dev", "--", "--host"]