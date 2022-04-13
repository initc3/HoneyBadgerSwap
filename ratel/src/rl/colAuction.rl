pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint constant public Fp = 2**16;


    uint public colAuctionCnt;


    
    mapping (uint => uint) public status; // created-1 submitted-2 closed-3
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;


    
    mapping (uint => string) public colres;
    mapping (address => string) public colresValue;
    mapping (string => uint) public colresCount;

    constructor() public {}

    function createAuction($uint StartPrice, $uint FloorPrice, $uint totalAmt) public{
        uint colAuctionId = ++colAuctionCnt;

        mpc(uint colAuctionId, $uint StartPrice, $uint FloorPrice, $uint totalAmt) {
            
            mpcInput(sint StartPrice, sint FloorPrice)

            valid = ((StartPrice.greater_equal(FloorPrice, bit_length=bit_length))).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                bids = [(0,0,0)]
                writeDB(f'bidsBoard_{colAuctionId}', bids, list)
                
                auc = {
                    'StartPrice': StartPrice,
                    'FloorPrice': FloorPrice,
                    'totalAmt': totalAmt,
                }
                print('**** auc', auc)
                writeDB(f'aucBoard_{colAuctionId}', auc, dict)

                curStatus = 1
                set(status, uint curStatus, uint colAuctionId)
        }
    }

    pureMpc scheduleCheck(curPrice, FloorPrice):
        mpcInput(sfix curPrice, sint FloorPrice)
        valid = ((sint(curPrice)).greater_equal(FloorPrice,bit_length = bit_length)).reveal()
        mpcOutput(cint valid)

        if valid == 0:
            res = 'Auction failed!!!'
            set(colres, string memory res, uint colAuctionId)
            curStatus = 3
            set(status, uint curStatus, uint colAuctionId)
            return

        bids = readDB(f'bidsBoard_{colAuctionId}', list)
        auc = readDB(f'aucBoard_{colAuctionId}',dict)
            
        totalAmt = auc['totalAmt']

        n = len(bids)
            
        amtSold = 0

        for i in range(n-1):
            (Xi,Pi,Amti) = bids[i+1]

            mpcInput(sint Xi, sfix curPrice)

            valid = ((sint(curPrice)).less_equal(Xi,bit_length = bit_length)).reveal()

            mpcOutput(cint valid)

            if valid == 1:
                mpcInput(sint Amti, sint amtSold, sint totalAmt)
                amtSold += Amti
                aucDone = (amtSold.greater_equal(totalAmt,bit_length = bit_length).reveal())
                mpcOutput(sint amtSold,cint aucDone)

                if aucDone == 1:
                    break

        mpcInput(sint amtSold, sint totalAmt)
        aucDone = (amtSold.greater_equal(totalAmt,bit_length = bit_length).reveal())
        mpcOutput(cint aucDone)

        if aucDone == 1:
            res = 'Auction success!!!'
            set(colres, string memory res, uint colAuctionId)
            curStatus = 3
            set(status, uint curStatus, uint colAuctionId)
            return

        mpcInput(sfix curPrice)
        curPrice = curPrice*0.99
        mpcOutput(sfix curPrice)

    function submitBids(uint colAuctionId, $uint price, $uint Amt) public {
        address P = msg.sender;

        mpc(uint colAuctionId, $uint price, address P, $uint Amt){
            bids = readDB(f'bidsBoard_{colAuctionId}', list)
            auc = readDB(f'aucBoard_{colAuctionId}', dict)

            FloorPrice = auc['FloorPrice']

            mpcInput(sint price, sint FloorPrice)

            valid = (price.greater_equal(FloorPrice, bit_length=bit_length)).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                bids.append((price,P,Amt))
                print('**** bids', bids)
                writeDB(f'bidsBoard_{colAuctionId}',bids,list)

                curStatus = 2
                set(status, uint curStatus, uint colAuctionId)
        }
    }
}
