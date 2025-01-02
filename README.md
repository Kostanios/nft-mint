# NFT Mint Project

# NFT Minting Project

This project involves the development of two smart contracts deployed on the Sepolia test network, a backend service written in GoLang, and a frontend interface using Next.js and TailwindCSS for minting and managing NFTs. Below is the detailed description of the project components and their functionalities.

## Smart Contracts

### Smart Contract 1: Mint NFT
- **Features:**
    - Allows minting of one NFT per wallet.
    - No associated minting cost.
    - Total supply capped at 100 NFTs.
    - Transaction fees are handled by the backend service.

### Smart Contract 2: Mint NFT with Additional Rules
- **Features:**
    - Minting is restricted to holders of an NFT from the first smart contract.
    - Maximum of 5 NFTs can be minted per transaction.
    - Total supply capped at 40 NFTs.
    - **Minting Costs:**
        - First 15 NFTs: 0.00000005 ETH each.
        - Remaining NFTs: 0.0006 ETH each.
    - Optional custom name for the NFT can be purchased for an additional cost of 0.0004 ETH.

## Backend Service
- **Technology:** GoLang
- **Features:**
    - API endpoint to facilitate minting of NFTs from the first smart contract.
    - Management of transaction fees.
    - Real-time monitoring of on-chain events from the second smart contract to:
        - Synchronize changes with the smart contract.

## Frontend Interface
- **Technology:** Next.js, TailwindCSS, Shadcn/ui
- **Features:**
    - User-friendly interface to interact with the backend and mint NFTs.
    - Display purchased NFTs, including details such as:
        - Name
        - Features
        - Updated customization status

## Additional Integrations
- **Pinata Integration:**
    - Upload and manage updated JSON metadata files for NFTs.

- **Event Management:**
    - Listen and respond to on-chain events related to NFT name changes and other updates.

## Development Stack
- **Backend:** GoLang
- **Frontend:** Next.js, TailwindCSS, Shadcn/ui

---
**Note:** This README serves as a high-level overview of the project requirements and features. Detailed implementation instructions and codebase documentation can be found within the repository.

# SIMPLE USAGE

1) create env files for chain, webapi, webapp with .example files and add your private key
2) run docker compose file and open localhost 3000
```
docker-compose up -d
```

# Advanced Setup

You can go with step-by-step configuration from CHAIN README.md file

