export const fetchNFTData = () => fetch(`${process.env.WEBAPI}/list-nft`, { method: 'POST' })
  .then(res => res.json())

export const mintNFT = () => fetch(`${process.env.WEBAPI}/nft-mint`, { method: 'POST' })

export const strictMintNFT = (payload: { quantity: string, names: string[] }) => fetch(`${process.env.WEBAPI}/strict-nft-mint`, { method: 'POST', body: JSON.stringify({ ...payload, quantity: Number(payload.quantity)}) })
