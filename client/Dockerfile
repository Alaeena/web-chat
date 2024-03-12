FROM node:18-alpine as build

WORKDIR /app/view
COPY . /app/view

COPY package.json package-lock.json ./
RUN npm install

EXPOSE 3000
CMD ["npm", "run", "dev"]
