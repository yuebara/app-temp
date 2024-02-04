# 開発環境向けステージ
FROM node:21-alpine AS development
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm", "run", "dev"]

# 本番環境向けステージ
FROM node:21-alpine AS build_prod
WORKDIR /app
COPY package*.json ./
RUN npm install --production
COPY . .
RUN npm run build

# 本番環境向けステージ (軽量なイメージ)
FROM node:21-alpine AS prod
WORKDIR /app
COPY --from=build_prod /app/package*.json ./
COPY --from=build_prod /app/.next ./.next
COPY --from=build_prod /app/public ./public
COPY --from=build_prod /app/next.config.js ./
RUN npm install --production
EXPOSE 3000
CMD ["npm", "start"]
