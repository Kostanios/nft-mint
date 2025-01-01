/* global ethers */
require('dotenv').config();
const { MINT_NFT_CONTRACT_ADDRESS } = process.env;

async function main() {
    const MySepolia = await ethers.getContractFactory("StrictMintNFT");
    const MySepoliaContract = await MySepolia.deploy(MINT_NFT_CONTRACT_ADDRESS);
    console.info("Contract deployed to address:", MySepoliaContract.address);
}

main().then(() =>
    process.exit(0)
).catch((error) => {
    console.error(error);
    process.exit(1);
});

