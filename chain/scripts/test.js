async function main() {
    const provider = new ethers.providers.JsonRpcProvider(process.env.SEPOLIA_API_URL);
    const network = await provider.getNetwork();
    console.log("Connected to network:", network);
}

main().catch((error) => {
    console.error(error);
    process.exit(1);
});