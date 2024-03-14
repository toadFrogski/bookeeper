FROM node:18-alpine

WORKDIR /frontend

COPY package.json .

RUN yarn install

COPY . .

ENTRYPOINT [ "yarn" ]
CMD [ "dev" ]