// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MintNFT is ERC721 {
    uint256 public constant MAX_SUPPLY = 100;
    uint256 private _tokenIdCounter = 0;
    mapping(address => bool) public hasMinted;

    constructor() ERC721("MintNFT", "MNT") {}

    function mint() external {
        require(!hasMinted[msg.sender], "One NFT per wallet allowed.");
        require(_tokenIdCounter < MAX_SUPPLY, "Max supply reached.");

        hasMinted[msg.sender] = true;
        uint256 tokenId = _tokenIdCounter;
        _tokenIdCounter++;
        _safeMint(msg.sender, tokenId);
    }
}