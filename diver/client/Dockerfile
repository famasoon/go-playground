FROM node:11.6.0-slim

WORKDIR /client

COPY . .
RUN npm install
EXPOSE 8080
CMD ["npm", "run", "dev"]
