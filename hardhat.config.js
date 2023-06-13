require("@nomicfoundation/hardhat-toolbox");
require('dotenv').config();

const accounts = [
  process.env.PRIVATE_KEY || "0x0000000000000000000000000000000000000000000000000000000000000000",
]

module.exports = {
  namedAccounts: {
    deployer: {
      default: 0,
    },
  },
  solidity: {
    compilers: [
      {
        version: "0.5.17",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
      {
        version: "0.8.17",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      }
    ],
  },
  networks: {
    hardhat: {
      forking: {
        url: process.env.MAINNET_RPC_URL,
        blockNumber: Number(process.env.FORKING_BLOCK_NUMBER),
        enabled: false,
      },
    },
    mainnet: {
      url: process.env.MAINNET_RPC_URL,
      accounts: accounts,
      saveDeployments: true,
      chainId: 1,
    },
    iotex: {
        url: 'https://babel-api.mainnet.iotex.io',
        accounts: accounts,
        chainId: 4689,
    },
    iotex_test: {
        url: 'https://babel-api.testnet.iotex.io',
        accounts: accounts,
        chainId: 4690,
    },
    bsc: {
      url: 'https://bsc-dataseed.binance.org',
      accounts: accounts,
      chainId: 56,
    },
    polygon: {
        url: 'https://polygon-rpc.com/',
        accounts: accounts,
        chainId: 137,
    },
  }
};
