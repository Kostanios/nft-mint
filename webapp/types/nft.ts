import {z} from "zod";

export interface NFTData {
  mintContractNFT: boolean
  strictMintContractCounter: number,
  strictMintContractTokenNames: string[]
}

export const StrictNftFormSchema = z.object({
  quantity: z.string().refine((val) => !Number.isNaN(parseInt(val, 10)), {
    message: "Expected number, received a string"
  }),
  names: z.string().array().min(1, {
    message: "CustomNames must consist at least 1 name.",
  }),
})
