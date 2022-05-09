pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public colAuctionCnt;

    mapping (uint => uint) public biddersCnt;

    mapping (uint => uint) public curPriceList;
    mapping (uint => uint) public floorPriceList;
    mapping (uint => uint) public startPriceList;
    
    mapping (uint => uint) public checkTime;

    mapping (uint => uint) public status; // closed-1 created-2 submitted --- bidders_num+2 
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;
    

    constructor() public {}

    function createAuction(uint StartPrice, uint FloorPrice, uint totalAmt) public{
        uint colAuctionId = ++colAuctionCnt;
        curPriceList[colAuctionId] = StartPrice;
        floorPriceList[colAuctionId] = FloorPrice;
        startPriceList[colAuctionId] = StartPrice;

        checkTime[colAuctionId] = block.number;

        status[colAuctionId] = 2;

        mpc(uint colAuctionId, uint StartPrice, uint FloorPrice, uint totalAmt) {
            auc = {
                'totalAmt': totalAmt,
                'StartPrice': StartPrice,
                'FloorPrice': FloorPrice,
            }
            writeDB(f'aucBoard_{colAuctionId}', auc, dict)
        }
    }

    function scheduleCheck(uint colAuctionId) public {
        
        uint lastTime = checkTime[colAuctionId];

        uint curTime = block.number;

        require(lastTime + 10 < curTime);

        checkTime[colAuctionId] = block.number;

        uint curPrice = curPriceList[colAuctionId]*99/100;
        curPriceList[colAuctionId] = curPrice;

        uint FloorPrice = floorPriceList[colAuctionId];

        mpc(uint colAuctionId, uint curPrice, uint FloorPrice){

            bids = readDB(f'bidsBoard_{colAuctionId}', list)
            auc = readDB(f'aucBoard_{colAuctionId}',dict)

            totalAmt = auc['totalAmt']

            print("**** curPrice:",curPrice)
            print("**** totalAmt:",totalAmt)

            if curPrice < FloorPrice:
                print(colAuctionId,'Auction failed!!!!!!!!!')

                curStatus = 1
                set(status, uint curStatus, uint colAuctionId)
            
            else:
            
                n = len(bids)
                print("**** n:",n)

                amtSold = 0

                for i in range(n):

                    (Xi,Pi,Amti) = bids[i]

                    mpcInput(sint Xi, sint curPrice, sint Amti, sint amtSold, sint totalAmt)

                    valid = (curPrice.less_equal(Xi,bit_length = bit_length))

                    amtSold += Amti*valid

                    mpcOutput(sint amtSold)


                mpcInput(sint amtSold, sint totalAmt)

                aucDone = (amtSold.greater_equal(totalAmt,bit_length = bit_length).reveal())
                
                mpcOutput(cint aucDone)

                print("**** aucDone", aucDone)

                if aucDone == 1:
                    print(colAuctionId,'Auction success!!!!!!!!!')
                    curStatus = 1
                    set(status, uint curStatus, uint colAuctionId)

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

            print_ln("**** price %s", price.reveal())
            print_ln("**** FloorPrice %s", FloorPrice.reveal())

            valid = (price.greater_equal(FloorPrice, bit_length=bit_length)).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                bids.append((price,P,Amt))
                print('**** bids', bids)

            writeDB(f'bidsBoard_{colAuctionId}',bids,list)
            
            curStatus = bidders_id+2
            set(status, uint curStatus, uint colAuctionId)

            print('submit end',colAuctionId,bidders_id)

        }
    }
}
