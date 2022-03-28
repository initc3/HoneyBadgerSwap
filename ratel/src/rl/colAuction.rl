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
                bids = [(0,0,0)]
                print('**** bids', bids)
                writeDB(f'bidsBoard_{colAuctionId}', bids, list)

                curStatus = 1
                set(status, uint curStatus, uint colAuctionId)
        }
    }

    function inputAuction(uint colAuctionId, $uint X, $uint Amt) public {
        address P = msg.sender;

        mpc(uint colAuctionId, $uint X, address P, $uint Amt){
            bids = readDB(f'bidsBoard_{colAuctionId}', list)

            mpcInput(sint X, sint Amt)

            valid = ((X.greater_equal(1, bit_length=bit_length)) * (Amt.greater_equal(1, bit_length=bit_length))).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                bids.append((X,P,Amt))
                print('**** bids', bids)
                writeDB(f'bidsBoard_{colAuctionId}',bids,list)

                curStatus = 2
                set(status, uint curStatus, uint colAuctionId)
        }
    }

    function dutchAuctionSettle(uint colAuctionId, $uint AmtToSell, $uint StartPrice, uint LowestPrice) public{
        address P = msg.sender;
        
        mpc(uint colAuctionId, $uint AmtToSell, $uint StartPrice, uint LowestPrice){
            bids = readDB(f'bidsBoard_{colAuctionId}', list)

            n = len(list)
            
            Cnt = 0

            for i in range(n-1):
                (Xi,Pi,Amti) = bids[i]
                (Xj,Pj,Amtj) = bids[i+1]
                mpcInput(sint Cnt, sint Xi, sint Xj)

                Cnt += (Xi.greater_equal(Xj,bit_length = bit_length))

                mpcOutput(sint Cnt)

            mpcInput(sint Cnt)

            Cnt = Cnt.reveal()

            mpcOutput(cint Cnt)

            if Cnt == n-1 :
                res = 'success'
            else:
                res = 'failed'

            set(colres, string memory res, uint colAuctionId)
        }
    }


}
