export const env = {
  chainId: Number(process.env.NEXT_PUBLIC_CHAIN_ID ?? 97),
  rpcUrl:
    process.env.NEXT_PUBLIC_RPC_URL ??
    "https://data-seed-prebsc-1-s1.binance.org:8545",
  backendUrl: process.env.NEXT_PUBLIC_BACKEND_URL ?? "http://localhost:8080",
  entryPoint: process.env.NEXT_PUBLIC_ENTRYPOINT_ADDRESS ?? "",
  paymaster: process.env.NEXT_PUBLIC_PAYMASTER_ADDRESS ?? "",
  tusdt: process.env.NEXT_PUBLIC_TUSDT_ADDRESS ?? process.env.NEXT_PUBLIC_USDT_ADDRESS ?? "",
  tusdc: process.env.NEXT_PUBLIC_TUSDC_ADDRESS ?? "",
  delegateLogic: process.env.NEXT_PUBLIC_DELEGATE_ACCOUNT_ADDRESS ?? "",
  recipient: process.env.NEXT_PUBLIC_RECIPIENT_ADDRESS ?? "",
  tusdtName: process.env.NEXT_PUBLIC_TUSDT_NAME ?? process.env.NEXT_PUBLIC_USDT_NAME ?? "Test USDT",
  tusdcName: process.env.NEXT_PUBLIC_TUSDC_NAME ?? "Test USDC",
};
