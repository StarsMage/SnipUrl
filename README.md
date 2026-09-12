
# 🌋🩸 Snip — High-Performance URL Shortener 🩸🌋

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-sqlite&logo=go)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18+-4169E1?style=flat-sqlite&logo=postgresql)](https://www.postgresql.org)
[![Docker Compose](https://img.shields.io/badge/Docker_Compose-v2+-2496ED?style=flat-sqlite&logo=docker)](https://www.docker.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)


**Snip** is a brutally fast, lightweight, and containerized URL shortener built with Go, PostgreSQL, and modern vanilla web technology. It features custom hash-based URL encoding, persistent click analytics, and effortless multi-container deployment via Docker Compose.

---

## 🌋 Features

- 🩸 **High Performance:** Powered by Go's high-concurrency standard runtime and non-blocking router.
- 🌋 **Custom Base62 Hashing:** Efficient unique short-code generation forged in code.
- 🩸 **Real-Time Click Tracking:** Tracks total link visits atomically stored in PostgreSQL.
- 🌋 **Full Containerization:** Pre-configured Docker Compose setup for instant localized deployment.
- 🩸 **Modern Minimal UI:** Clean, responsive, dependency-free frontend HTML5/CSS3 interface.

---

## 🛠 Tech Stack

- **Backend:** Go (`net/http`, `chi` router)
- **Database:** PostgreSQL 18+ with `pgxpool` connection pooling
- **Frontend:** Plain HTML5, Modern CSS Variables
- **DevOps:** Docker, Docker Compose

---

## 🚀 Quick Start

### Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running.

### Installation & Run

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/StarsMage/SnipUrl.git](https://github.com/StarsMage/SnipUrl.git)
   cd SnipUrl

   ```

2. **Start the application with Docker Compose:**
   ```bash
   docker compose up --build -d

   ```


3. **Access the Web Interface:**
Open your browser and navigate to `http://localhost:8080`.
4. **Stop the services:**
   ```bash
   docker compose down -v

   ```



---

## 📂 Project Structure

```text
.
├── db/
│   └── postgres.sql       # Database schema initialization script
├── static/
│   └── style.css          # Frontend styling
├── storage/
│   └── postgres.go        # PostgreSQL connection pool & queries
├── .dockerignore
├── .gitignore
├── alg.go                 # Base62 hash encoder logic
├── Dockerfile             # Multi-stage Go build container definition
├── docker-compose.yml     # Service orchestration (App + Postgres)
├── go.mod
├── go.sum
├── index.html             # Main frontend page
└── main.go                # HTTP handlers & entrypoint

```

---

## 📄 License

Distributed under the MIT License. See [`LICENSE`](https://www.google.com/search?q=LICENSE) for more information.

```

