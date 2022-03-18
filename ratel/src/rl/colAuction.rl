pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public toyCnt;

    uint public colAuctionCnt;

    ///////////////constant///////////////
    uint public tau = 2 days;   // 2 days total auction length  [seconds]
    uint public beg = 1.05E18;  // 5% minimum bid increase
    uint public ttl = 3 hours;  // 3 hours bid duration         [seconds]

    ///////////////for bids///////////
    mapping(uint=>uint) bids_bid;  // amount of DAI a bidder would like to pay
    mapping(uint=>uint) bids_lot;  // amount of collateral for sell
    mapping(uint=>address) bids_guy;  // high bidder address of the bidder with current highest price pay DAI receive collateral
    mapping(uint=>uint)  bids_tic;  // bid expiry time          [unix epoch time]
    mapping(uint=>uint)  bids_end;  // auction expiry time      [unix epoch time]
    mapping(uint=>address) bids_usr; // usr: address to receive residual collateral after the auction
    mapping(uint=>address) bids_gal; // gal: address to receive raised DAI
    mapping(uint=>uint) bids_tab;  // total dai wanted         [rad]
    
    ////////////////for dai////////////
//    mapping(address=>uint) balance_DAI;// Q4 :???

    ///////////////for status///////////
    mapping (uint => uint) public status; // kick success-1, ready-2, completed-3
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;


    constructor() public {}

    // func kick(): initiate a new auction
    // tab: amount of DAI to raise;    lot: amount of collateral for sell
    // usr: address to receive residual collateral after the auction
    // gal: address to receive raised DAI
    // bid: amount of DAI a bidder would like to pay
    function kick(uint tab, $uint lot) public {
        address usr = msg.sender;
        address gal = msg.sender;
        address guy = address(0);
        uint colAuctionId = ++colAuctionCnt;
        uint end = uint(now) + tau;
        uint tic = 0;
        uint bid = 0;

        bids_tab[colAuctionId] = tab;
//        bids_lot[colAuctionId] = lot;//Q3 maybe not here //does lot need to be secret? 
        bids_usr[colAuctionId] = usr;
        bids_gal[colAuctionId] = gal;
        bids_end[colAuctionId] = end; ///Q2 need change?

        mpc(uint colAuctionId, uint tab, $uint lot, address usr, address gal, uint end, uint tic, address guy, uint bid) {
            colAuc = {
                'colAuctionId' : colAuctionId,
                'tab': tab,
                'lot': lot,
                'usr': usr,
                'gal': gal,
                'end': end,
                'tic': tic,
                'guy': guy,
                'bid': bid,
            }
            print('**** new colAuction', colAuc)
            writeDB(f'colAuctionBoard_{colAuctionId}', colAuc, dict)
            writeDB(f'balanceDAIBoard_{guy}',bid,int)

            curStatus = 1
            set(status, uint curStatus, uint colAuctionId)
        }
    }

    //func tend() raise the amount of DAI paid to get all collatera
    function tend(uint colAuctionId,$uint lot,$uint bid) public{
        require(status[colAuctionId] == 1 || status[colAuctionId] == 2);
        require(bids_tic[colAuctionId] == 0 || bids_tic[colAuctionId] > uint(now),"colAuction bids_tic finish");
        require(bids_end[colAuctionId] > uint(now), "colAuction bids_end finish");

        address guy = msg.sender;

        mpc(uint colAuctionId, $uint lot, $uint bid, address guy) {
            colAuc = readDB(f'colAuctionBoard_{colAuctionId}', dict)
            
            previous_lot = colAuc['lot']
            previous_bid = colAuc['bid']
            previous_guy = colAuc['guy']

            mpcInput(sint lot,sint previous_lot,sint previous_bid)
            # lot == previous lot
            valid_lot = ((lot.greater_equal(previous_lot, bit_length=bit_length)) * (lot.less_equal(previous_bid, bit_length=bit_length)))
#            valid_bid = ((bid.greater_equal))
            valid = valid_lot.reveal()
            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                colAuc['guy'] = guy
                colAuc['bid'] = bid
                colAuc['tic'] = uint(now)+ttl

                print('**** colAuc', colAuc)

                tmp = readDB(f'balanceDAIBoard_{previous_guy}',int)
                writeDB(f'balanceDAIBoard_{previous_guy}',tmp+previous_bid,int)
                tmp = readDB(f'balanceDAIBoard_{previous_guy}',int)
                writeDB(f'balanceDAIBoard_{previous_guy}',tmp+previous_bid,int)



                writeDB(f'colAuctionBoard_{colAuctionId}', colAuc, dict)

                curStatus = 2
                set(status, uint curStatus, uint culAuctionId)
        }
    } 

    function toyGame($uint value1) public {
        uint toyId = ++toyCnt;

        mpc(uint toyId, $uint value1) {
            mpcInput(sint value1)

            valid = ((value1.greater_equal(1, bit_length=bit_length)) * (value1.less_equal(10, bit_length=bit_length))).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                toy = {
                    'toyId': toyId,
                    'value1': value1,
                }
                print('**** toy', toy)
                writeDB(f'toyBoard_{toyId}', toy, dict)

                curStatus = 1
                set(status, uint curStatus, uint toyId)
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