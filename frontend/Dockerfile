# этап сборки (build stage)
FROM node:16.3-alpine3.12 AS build-stage
WORKDIR /app
#COPY package*.json ./
#RUN npm install
COPY dist dist

# RUN npm run build

# этап production (production-stage)
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]