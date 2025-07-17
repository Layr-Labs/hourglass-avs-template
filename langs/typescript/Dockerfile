FROM node:24-bookworm-slim

ADD . /app

WORKDIR /app

RUN npm install --production

CMD ["./node_modules/.bin/ts-node", "index.ts"]
