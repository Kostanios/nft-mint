// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./MintNFT.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract StrictMintNFT is ERC721, Ownable {
    uint256 public constant MAX_SUPPLY = 40;

    uint256 public firstTierPrice = 0.00000005 ether;
    uint256 public secondTierPrice = 0.0006 ether;
    uint256 public customNamePrice = 0.0004 ether;
    uint256 public tokenIdCounter = 0;

    uint256 public firstTierRemaining = 15;
    mapping(uint256 => string) public customNames;
    mapping(address => uint256) public mintCount;

    MintNFT public mintNFTContract;

    event MintedNFT(uint256 indexed tokenId, address indexed owner, string customName);

    constructor(address mintNFTContractAddress) Ownable(msg.sender) ERC721("StrictMintNFT", "MNTR") {
        mintNFTContract = MintNFT(mintNFTContractAddress);
    }

    function mint(uint256 quantity, string[] calldata names) external payable {
        require(quantity > 0 && quantity <= 5, "You can mint between 1 and 5 NFTs.");
        require(tokenIdCounter + quantity <= MAX_SUPPLY, "Exceeds max supply.");
        require(mintNFTContract.balanceOf(msg.sender) > 0, "You must own an NFT from the first contract.");

        uint256 totalCost = 0;

        for (uint256 i = 0; i < quantity; i++) {
            if (firstTierRemaining > 0) {
                totalCost += firstTierPrice;
                firstTierRemaining--;
            } else {
                totalCost += secondTierPrice;
            }
        }

        for (uint256 j = 0; j < names.length; j++) {
            totalCost += customNamePrice;
        }

        require(msg.value >= totalCost, "Insufficient funds sent.");

        for (uint256 k = 0; k < quantity; k++) {
            uint256 tokenId = tokenIdCounter;
            tokenIdCounter++;
            _safeMint(msg.sender, tokenId);

            string memory customName = k < names.length ? names[k] : "";
            customNames[tokenId] = customName;

            emit MintedNFT(tokenId, msg.sender, customName);
        }
    }

    function setPrices(uint256 _firstTierPrice, uint256 _secondTierPrice, uint256 _customNamePrice) external onlyOwner {
        firstTierPrice = _firstTierPrice;
        secondTierPrice = _secondTierPrice;
        customNamePrice = _customNamePrice;
    }

    function setCustomName(uint256 tokenId, string calldata name) external payable {
        require(ownerOf(tokenId) == msg.sender, "You do not own this NFT.");
        require(msg.value >= customNamePrice, "Insufficient funds sent for name change.");

        customNames[tokenId] = name;

        emit MintedNFT(tokenId, msg.sender, name);
    }
}
