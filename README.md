# PopSenzawa Echo

Welcome to use PopSenzawa Echo!

This is the server-side reproduction,
similar the one of [https://popcat.click](https://popcat.click),
improve the performance and speed.

The software provides the high-performance,
high-reliability, and high-security,
with GoLang, Redis, and MariaDB.

It's licensed under the [MIT License](LICENSE).

## How to use

### Prerequisites

- Docker
- Docker Compose

### Run

```bash
git clone https://github.com/supersonictw/popsenzawa-echo-deploy.git
cd popsenzawa-echo-deploy
sh initialize_mmdb.sh
docker-compose up -d
```

The server will start on <https://localhost:8000>.
