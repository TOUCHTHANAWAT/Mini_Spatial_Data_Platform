# FROM node:22-alpine

# WORKDIR /app

# ARG VITE_API_URL
# ENV VITE_API_URL=$VITE_API_URL

# COPY package*.json ./

# RUN npm install

# COPY . .

# RUN npm run build

# EXPOSE 4173

# CMD ["npm", "run", "preview", "--", "--host", "0.0.0.0"]

FROM node:22-alpine AS builder

WORKDIR /app

ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]