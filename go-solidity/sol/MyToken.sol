pragma solidity ^0.4.11;

import "./ERC223.sol";

contract MyToken is ERC223Token {
    string public name = "MyToken";
    string public symbol = "MTK";
    uint public decimals = 6;
    uint public INITIAL_SUPPLY = 10000000000;
    
    function MyToken() {
        totalSupply = INITIAL_SUPPLY;
        balances[msg.sender] = INITIAL_SUPPLY;
    }
}
