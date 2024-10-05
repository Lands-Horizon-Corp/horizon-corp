FROM node:20.17.0

WORKDIR /client

COPY package.json pnpm-lock.yaml ./

RUN npm install -g npm@10.8.2

RUN npm install -g pnpm@9.10.0

RUN pnpm install

COPY . .

EXPOSE 80

CMD ["pnpm", "dev-docker"]
