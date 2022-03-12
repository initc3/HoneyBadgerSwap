pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public auctionCnt;


    uint public tau = 2 days;   // 2 days total auction length  [seconds]

    struct Bid {
//Q1 : is guy needed?

        uint bid;  // amount of DAI a bidder would like to pay
        uint lot;  // amount of collateral for sell
        address guy;  // high bidder address of the bidder with current highest price pay DAI receive collateral

        uint  tic;  // bid expiry time          [unix epoch time]
        uint  end;  // auction expiry time      [unix epoch time]
        address usr; // usr: address to receive residual collateral after the auction
        address gal; // gal: address to receive raised DAI
        uint tab;  // total dai wanted         [rad]
    }

    mapping(uint => Bid) public bids; //storage of all bids

    constructor() public {
        auctionCnt = 0;
    }

    //=======================math======================
    function col_add(uint x, uint y) public returns(uint z){
        require((z=x+y) >= x);
    }

    // func kick(): initiate a new auction
    // tab: amount of DAI to raise;    lot: amount of collateral for sell
    // usr: address to receive residual collateral after the auction
    // gal: address to receive raised DAI
    // bid: amount of DAI a bidder would like to pay
    function kick(uint tab, uint lot, address usr, address gal, uint bid) public { 
        address P = msg.sender;
        uint auctionId = ++auctionCnt;

        bids[auctionId].tab = tab;
        bids[auctionId].lot = lot;
        bids[auctionId].usr = usr;
        bids[auctionId].gal = gal;
        bids[auctionId].bid = bid;

        bids[auctionId].guy = msg.sender;
        bids[auctionId].tic = col_add(uint(now),tau);
    }


}
