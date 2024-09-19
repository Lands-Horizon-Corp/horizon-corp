FROM node:20.17.0
WORKDIR /client
COPY package.json ./
COPY pnpm-lock.yaml ./
RUN npm i -g pnpm 
RUN pnpm install

COPY . .

EXPOSE 80

CMD ["pnpm", "dev"]
