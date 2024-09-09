FROM node:20 AS development

WORKDIR /app

RUN npm install -g bun

COPY package.json bun.lockb ./

RUN bun install

COPY . .

EXPOSE 3000

CMD ["bun", "run", "dev"]
