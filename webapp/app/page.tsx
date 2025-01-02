"use client"

import { Button, buttonVariants } from "@/components/ui/button"
import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious } from "@/components/ui/carousel";
import { Card, CardContent } from "@/components/ui/card";
import { useCallback, useEffect, useState } from "react";
import { NFTData, StrictNftFormSchema } from "@/types/nft";
import {fetchNFTData, mintNFT, strictMintNFT} from "@/fetchers/nft";
import { FormProvider, useForm } from "react-hook-form";
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod";
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";


export default function IndexPage() {
  const [nftData, setNftData] = useState<null | NFTData>(null);
  const [nftDataLoading, setNftDataLoading] = useState(false);
  const [mintLoading, setMintLoading] = useState(false);
  const [mintError, setMintError] = useState(false);
  const [strictMintError, setStrictMintError] = useState(false)

  const form = useForm<z.infer<typeof StrictNftFormSchema>>({
    resolver: zodResolver(StrictNftFormSchema),
    defaultValues: {
      quantity: '1',
      names: []
    },
  })

  const fetchAfterSleep = () => setTimeout(() => fetchNFTData().then(data => {
    setNftData(data);
    setNftDataLoading(false);
  }), 8000)

  function onSubmit(data: z.infer<typeof StrictNftFormSchema>) {
    setMintLoading(true);
    strictMintNFT(data).then((res) => {
      if (res.ok) {
        setNftDataLoading(true);
        fetchAfterSleep();
        setStrictMintError(false);
      } else {
        setStrictMintError(true);
      }
      setMintLoading(false);
    })
  }

  useEffect(() => {
    setNftDataLoading(true)
    fetchNFTData().then(data => {
      setNftData(data);
      setNftDataLoading(false);
    });

  }, [])

  const mintCallback = useCallback(async () => {
    setMintLoading(true);
    await mintNFT().then((res) => {
      if (res.ok) {
        fetchAfterSleep();
        setMintError(false);
        setNftDataLoading(true);
      } else {
        setMintError(true);
      }
    });
    setMintLoading(false);
  } ,[])

  const strictMintContractCounter = nftData?.strictMintContractCounter || 0;

  return (
    <section className="container grid items-center gap-6 pb-8 pt-6 md:py-10 min-h-80">
      <div className="flex flex-col items-center gap-2">
        <h1 className="text-3xl font-bold leading-tight tracking-tighter md:text-4xl pb-10">
          {`Mint NFT Contract${nftData?.mintContractNFT ? ' (Received) ' : ''}${nftDataLoading && !nftData?.mintContractNFT ? ': Awaiting Transaction' : ''}`}
        </h1>
        <Button disabled={mintLoading || nftDataLoading} className={buttonVariants()} onClick={mintCallback}>
          {!mintLoading ? "Mint NFT" : "Loading"}
        </Button>
        <p className="p-5 h-5 text-red-600 text-center max-w-96">{mintError ? 'Transaction error, please check your funds for gas or your wallet already has a NFT' : ''}</p>
      </div>

      <div className="h-1 bg-amber-50 w-full"/>

      {nftData?.mintContractNFT && <div className="flex flex-col items-center gap-2">
        <h1 className="text-3xl font-bold leading-tight tracking-tighter md:text-4xl pb-10">
          {`Strict Mint Contract NFT${nftDataLoading && strictMintContractCounter > 0 ? ': Awaiting Transaction' : ''}`}
        </h1>
        {strictMintContractCounter > 0
          ? <Carousel className="w-full max-w-xl min-h-40">
            <CarouselContent>
              {Array.from({length: strictMintContractCounter}).map((_, index) => (
                <CarouselItem className="basis-1/4" key={index}>
                  <div className="p-1">
                    <Card>
                      <CardContent className="flex flex-col aspect-square items-center justify-center p-4">
                        <div className="text-xl font-semibold">{`NFT ${index}`}</div>
                        <div
                          className="font-semibold text-xs text-ellipsis overflow-hidden max-w-16">{nftData?.strictMintContractTokenNames[index]}</div>
                      </CardContent>
                    </Card>
                  </div>
                </CarouselItem>
              ))}
            </CarouselContent>
            <CarouselPrevious/>
            <CarouselNext/>
          </Carousel>
          : <p className="min-h-40 flex items-center text-5xl"><span>{nftDataLoading ? 'Loading' : 'Nothing'}</span></p>
        }
        <div className="flex justify-center">
          <FormProvider {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6 flex gap-5 items-end">
              <FormField
                control={form.control}
                name="quantity"
                render={({field}) => (
                  <FormItem>
                    <FormLabel>Quantity</FormLabel>
                    <FormControl>
                      <Input min={1} max={5} type="number" placeholder="shadcn" {...field}
                             onClick={() => form.trigger('names')}/>
                    </FormControl>
                    <FormMessage/>
                  </FormItem>
                )}
              />
              {Array.from({length: Number(form.getValues("quantity").valueOf())}).map((_, index) => (
                <FormItem key={index} className="min-w-6">
                  <FormLabel>{`NFT ${index}`}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={`NFT ${index} name`}
                      onChange={(event) => {
                        const names = form.getValues("names").valueOf() as string[];
                        names[index] = event.target.value;
                        form.setValue("names", names);
                      }}
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              ))}
              <Button type="submit" disabled={mintLoading} className={buttonVariants()}>
                {!mintLoading ? "Strict Mint NFT" : "Loading"}
              </Button>
            </form>
          </FormProvider>
        </div>
        <p
          className="p-5 h-5 text-red-600 text-center max-w-96"
        >
          {strictMintError ? 'Transaction error, your have not enough funds for gas/fee or nft limit reached' : ''}
        </p>
      </div>}
    </section>
  )
}
