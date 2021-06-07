pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract VolumeMatching {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public bidCnt;

    uint constant public Decimals = 10**15;
    uint constant public Fp = 2**16;

    mapping (address => mapping (address => uint)) public publicBalance;

    uint public secretWithdrawCnt;
    mapping (address => mapping (uint => uint)) public secretWithdrawValue;
    mapping (uint => bool) public secretWithdrawFinish;
    mapping (uint => mapping (uint => uint)) public secretWithdrawCount;

    constructor() public {}

    function publicDeposit(address token, uint amt) payable public {
        address user = msg.sender;
        require(amt > 0);
        if (token == address(0x0)) {
            require(msg.value * Fp == amt * Decimals); // take care: unit conversion
        } else {
            IERC20(token).safeTransferFrom(user, address(this), amt / Fp * Decimals); // take care: unit conversion
        }
        publicBalance[token][user] += amt;
    }

    function secretDeposit(address token, uint amt) public {
        address user = msg.sender;
        require(amt > 0 && publicBalance[token][user] >= amt);
        publicBalance[token][user] -= amt;

        mpc(address user, address token, uint amt) {
            secretBalance = readDB(f'balance_{token}_{user}', int)

            mpcInput(sfix secretBalance, sfix amt)

            secretBalance += amt

            mpcOutput(sfix secretBalance)

            print('**** secretBalance', token, secretBalance)
            writeDB(f'balance_{token}_{user}', secretBalance, int)
        }
    }

    function publicWithdraw(address token, uint amt) public {
        address payable user = msg.sender;
        require(amt > 0 && publicBalance[token][user] >= amt);

        if (token == address(0x0)) {
           user.transfer(amt / Fp * Decimals);
        } else {
           IERC20(token).safeTransfer(user, amt / Fp * Decimals);
        }
        publicBalance[token][user] -= amt;
    }

    function secretWithdraw(address token, uint amt) public {
        address user = msg.sender;
        require(amt > 0);
        uint secretWithdrawSeq = ++secretWithdrawCnt;

        mpc(uint secretWithdrawSeq, address user, address token, uint amt) {
            balance = readDB(f'balance_{token}_{user}', int)

            mpcInput(sfix balance, sfix amt)

            enough = (balance >= amt).reveal()

            if_then(enough)
            balance -= amt
            end_if()

            mpcOutput(sfix balance, cint enough)

            print('****', balance, enough)
            if enough == 1:
                add(publicBalance, uint secretWithdrawSeq, uint amt, address token, address user)
                writeDB(f'balance_{token}_{user}', balance, int)
        }
    }

    function submitBid(address tokenA, address tokenB, $uint amtB) public {
        require(tokenA < tokenB);
        address user = msg.sender;
        uint bidSeq = ++bidCnt;

        mpc(uint bidSeq, address user, address tokenA, address tokenB, $uint amtB) {
            bidList = readDB(f'bidList_{tokenA}_{tokenB}', list)

            mpcInput(sfix amtB)

            validBid = (amtB != 0).reveal()

            mpcOutput(cint validBid)

            print('**** validBid', validBid)
            if validBid == 1:
                bidList.append((bidSeq, user, amtB))
                print('**** bidList', bidList)
                writeDB(f'bidList_{tokenA}_{tokenB}', bidList, list)
        }
    }

    function volumeMatch(address tokenA, address tokenB, uint price) public {
        require(tokenA < tokenB);
        require(price > 0);

        mpc(address tokenA, address tokenB, uint price) {
            bidList = readDB(f'bidList_{tokenA}_{tokenB}', list)

            buySum, sellSum = 0, 0
            for bidSeq, user, amtB in bidList:
                balanceA = readDB(f'balance_{tokenA}_{user}', int)
                balanceB = readDB(f'balance_{tokenB}_{user}', int)

                mpcInput(sfix price, sfix amtB, sfix balanceA, sfix balanceB, sfix buySum, sfix sellSum)
                print_ln('**** price %s', price.reveal())
                print_ln('**** amtB %s', amtB.reveal())
                print_ln('**** balanceA %s', balanceA.reveal())
                print_ln('**** balanceB %s', balanceB.reveal())
                print_ln('**** buySum %s', buySum.reveal())
                print_ln('**** sellSum %s', sellSum.reveal())

                amtA = amtB / price

                buyB = amtB > 0
                enoughA = amtA <= balanceA
                valid = buyB * enoughA
                print_ln('**** valid %s', valid.reveal())
                changeA = valid * amtA
                changeB = valid * amtB
                buySum += changeB
                balanceA -= changeA
                balanceB += changeB

                print_ln('**** buySum %s', buySum.reveal())
                print_ln('**** balanceA %s', balanceA.reveal())
                print_ln('**** balanceB %s', balanceB.reveal())

                sellB = 1 - buyB
                enoughB = -amtB <= balanceB
                valid = sellB * enoughB
                print_ln('**** valid %s', valid.reveal())
                changeA = valid * amtA
                changeB = valid * amtB
                sellSum -= changeB
                balanceA -= changeA
                balanceB += changeB

                print_ln('**** sellSum %s', sellSum.reveal())
                print_ln('**** balanceA %s', balanceA.reveal())
                print_ln('**** balanceB %s', balanceB.reveal())

                mpcOutput(sfix buySum, sfix sellSum, sfix balanceA, sfix balanceB)

                writeDB(f'balance_{tokenA}_{user}', balanceA, int)
                writeDB(f'balance_{tokenB}_{user}', balanceB, int)

            mpcInput(sfix buySum, sfix sellSum)

            f = sellSum > buySum
            tradeAmt = f * (buySum - sellSum) + sellSum

            print_ln('**** tradeAmt %s', tradeAmt.reveal())
            mpcOutput(sfix tradeAmt)

            sellSum, buySum = tradeAmt, tradeAmt
            for bidSeq, user, amtB in bidList:
                balanceA = readDB(f'balance_{tokenA}_{user}', int)
                balanceB = readDB(f'balance_{tokenB}_{user}', int)

                mpcInput(sfix price, sfix amtB, sfix balanceA, sfix balanceB, sfix buySum, sfix sellSum)
                print_ln('**** price %s', price.reveal())
                print_ln('**** amtB %s', amtB.reveal())
                print_ln('**** balanceA %s', balanceA.reveal())
                print_ln('**** balanceB %s', balanceB.reveal())
                print_ln('**** buySum %s', buySum.reveal())
                print_ln('**** sellSum %s', sellSum.reveal())

                amtA = amtB / price

                buyB = amtB > 0
                z1 = buySum <= 0
                z2 = buySum < amtB
                actualAmtB = buyB * (1 - z1) * (z2 * (buySum - amtB) + amtB)
                buySum -= actualAmtB
                actualAmtA = actualAmtB / price
                balanceA += buyB * (amtA - actualAmtA)
                balanceB -= buyB * (amtB - actualAmtB)

                print_ln('**** actualAmtA %s', actualAmtA.reveal())
                print_ln('**** actualAmtB %s', actualAmtB.reveal())
                print_ln('**** buySum %s', buySum.reveal())
                print_ln('**** balanceA %s', balanceA.reveal())
                print_ln('**** balanceB %s', balanceB.reveal())

                sellB = 1 - buyB
                z1 = sellSum <= 0
                z2 = sellSum < -amtB
                actualAmtB = sellB * (1 - z1) * (z2 * (-sellSum - amtB) + amtB)
                sellSum += actualAmtB
                actualAmtA = actualAmtB / price
                balanceA += sellB * (amtA - actualAmtA)
                balanceB -= sellB * (amtB - actualAmtB)

                print_ln('**** actualAmtA %s', actualAmtA.reveal())
                print_ln('**** actualAmtB %s', actualAmtB.reveal())
                print_ln('**** sellSum %s', sellSum.reveal())
                print_ln('**** balanceA %s', balanceA.reveal())
                print_ln('**** balanceB %s', balanceB.reveal())

                mpcOutput(sfix balanceA, sfix balanceB, sfix buySum, sfix sellSum)

                writeDB(f'balance_{tokenA}_{user}', balanceA, int)
                writeDB(f'balance_{tokenB}_{user}', balanceB, int)

            bidList = []
            writeDB(f'bidList_{tokenA}_{tokenB}', bidList, list)

        }
    }
}
