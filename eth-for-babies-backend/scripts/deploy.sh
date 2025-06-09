#!/bin/bash

# Exit on error
set -e

# Print commands
set -x

# Set variables
APP_NAME="eth-for-babies-backend"
BUILD_DIR="./build"
DEPLOY_DIR="/opt/eth-for-babies"
CONFIG_DIR="/etc/eth-for-babies"
LOG_DIR="/var/log/eth-for-babies"
SERVICE_NAME="eth-for-babies"

# Check if running as root
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root or with sudo"
  exit 1
fi

# Build the application
echo "Building application..."
./scripts/build.sh

# Create directories if they don't exist
echo "Creating directories..."
mkdir -p $DEPLOY_DIR
mkdir -p $CONFIG_DIR
mkdir -p $LOG_DIR

# Copy binary to deploy directory
echo "Copying binary to deploy directory..."
cp $BUILD_DIR/$APP_NAME $DEPLOY_DIR/

# Copy config files
echo "Copying config files..."
if [ -f ".env" ]; then
  cp .env $CONFIG_DIR/
fi

# Create systemd service file
echo "Creating systemd service file..."
cat > /etc/systemd/system/$SERVICE_NAME.service << EOF
[Unit]
Description=Eth For Babies Backend Service
After=network.target

[Service]
ExecStart=$DEPLOY_DIR/$APP_NAME
WorkingDirectory=$DEPLOY_DIR
EnvironmentFile=$CONFIG_DIR/.env
Restart=always
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=$SERVICE_NAME
User=root
Group=root

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd
echo "Reloading systemd..."
systemctl daemon-reload

# Enable and start service
echo "Enabling and starting service..."
systemctl enable $SERVICE_NAME
systemctl restart $SERVICE_NAME

# Check service status
echo "Service status:"
systemctl status $SERVICE_NAME --no-pager

echo "Deployment completed successfully!" 