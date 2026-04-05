"use client";

import { useEffect, useMemo, useRef, useState } from "react";
import {
  createPublicClient,
  createWalletClient,
  custom,
  encodeFunctionData,
  http,
  parseUnits,
  toHex,
  type Address,
  type Hex,
} from "viem";
import { bscTestnet } from "viem/chains";
import { delegateAccountAbi, entryPointAbi, erc20Abi } from "@/lib/abis";
import { env } from "@/lib/env";

type EIP1193Provider = {
  request: (args: { method: string; params?: any[] | object }) => Promise<any>;
  isMetaMask?: boolean;
  isSafe?: boolean;
  providers?: EIP1193Provider[];
};

type EIP6963ProviderDetail = {
  info: {
    uuid: string;
    name: string;
    icon: string;
    rdns: string;
  };
  provider: EIP1193Provider;
};

type SponsorResp = {
  paymasterAndData: Hex;
};

type TokenOption = {
  key: "tusdt" | "tusdc";
  name: string;
  address: Address;
};

const chain = {
  ...bscTestnet,
  id: env.chainId,
  rpcUrls: {
    ...bscTestnet.rpcUrls,
    default: { http: [env.rpcUrl] },
    public: { http: [env.rpcUrl] },
  },
};

const publicClient = createPublicClient({
  chain,
  transport: http(env.rpcUrl),
});

function short(addr?: string) {
  if (!addr) return "-";
  return `${addr.slice(0, 6)}...${addr.slice(-4)}`;
}

async function postJSON<T>(url: string, body: unknown): Promise<T> {
  const resp = await fetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  const data = await resp.json();
  if (!resp.ok || data.error) {
    let message = `HTTP ${resp.status}`;
    if (typeof data?.error === "string") {
      message = data.error;
    } else if (data?.error && typeof data.error === "object") {
      const errObj = data.error as { code?: unknown; message?: unknown; data?: unknown };
      const msg = String(errObj.message ?? "");
      const code = errObj.code !== undefined ? `code=${String(errObj.code)} ` : "";
      const detail = errObj.data !== undefined ? ` data=${JSON.stringify(errObj.data)}` : "";
      message = `${code}${msg}${detail}`.trim();
      if (!message) {
        message = JSON.stringify(data.error);
      }
    }
    throw new Error(message);
  }
  return data as T;
}

const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

function getMetaMaskProvider(): EIP1193Provider | null {
  const eth = (window as any).ethereum as EIP1193Provider | undefined;
  if (!eth) return null;

  if (Array.isArray(eth.providers) && eth.providers.length > 0) {
    const mm = eth.providers.find(
      (p) => p && p.isMetaMask && !p.isSafe,
    );
    return mm ?? null;
  }

  return eth.isMetaMask && !eth.isSafe ? eth : null;
}

export function AADemo() {
  const providerRef = useRef<EIP1193Provider | null>(null);
  const [eip6963Providers, setEip6963Providers] = useState<EIP6963ProviderDetail[]>([]);
  const [walletClient, setWalletClient] = useState<ReturnType<typeof createWalletClient> | null>(null);
  const [address, setAddress] = useState<Address | undefined>(undefined);
  const [chainId, setChainId] = useState<number | undefined>(undefined);

  const isConnected = Boolean(address);

  const [to, setTo] = useState(env.recipient);
  const [amount, setAmount] = useState("1");
  const [transferTokenKey, setTransferTokenKey] = useState<"tusdt" | "tusdc">("tusdt");
  const [gasTokenKey, setGasTokenKey] = useState<"tusdt" | "tusdc">("tusdt");
  const [faucetTokenKey, setFaucetTokenKey] = useState<"tusdt" | "tusdc">("tusdt");
  const [faucetTo, setFaucetTo] = useState("");
  const [faucetAmount, setFaucetAmount] = useState("1000");
  const [loading, setLoading] = useState(false);
  const [logs, setLogs] = useState<string[]>([]);

  const tokenOptions = useMemo<TokenOption[]>(
    () => [
      { key: "tusdt", name: env.tusdtName, address: env.tusdt as Address },
      { key: "tusdc", name: env.tusdcName, address: env.tusdc as Address },
    ],
    [],
  );

  const transferToken = useMemo(
    () => tokenOptions.find((t) => t.key === transferTokenKey),
    [tokenOptions, transferTokenKey],
  );
  const gasToken = useMemo(
    () => tokenOptions.find((t) => t.key === gasTokenKey),
    [tokenOptions, gasTokenKey],
  );
  const faucetToken = useMemo(
    () => tokenOptions.find((t) => t.key === faucetTokenKey),
    [tokenOptions, faucetTokenKey],
  );

  useEffect(() => {
    if (address && !faucetTo) {
      setFaucetTo(address);
    }
  }, [address, faucetTo]);

  useEffect(() => {
    const providers = new Map<string, EIP6963ProviderDetail>();

    const onAnnounce = (event: Event) => {
      const detail = (event as CustomEvent<EIP6963ProviderDetail>).detail;
      if (!detail?.info?.uuid || !detail?.provider) return;
      providers.set(detail.info.uuid, detail);
      setEip6963Providers(Array.from(providers.values()));
    };

    window.addEventListener("eip6963:announceProvider", onAnnounce as EventListener);
    window.dispatchEvent(new Event("eip6963:requestProvider"));

    return () => {
      window.removeEventListener("eip6963:announceProvider", onAnnounce as EventListener);
    };
  }, []);

  const configCheck = useMemo(() => {
    const required = [
      ["NEXT_PUBLIC_BACKEND_URL", env.backendUrl],
      ["NEXT_PUBLIC_ENTRYPOINT_ADDRESS", env.entryPoint],
      ["NEXT_PUBLIC_PAYMASTER_ADDRESS", env.paymaster],
      ["NEXT_PUBLIC_TUSDT_ADDRESS", env.tusdt],
      ["NEXT_PUBLIC_TUSDC_ADDRESS", env.tusdc],
      ["NEXT_PUBLIC_DELEGATE_ACCOUNT_ADDRESS", env.delegateLogic],
    ] as const;
    const missing = required.filter((item) => !item[1]).map((item) => item[0]);
    return { ok: missing.length === 0, missing };
  }, []);

  const walletHints = useMemo(() => {
    if (eip6963Providers.length === 0) return "";
    const names = eip6963Providers.map((p) => `${p.info.name} (${p.info.rdns})`);
    const hasMetaMask = eip6963Providers.some((p) => p.info.rdns === "io.metamask");
    return hasMetaMask
      ? `已检测钱包: ${names.join(" | ")}`
      : `未检测到 MetaMask，当前钱包: ${names.join(" | ")}`;
  }, [eip6963Providers]);

  const appendLog = (message: string) => {
    setLogs((prev) => [`${new Date().toLocaleTimeString()} ${message}`, ...prev]);
  };

  const ensureChain = async () => {
    const provider = providerRef.current;
    if (!provider) throw new Error("MetaMask provider 不可用");

    const targetHex = toHex(env.chainId);
    const currentHex = (await provider.request({ method: "eth_chainId" })) as string;
    const current = Number.parseInt(currentHex, 16);

    if (current === env.chainId) {
      setChainId(current);
      return;
    }

    try {
      await provider.request({
        method: "wallet_switchEthereumChain",
        params: [{ chainId: targetHex }],
      });
    } catch {
      await provider.request({
        method: "wallet_addEthereumChain",
        params: [
          {
            chainId: targetHex,
            chainName: "BSC Testnet",
            rpcUrls: [env.rpcUrl],
            nativeCurrency: { name: "tBNB", symbol: "tBNB", decimals: 18 },
            blockExplorerUrls: ["https://testnet.bscscan.com"],
          },
        ],
      });
    }

    setChainId(env.chainId);
    appendLog(`已切换到链 ${env.chainId}`);
  };

  const connectWallet = async () => {
    const metaMaskFrom6963 =
      eip6963Providers.find((p) => p.info.rdns === "io.metamask") ??
      eip6963Providers.find((p) => p.info.name.toLowerCase().includes("metamask"));

    const provider = metaMaskFrom6963?.provider ?? getMetaMaskProvider();
    if (!provider) {
      throw new Error("未检测到 MetaMask（请关闭 Safe 扩展干扰或仅保留 MetaMask）");
    }

    const accounts = (await provider.request({ method: "eth_requestAccounts" })) as string[];
    if (!accounts || accounts.length === 0) {
      throw new Error("MetaMask 未返回账号");
    }

    providerRef.current = provider;
    const wc = createWalletClient({
      chain,
      transport: custom(provider as any),
    });

    const addr = accounts[0] as Address;
    setWalletClient(wc);
    setAddress(addr);

    const cidHex = (await provider.request({ method: "eth_chainId" })) as string;
    setChainId(Number.parseInt(cidHex, 16));

    appendLog(`已连接 MetaMask: ${addr}`);
  };

  const disconnectWallet = () => {
    setWalletClient(null);
    setAddress(undefined);
    setChainId(undefined);
    providerRef.current = null;
    appendLog("已在页面状态中断开连接");
  };

  const upgradeBy7702 = async () => {
    if (!isConnected || !address) throw new Error("请先连接钱包");
    if (!configCheck.ok) throw new Error(`缺少配置: ${configCheck.missing.join(", ")}`);

    await ensureChain();
    appendLog("MetaMask 暂不支持 7702 capability，改为后端私钥发起链上升级...");
    const result = await postJSON<{
      owner: string;
      delegateTo: string;
      txHash: string;
      txNonce: number;
      authNonce: number;
    }>(`${env.backendUrl}/api/v1/7702/upgrade`, { owner: address });
    appendLog(
      `7702 已由后端发起: owner=${result.owner}, delegate=${result.delegateTo}, tx=${result.txHash}, txNonce=${result.txNonce}, authNonce=${result.authNonce}`,
    );
  };

  const runTransfer = async () => {
    if (!isConnected || !address || !walletClient) {
      throw new Error("请先连接钱包");
    }
    if (!to) throw new Error("请输入收款地址");
    if (!transferToken || !gasToken) throw new Error("token 配置缺失");
    if (!configCheck.ok) throw new Error(`缺少配置: ${configCheck.missing.join(", ")}`);

    await ensureChain();

    const sender = address as Address;
    const recipient = to as Address;

    const senderCode = await publicClient.getBytecode({ address: sender });
    if (!senderCode || senderCode === "0x") {
      throw new Error(
        `当前地址 ${sender} 还不是 7702 智能钱包（无合约代码）。请先执行第 2 步升级，或切回已升级的钱包地址。`,
      );
    }

    const { nonce, transferAmount, accountCall, gasPrice, callGasLimit, verificationGasLimit, preVerificationGas } =
      await buildOpDraft(sender, recipient, transferToken.address, gasToken.address);

    const quotePayload = {
      sender,
      gasToken: gasToken.address,
      nonce: nonce.toString(),
      callData: accountCall,
      callGasLimit: callGasLimit.toString(),
      verificationGasLimit: verificationGasLimit.toString(),
      preVerificationGas: preVerificationGas.toString(),
      maxFeePerGas: gasPrice.toString(),
      maxPriorityFeePerGas: gasPrice.toString(),
    };

    const quote = await postJSON<{ tokenAmount: string }>(
      `${env.backendUrl}/api/v1/paymaster/quote`,
      quotePayload,
    );

    const maxTokenCharge = BigInt(quote.tokenAmount);
    const gasTokenBalance = (await publicClient.readContract({
      address: gasToken.address,
      abi: erc20Abi,
      functionName: "balanceOf",
      args: [sender],
    })) as bigint;
    const transferTokenBalance = (await publicClient.readContract({
      address: transferToken.address,
      abi: erc20Abi,
      functionName: "balanceOf",
      args: [sender],
    })) as bigint;

    if (transferToken.address === gasToken.address) {
      const minRequired = transferAmount + maxTokenCharge;
      if (transferTokenBalance < minRequired) {
        throw new Error(
          `${transferToken.name} 余额不足。当前余额=${transferTokenBalance}，至少需要=${minRequired}（转账金额 + gas 代付预留）`,
        );
      }
    } else {
      if (transferTokenBalance < transferAmount) {
        throw new Error(
          `${transferToken.name} 余额不足。当前余额=${transferTokenBalance}，至少需要=${transferAmount}`,
        );
      }
      if (gasTokenBalance < maxTokenCharge) {
        throw new Error(
          `${gasToken.name} 余额不足。当前余额=${gasTokenBalance}，至少需要=${maxTokenCharge}（gas 代付预留）`,
        );
      }
    }

    const sponsor = await postJSON<SponsorResp>(`${env.backendUrl}/api/v1/paymaster/sponsor`, quotePayload);

    const userOp = {
      sender,
      nonce: toHex(nonce),
      initCode: "0x",
      callData: accountCall,
      callGasLimit: toHex(callGasLimit),
      verificationGasLimit: toHex(verificationGasLimit),
      preVerificationGas: toHex(preVerificationGas),
      maxFeePerGas: toHex(gasPrice),
      maxPriorityFeePerGas: toHex(gasPrice),
      paymasterAndData: sponsor.paymasterAndData,
      signature: "0x",
    };

    const computedUserOpHash = (await publicClient.readContract({
      address: env.entryPoint as Address,
      abi: entryPointAbi,
      functionName: "getUserOpHash",
      args: [userOp as any],
    })) as Hex;

    const userOpSignature = await walletClient.signMessage({
      account: sender,
      message: { raw: computedUserOpHash },
    });

    const sendResp = await postJSON<any>(`${env.backendUrl}/api/v1/userop/send`, {
      entryPoint: env.entryPoint,
      userOperation: {
        ...userOp,
        signature: userOpSignature,
      },
    });

    if (sendResp.error) {
      throw new Error(sendResp.error.message ?? JSON.stringify(sendResp.error));
    }

    const txOrUserOpHash = String(sendResp.result ?? "");
    appendLog(`已提交，hash: ${txOrUserOpHash || "(empty)"}`);

    if (!txOrUserOpHash || !txOrUserOpHash.startsWith("0x")) return;

    appendLog("开始轮询回执...");
    for (let i = 0; i < 20; i++) {
      const receiptResp = await fetch(
        `${env.backendUrl}/api/v1/userop/receipt?hash=${encodeURIComponent(txOrUserOpHash)}`,
      );
      const receiptData = await receiptResp.json();

      if (receiptData?.result) {
        const txHash =
          receiptData.result.receipt?.transactionHash ??
          receiptData.result.transactionHash ??
          txOrUserOpHash;
        appendLog(`已上链，txHash: ${txHash}`);
        return;
      }
      await sleep(3000);
    }
    appendLog("回执轮询超时，请稍后在区块浏览器确认。");
  };

  const buildOpDraft = async (
    sender: Address,
    recipient: Address,
    transferTokenAddress: Address,
    gasTokenAddress: Address,
  ) => {
    const nonce = (await publicClient.readContract({
      address: sender,
      abi: delegateAccountAbi,
      functionName: "nonce",
    })) as bigint;

    const transferAmount = parseUnits(amount || "0", 6);
    const approveCall = encodeFunctionData({
      abi: erc20Abi,
      functionName: "approve",
      args: [env.paymaster as Address, (1n << 256n) - 1n],
    });
    const transferCall = encodeFunctionData({
      abi: erc20Abi,
      functionName: "transfer",
      args: [recipient, transferAmount],
    });
    const accountCall = encodeFunctionData({
      abi: delegateAccountAbi,
      functionName: "executeBatch",
      args: [
        [gasTokenAddress, transferTokenAddress],
        [0n, 0n],
        [approveCall, transferCall],
      ],
    });

    const rpcGasPrice = await publicClient.getGasPrice();
    const minBundlerGasPrice = 1_000_000_000n; // 1 gwei
    const baseGasPrice = rpcGasPrice > minBundlerGasPrice ? rpcGasPrice : minBundlerGasPrice;
    const gasPrice = (baseGasPrice * 12n) / 10n; // 上浮 20%，提升打包概率

    return {
      nonce,
      transferAmount,
      accountCall,
      gasPrice,
      callGasLimit: 420000n,
      verificationGasLimit: 110000n,
      preVerificationGas: 90000n,
    };
  };


  const mintFromFaucet = async () => {
    if (!faucetToken) throw new Error("faucet token 配置缺失");
    if (!faucetTo) throw new Error("请输入领取地址");
    if (!faucetAmount) throw new Error("请输入领取数量");
    const mintAmount = parseUnits(faucetAmount, 6);
    if (mintAmount <= 0n) throw new Error("领取数量必须大于 0");

    const resp = await postJSON<{ txHash: string }>(`${env.backendUrl}/api/v1/faucet/mint`, {
      token: faucetToken.address,
      to: faucetTo,
      amount: mintAmount.toString(),
    });
    appendLog(`水龙头已提交: ${faucetToken.name} -> ${faucetTo}, tx=${resp.txHash}`);
  };

  const handleAction = async (fn: () => Promise<void>) => {
    setLoading(true);
    try {
      await fn();
    } catch (err) {
      appendLog(`失败: ${(err as Error).message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <main className="page">
      <section className="card">
        <h1>EIP-7702 + EIP-4337 BSC Testnet Demo</h1>
        <p className="hint">目标: 用 tUSDT / tUSDC 支付 gas（单 Paymaster 多 ERC20），单笔 UserOp 执行 approve + transfer。</p>

        <div className="meta">
          <span>钱包: {short(address)}</span>
          <span>链: {chainId ?? "-"}</span>
        </div>

        <div className="actions">
          {!isConnected ? (
            <button disabled={loading} onClick={() => handleAction(connectWallet)}>
              1. 连接 MetaMask
            </button>
          ) : (
            <button disabled={loading} onClick={() => disconnectWallet()}>
              断开钱包
            </button>
          )}

          <button disabled={loading || !isConnected} onClick={() => handleAction(upgradeBy7702)}>
            2. 升级为智能钱包 (EIP-7702)
          </button>

          <button disabled={loading || !isConnected} onClick={() => handleAction(runTransfer)}>
            3. 4337 转账（approve + ERC20 代付 gas）
          </button>
          <button disabled={loading} onClick={() => handleAction(mintFromFaucet)}>
            4. 水龙头领取
          </button>
        </div>

        <div className="form">
          <label>
            转账 Token
            <select value={transferTokenKey} onChange={(e) => setTransferTokenKey(e.target.value as "tusdt" | "tusdc")}>
              {tokenOptions.map((token) => (
                <option key={token.key} value={token.key}>
                  {token.name}
                </option>
              ))}
            </select>
          </label>
          <label>
            Gas Token
            <select value={gasTokenKey} onChange={(e) => setGasTokenKey(e.target.value as "tusdt" | "tusdc")}>
              {tokenOptions.map((token) => (
                <option key={token.key} value={token.key}>
                  {token.name}
                </option>
              ))}
            </select>
          </label>
          <label>
            收款地址
            <input value={to} onChange={(e) => setTo(e.target.value)} placeholder="0x..." />
          </label>
          <label>
            转账数量
            <input value={amount} onChange={(e) => setAmount(e.target.value)} placeholder="1" />
          </label>
          <label>
            水龙头 Token
            <select value={faucetTokenKey} onChange={(e) => setFaucetTokenKey(e.target.value as "tusdt" | "tusdc")}>
              {tokenOptions.map((token) => (
                <option key={token.key} value={token.key}>
                  {token.name}
                </option>
              ))}
            </select>
          </label>
          <label>
            领取地址
            <input value={faucetTo} onChange={(e) => setFaucetTo(e.target.value)} placeholder="0x..." />
          </label>
          <label>
            领取数量
            <input value={faucetAmount} onChange={(e) => setFaucetAmount(e.target.value)} placeholder="1000" />
          </label>
        </div>

        {!configCheck.ok ? <p className="error">缺少前端配置: {configCheck.missing.join(", ")}</p> : null}
        {walletHints ? <p className="hint">{walletHints}</p> : null}

        <div className="logBox">
          {logs.length === 0 ? <p className="hint">操作日志会显示在这里。</p> : null}
          {logs.map((log) => (
            <p key={log}>{log}</p>
          ))}
        </div>
      </section>
    </main>
  );
}
