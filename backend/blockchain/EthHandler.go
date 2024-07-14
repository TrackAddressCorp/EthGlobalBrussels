package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

/**
  !!! UNSAFE !!! ONLY FOR TESTING PURPOSES !!!
  wrapper around ethclient.Client
*/
type Handler struct {
    *ethclient.Client
    ChainID             *big.Int
    PrivateKey          *ecdsa.PrivateKey
    PrivateKeyString    string
    PublicKey           ecdsa.PublicKey
    PublicKeyString     string
    PublicAddress       common.Address
    PublicAddressString string
}

/**
    !!! UNSAFE !!! ONLY FOR TESTING PURPOSES !!!
    Creates a new handler for the sepolia eth test network
    * @param ctx context.Context
    * @param rpc string
    * @return *Client
    * @return error
*/
func NewHandler(ctx context.Context, rpc string, privateKeyString string) (*Handler, error) {
    ethClient, err := connectToEth(
        ctx,
        rpc,
    )
    if err != nil {
        fmt.Println("Error connecting to eth")
        return &Handler{}, err
    }
    defer ethClient.Close()

    wallet, err := loadWallet(privateKeyString)
    if err != nil {
        fmt.Println("Error loading wallet")
        return &Handler{}, err
    }

    chainID, err := ethClient.ChainID(ctx)
    if err != nil {
        return &Handler{}, err
    }

    return &Handler{
        Client: ethClient,
        ChainID: chainID,
        PrivateKey: wallet.PrivateKey,
        PrivateKeyString: wallet.PrivateKeyString,
        PublicKey: wallet.PublicKey,
        PublicKeyString: wallet.PublicKeyString,
        PublicAddress: wallet.PublicAddress,
        PublicAddressString: wallet.PublicAddressString,
        }, nil
}

/**
    !!! UNSAFE !!! ONLY FOR TESTING PURPOSES !!!
    Connects to the sepolia eth test network
    * @param ctx context.Context
    * @param rpc string
    * @return *ethclient.Client
    * @return error
*/
func connectToEth(ctx context.Context, rpc string) (*ethclient.Client, error) {
    client, err := ethclient.DialContext(
        ctx,
        rpc,
    )
    if err != nil {
        return nil, err
    }

    return client, nil
}

type EthWallet struct {
    PrivateKey          *ecdsa.PrivateKey
    PrivateKeyString    string
    PublicKey           ecdsa.PublicKey
    PublicKeyString     string
    PublicAddress       common.Address
    PublicAddressString string
}

/**
    !!! UNSAFE !!! ONLY FOR TESTING PURPOSES !!!
    Loads the metamask wallet from the environment variables
    * @return *EthWallet
    * @return error
*/
func loadWallet(privateKeyString string) (*EthWallet, error) {
    privateKey, err := crypto.HexToECDSA(privateKeyString)
    if err != nil {
        return &EthWallet{}, err
    }

    publicKey := privateKey.PublicKey
    publicKeyString := hexutil.Encode(crypto.FromECDSAPub(&publicKey))[4:]

    hash := sha3.NewLegacyKeccak256()
    hash.Write(crypto.FromECDSAPub(&publicKey)[1:])
    publicAddressString := hexutil.Encode(hash.Sum(nil)[12:])[2:]

    return &EthWallet{
        PrivateKey:             privateKey,
        PrivateKeyString:       privateKeyString,
        PublicKey:              publicKey,
        PublicKeyString:        publicKeyString,
        PublicAddress:          crypto.PubkeyToAddress(publicKey),
        PublicAddressString:    publicAddressString,
    }, nil
}
