pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public colAuctionCnt;



    mapping(uint=>uint) public bids_cnt;
    mapping (uint => mapping (uint => address)) public bids_P;  // address 
    mapping (uint => mapping (uint => uint)) public bids_Amt;  
    
     
    mapping (uint => uint) public status; // init success-1 input success-2 settle success-3
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;


    
    mapping (uint => string) public colres;
    mapping (address => string) public colresValue;
    mapping (string => uint) public colresCount;

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

    function inputAuction(uint colAuctionId, $uint X, uint Amt) public {
        address P = msg.sender;
        uint cur_num = ++bids_cnt[colAuctionId];
        bids_P[colAuctionId][cur_num] = P;
        bids_Amt[colAuctionId][cur_num] = Amt;

        mpc(uint colAuctionId, $uint X, address P, uint Amt){
            bids = readDB(f'bidsBoard_{colAuctionId}', list)
            bids.append((X,P,Amt))
            writeDB(f'bidsBoard_{colAuctionId}',bids,list)

            curStatus = 2
            set(status, uint curStatus, uint colAuctionId)
        }
    }

    function dutchAuctionSettle(uint colAuctionId, $uint AmtToSell, $uint StartPrice, uint LowestPrice) public{
        uint n = bids_cnt[colAuctionId];
        
        mpc(uint n, uint colAuctionId, $uint AmtToSell, $uint StartPrice, uint LowestPrice){
            bids = readDB(f'bidsBoard_{colAuctionId}', list)

            for i in range(n):
                for j in range(i) :
                    (Xi,Pi,Amti) = bids[i]
                    (Xj,Pj,Amtj) = bids[j]
                    mpcInput(sint Xi, sint Xj)
                    need_swap = (Xi.less_equal(Xj,bit_length=bit_length)).reveal()
                    mpcOutput(cint need_swap)

                    if need_swap == 1:
                        tmp = bids[i]
                        bids[i] = bids[j]
                        bids[j] = tmpX

            amtSold = 0

            mpcInput(sint StartPrice)
            curPrice = StartPrice
            mpcOutput(sint curPrice)
            
            for i in range(n):
                Xi,Pi,Amti = bids[i]
                mpcInput(sint Xi, sint Amti,sint curPrice, sint amtSold)
                curPrice = Xi
                amtSold += Amti
                mpcOutput(sint curPrice, sint amtSold)
                
                mpcInput(sint amtSold, sint AmtToSell)
                need_break = (amtSold.greater_equal(AmtToSell,bit_length = bit_length)).reveal()
                mpcOutput(cint need_break)

                if need_break == 1:
                    break
            
            mpcInput(sint amtSold, sint curPrice)
            amtSold = amtSold.reveal()
            curPrice = curPrice.reveal()
            mpcOutput(cint amtSold, cint curPrice)

            res = 'amtSold: {amtSold} curPrice:{curPrice}'
            set(colres, string memory res, uint colAuctionId)
        }
    }
}
