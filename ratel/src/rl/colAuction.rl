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

    function initAuction($uint val1) public{
        uint colAuctionId = ++colAuctionCnt;

        mpc(uint colAuctionId, $uint val1) {
            
            mpcInput(sint val1)

            valid = ((val1.greater_equal(1, bit_length=bit_length)) * (val1.less_equal(3, bit_length=bit_length))).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                bids = {
                    '0':0,
                }
                print('**** bids', bids)
                writeDB(f'bidsBoard_{colAuctionId}', bids, dict)

                curStatus = 1
                set(status, uint curStatus, uint colAuctionId)
        }
    }

    function inputAuction(uint colAuctionId, $uint X, uint Amt) public {
        address P = msg.sender;

        mpc(uint colAuctionId, $uint X, address P, uint Amt){
            bids = readDB(f'bidsBoard_{colAuctionId}', dict)

            mpcInput(sint X)

            valid = ((X.greater_equal(1, bit_length=bit_length)) * (X.less_equal(100, bit_length=bit_length))).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                bids['{P}'] = X
                print('**** bids', bids)

                curStatus = 2
                set(status, uint curStatus, uint colAuctionId)
        }
    }

}
