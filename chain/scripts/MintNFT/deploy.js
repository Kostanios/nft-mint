/* global ethers */

async function main() {
    const MySepolia = await ethers.getContractFactory("MintNFT");
    const MySepoliaContract = await MySepolia.deploy();
    console.info("Contract deployed to address:", MySepoliaContract.address);
}

main().then(() =>
    process.exit(0)
).catch((error) => {
    console.error(error);
    process.exit(1);
});

