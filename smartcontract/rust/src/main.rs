#![no_main]
// #![no_std]
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
    alloy_primitives::{Uint, U32, U256},
    storage::{StorageString, StorageType, StorageU256, StorageU32, StorageVec},
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
  petition_text:    StorageString,                     // Name of the petition                                          |  Assigned at Construction
  petition_title:   StorageString,                     // Title of the petition                                         |  Assigned at Construction
  signers_count:    StorageU32,                        // How much signers the petition currently has
  signers:          StorageVec<StorageSignersType>,    // List of all Signers w/ their corresponding informations
}


#[allow(dead_code)]
#[solidity_storage]
pub struct StorageSignersType {
  world_id:         StorageU256
}




#[external]
impl Petition {
    // Constructor
    pub fn init(&mut self, name: String, title: String) -> Result<(), Vec<u8>> {
        // If already initialized => Do nothing and return
        if !self.petition_text.is_empty() {
            return Ok(()) ;
        }
        
        unsafe { 
            // Initializes the petition name & title
            self.petition_text = StorageString::new(U256::from(0), 0);
            self.petition_title = StorageString::new(U256::from(0), 0);

            // Assign the provided name and title
            self.petition_text.set_str(name);
            self.petition_title.set_str(title);

            // Default initialization of signers informations
            self.signers_count = StorageU32::new(U256::from(2), 0);
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

    pub fn get_signers_count(&self) -> Result<U32, Vec<u8>> {
        Ok(self.signers_count.get())
    }


    // Sign the petition, saving the signer's informations
    pub fn sign(&mut self, world_id: U256) -> Result<(), Vec<u8>> {
        let signer = StorageSignersType {
            world_id: StorageU256::new(world_id, 0),
        };
        
        unsafe {
            self.signers.push(signer);
            self.signers_count.set(self.signers_count.get() + U32::from(1));
        }
        
        Ok(())
    }
}

