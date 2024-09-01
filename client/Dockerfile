FROM node:20 AS development

WORKDIR /app

RUN npm install -g bun

RUN bun --version || echo "Bun not found"

COPY package.json bun.lockb ./

RUN bun install

COPY . .

EXPOSE 3000

CMD ["bun", "run", "dev"]