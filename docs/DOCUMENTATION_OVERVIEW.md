# Documentation Overview

Comprehensive documentation for Podgrab covering architecture, API, deployment,
development, and testing.

## 📁 Documentation Structure

```
docs/
├── README.md                           # Documentation index and navigation
├── TESTING.md                          # Comprehensive testing guide
├── CI_CD.md                            # CI/CD pipeline documentation
├── TESTING_PROGRESS.md                 # Testing implementation progress
├── ubuntu-install.md                   # Native Ubuntu installation
│
├── architecture/                       # System Architecture (4 files)
│   ├── overview.md                    # High-level architecture with diagrams
│   ├── system-design.md               # Detailed design patterns
│   ├── data-flow.md                   # Request/response flows
│   └── database-schema.md             # Complete database documentation
│
├── api/                               # API Reference (2 files)
│   ├── rest-api.md                    # All 40+ REST endpoints documented
│   └── websocket.md                   # Real-time WebSocket API
│
├── deployment/                        # Deployment Guides (2 files)
│   ├── docker.md                      # Docker/docker-compose deployment
│   └── production.md                  # Production best practices
│
├── development/                       # Developer Guides (2 files)
│   ├── setup.md                       # Development environment setup
│   └── pre-commit.md                  # Pre-commit hooks setup
│
└── guides/                            # User Guides (2 files)
    ├── user-guide.md                  # Complete user manual
    └── configuration.md               # Configuration reference
```

**Root Documentation**:

- `CONTRIBUTING.md` - Contributing guidelines
- `CLAUDE.md` - AI assistant instructions
- `README.md` - Project overview and quick start

**Test Documentation**:

- `integration_test/README.md` - Integration testing guide
- `e2e_test/README.md` - E2E testing guide
- `.github/workflows/README.md` - CI/CD workflow documentation

## 📊 Documentation Statistics

| Category         | Files  | Key Content                               |
| ---------------- | ------ | ----------------------------------------- |
| **Testing**      | 4      | Testing guide, CI/CD, progress tracking   |
| **Architecture** | 4      | System design, data flow, database schema |
| **API**          | 2      | REST endpoints, WebSocket protocol        |
| **Deployment**   | 2      | Docker, reverse proxy, production setup   |
| **Development**  | 2      | Setup, pre-commit hooks                   |
| **User Guides**  | 2      | User manual, configuration                |
| **Total**        | **16** | **Comprehensive coverage**                |

## 🎯 Quick Navigation by Role

### For Users

**Getting Started**:

- [User Guide](guides/user-guide.md) - Complete user manual
- [Configuration](guides/configuration.md) - All configuration options
- [README](../README.md) - Project overview and installation

**Installation**:

- [Docker Deployment](deployment/docker.md) - Docker/docker-compose setup
- [Ubuntu Installation](ubuntu-install.md) - Native Ubuntu install
- [Production Setup](deployment/production.md) - Production best practices

### For Developers

**Development Setup**:

- [Development Setup](development/setup.md) - Local development environment
- [Contributing](../CONTRIBUTING.md) - Contribution guidelines
- [Pre-commit Hooks](development/pre-commit.md) - Code quality automation

**Testing**:

- [Testing Guide](TESTING.md) - Comprehensive testing documentation
- [CI/CD Documentation](CI_CD.md) - GitHub Actions pipeline
- [Testing Progress](TESTING_PROGRESS.md) - Implementation details
- [Integration Tests](../integration_test/README.md) - Integration testing
- [E2E Tests](../e2e_test/README.md) - Browser automation tests

**Architecture**:

- [Architecture Overview](architecture/overview.md) - High-level system design
- [System Design](architecture/system-design.md) - Detailed patterns
- [Data Flow](architecture/data-flow.md) - Request/response flows
- [Database Schema](architecture/database-schema.md) - Complete DB documentation

### For API Consumers

**API Documentation**:

- [REST API](api/rest-api.md) - All 40+ REST endpoints
- [WebSocket API](api/websocket.md) - Real-time communication

## 📚 Documentation by Topic

### Testing (NEW)

**Primary Documentation**:

- [TESTING.md](TESTING.md) - Comprehensive testing guide

  - Unit tests (140+ tests)
  - Integration tests (21 tests)
  - E2E tests (21 tests)
  - Coverage reports and requirements
  - Writing tests and best practices

- [CI_CD.md](CI_CD.md) - CI/CD pipeline

  - GitHub Actions workflows
  - Code quality gates
  - Automated testing
  - Docker builds and deployment

- [TESTING_PROGRESS.md](TESTING_PROGRESS.md) - Implementation tracking

  - Phase-by-phase progress
  - Coverage statistics
  - Implementation details

### Architecture

- **[Overview](architecture/overview.md)**: High-level system architecture
- **[System Design](architecture/system-design.md)**: Design patterns and
  principles
- **[Data Flow](architecture/data-flow.md)**: Request/response flows
- **[Database Schema](architecture/database-schema.md)**: Complete database
  documentation

### API

- **[REST API](api/rest-api.md)**: All HTTP endpoints with examples
- **[WebSocket](api/websocket.md)**: Real-time updates and player protocol

### Deployment

- **[Docker](deployment/docker.md)**: Docker and docker-compose setup
- **[Production](deployment/production.md)**: Production deployment best
  practices

### Development

- **[Setup](development/setup.md)**: Development environment configuration
- **[Pre-commit](development/pre-commit.md)**: Pre-commit hooks and code quality

### User Guides

- **[User Guide](guides/user-guide.md)**: Complete user manual
- **[Configuration](guides/configuration.md)**: Configuration reference

## 🔍 Finding Information

### Quick Search by Keyword

- **Installation**: README, docker.md, ubuntu-install.md, production.md
- **Configuration**: configuration.md, docker.md, deployment guides
- **API Usage**: rest-api.md, websocket.md
- **Database**: database-schema.md, system-design.md
- **Testing**: TESTING.md, CI_CD.md, integration_test/README.md,
  e2e_test/README.md
- **Contributing**: CONTRIBUTING.md, development/setup.md,
  development/pre-commit.md
- **Architecture**: architecture/\* (4 comprehensive files)
- **Deployment**: deployment/\* (2 comprehensive files)
- **CI/CD**: CI_CD.md, .github/workflows/README.md

## 📝 Documentation Standards

All documentation follows these standards:

- **Clear Structure**: Table of contents for documents >500 lines
- **Code Examples**: Practical, tested examples included
- **Cross-References**: Links to related documentation
- **Accuracy**: Technical accuracy verified
- **Completeness**: Comprehensive coverage of topics
- **Maintenance**: Updated with code changes

## 🚀 Recent Updates

**Testing & CI/CD** (Latest):

- Added comprehensive testing guide (TESTING.md)
- Added CI/CD pipeline documentation (CI_CD.md)
- Added testing implementation progress (TESTING_PROGRESS.md)
- Updated README with test coverage badges
- Created CONTRIBUTING.md with detailed guidelines
- Added test-specific READMEs for integration and E2E tests

**Architecture & Design**:

- Complete architecture documentation
- System design patterns
- Data flow diagrams
- Database schema with relationships

**API Documentation**:

- All 40+ REST endpoints documented
- WebSocket protocol specification
- Request/response examples

## 📖 Contributing to Documentation

See [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines on:

- Documentation standards
- Writing style
- Code examples
- Cross-referencing
- Review process

## 🔗 External Resources

- [Go Documentation](https://golang.org/doc/)
- [Gin Web Framework](https://gin-gonic.com/docs/)
- [GORM ORM](https://gorm.io/docs/)
- [Docker Documentation](https://docs.docker.com/)
- [GitHub Actions](https://docs.github.com/en/actions)
- [Testing in Go](https://golang.org/doc/code.html#Testing)

## 📧 Documentation Feedback

Found an error or have suggestions? Please:

1. Check existing [issues](https://github.com/akhilrex/podgrab/issues)
1. Create a new issue with the "documentation" label
1. Submit a pull request with improvements

______________________________________________________________________

**Last Updated**: February 2026 **Documentation Version**: 2.0 (with
comprehensive testing coverage)
