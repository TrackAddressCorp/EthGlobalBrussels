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

    string private petition_name;
    string private petition_title;
    Signer[] private signers;
    uint256 private signers_count;

    constructor(string memory _petition_name, string memory _petition_title) {
        petition_name = _petition_name;
        petition_title = _petition_title;
        signers_count = 0;
    }

    function sign(string _world_id) public {
        for (uint i = 0; i < signers.length; i++) {
            string id = signers[i].world_id;
            if (_world_id == msg.sender) {
                revert("You have already signed this petition.");
            }
        }
        signers.push(msg.sender);
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
        return petition_name;
    }
}
