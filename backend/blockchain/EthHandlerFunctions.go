package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type GetNextTransactionParams struct {
    PrivateKey    *ecdsa.PrivateKey `json:"private_key"`
    PublicAddress common.Address    `json:"public_address"`
    ChainID       *big.Int          `json:"chain_id"`
}

func (h *Handler) GetNextTransaction(ctx context.Context, params *GetNextTransactionParams) (*bind.TransactOpts, error) {
    nonce, err := h.Client.PendingNonceAt(ctx, params.PublicAddress)
    if err != nil {
        return &bind.TransactOpts{}, err
    }


    // auth, err := bind.NewKeyedTransactorWithChainID(params.PrivateKey, params.ChainID)
    // if err != nil {
    // 	return &bind.TransactOpts{}, err
    // }
    auth, err := bind.NewKeyedTransactorWithChainID(
        params.PrivateKey,
        params.ChainID,
    )

    // ---- eipp-1559 standard ----
    auth.From = params.PublicAddress
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0) // amount of eth to send along with the transaction
    auth.GasLimit = uint64(3000000) // max amount of gas client is willing to pay for this tx
    auth.GasFeeCap = big.NewInt(300000000000) // Set a higher max fee per gas
    auth.GasTipCap = big.NewInt(2000000000) // Set a max priority fee per gas
    auth.Context = ctx
        
    // ---- older standard (for ganache) ----
    // gasPrice, err := h.Client.SuggestGasPrice(ctx)
    // if err != nil {
    //     return &bind.TransactOpts{}, err
    // }
    // auth.From = params.PublicAddress
    // auth.Nonce = big.NewInt(int64(nonce))
    // auth.Value = big.NewInt(0) // amount of eth to send along with the transaction
    // auth.GasLimit = uint64(3000000) // max amount of gas client is willing to pay for this tx
    // auth.Context = ctx
    // auth.GasPrice = gasPrice // price of gas in wei

    return auth, nil
}

type DeployParams struct {
    PrivateKey      *ecdsa.PrivateKey   `json:"private_key"`
    PetitionTitle   string              `json:"petition_title"`
    PetitionText    string              `json:"petition_text"`
}

func (h *Handler) DeployPetition(ctx context.Context, params DeployParams) (string, error) {
    chainID, err := h.Client.ChainID(ctx)
    if err != nil {
        return "", err
    }
    auth, err := h.GetNextTransaction(
		ctx, &GetNextTransactionParams{
			Hash: params.Hash,
			PrivateKey: params.PrivateKey,
			PublicAddress: params.PublicAddress,
			ChainID: chainID,
		},
	)
	if err != nil {
		fmt.Println("error 2:", err)
		return &UploadEthCertificateResult{}, err
	}

	address, transaction, contract, err := DeployCertificateContract(
		auth,
		h.client,
		params.Hash,
		params.IssuerName,
		params.RecipientName,
	)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice
    DeployPetitionContract(eth, params.PetitionTitle, params.PetitionText)
    return "", nil
}