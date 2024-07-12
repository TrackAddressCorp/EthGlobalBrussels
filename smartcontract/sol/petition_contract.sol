// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Petition {
    struct Signer {
        string  world_id;
        uint256 timestamp;
        string  location;
        uint16  age;
        string  gender;
    }

    string private petition_title;
    string private petition_text;
    Signer[] private signers;
    uint256 private signers_count;

    constructor(string memory _petition_title, string memory _petition_text) {
        petition_title = _petition_title;
        petition_text = _petition_text;
        signers_count = 0;
    }

    function sign(string memory _world_id) public {
        for (uint i = 0; i < signers.length; i++) {
            if (keccak256(abi.encodePacked(signers[i].world_id)) == keccak256(abi.encodePacked(_world_id))) {
                revert("You have already signed this petition.");
            }
        }
        signers.push(
            Signer(
                _world_id,
                block.timestamp,
                "N/A",
                0,
                "N/A"));
        signers_count++;
    }

    function get_signers() public view returns (Signer[] memory) {
        return signers;
    }

    function get_signers_count() public view returns (uint256) {
        return signers_count;
    }

    function get_petition_title() public view returns (string memory) {
        return petition_title;
    }

    function get_petition_text() public view returns (string memory) {
        return petition_text;
    }
}
