package blockchain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (h *Handler) getAuth(ctx context.Context) (*bind.TransactOpts, error) {
    nonce, err := h.Client.PendingNonceAt(ctx, h.PublicAddress)
    if err != nil {
        return &bind.TransactOpts{}, err
    }

    auth, err := bind.NewKeyedTransactorWithChainID(
        h.PrivateKey,
        h.ChainID,
    )
    if err != nil {
        return &bind.TransactOpts{}, err
    }

    // ---- eipp-1559 standard ----
    auth.From = h.PublicAddress
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
    PetitionTitle   string              `json:"petition_title"`
    PetitionText    string              `json:"petition_text"`
}

type DeployResult struct {
    Address     common.Address      `json:"address"`
    Transaction string              `json:"transaction"`
    Contract    *PetitionContract   `json:"contract"`
}

func (h *Handler) DeployPetition(ctx context.Context, params DeployParams) (*DeployResult, error) {
    auth, err := h.getAuth(ctx)
    if err != nil {
        return &DeployResult{}, err
    }

    address, transaction, contract, err := DeployPetitionContract(
        auth,
        h.Client,
        params.PetitionTitle,
        params.PetitionText,
    )
    if err != nil {
        return &DeployResult{}, nil
    }

    return &DeployResult{
        Address: address,
        Transaction: transaction.Hash().Hex(),
        Contract: contract,
    }, nil
}

type SignParams struct {
    PetitionAddress common.Address `json:"petition_address"`
    WorldID         string         `json:"world_id"`
}

func (h *Handler) SignPetition(ctx context.Context, params SignParams) error {
    auth, err := h.getAuth(ctx)
    if err != nil {
        return err
    }

    contract, err := NewPetitionContractTransactor(
        params.PetitionAddress,
        h.Client,
    )
    if err != nil {
        return err
    }

    _, err = contract.Sign(auth, params.WorldID)
    return err
}

func (h *Handler) GetSignerCount(ctx context.Context, petitionAddress common.Address) (uint64, error) {
    contract, err := NewPetitionContractCaller(
        petitionAddress,
        h.Client,
    )
    if err != nil {
        return 0, err
    }

    count, err := contract.GetSignersCount(
        &bind.CallOpts{
            Context: ctx,
        },
    )
    if err != nil {
        return 0, err
    }

    return count.Uint64(), nil
}

func (h *Handler) GetTitle(ctx context.Context, petitionAddress common.Address) (string, error) {
    contract, err := NewPetitionContractCaller(
        petitionAddress,
        h.Client,
    )
    if err != nil {
        return "", err
    }

    title, err := contract.GetPetitionTitle(
        &bind.CallOpts{
            Context: ctx,
        },
    )
    if err != nil {
        return "", err
    }

    return title, nil
}

func (h *Handler) GetText(ctx context.Context, petitionAddress common.Address) (string, error) {
    contract, err := NewPetitionContractCaller(
        petitionAddress,
        h.Client,
    )
    if err != nil {
        return "", err
    }

    text, err := contract.GetPetitionText(
        &bind.CallOpts{
            Context: ctx,
        },
    )
    if err != nil {
        return "", err
    }

    return text, nil
}
