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
            secretBalance = int.from_bytes(readDB(f'balance_{token}_{user}'), 'big')
            mpcInput(secretBalance, amt)
            secretBalance = sfix._new(secretBalance)
            amt = sfix._new(amt)
            secretBalance += amt
            secretBalance = secretBalance.v
            mpcOutput(secretBalance)
            print('after secretDeposit', token, secretBalance)
            writeDB(f'balance_{token}_{user}', secretBalance.to_bytes((secretBalance.bit_length() + 7) // 8, 'big'))
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
            balance = int.from_bytes(readDB(f'balance_{token}_{user}'), 'big')

            mpcInput(balance, amt)
            balance = sfix._new(balance)
            amt = sfix._new(amt)

            enough = (balance >= amt).reveal()

            if_then(enough)
            balance -= amt
            end_if()

            balance = balance.v
            enough = sint(enough)
            mpcOutput(balance, enough)

            print(balance, enough)
            if enough == 1:
                add(publicBalance, uint secretWithdrawSeq, uint amt, address token, address user)
                writeDB(f'balance_{token}_{user}', balance.to_bytes((balance.bit_length() + 7) // 8, 'big'))
        }
    }

    function submitBid(address tokenA, address tokenB, $uint amtB) public {
        require(tokenA < tokenB);
        address user = msg.sender;
        uint bidSeq = ++bidCnt;

        mpc(uint bidSeq, address user, address tokenA, address tokenB, $uint amtB) {
            import ast
            bidList = readDB(f'bidList_{tokenA}_{tokenB}')
            try:
                bidList = bidList.decode(encoding='utf-8')
                bidList = list(ast.literal_eval(bidList))
            except:
                bidList = []

            mpcInput(amtB)
            amtB = sfix._new(amtB)

            validBid = (amtB != 0).reveal()

            validBid = sint(validBid)
            mpcOutput(validBid)

            print(validBid)
            if validBid == 1:
                bidList.append((bidSeq, user, amtB))
                print(bidList)
                bidList = str(bidList)
                bidList = bytes(bidList, encoding='utf-8')
                writeDB(f'bidList_{tokenA}_{tokenB}', bidList)
        }
    }

    function volumeMatch(address tokenA, address tokenB, uint price) public {
        require(tokenA < tokenB);

        mpc(address tokenA, address tokenB, uint price) {
            import ast
            bidList = readDB(f'bidList_{tokenA}_{tokenB}')
            try:
                bidList = bidList.decode(encoding='utf-8')
                bidList = list(ast.literal_eval(bidList))
            except:
                bidList = []

            buySum, sellSum = 0, 0
            for bidSeq, user, amtB in bidList:
                balanceA = int.from_bytes(readDB(f'balance_{tokenA}_{user}'), 'big')
                balanceB = int.from_bytes(readDB(f'balance_{tokenB}_{user}'), 'big')

                mpcInput(price, amtB, balanceA, balanceB, buySum, sellSum)
                price = sfix._new(price)
                amtB = sfix._new(amtB)
                balanceA = sfix._new(balanceA)
                balanceB = sfix._new(balanceB)
                buySum = sfix._new(buySum)
                sellSum = sfix._new(sellSum)

                print_ln('price %s', price.reveal())
                print_ln('amtB %s', amtB.reveal())
                print_ln('balanceA %s', balanceA.reveal())
                print_ln('balanceB %s', balanceB.reveal())

                amtA = amtB / price

                buyB = amtB > 0
                enoughA = amtA <= balanceA
                valid = buyB * enoughA
                print_ln('valid %s', valid.reveal())
                changeA = valid * amtA
                changeB = valid * amtB
                buySum += changeB
                balanceA -= changeA
                balanceB += changeB

                print_ln('buySum %s', buySum.reveal())
                print_ln('sellSum %s', sellSum.reveal())
                print_ln('balanceA %s', balanceA.reveal())
                print_ln('balanceB %s', balanceB.reveal())

                sellB = 1 - buyB
                enoughB = -amtB <= balanceB
                valid = sellB * enoughB
                print_ln('valid %s', valid.reveal())
                changeA = valid * amtA
                changeB = valid * amtB
                sellSum -= changeB
                balanceA -= changeA
                balanceB += changeB

                print_ln('buySum %s', buySum.reveal())
                print_ln('sellSum %s', sellSum.reveal())
                print_ln('balanceA %s', balanceA.reveal())
                print_ln('balanceB %s', balanceB.reveal())

                buySum = buySum.v
                sellSum = sellSum.v
                balanceA = balanceA.v
                balanceB = balanceB.v
                mpcOutput(buySum, sellSum, balanceA, balanceB)

                writeDB(f'balance_{tokenA}_{user}', balanceA.to_bytes((balanceA.bit_length() + 7) // 8, 'big'))
                writeDB(f'balance_{tokenB}_{user}', balanceB.to_bytes((balanceB.bit_length() + 7) // 8, 'big'))

            mpcInput(buySum, sellSum)
            buySum = sfix._new(buySum)
            sellSum = sfix._new(sellSum)

            f = sellSum > buySum
            tradeAmt = f * (buySum - sellSum) + sellSum

            print_ln('tradeAmt %s', tradeAmt.reveal())

            tradeAmt = tradeAmt.v
            mpcOutput(tradeAmt)

            sellSum, buySum = tradeAmt, tradeAmt
            for bidSeq, user, amtB in bidList:
                balanceA = int.from_bytes(readDB(f'balance_{tokenA}_{user}'), 'big')
                balanceB = int.from_bytes(readDB(f'balance_{tokenB}_{user}'), 'big')

                mpcInput(price, amtB, balanceA, balanceB, buySum, sellSum)
                price = sfix._new(price)
                amtB = sfix._new(amtB)
                balanceA = sfix._new(balanceA)
                balanceB = sfix._new(balanceB)
                buySum = sfix._new(buySum)
                sellSum = sfix._new(sellSum)

                amtA = amtB / price

                print_ln('price %s', price.reveal())
                print_ln('amtA %s', amtA.reveal())
                print_ln('amtB %s', amtB.reveal())
                print_ln('balanceA %s', balanceA.reveal())
                print_ln('balanceB %s', balanceB.reveal())
                print_ln('buySum %s', buySum.reveal())
                print_ln('sellSum %s', sellSum.reveal())

                buyB = amtB > 0
                z1 = buySum <= 0
                z2 = buySum < amtB
                actualAmtB = buyB * (1 - z1) * (z2 * (buySum - amtB) + amtB)
                print_ln('actualAmtB %s', actualAmtB.reveal())
                buySum -= actualAmtB
                print_ln('buySum %s', buySum.reveal())
                actualAmtA = actualAmtB / price
                print_ln('actualAmtA %s', actualAmtA.reveal())
                balanceA += buyB * (amtA - actualAmtA)
                balanceB -= buyB * (amtB - actualAmtB)

                print_ln('balanceA %s', balanceA.reveal())
                print_ln('balanceB %s', balanceB.reveal())

                sellB = 1 - buyB
                z1 = sellSum <= 0
                z2 = sellSum < -amtB
                actualAmtB = sellB * (1 - z1) * (z2 * (buySum - amtB) + amtB)
                print_ln('actualAmtB %s', actualAmtB.reveal())
                sellSum += actualAmtB
                print_ln('sellSum %s', sellSum.reveal())
                actualAmtA = actualAmtB / price
                print_ln('actualAmtA %s', actualAmtA.reveal())
                balanceA += sellB * (amtA - actualAmtA)
                balanceB -= sellB * (amtB - actualAmtB)

                print_ln('balanceA %s', balanceA.reveal())
                print_ln('balanceB %s', balanceB.reveal())

                balanceA = balanceA.v
                balanceB = balanceB.v
                buySum = buySum.v
                sellSum = sellSum.v
                mpcOutput(balanceA, balanceB, buySum, sellSum)

                writeDB(f'balance_{tokenA}_{user}', balanceA.to_bytes((balanceA.bit_length() + 7) // 8, 'big'))
                writeDB(f'balance_{tokenB}_{user}', balanceB.to_bytes((balanceB.bit_length() + 7) // 8, 'big'))

            writeDB(f'bidList_{tokenA}_{tokenB}', bytes(0))

        }
    }
}
