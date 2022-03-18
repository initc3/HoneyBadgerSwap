pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public colAuctionCnt;

    ///////////////for bids///////////
    mapping(uint=>uint) bids_cnt;

    mapping (uint => mapping (uint => address)) bids_P;  // address 

    mapping (uint => mapping (uint => uint)) bids_Amt;  
    
    ///////////////for status///////////
    mapping (uint => uint) public status; // init success-1 input success-2 settle success-3
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;

    mapping (uint => uint) public ret_curPrice; 
    mapping (address => uint) public ret_curPriceValue;
    mapping (uint => uint) public ret_curPriceCount;

    mapping (uint => uint) public ret_amtSold; // init success-1 input success-2 settle success-3
    mapping (address => uint) public ret_amtSoldValue;
    mapping (uint => uint) public ret_amtSoldCount;

    constructor() public {}

    function initAuction() public{
        uint colAuctionId = ++colAuctionCnt;
        bids_cnt[colAuctionId] = 0;

        mpc(uint colAuctionId) {
            bids = []
            writeDB(f'bidsBoard_{colAuctionId}', bids, list)

            curStatus = 1
            set(status, uint curStatus, uint colAuctionId)
        }
    }

}
/*
pseudocode for batch dutch auction:
bids := []
on input Bid($X, Amt)  from P    // means I'll buy up to Amt if the prices reaches $X or below
      append ($X, Amt, P) to bids
on DutchAuctionSettle(AmtToSell, StartPrice, LowestPrice):
         amtSold := 0
         sort (bids in decreasing order by $X)
         curPrice := StartPrice
         for each ($X, Amt, P):
                  curPrice := $X
                 amtSold += amt
                 amt amtSold >= AmtToSell: break
        output amtSold, $X last price
*/