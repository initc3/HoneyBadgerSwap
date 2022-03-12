pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public auctionCnt;

    struct Bid {
        uint bid;  // amount of DAI a bidder would like to pay
        uint lot;  // amount of collateral for sell
        address guy;  // high bidder address of the bidder with current highest price pay DAI receive collateral

        uint  tic;  // bid expiry time          [unix epoch time]
        uint  end;  // auction expiry time      [unix epoch time]
        address usr;
        address gal;
        uint tab;  // total dai wanted         [rad]
    }

    mapping(uint => Bid) public bids; //storage of all bids

    constructor() public {
        auctionCnt = 0;
    }


}
