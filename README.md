# Pepetitions
With PetitionVault we make petitions safe from malicious activities by both the signers and the creators.
<div align="center">
    <img src="/frontend/public/Cover.png" width="500" height="300" alt="image">
</div>

## Team
- Johannes: Frontend
- Leon:     Backend & Smartcontract
- Louen:    Smartcontract
- Paul:     Frontend & Backend

## Purpose
- make petitions **tamper-proof** ...:
  - petitions can only be signed **once**
  - ... by the same **human**
- ... and more **anonymous** ...
- ... while making **source information** accessible in a **decentralized, hard-to-tamper-with** storage for **free**

## Components
- Frontend:
  - Authenticate for signing petitions using <a href="https://worldcoin.org/world-id">WorldID</a> <img src="https://github.com/user-attachments/assets/8fb7a853-20b8-4be9-ba57-eef47705322c" width="20" height="20" alt="WorldID">
  - Create and sign petitions
  - Upload source documents
- Backend:
  - Store addresses and other safe-to-store, non-privacy-invasive information in sqlite-database
  - Upload PDF documents to lighthouse storage powered by <a href="https://filecoin.io/">Filecoin</a> <img src="https://github.com/user-attachments/assets/0a33c99a-5e24-4cf7-8da2-2cf3cc7a613b" width="20" height="20" alt="Filecoin">
  - Missing: query smartcontract
- Smart Contract:
  - Ensure immutability of petitions
  - Allow authenticated users to sign
  - Store link(s) to sources
  - Increased efficiency thanks to <a href="https://arbitrum.io/stylus">Arbitrum Stylus</a> <img src="https://github.com/user-attachments/assets/c16383f3-732f-4329-a840-52c5a7f4f5c0" width="20" height="20" alt="Arbitrum Stylus">


## Next Steps

- **Additional Identification Methods:**
  - Explore and integrate other forms of identification that align with the concept of "Certified Anonymous Unique Human" users.

- **Enhanced User Information:**
  - Some additional anonymized informations that are relevant for a specific petition could be voluntarily submitted alongside the signature.
  - Some identification methods may allow for some reliable general informations about the user such as General location (on the scale of the country, region or even town).

- **Factory Smart Contract:**
  - Implement a Factory smart contract to facilitate easier interfacing with individual petition smart contracts.
  - This would streamline the creation and management of new petitions.

## Try it out
- copy and rename the frontend/.env.example and backend/.env.example to .env
- get API keys
  - WorldCoin
    - go to https://developer.worldcoin.org
    - create a team
    - go to App Profile
    - copy the app_staging ID and fill out both .env files
    - click on the 0 in the upper right corner -> Overview
    - go to API keys
    - create a new one
    - click on the three dots and reset the key
    - copy the key into the backend/.env file
  - LightHouse
    - login/register at https://files.lighthouse.storage/
    - click on API Key
    - generate a new one and put it into the backend/.env file
- run
  - on your first run, run `make init` in the base dir
  - after that you can always start with `make`
