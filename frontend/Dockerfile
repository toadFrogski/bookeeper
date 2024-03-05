FROM node:18-alpine

WORKDIR /frontend

COPY . .

RUN yarn install

EXPOSE 80

ENTRYPOINT [ "yarn" ]
CMD [ "dev", "--port", "80" ]