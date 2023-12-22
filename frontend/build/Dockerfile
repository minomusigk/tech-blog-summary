# local
FROM node:21-alpine AS local
COPY package.json package-lock.json ./
RUN yarn install
CMD ["yarn", "start"]

FROM local AS build
RUN yarn build

FROM nginx:alpine as production
RUN apk update && apk upgrade --no-cache
COPY --from=build /app/dist/ /var/www/tech-blog-summary/
COPY ./docker/nginx.conf /etc/nginx/conf.d/default.conf