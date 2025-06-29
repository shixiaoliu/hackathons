# ETH for Babies Backend Deployment Guide

This document provides instructions for deploying the ETH for Babies backend service.

## Prerequisites

Before deploying, ensure you have the following:

- Go 1.19 or later
- SQLite 3
- Git
- Access to an Ethereum node (Infura, Alchemy, or your own node)
- Private key for deploying contracts

## Environment Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/eth-for-babies-backend.git
cd eth-for-babies-backend
```

2. Create a `.env` file in the root directory with the following variables:

```
# Server Configuration
PORT=8080
ENV=production
LOG_LEVEL=info
JWT_SECRET=your-jwt-secret-key
JWT_EXPIRATION=24h

# Database Configuration
DB_PATH=./data/database.db

# Blockchain Configuration
BLOCKCHAIN_RPC_URL=https://sepolia.infura.io/v3/your-project-id
BLOCKCHAIN_PRIVATE_KEY=your-private-key
BLOCKCHAIN_CHAIN_ID=11155111  # Sepolia testnet

# Contract Addresses (leave empty for initial deployment)
CONTRACT_TASK_REGISTRY=
CONTRACT_FAMILY_REGISTRY=
CONTRACT_REWARD_TOKEN=
```

Replace the placeholder values with your actual configuration.

## Building the Application

Use the provided build script to compile the application:

```bash
chmod +x ./scripts/build.sh
./scripts/build.sh
```

This will create a binary in the `./build` directory.

## Database Initialization

The application will automatically create and initialize the database on first run. However, if you want to manually initialize the database, you can use the following commands:

```bash
# Create database directory
mkdir -p ./data

# Initialize database with schema
sqlite3 ./data/database.db < ./migrations/001_create_users_table.sql
sqlite3 ./data/database.db < ./migrations/002_create_families_table.sql
sqlite3 ./data/database.db < ./migrations/003_create_children_table.sql
sqlite3 ./data/database.db < ./migrations/004_create_tasks_table.sql
```

## Smart Contract Deployment

If you're deploying for the first time and need to deploy the smart contracts:

1. Run the application with the `--deploy-contracts` flag:

```bash
./build/eth-for-babies-backend --deploy-contracts
```

2. After successful deployment, the contract addresses will be printed in the logs. Update your `.env` file with these addresses.

## Running the Application

### Local Development

For local development, you can run the application directly:

```bash
go run cmd/server/main.go
```

### Production Deployment

For production deployment, use the provided deploy script (requires root/sudo privileges):

```bash
sudo ./scripts/deploy.sh
```

This will:
1. Build the application
2. Copy the binary to `/opt/eth-for-babies/`
3. Create a systemd service
4. Start the service

### Docker Deployment

You can also deploy using Docker:

1. Build the Docker image:

```bash
docker build -t eth-for-babies-backend .
```

2. Run the container:

```bash
docker run -d \
  --name eth-for-babies \
  -p 8080:8080 \
  -v $(pwd)/data:/app/data \
  -v $(pwd)/.env:/app/.env \
  eth-for-babies-backend
```

Or use Docker Compose:

```bash
docker-compose up -d
```

## Health Check

To verify that the application is running correctly, make a request to the health check endpoint:

```bash
curl http://localhost:8080/health
```

You should receive a response like:

```json
{
  "status": "ok",
  "version": "1.0.0",
  "timestamp": "2023-06-01T12:00:00Z"
}
```

## Troubleshooting

### Service Not Starting

If the service doesn't start, check the logs:

```bash
journalctl -u eth-for-babies.service
```

### Database Issues

If you encounter database issues, you can reset the database:

```bash
rm ./data/database.db
```

The application will recreate the database on the next start.

### Contract Deployment Failures

If contract deployment fails:

1. Check that your RPC URL is correct and the node is accessible
2. Ensure your private key has enough ETH for gas fees
3. Verify that the chain ID matches the network you're deploying to

## Backup and Restore

### Database Backup

To backup the database:

```bash
cp ./data/database.db ./data/database.db.backup
```

### Database Restore

To restore from a backup:

```bash
cp ./data/database.db.backup ./data/database.db
```

## Updating the Application

To update the application:

1. Pull the latest changes:

```bash
git pull origin main
```

2. Rebuild and redeploy:

```bash
./scripts/build.sh
sudo ./scripts/deploy.sh
```

## Security Considerations

- Store your `.env` file securely and restrict access
- Use a strong JWT secret
- Keep your blockchain private key secure
- Consider using a hardware wallet for production deployments
- Regularly update dependencies to patch security vulnerabilities 