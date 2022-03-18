pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public toyCnt;

    uint public colAuctionCnt;

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

    mapping (uint => uint) public status; // active-1, ready-2, completed-3
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;


    constructor() public {}

    // func kick(): initiate a new auction
    // tab: amount of DAI to raise;    lot: amount of collateral for sell
    // usr: address to receive residual collateral after the auction
    // gal: address to receive raised DAI
    // bid: amount of DAI a bidder would like to pay
    function kick(uint tab, uint lot, $uint bid) public {
        address P = msg.sender;
        address usr = msg.sender;
        address gal = msg.sender;
        uint colAuctionId = ++colAuctionCnt;

        bids[colAuctionId].tab = tab;
        bids[colAuctionId].lot = lot;
        bids[colAuctionId].usr = usr;
        bids[colAuctionId].gal = gal;
//        bids[colAuctionId].bid = bid;

        bids[colAuctionId].guy = msg.sender;
        bids[colAuctionId].end = uint(now) + tau; ///Q2 need change?

        mpc(uint colAuctionId, uint tab, uint lot, address usr, address gal, $uint bid) {
            mpcInput(sint bid)

            valid = (bid.greater_equal(1, bit_length=bit_length)).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                colAuc = {
                    'colAuctionId' : colAuctionId,
                    'tab': tab,
                    'lot': lot,
                    'usr': usr,
                    'gal': gal,
                    'bid': bid,
                }
                print('**** new colAuction', colAuc)
                writeDB(f'colAuctionBoard_{colAuctionId}', colAuc, dict)

                curStatus = 1
                set(status, uint curStatus, uint colAuctionId)
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