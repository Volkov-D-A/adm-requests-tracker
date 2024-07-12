FROM node:22.4-alpine3.19 AS build
WORKDIR /app
ENV PATH=/app/node_modules/.bin:$PATH
COPY package.json /app/package.json
RUN npm install
COPY . /app
RUN npm run build

FROM nginx:stable-alpine
RUN mkdir /public_html
COPY --from=build /app/dist /public_html
RUN rm /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]