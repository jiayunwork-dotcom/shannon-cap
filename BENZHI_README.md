Go AWGN 信道容量核算命令行工具：由信噪比与带宽算出 Shannon 容量和频谱效率，或反解达到目标容量所需 SNR/带宽；命令行用 capacity / tradeoff 读 JSON，也可用 serve 在 :8080 提供同一套 HTTP。

## 构建 / 运行 / 测试

```text
go build ./...
go run . capacity --table example/lte-10mhz.json
go run . serve --http :8080
go test ./...
```

## 评测镜像

本目录评测专用文件（勿覆盖项目自带 Dockerfile/README）：

- `benzhi.Dockerfile`
- `build_benzhi_docker.sh`
- `BENZHI_README.md`（本文件）

两种架构都要构建并进容器验证：

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh <image-name> linux/arm64
./build_benzhi_docker.sh <image-name> linux/amd64
docker run -it <image-name>:latest
```
