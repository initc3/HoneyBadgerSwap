pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint constant public Fp = 2**16;


    uint public colAuctionCnt;

    mapping (uint => uint) public biddersCnt;

    mapping (uint => uint) public curPriceList;
    mapping (uint => uint) public floorPriceList;
    mapping (uint => uint) public startPriceList;
    
    mapping (uint => uint) public checkTime;

    mapping (uint => uint) public checkCnt;
    mapping (uint => uint) public checkNum; // closed-1 created-2 submitted --- bidders_num+2 
    mapping (address => uint) public checkNumValue;
    mapping (uint => uint) public checkNumCount;

    mapping (uint => uint) public status; // closed-1 created-2 submitted --- bidders_num+2 
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;


    
    mapping (uint => string) public colres;
    mapping (address => string) public colresValue;
    mapping (string => uint) public colresCount;

    constructor() public {}

    function createAuction(uint StartPrice, uint FloorPrice, $uint totalAmt) public{
        uint colAuctionId = ++colAuctionCnt;
        curPriceList[colAuctionId] = StartPrice * Fp;
        floorPriceList[colAuctionId] = FloorPrice;
        startPriceList[colAuctionId] = StartPrice;

        checkTime[colAuctionId] = block.timestamp;
        checkCnt[colAuctionId] = 0;

        biddersCnt[colAuctionId] = 0;

        mpc(uint colAuctionId, uint StartPrice, uint FloorPrice, $uint totalAmt) {
                
            bids = [(0,0,0)]
            writeDB(f'bidsBoard_{colAuctionId}', bids, list)
            
            auc = {
                'totalAmt': totalAmt,
                'StartPrice': StartPrice,
                'FloorPrice': FloorPrice,
            }
            print('**** auc', auc)
            writeDB(f'aucBoard_{colAuctionId}', auc, dict)

            curStatus = 2
            set(status, uint curStatus, uint colAuctionId)
        }
    }

    function scheduleCheck(uint colAuctionId) public {
        
        uint lastTime = checkTime[colAuctionId];

        require(block.timestamp < lastTime + 10 seconds);
        
        checkTime[colAuctionId] = block.timestamp;

        uint curPrice = curPriceList[colAuctionId]*100/99;
        curPriceList[colAuctionId] = curPrice;

        uint FloorPrice = floorPriceList[colAuctionId];

        uint curCheckNum = checkCnt[colAuctionId]+1;
        checkCnt[colAuctionId] = curCheckNum;

        mpc(uint colAuctionId, uint curCheckNum, uint curPrice, uint FloorPrice){

            bids = readDB(f'bidsBoard_{colAuctionId}', list)
            auc = readDB(f'aucBoard_{colAuctionId}',dict)

            totalAmt = auc['totalAmt']

            if curPrice/fp < FloorPrice:
                res = 'Auction failed!!!'
                set(colres, string memory res, uint colAuctionId)
                
                set(checkNum, uint curCheckNum, uint colAuctionId)

                curStatus = 1
                set(status, uint curStatus, uint colAuctionId)
                return
            

            n = len(bids)

            amtSold = 0

            for i in range(n-1):
                (Xi,Pi,Amti) = bids[i+1]

                mpcInput(sint Xi, sfix curPrice)

                valid = ((sint(curPrice/65536)).less_equal(Xi,bit_length = bit_length)).reveal()

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
                set(checkNum, uint curCheckNum, uint colAuctionId)
                curStatus = 1
                set(status, uint curStatus, uint colAuctionId)
                return

        }
    }

    function submitBids(uint colAuctionId, $uint price, $uint Amt) public {
        address P = msg.sender;

        uint bidders_id = biddersCnt[colAuctionId]+1;
        biddersCnt[colAuctionId] = bidders_id;

        uint FloorPrice = floorPriceList[colAuctionId];

        mpc(uint colAuctionId, uint bidders_id, uint FloorPrice, $uint price, address P, $uint Amt){
            bids = readDB(f'bidsBoard_{colAuctionId}', list)
            auc = readDB(f'aucBoard_{colAuctionId}', dict)

            mpcInput(sint price, sint FloorPrice)

            valid = (price.greater_equal(FloorPrice, bit_length=bit_length)).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                bids.append((price,P,Amt))
                print('**** bids', bids)
                writeDB(f'bidsBoard_{colAuctionId}',bids,list)

                curStatus = bidders_id+2
                set(status, uint curStatus, uint colAuctionId)
        }
    }
}
