# Pinned Employee

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/pinned-employee)](https://goreportcard.com/report/github.com/patricksferraz/pinned-employee)
[![GoDoc](https://godoc.org/github.com/patricksferraz/pinned-employee?status.svg)](https://godoc.org/github.com/patricksferraz/pinned-employee)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, scalable employee management system built with Go, featuring a clean architecture design and robust API endpoints.

## 🚀 Features

- RESTful API built with Fiber framework
- Clean Architecture implementation
- PostgreSQL database integration
- Docker and Kubernetes support
- Swagger API documentation
- Environment-based configuration
- Database migrations support
- Hot-reload development environment

## 🛠️ Tech Stack

- **Language:** Go 1.18
- **Framework:** Fiber v2
- **Database:** PostgreSQL
- **ORM:** GORM
- **API Documentation:** Swagger
- **Containerization:** Docker
- **Orchestration:** Kubernetes
- **Development:** Air (hot-reload)

## 📋 Prerequisites

- Go 1.18 or higher
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## 🚀 Getting Started

1. Clone the repository:
```bash
git clone https://github.com/patricksferraz/pinned-employee.git
cd pinned-employee
```

2. Copy the environment file and configure it:
```bash
cp .env.example .env
```

3. Start the application using Docker Compose:
```bash
docker-compose up -d
```

4. For local development with hot-reload:
```bash
make dev
```

## 📚 API Documentation

Once the application is running, you can access the Swagger documentation at:
```
http://localhost:8080/swagger/
```

## 🏗️ Project Structure

```
.
├── app/            # Application layer
│   └── rest/       # REST API handlers
├── cmd/            # Command line interface
├── domain/         # Domain layer
│   ├── entity/     # Domain entities
│   ├── repo/       # Repository interfaces
│   └── service/    # Business logic
├── infra/          # Infrastructure layer
├── k8s/            # Kubernetes configurations
└── utils/          # Utility functions
```

## 🧪 Testing

Run the test suite:
```bash
make test
```

## 📦 Deployment

The project includes Kubernetes configurations in the `k8s/` directory. Deploy to your cluster using:
```bash
kubectl apply -f k8s/
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Patrick Sferraz** - *Initial work* - [GitHub Profile](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://github.com/go-gorm/gorm)
- [Swagger](https://github.com/swaggo/swag)
- [Air](https://github.com/cosmtrek/air)
