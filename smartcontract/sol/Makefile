BLOCKCHAIN_DIR = `pwd`
BUILD_DIR = $(BLOCKCHAIN_DIR)/build
SOLC_IMAGE = ethereum/solc:0.8.26-alpine
CONTRACT_FILE = petition_contract
CONTRACT_NAME = Petition
DEP_CONTRACT = $(CONTRACT_NAME)Deploy
CONTRACT_SOL = $(BLOCKCHAIN_DIR)/$(CONTRACT_FILE).sol
PKG = blockchain
TYPE = petit_contract

# run outside of devcontainer
solc:
	docker run -v $(BLOCKCHAIN_DIR):$(BLOCKCHAIN_DIR) $(SOLC_IMAGE) -o $(BUILD_DIR) --abi --bin --overwrite $(CONTRACT_SOL)

# run inside of devcontainer
abigen:
	abigen --abi=./build/$(CONTRACT_NAME).abi --pkg $(PKG) --type $(TYPE) --out ../../backend/eth/$(CONTRACT_NAME).go

abigen_deploy:
	abigen --bin=./build/$(CONTRACT_NAME).bin --abi=./build/$(CONTRACT_NAME).abi --pkg=$(PKG) --type=$(TYPE) --out=../../backend/eth/$(DEP_CONTRACT).go
