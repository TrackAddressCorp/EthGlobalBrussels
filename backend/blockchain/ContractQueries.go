package blockchain

// import (
// 	"context"
// 	"fmt"
// 	"math/big"
// 	"strings"

// 	ethereum "github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// )

// var PetitionMetaData = &bind.MetaData{
// 	ABI: "[{\"inputs\":[],\"name\":\"getPetitionText\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPetitionTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignersCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSourceLinks\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"links\",\"type\":\"string[]\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"world_id\",\"type\":\"uint256\"}],\"name\":\"sign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
// 	Bin: "",
// }
// var PetitionAddr = "0xc9f1af3f3119f81e9fef59399e0ff96a62b86f7a"

// var FactoryMetaData = &bind.MetaData{
// 	ABI: "",
// 	Bin: "",
// }
// var FactoryAddr = ""

// /**
// 	* @dev DeployPetition is a function that deploys a new instance of the petition contract.
// 	* @param petitionTitle The title of the petition.
// 	* @param petitionText The text of the petition.
// 	* @return address The address of the deployed contract.
// 	* @return transaction The transaction hash of the deployment.
// 	* @return contract The contract instance.
// */
// func (h *Handler) DeployPetition() {}

// /**
// 	* @dev SignPetition is a function that signs a petition.
// 	* @param petitionAddress The address of the petition contract.
// 	* @param worldID.
// */
// func (h *Handler) SignPetition(worldID *big.Int) error {
// 	contractABI, err := abi.JSON(strings.NewReader(PetitionMetaData.ABI))
// 	if err != nil {
// 		fmt.Println("Error1: ", err)
// 		return err
// 	}

// 	contractAddress := common.HexToAddress(PetitionAddr)

// 	data, err := contractABI.Pack("sign", worldID)
// 	if err != nil {
// 		fmt.Println("Error2: ", err)
// 		return err
// 	}

// 	msg := ethereum.CallMsg {
// 		From: h.PublicAddress,
// 		To:   &contractAddress,
// 		Data: data,
// 	}

// 	result, err := h.Client.CallContract(context.Background(), msg, nil)
// 	if err != nil {
// 		fmt.Println("Error3: ", err)
// 		return err
// 	}

// 	unpackedResult, err := contractABI.Unpack("sign", result)
// 	if err != nil {
// 		fmt.Println("Error4: ", err)
// 		return err
// 	}

// 	fmt.Println("Print result: ", unpackedResult)
// 	if len(unpackedResult) == 0 {
// 		fmt.Println("Error5: ", err)
// 		return fmt.Errorf("failed to sign petition")
// 	}

// 	fmt.Println("Signer: ", unpackedResult[0].(bool))
// 	_, ok := unpackedResult[0].(bool)
// 	if !ok {
// 		return fmt.Errorf("failed to sign petition")
// 	}
// 	return nil
// }

// /**
// 	* @dev GetSignerCount is a function that returns the number of signers of a petition.
// 	* @param petitionAddress The address of the petition contract.
// 	* @return uint64 The number of signers.
// */
// func (h *Handler) GetSignerCount() {}