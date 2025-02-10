FROM node:16
WORKDIR /app
COPY package*.json ./
# COPY ./public ./public

# Устанавливаем зависимости
RUN npm install

# Копируем исходный код
# COPY ./src/components ./src/components
# COPY ./src/App.tsx ./src/App.tsx
# COPY ./src/index.css ./src/index.css 
# COPY ./src/index.tsx ./src/index.tsx 
COPY . .

# Собираем приложение
RUN rm -rf src/logo.svg src/App.css src/App.test.js src/reportWebVitals.js src/setupTests.js
RUN npm run build

# Указываем порт, который будет использовать сервис
EXPOSE 3000

# Запускаем приложение
CMD ["npm", "start"]