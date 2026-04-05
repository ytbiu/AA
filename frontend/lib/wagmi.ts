import { QueryClient } from "@tanstack/react-query";
import { bscTestnet } from "viem/chains";
import { createConfig, http } from "wagmi";
import { injected } from "wagmi/connectors";
import { env } from "./env";

export const queryClient = new QueryClient();

const chain = {
  ...bscTestnet,
  id: env.chainId,
  rpcUrls: {
    ...bscTestnet.rpcUrls,
    default: { http: [env.rpcUrl] },
    public: { http: [env.rpcUrl] },
  },
};

export const wagmiConfig = createConfig({
  chains: [chain],
  transports: {
    [chain.id]: http(env.rpcUrl),
  },
  connectors: [
    injected({
      target: "metaMask",
      shimDisconnect: true,
    }),
  ],
  ssr: false,
});

export { chain };
