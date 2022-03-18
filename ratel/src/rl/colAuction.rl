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

    constructor() public {}

    function initAuction() public{
        uint colAuctionId = ++colAuctionCnt;
        bids_cnt[colAuctionId] = 0;

        mpc(uint colAuctionId) {
            bidsX = []
            writeDB(f'bidsXBoard_{colAuctionId}', bidsX, list)
            bidsP = []
            writeDB(f'bidsPBoard_{colAuctionId}', bidsP, list)
            bidsAmt = []
            writeDB(f'bidsAmtBoard_{colAuctionId}', bidsAmt, list)

            curStatus = 1
            set(status, uint curStatus, uint colAuctionId)
        }
    }

    function inputAuction(uint colAuctionId, $uint X, uint Amt) public {
        address P = msg.sender;
        uint cur_num = ++bids_cnt[colAuctionId];
        bids_P[colAuctionId][cur_num] = P;
        bids_Amt[colAuctionId][cur_num] = Amt;

        mpc(uint colAuctionId, $uint X, address P, uint Amt){
            bidsX = readDB(f'bidsXBoard_{colAuctionId}', list)
            bidsX.append(X)
            writeDB(f'bidsXBoard_{colAuctionId}',bidsX,list)

            bidsP = readDB(f'bidsPBoard_{colAuctionId}',list)
            bidsP.append(P)
            writeDB(f'bidsPBoard_{colAuctionId}',bidsP,list)

            bidsAmt = readDB(f'bidsAmtBoard_{colAuctionId}',list)
            bidsAmt.append(Amt)
            writeDB(f'bidsAmtBoard_{colAuctionId}',bidsAmt,list)

            curStatus = 2
            set(status, uint curStatus, uint colAuctionId)
        }
    }

    function dutchAuctionSettle(uint colAuctionId, uint AmtToSell, uint StartPrice, uint LowestPrice) public{
        uint n = bids_cnt[colAuctionId];

        mpc(uint colAuctionId, uint n, uint AmtToSell, uint StartPrice, uint LowestPrice){
            bidsX = readDB(f'bidsXBoard_{colAuctionId}', list)
            bidsP = readDB(f'bidsPBoard_{colAuctionId}',list)
            bidsAmt = readDB(f'bidsAmtBoard_{colAuctionId}',list)

            for i in range(n) :
                for j in range(i) :
                    mpcInput(sint bidsX[i], sint bidsX[j])
                    need_swap = (bidsX[i].less_equal(bidsX[j],bit_length=bit_length)).reveal()
                    mpcOutput(cint need_swap)

                    if need_swap == 1:
                        tmpX = bidsX[i]
                        bidsX[i] = bidsX[j]
                        bidsX[j] = tmpX

                        tmpP = bidsP[i]
                        bidsP[i] = bidsP[j]
                        bidsP[j] = tmpP
                        
                        tmpAmt = bidsAmt[i]
                        bidsAmt[i] = bidsAmt[j]
                        bidsAmt[j] = tmpAmt

            amtSold = 0
            curPrice = StartPrice
            for i in range(n):
                mpcInput(sint bidsX[i], sint bids_Amt[i],sint curPrice, sint amtSold)
                curPrice = bidsX[i].reveal()
                amtSold += bids_Amt[i].reveal()
                mpcOutput(sint curPrice, sint amtSold)
                
                mpcInput(sint amtSold)
                need_break = (amtSold.greater_equal(AmtToSell,bit_length = bit_length)).reveal()
                mpcOutput(cint need_break)

                if need_break == 1:
                    break
            
            mpcInput(sint amtSold, sint curPrice)
            amtSold = amtSold.reveal()
            curPrice = curPrice.reveal()
            mpcOutput(cint amtSold, cint curPrice)

            curStatus = 3
            set(status, uint curStatus, uint colAuctionId)
            set(amtSold, uint amtSold, uint colAuctionId)
            set(curPrice, uint curPrice, uint colAuctionId)
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