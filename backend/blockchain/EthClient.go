package blockchain

import (
	"context"
	"crypto/ecdsa"
	"os"

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
    PrivateKey      *ecdsa.PrivateKey
    PublicKey       ecdsa.PublicKey
    PublicKeyString string
    PublicAddress   string
}

/**
    !!! UNSAFE !!! ONLY FOR TESTING PURPOSES !!!
    Creates a new handler for the sepolia eth test network
    * @param ctx context.Context
    * @param rpc string
    * @return *Client
    * @return error
*/
func NewHandler(ctx context.Context, rpc string) (*Handler, error) {
    ethClient, err := connectToEth(
        ctx,
        rpc,
    )
    if err != nil {
        return nil, err
    }
    defer ethClient.Close()

    wallet, err := loadWallet()
    if err != nil {
        return nil, err
    }

    return &Handler{
        Client: ethClient,
        PrivateKey: wallet.PrivateKey,
        PublicKey: wallet.PublicKey,
        PublicKeyString: wallet.PublicKeyString,
        PublicAddress: wallet.PublicAddress,
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
    PrivateKey      *ecdsa.PrivateKey
    PublicKey       ecdsa.PublicKey
    PublicKeyString string
    PublicAddress   string
}

/**
    !!! UNSAFE !!! ONLY FOR TESTING PURPOSES !!!
    Loads the metamask wallet from the environment variables
    * @return *EthWallet
    * @return error
*/
func loadWallet() (*EthWallet, error) {
    privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
    if err != nil {
        return &EthWallet{}, err
    }

    publicKey := privateKey.PublicKey
    publicKeyString := hexutil.Encode(crypto.FromECDSAPub(&publicKey))[4:]

    hash := sha3.NewLegacyKeccak256()
    hash.Write(crypto.FromECDSAPub(&publicKey)[1:])
    publicAddress := hexutil.Encode(hash.Sum(nil)[12:])[2:]

    return &EthWallet{
        PrivateKey:         privateKey,
        PublicKey:          publicKey,
        PublicKeyString:    publicKeyString,
        PublicAddress:      publicAddress,
    }, nil
}
