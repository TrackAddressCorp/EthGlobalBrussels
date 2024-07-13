// Only run this as a WASM if the export-abi feature is not set.
#![cfg_attr(not(feature = "export-abi"), no_main)]
extern crate alloc;

extern crate stylus_sdk;

/// Use an efficient WASM allocator.
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

use alloc::{
    string::String,
    vec::Vec,
};

/// Import items from the SDK. The prelude contains common traits and macros.
use stylus_sdk::{
    prelude::*,
    alloy_primitives::U256,
    storage::{StorageString, StorageU256, StorageUint, StorageVec, StorageGuard, StorageType, SimpleStorageType, Erase},
};

/// The solidity_storage macro allows this struct to be used in persistent
/// storage. It accepts fields that implement the StorageType trait. Built-in
/// storage types for Solidity ABI primitives are found under
/// stylus_sdk::storage.
#[solidity_storage]
/// The entrypoint macro defines where Stylus execution begins. External methods
/// are exposed by annotating an impl for this struct with #[external] as seen
/// below.
#[entrypoint]
pub struct Petition {
    petition_text:    StorageString,                     // Name of the petition
    petition_title:   StorageString,                     // Title of the petition
    signers_count:    StorageUint<32, 1>,                // How many signers the petition currently has
    signers:          StorageVec<StorageSignersType>,    // List of all Signers with their corresponding information
}

#[allow(dead_code)]
#[solidity_storage]
pub struct StorageSignersType {
    world_id: StorageU256
}

// Implement Erase for StorageSignersType
impl Erase for StorageSignersType {
    fn erase(&mut self) {
        self.world_id.erase();
    }
}

// Implement SimpleStorageType for StorageSignersType
impl<'a> SimpleStorageType<'a> for StorageSignersType {
    type Wraps<'b> = StorageSignersType where 'a: 'b;

    fn set_by_wrapped(&mut self, value: Self::Wraps<'a>) {
        *self = value;
    }
}

#[external]
impl Petition {
    // Constructor
    pub fn init(&mut self, name: String, title: String) -> Result<(), Vec<u8>> {
        // If already initialized => Do nothing and return
        if !self.petition_text.is_empty() {
            return Ok(());
        }

        unsafe {
            // Initializes the petition name & title
            self.petition_text = StorageString::new(U256::from(0), 0);
            self.petition_title = StorageString::new(U256::from(0), 0);

            // Assign the provided name and title
            self.petition_text.set_str(name);
            self.petition_title.set_str(title);

            // Default initialization of signers information
            self.signers_count = StorageUint::new(U256::from(2), 0);
            self.signers = StorageVec::new(U256::from(3), 0);
        }

        Ok(())
    }

    // Getters
    pub fn get_petition_text(&self) -> Result<String, Vec<u8>> {
        Ok(self.petition_text.get_string())
    }

    pub fn get_petition_title(&self) -> Result<String, Vec<u8>> {
        Ok(self.petition_title.get_string())
    } 

    pub fn get_signers_count(&self) -> Result<u32, Vec<u8>> {        
        Ok(self.signers_count.to::<u32>())
    }

    // Sign the petition, saving the signer's information
    pub fn sign(&mut self, world_id: U256) -> Result<(), Vec<u8>> {
        let signer = StorageSignersType {
            world_id: StorageU256::new(world_id, 0),
        };

        self.signers.push(signer);
        let new_count = self.signers_count.get() + stylus_sdk::alloy_primitives::Uint::<32, 1>::from(1u32);
        self.signers_count.set(new_count);

        Ok(())
    }
}
