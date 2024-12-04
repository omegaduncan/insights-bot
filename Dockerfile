# 使用官方 Node.js 長期支持版作為基礎鏡像
FROM node:20-alpine3.18 AS base

# 設置工作目錄
WORKDIR /app

# 設置非 root 用戶
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 複製依賴文件
COPY package*.json ./

# 構建階段
FROM base AS build

# 安裝構建依賴
RUN apk add --no-cache python3 make gcc g++

# 安裝項目依賴
RUN npm ci --only=production

# 最終階段
FROM base AS production

# 複製編譯後的依賴
COPY --from=build /app/node_modules ./node_modules
COPY . .

# 修改文件權限
RUN chown -R appuser:appgroup /app

# 切換到非 root 用戶
USER appuser

# 暴露應用端口
EXPOSE 7069

# 設置健康檢查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s \
  CMD wget -q -O- http://localhost:3000/health || exit 1

# 設置容器啟動命令
CMD ["npm", "start"]
