# Airbnb Booking Service

A robust and scalable booking management service for Airbnb-like platforms, built with Express.js, TypeScript, and modern backend technologies.

## Overview

This service handles booking operations including creation, confirmation, and management of hotel reservations. It implements idempotency patterns, distributed locking, and queue-based email notifications for reliable booking processing.

## Features

- **Booking Management**: Create and confirm bookings with proper validation
- **Idempotency**: Prevent duplicate booking operations using idempotency keys
- **Distributed Locking**: Redis-based locking mechanism to prevent race conditions
- **Queue System**: BullMQ integration for asynchronous email notifications
- **Logging**: Winston-based structured logging with daily rotation
- **Database**: MySQL with Prisma ORM for type-safe database operations
- **API Versioning**: Support for multiple API versions (v1, v2)
- **Type Safety**: Full TypeScript implementation with Zod validation

## Tech Stack

- **Runtime**: Node.js with TypeScript
- **Framework**: Express.js
- **Database**: MySQL with Prisma ORM
- **Cache & Locking**: Redis with Redlock
- **Queue**: BullMQ with Redis
- **Logging**: Winston with daily rotation
- **Validation**: Zod schemas
- **Development**: Nodemon with tsx

## Prerequisites

- Node.js (v18 or higher)
- MySQL database
- Redis server
- Git

## Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url> BookingService
   cd BookingService
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the root directory with the following variables:
   ```env
   PORT=3000
   DATABASE_URL="mysql://username:password@localhost:3306/booking_db"
   REDIS_SERVER_URL="redis://localhost:6379"
   LOCK_TTL=5000
   ```

4. **Set up the database**
   ```bash
   npx prisma generate
   npx prisma db push
   ```

5. **Start the development server**
   ```bash
   npm run dev
   ```

## API Endpoints

### Health Check
- `GET /api/v1/ping` - Service health check

### Bookings
- `POST /api/v1/bookings` - Create a new booking
- `POST /api/v1/bookings/confirm/:idempotencyKey` - Confirm a booking

## API Documentation

### Create Booking
```http
POST /api/v1/bookings
Content-Type: application/json

{
  "userId": 1,
  "hotelId": 1,
  "bookingAmount": 10000,
  "totalGuests": 2,
  "idempotencyKey": "unique-uuid-string"
}
```

### Confirm Booking
```http
POST /api/v1/bookings/confirm/{idempotencyKey}
```

## Workflow

### Booking Creation Flow

1. **Request Reception**
   - Client sends POST request to `/api/v1/bookings` with booking details
   - Request passes through correlation ID middleware for tracking
   - Request body is validated using Zod schemas

2. **Idempotency Check**
   - System checks if the provided idempotency key already exists
   - If key exists, returns the existing booking response (prevents duplicates)
   - If key is new, proceeds with booking creation

3. **Distributed Locking**
   - Acquires Redis-based lock on the booking resource
   - Prevents race conditions when multiple requests target the same booking
   - Lock timeout is configurable via `LOCK_TTL` environment variable

4. **Database Operations**
   - Creates booking record with PENDING status
   - Creates corresponding idempotency key record
   - All operations are transactional to ensure data consistency

5. **Queue Integration**
   - Adds email notification to BullMQ queue for asynchronous processing
   - Email contains booking details and confirmation instructions
   - Queue processing happens in the background

6. **Response**
   - Returns booking details with unique booking ID
   - Includes idempotency key for future reference
   - Logs the entire operation with correlation ID

### Booking Confirmation Flow

1. **Request Reception**
   - Client sends POST request to `/api/v1/bookings/confirm/{idempotencyKey}`
   - System validates the idempotency key format and existence

2. **Idempotency Validation**
   - Retrieves the idempotency key record from database
   - Checks if the booking has already been confirmed
   - Prevents duplicate confirmations

3. **Status Update**
   - Updates booking status from PENDING to CONFIRMED
   - Marks idempotency key as finalized
   - Timestamps the confirmation operation

4. **Notification Queue**
   - Adds confirmation email to queue
   - Includes booking confirmation details and next steps
   - Processes asynchronously for better performance

5. **Response**
   - Returns confirmation details with updated status
   - Provides booking reference information

### Error Handling and Recovery

- **Validation Errors**: Immediate response with detailed error messages
- **Database Errors**: Automatic rollback and error logging
- **Queue Failures**: Retry mechanism with exponential backoff
- **Lock Timeouts**: Automatic lock release and retry suggestion
- **Network Issues**: Circuit breaker pattern for external dependencies

### Logging and Monitoring

- Every request is logged with a unique correlation ID
- Structured logging includes timestamps, request details, and outcomes
- Daily log rotation prevents disk space issues
- Error alerts are sent to monitoring systems

## Database Schema

The service uses the following main entities:

- **Booking**: Stores booking information with status tracking
- **IdempotencyKey**: Ensures idempotent operations
- **BookingStatus**: Enum (PENDING, CONFIRMED, CANCELLED)

## Project Structure

```
src/
├── config/          # Configuration files
├── controllers/     # Request handlers
├── dto/            # Data transfer objects
├── middlewares/    # Express middlewares
├── prisma/         # Database schema and migrations
├── producers/      # Queue producers
├── queues/         # Queue configurations
├── repositories/   # Data access layer
├── routers/        # API routes
├── services/       # Business logic
├── types/          # TypeScript type definitions
├── utils/          # Utility functions
└── validators/     # Request validation schemas
```

## Available Scripts

- `npm run dev` - Start development server with hot reload
- `npm run build` - Compile TypeScript to JavaScript
- `npm start` - Start production server

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Server port | 3000 |
| DATABASE_URL | MySQL connection string | - |
| REDIS_SERVER_URL | Redis connection string | redis://localhost:6379 |
| LOCK_TTL | Lock time-to-live in milliseconds | 5000 |

## Development

The service includes comprehensive error handling, logging, and validation. All operations are logged with correlation IDs for better debugging and monitoring.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the ISC License.