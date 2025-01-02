# Chain
### Smart Contract Repository

Usage:
1) fill env file (you can take mine except Private Key from your sepolia wallet)
2) test your connection to sepolia 
```npm
npm run sepoliatest
```
3) Compile the smart contracts
```npm
npm run compile
```
4) deploy first contract, get it's address and put into MINT_NFT_CONTRACT_ADDRESS var
```npm
npm run deploymintnft
```
5) deploy second contract and save it's address somewhere - we need that to fill golang repo env
```npm
npm run deploystrictmintnft
```

6) pls go to go repo readme file

My deployed contracts
MintNFT Contract
```
0x4C953F0a78b313087C6196EEd372d0620B1E1356
```

StrictMintNFT Contract
```
0x99bc37b6Fd0b88CEA0541a5435f4e5da26982B67
```


