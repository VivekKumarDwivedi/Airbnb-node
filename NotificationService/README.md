# Notification Service

A robust email notification service built with Node.js, Express, TypeScript, and Redis Queue (BullMQ) for handling asynchronous email processing.

## 🚀 Features

- **Asynchronous Email Processing**: Uses BullMQ with Redis for reliable queue-based email processing
- **Template Engine**: Handlebars for dynamic email templates
- **TypeScript Support**: Full type safety with TypeScript
- **Structured Logging**: Winston logger with daily rotation
- **Error Handling**: Comprehensive error handling middleware
- **Correlation IDs**: Request tracking with correlation IDs
- **Environment Configuration**: Secure environment-based configuration

## 📋 Prerequisites

- Node.js (v18 or higher)
- Redis server
- SMTP email service (Gmail, SendGrid, etc.)

## 🛠️ Installation

```bash
# Install dependencies
npm install

# Copy environment variables
cp .env.example .env

# Update .env with your configuration
```

## ⚙️ Environment Variables

Create a `.env` file with the following variables:

```env
PORT=3000
REDIS_HOST=localhost
REDIS_PORT=6379
MAIL_USER=your-email@gmail.com
MAIL_PASS=your-app-password
```

## 🏗️ Architecture

### Core Components

1. **Server** (`src/server.ts`): Express server setup with middleware and routes
2. **Queue System**: BullMQ with Redis for job processing
3. **Email Service**: Nodemailer integration for SMTP sending
4. **Template Engine**: Handlebars for dynamic email templates
5. **Logger**: Winston for structured logging


## 🔄 Workflows

### Project Structure Workflow

For a detailed understanding of the project structure and navigation patterns, see the [Project Structure Workflow]

#### Quick Structure Overview

```
NotificationService/
├── src/
│   ├── config/           # Configuration files (server, logger, mailer, redis)
│   ├── controllers/      # Route controllers
│   ├── dto/             # Data Transfer Objects
│   ├── middlewares/     # Express middlewares (error, correlation)
│   ├── processors/      # Queue job processors
│   ├── producers/       # Queue job producers
│   ├── queues/          # Queue configurations
│   ├── routers/         # API routes (v1, v2)
│   ├── services/        # Business logic services
│   ├── templates/       # Email templates (Handlebars)
│   ├── utils/           # Utility functions
│   ├── validators/      # Input validation
│   └── server.ts        # Main application entry point
├── .windsurf/
│   └── workflows/       # Development workflows
├── logs/                # Application logs (created at runtime)
├── package.json         # Dependencies and scripts
├── tsconfig.json        # TypeScript configuration
└── README.md           # This file
```

#### Data Flow

1. **Request Flow**: `Client → Router → Controller → Producer → Queue → Processor → Service → Email`
2. **Queue Processing**: `Producer → Redis Queue → Worker → Processor → Service → Email`

#### Key Components

- **NotificationDto**: Core data structure for email notifications
- **Queue System**: BullMQ with Redis for async processing
- **Template Engine**: Handlebars for dynamic email content
- **Logger**: Winston with daily rotation
- **Error Handling**: Comprehensive middleware stack

#### Development Patterns

- **Adding Templates**: Create `.hbs` files in `src/templates/mailer/`
- **Adding Queue Jobs**: Define DTO → Create Producer → Implement Processor → Configure Queue
- **Adding API Endpoints**: Create Controller → Add Routes → Add Validation

## 📧 Email Templates

Email templates are stored in `src/templates/mailer/` using Handlebars syntax:

Example template (`welcome.hbs`):
```handlebars
Hi {{name}},

Welcome to {{appName}}!
We are excited to have you on board...

Thanks for joining us!
Best regards.
```

## 🔄 Queue Processing

The service uses a producer-consumer pattern:

1. **Producer**: Adds email jobs to the queue
2. **Consumer**: Processes jobs asynchronously
3. **Worker**: Handles email sending with retry logic

### Notification DTO

```typescript
interface NotificationDto {
    to: string;           // Recipient email
    subject: string;      // Email subject
    templateId: string;   // Template identifier
    params: Record<string, any>; // Template parameters
}
```

## 🚀 Usage

### Starting the Server

```bash
# Development mode
npm run dev

# Production mode
npm start

# Build TypeScript
npm run build
```

### Sending Emails

```typescript
import { addEmailToQueue } from './producers/email.producer';

// Add email to queue
await addEmailToQueue({
    to: "user@example.com",
    subject: "Welcome to our service",
    templateId: "welcome",
    params: {
        name: "John Doe",
        appName: "Booking App"
    }
});
```

## 🔧 Configuration

### Redis Configuration
- Host: Configured via `REDIS_HOST` environment variable
- Port: Configured via `REDIS_PORT` environment variable
- Default: `localhost:6379`

### Mailer Configuration
- Uses Nodemailer with SMTP transport
- Configured via `MAIL_USER` and `MAIL_PASS` environment variables
- Supports any SMTP provider (Gmail, SendGrid, etc.)

## 📊 Monitoring & Logging

- **Winston Logger**: Structured logging with daily rotation
- **Queue Events**: Job completion and failure events
- **Correlation IDs**: Request tracking across the system

Log files are stored in the `logs/` directory with daily rotation.

## 🔒 Security

- Environment variables for sensitive configuration
- Input validation with Zod schemas
- Error handling without information leakage
- Correlation ID middleware for request tracking

## 📝 API Endpoints

### Health Check
```
GET /api/v1/ping
```

Returns service health status.

## 🧪 Development

### Adding New Templates

1. Create `.hbs` file in `src/templates/mailer/`
2. Use Handlebars syntax for dynamic content
3. Reference by filename (without extension) in `templateId`

### Adding New Queue Jobs

1. Define DTO in `src/dto/`
2. Create producer in `src/producers/`
3. Implement processor in `src/processors/`
4. Configure queue in `src/queues/`

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the ISC License.

## 🔗 Dependencies

- **Express**: Web framework
- **BullMQ**: Queue system
- **Nodemailer**: Email sending
- **Handlebars**: Template engine
- **Winston**: Logging
- **TypeScript**: Type safety
- **Zod**: Schema validation
- **ioredis**: Redis client

## 📞 Support

For issues and questions, please create an issue in the repository.
