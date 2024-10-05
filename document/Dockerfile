
FROM node:20.17.0

WORKDIR /document

COPY package.json pnpm-lock.yaml ./

RUN npm install -g npm@10.8.2

RUN npm install -g pnpm@9.10.0

RUN pnpm install

COPY . .

EXPOSE 5321

CMD ["pnpm", "docs:dev"]
