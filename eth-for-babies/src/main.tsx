import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { RainbowKitProvider } from '@rainbow-me/rainbowkit';
import { WagmiProvider, createConfig } from 'wagmi';
import { mainnet, sepolia } from 'wagmi/chains';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { AuthProvider } from './context/AuthContext';
import App from './App';
import './index.css';
import '@rainbow-me/rainbowkit/styles.css';
import { http } from 'wagmi';
import { getDefaultWallets } from '@rainbow-me/rainbowkit';

const { wallets } = getDefaultWallets();

const config = createConfig({
  chains: [mainnet, sepolia],
  transports: {
    [mainnet.id]: http(),
    [sepolia.id]: http(),
  },
  ssr: false,
});

const queryClient = new QueryClient();

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <RainbowKitProvider wallets={wallets} appInfo={{
          appName: 'FamilyChain Tasks',
          projectId: import.meta.env.VITE_WALLETCONNECT_PROJECT_ID || 'YOUR_WALLETCONNECT_PROJECT_ID',
        }}>
          <AuthProvider>
            <BrowserRouter>
              <App />
            </BrowserRouter>
          </AuthProvider>
        </RainbowKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  </StrictMode>
);