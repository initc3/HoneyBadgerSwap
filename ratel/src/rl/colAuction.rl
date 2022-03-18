pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public auctionCnt;
    uint public toyCnt;


    mapping (uint => uint) public status; // active-1, ready-2, completed-3
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;


    constructor() public {
        auctionCnt = 0;
        toyCnt = 0;
    }

    function toy($uint value) public { 
        address P = msg.sender;
        uint toyId = ++toyCnt;

        mpc(uint toyId, $uint value) {
            mpcInput(sint value)

            valid =  value.greater_equal(1,bit_length = bit_length)

            mpcOutput(cint valid)

            if valid == 1:
                toy = {
                    'id': toyId,
                }
                print('**** toy', toy)
                writeDB(f'toyBoard_{toyId}', toy, dict)
                curStatus = 1
                set(status, uint curStatus, uint toyId)
        }
    }

}
