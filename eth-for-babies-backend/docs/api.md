# ETH for Babies API Documentation

This document describes the API endpoints available in the ETH for Babies backend.

## Base URL

All API endpoints are prefixed with `/api/v1`.

## Authentication

Most endpoints require authentication. To authenticate, include an `Authorization` header with a bearer token:

```
Authorization: Bearer <token>
```

You can obtain a token by logging in via the `/api/v1/auth/login` endpoint.

## Error Handling

All API responses follow a consistent format:

```json
{
  "success": true|false,
  "message": "Description of the result or error",
  "data": { ... } // Optional data object
}
```

HTTP status codes are used appropriately:
- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error

## Endpoints

### Authentication

#### Register User

```
POST /api/v1/auth/register
```

Create a new user account.

**Request Body:**
```json
{
  "username": "string",
  "password": "string",
  "email": "string",
  "walletAddress": "string"
}
```

**Response:**
```json
{
  "success": true,
  "message": "User registered successfully",
  "user": {
    "id": 1,
    "username": "string",
    "email": "string",
    "walletAddress": "string",
    "createdAt": "timestamp"
  }
}
```

#### Login

```
POST /api/v1/auth/login
```

Authenticate and get a token.

**Request Body:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "token": "jwt-token-string",
  "user": {
    "id": 1,
    "username": "string",
    "email": "string",
    "walletAddress": "string",
    "role": "string"
  }
}
```

### Family Management

#### Create Family

```
POST /api/v1/families
```

Create a new family.

**Request Body:**
```json
{
  "name": "string"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Family created successfully",
  "family": {
    "id": 1,
    "name": "string",
    "parentId": 1,
    "createdAt": "timestamp"
  }
}
```

#### Get Families

```
GET /api/v1/families
```

Get all families for the authenticated user.

**Response:**
```json
{
  "success": true,
  "families": [
    {
      "id": 1,
      "name": "string",
      "parentId": 1,
      "createdAt": "timestamp"
    }
  ]
}
```

#### Get Family by ID

```
GET /api/v1/families/:id
```

Get a specific family by ID.

**Response:**
```json
{
  "success": true,
  "family": {
    "id": 1,
    "name": "string",
    "parentId": 1,
    "createdAt": "timestamp",
    "children": [
      {
        "id": 1,
        "name": "string",
        "age": 10,
        "walletAddress": "string"
      }
    ]
  }
}
```

### Child Management

#### Add Child to Family

```
POST /api/v1/children
```

Add a child to a family.

**Request Body:**
```json
{
  "name": "string",
  "age": 10,
  "familyId": 1,
  "walletAddress": "string",
  "avatar": "string" // optional
}
```

**Response:**
```json
{
  "success": true,
  "message": "Child added successfully",
  "child": {
    "id": 1,
    "name": "string",
    "age": 10,
    "familyId": 1,
    "walletAddress": "string",
    "avatar": "string",
    "createdAt": "timestamp"
  }
}
```

#### Get Children by Family

```
GET /api/v1/families/:familyId/children
```

Get all children in a family.

**Response:**
```json
{
  "success": true,
  "children": [
    {
      "id": 1,
      "name": "string",
      "age": 10,
      "familyId": 1,
      "walletAddress": "string",
      "avatar": "string",
      "createdAt": "timestamp"
    }
  ]
}
```

#### Get Child by ID

```
GET /api/v1/children/:id
```

Get a specific child by ID.

**Response:**
```json
{
  "success": true,
  "child": {
    "id": 1,
    "name": "string",
    "age": 10,
    "familyId": 1,
    "walletAddress": "string",
    "avatar": "string",
    "createdAt": "timestamp"
  }
}
```

#### Get Child by Wallet Address

```
GET /api/v1/children/wallet/:address
```

Get a child by wallet address.

**Response:**
```json
{
  "success": true,
  "child": {
    "id": 1,
    "name": "string",
    "age": 10,
    "familyId": 1,
    "walletAddress": "string",
    "avatar": "string",
    "createdAt": "timestamp"
  }
}
```

### Task Management

#### Create Task

```
POST /api/v1/tasks
```

Create a new task.

**Request Body:**
```json
{
  "title": "string",
  "description": "string",
  "reward": "0.01",
  "familyId": 1
}
```

**Response:**
```json
{
  "success": true,
  "message": "Task created successfully",
  "task": {
    "id": 1,
    "title": "string",
    "description": "string",
    "reward": "0.01",
    "status": "available",
    "creatorId": 1,
    "familyId": 1,
    "createdAt": "timestamp"
  }
}
```

#### Get Tasks by Family

```
GET /api/v1/families/:familyId/tasks
```

Get all tasks in a family.

**Response:**
```json
{
  "success": true,
  "tasks": [
    {
      "id": 1,
      "title": "string",
      "description": "string",
      "reward": "0.01",
      "status": "string",
      "creatorId": 1,
      "assignedTo": 1,
      "familyId": 1,
      "createdAt": "timestamp",
      "updatedAt": "timestamp"
    }
  ]
}
```

#### Get Task by ID

```
GET /api/v1/tasks/:id
```

Get a specific task by ID.

**Response:**
```json
{
  "success": true,
  "task": {
    "id": 1,
    "title": "string",
    "description": "string",
    "reward": "0.01",
    "status": "string",
    "creatorId": 1,
    "assignedTo": 1,
    "familyId": 1,
    "proof": "string",
    "createdAt": "timestamp",
    "updatedAt": "timestamp"
  }
}
```

#### Assign Task

```
PUT /api/v1/tasks/:id/assign
```

Assign a task to a child.

**Request Body:**
```json
{
  "childId": 1
}
```

**Response:**
```json
{
  "success": true,
  "message": "Task assigned successfully",
  "task": {
    "id": 1,
    "status": "in-progress",
    "assignedTo": 1,
    "updatedAt": "timestamp"
  }
}
```

#### Complete Task

```
PUT /api/v1/tasks/:id/complete
```

Mark a task as completed by a child.

**Request Body:**
```json
{
  "proof": "string" // optional
}
```

**Response:**
```json
{
  "success": true,
  "message": "Task completed successfully",
  "task": {
    "id": 1,
    "status": "completed",
    "proof": "string",
    "updatedAt": "timestamp"
  }
}
```

#### Approve Task

```
PUT /api/v1/tasks/:id/approve
```

Approve a completed task and transfer reward.

**Response:**
```json
{
  "success": true,
  "message": "Task approved and reward transferred",
  "task": {
    "id": 1,
    "status": "approved",
    "updatedAt": "timestamp"
  },
  "transaction": {
    "hash": "string",
    "amount": "0.01"
  }
}
```

#### Reject Task

```
PUT /api/v1/tasks/:id/reject
```

Reject a completed task.

**Request Body:**
```json
{
  "reason": "string"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Task rejected",
  "task": {
    "id": 1,
    "status": "rejected",
    "updatedAt": "timestamp"
  }
}
```

### Contract Interaction

#### Get Contract Addresses

```
GET /api/v1/contracts/addresses
```

Get the addresses of deployed smart contracts.

**Response:**
```json
{
  "success": true,
  "addresses": {
    "TaskRegistry": "0x...",
    "FamilyRegistry": "0x...",
    "RewardToken": "0x..."
  }
}
```

#### Get Child Balance

```
GET /api/v1/children/:id/balance
```

Get the ETH balance of a child's wallet.

**Response:**
```json
{
  "success": true,
  "child": {
    "id": 1,
    "name": "string",
    "walletAddress": "0x..."
  },
  "balance": "0.05"
}
``` 