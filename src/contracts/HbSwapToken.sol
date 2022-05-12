pragma solidity ^0.5.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20Detailed.sol";

contract HbSwapToken is ERC20, ERC20Detailed {
    uint8 public constant DECIMALS = 18;
    uint256 public constant INITIAL_SUPPLY = 1e28;

    constructor () public ERC20Detailed("HbSwapToken", "HBS", DECIMALS) {
        _mint(msg.sender, INITIAL_SUPPLY);
    }
}