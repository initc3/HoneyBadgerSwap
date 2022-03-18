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


    function toyGame($uint value1) public {
        address player1 = msg.sender;
        uint toyId = ++toyCnt;

        mpc(uint toyId, address player1, $uint value1) {
            mpcInput(sint value1)

            valid = ((value1.greater_equal(1, bit_length=bit_length)) * (value1.less_equal(3, bit_length=bit_length))).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                toy = {
                    'player1': player1,
                    'value1': value1,
                }
                print('**** toy', toy)
                writeDB(f'toyBoard_{toyId}', toy, dict)

                curStatus = 1
                set(status, uint curStatus, uint toyId)
        }
    }

}
