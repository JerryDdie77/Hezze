FROM node:22-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5173

CMD ["sh", "-c", "echo '' && echo '=========================================' && echo '🔧 Starting Vite development server...' && echo '=========================================' && echo '' && echo '📍  Application URL: http://localhost:5173' && echo '📝  Access your project at the URL above' && echo '🔄  Hot reload is enabled' && echo '' && echo '=========================================' && echo '' && npm run dev -- --host 0.0.0.0"]