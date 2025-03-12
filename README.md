# DeFi Swap

A decentralized finance application for swapping tokens on the blockchain.

## Overview

DeFi Swap is a comprehensive decentralized exchange platform that combines Ethereum smart contracts with a Go backend service. It enables users to swap various ERC20 tokens in a secure, non-custodial environment.

## Features

- Token swapping with automatic price discovery
- ERC20 token implementation
- Secure transaction processing
- API endpoints for blockchain interactions
- User authentication and authorization

## Technology Stack

### Smart Contracts
- Solidity ^0.8.20
- Hardhat development environment
- OpenZeppelin contract libraries

### Backend
- Go programming language
- RESTful API architecture
- Environment-based configuration

## Requirements

- Node.js and npm
- Go 1.x
- Ethereum wallet (MetaMask recommended)
- Hardhat

## Installation

### Smart Contract Setup

1. Install dependencies:
   ```bash
   npm install
   ```

2. Compile the contracts:
   ```bash
   npx hardhat compile
   ```

3. Run tests:
   ```bash
   npx hardhat test
   ```

4. Deploy contracts:
   ```bash
   npx hardhat ignition deploy ./ignition/modules/Lock.js
   ```

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Configure environment variables:
   - Copy `.env.example` to `.env`
   - Update the variables with your configuration

4. Run the server:
   ```bash
   go run main.go
   ```
   The server will run on port 8080 by default.

## Project Structure
├── contracts/ # Smart contract code
│ └── Token.sol # ERC20 token implementation
├── backend/ # Go backend service
│ ├── config/ # Configuration handling
│ ├── handlers/ # HTTP request handlers
│ ├── routes/ # API route definitions
│ ├── services/ # Business logic
│ └── utils/ # Utility functions
├── scripts/ # Deployment and utility scripts
└── test/ # Contract test files



## Usage

1. Connect your Ethereum wallet to the application
2. Approve tokens for swapping
3. Select the token pair and amount for swapping
4. Confirm the transaction in your wallet
5. View transaction history and token balances

## Development

To contribute to this project:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

ISC