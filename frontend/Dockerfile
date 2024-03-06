FROM node:18-alpine

WORKDIR /frontend

COPY . .

RUN yarn install

ENTRYPOINT [ "yarn" ]
CMD [ "dev" ]