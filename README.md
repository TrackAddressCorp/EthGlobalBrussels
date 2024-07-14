# Project: PetitionVault
With PetitionVault we make petitions safe from malicious activities by both the signers and the creators.
<img src="/Users/leon/Downloads/Cover.png" width="100" height="100" alt="image">

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
  - Authenticate for signing petitions using [WorldID]("https://worldcoin.org/world-id")
  - create and ...
  - sign petitions
  - upload source documents
- Backend:
  - store addresses and other safe-to-store, non-privacy-invasive information in sqlite-database
  - upload pdf documents to lighthouse storeage powered by [Filecoin]("https://filecoin.io/")
  - missing (:/): query smartcontract
- Smart Contract:
  - ensure immutability of petitions
  - allow authenticate users to sign
  - store link(s) to sources
  - increased efficiency thanks to [Arbitrum Stylus]("https://arbitrum.io/stylus")

## Technologies
<a href="https://filecoin.io/">
  <img src="https://github.com/user-attachments/assets/0a33c99a-5e24-4cf7-8da2-2cf3cc7a613b" width="100" height="100" alt="image">
</a>
<a href="https://arbitrum.io/stylus">
  <img src="https://github.com/user-attachments/assets/c16383f3-732f-4329-a840-52c5a7f4f5c0" width="100" height="100" alt="image">
</a>
<a href="https://worldcoin.org/world-id">
  <img src="https://github.com/user-attachments/assets/8fb7a853-20b8-4be9-ba57-eef47705322c" width="100" height="100" alt="image">
</a>

## 
