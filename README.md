# Project: PetitionVault
With PetitionVault we make petitions safe from malicious activities by both the signers and the creators.
<img src="frontend/public/Banner.png" width="100" height="100" alt="image">

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

## 
