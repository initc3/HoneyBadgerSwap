pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public colAuctionCnt;

     
    mapping (uint => uint) public status; // init success-1 input success-2 settle success-3
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;


    
    mapping (uint => string) public colres;
    mapping (address => string) public colresValue;
    mapping (string => uint) public colresCount;

    constructor() public {}
    function initAuction() public{
        uint colAuctionId = ++colAuctionCnt;

        mpc(uint colAuctionId) {
            bids = [(0,0,0)]
            writeDB(f'bidsBoard_{colAuctionId}', bids, list)

            curStatus = 1
            set(status, uint curStatus, uint colAuctionId)
        }
    }

    function inputAuction(uint colAuctionId, $uint X, uint Amt) public {
        address P = msg.sender;

        mpc(uint colAuctionId, $uint X, address P, uint Amt){
            bids = readDB(f'bidsBoard_{colAuctionId}', list)

            curStatus = 2
            set(status, uint curStatus, uint colAuctionId)
        }
    }

}
