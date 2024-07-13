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
    uint104 private constant max_signers = 10;
    uint104 private signers_count;
    Signer[max_signers] private signers;
    // mapping(string => bool) private hasSigned;

    constructor(string memory _petition_title, string memory _petition_text) {
        petition_title = _petition_title;
        petition_text = _petition_text;
        signers_count = 0;
    }

    function sign(string memory _world_id) public {
        // require(!hasSigned[_world_id], "You have already signed this petition.");
        if (signers_count >= max_signers) {
            revert("The petition has reached the maximum number of signers.");
        }
        signers[signers_count] = Signer(
            _world_id,
            block.timestamp,
            "N/A",
            0,
            "N/A");
        signers_count++;
        // hasSigned[_world_id] = true;
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
