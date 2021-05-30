pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract Test {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint constant public Decimals = 10**15;
    uint constant public Fp = 2**16;

    mapping (address => mapping (address => uint)) public publicBalance;

    mapping (address => mapping (address => string)) public estimatedPrice;
    mapping (address => string) public valueEstimatedPrice;
    mapping (string => uint) public countEstimatedPrice;

    uint public secretWithdrawCnt;
    mapping (address => mapping (uint => uint)) public secretWithdrawValue;
    mapping (uint => bool) public secretWithdrawFinish;
    mapping (uint => mapping (uint => uint)) public secretWithdrawCount;

    uint public tradeCnt;

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

    function initPool(address tokenA, address tokenB, uint amtA, uint amtB) public {
        require(tokenA < tokenB && amtA > 0 && amtB > 0);
        address user = msg.sender;

        mpc(address user, address tokenA, address tokenB, uint amtA, uint amtB) {
            balanceA = int.from_bytes(readDB(f'balance_{tokenA}_{user}'), 'big')
            balanceB = int.from_bytes(readDB(f'balance_{tokenB}_{user}'), 'big')
            balanceLT = int.from_bytes(readDB(f'balance_{tokenA}+{tokenB}_{user}'), 'big')
            poolA = int.from_bytes(readDB(f'pool_{tokenA}_{tokenB}_{tokenA}'), 'big')
            poolB = int.from_bytes(readDB(f'pool_{tokenA}_{tokenB}_{tokenB}'), 'big')
            totalSupplyLT = int.from_bytes(readDB(f'total_supply_{tokenA}_{tokenB}'), 'big')

            import math
            amtLT = math.floor(math.sqrt(amtA * amtB))

            print('before initPool', balanceA, amtA, balanceB, amtB, totalSupplyLT, balanceLT, poolA, poolB, amtLT)
            mpcInput(balanceA, amtA, balanceB, amtB, totalSupplyLT, balanceLT, poolA, poolB, amtLT)
            balanceA = sfix._new(balanceA)
            amtA = sfix._new(amtA)
            balanceB = sfix._new(balanceB)
            amtB = sfix._new(amtB)
            totalSupplyLT = sfix._new(totalSupplyLT)
            balanceLT = sfix._new(balanceLT)
            poolA = sfix._new(poolA)
            poolB = sfix._new(poolB)
            amtLT = sfix._new(amtLT)

            enoughA = (balanceA >= amtA)
            enoughB = (balanceB >= amtB)
            zeroTotalLT = (totalSupplyLT == 0)
            validOrder = (enoughA * enoughB * zeroTotalLT).reveal()

            amtA *= validOrder
            amtB *= validOrder
            amtLT *= validOrder

            balanceA -= amtA
            balanceB -= amtB
            balanceLT += amtLT
            poolA += amtA
            poolB += amtB
            totalSupplyLT += validOrder * amtLT

            balanceA = balanceA.v
            balanceB = balanceB.v
            balanceLT = balanceLT.v
            poolA = poolA.v
            poolB = poolB.v
            totalSupplyLT = totalSupplyLT.v
            mpcOutput(balanceA, balanceB, balanceLT, poolA, poolB, totalSupplyLT)

            print('after initPool', balanceA, balanceB, balanceLT, poolA, poolB, totalSupplyLT)

            writeDB(f'balance_{tokenA}_{user}', balanceA.to_bytes((balanceA.bit_length() + 7) // 8, 'big'))
            writeDB(f'balance_{tokenB}_{user}', balanceB.to_bytes((balanceB.bit_length() + 7) // 8, 'big'))
            writeDB(f'balance_{tokenA}+{tokenB}_{user}', balanceLT.to_bytes((balanceLT.bit_length() + 7) // 8, 'big'))
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenA}', poolA.to_bytes((poolA.bit_length() + 7) // 8, 'big'))
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenB}', poolB.to_bytes((poolB.bit_length() + 7) // 8, 'big'))
            writeDB(f'total_supply_{tokenA}_{tokenB}', totalSupplyLT.to_bytes((totalSupplyLT.bit_length() + 7) // 8, 'big'))

            initPrice = str(1. * amtB / amtA)
            print(initPrice, tokenA, tokenB)
            set(estimatedPrice, string memory initPrice, address tokenA, address tokenB)
        }
    }

    function addLiquidity(address tokenA, address tokenB, $uint amtA, $uint amtB) public {
        require(tokenA < tokenB);
        address user = msg.sender;

        mpc(address user, address tokenA, address tokenB, $uint amtA, $uint amtB) {
            balanceA = int.from_bytes(readDB(f'balance_{tokenA}_{user}'), 'big')
            balanceB = int.from_bytes(readDB(f'balance_{tokenB}_{user}'), 'big')
            balanceLT = int.from_bytes(readDB(f'balance_{tokenA}+{tokenB}_{user}'), 'big')
            poolA = int.from_bytes(readDB(f'pool_{tokenA}_{tokenB}_{tokenA}'), 'big')
            poolB = int.from_bytes(readDB(f'pool_{tokenA}_{tokenB}_{tokenB}'), 'big')
            totalSupplyLT = int.from_bytes(readDB(f'total_supply_{tokenA}_{tokenB}'), 'big')

            print('before addLiquidity', balanceA, amtA, balanceB, amtB, totalSupplyLT, balanceLT, poolA, poolB)

            mpcInput(balanceA, amtA, balanceB, amtB, totalSupplyLT, balanceLT, poolA, poolB)
            balanceA = sfix._new(balanceA)
            amtA = sfix._new(amtA)
            balanceB = sfix._new(balanceB)
            amtB = sfix._new(amtB)
            totalSupplyLT = sfix._new(totalSupplyLT)
            balanceLT = sfix._new(balanceLT)
            poolA = sfix._new(poolA)
            poolB = sfix._new(poolB)

            enoughA = (balanceA >= amtA)
            positiveA = (amtA > 0)
            enoughB = (balanceB >= amtB)
            positiveB = (amtB > 0)
            positiveTotalLT = (totalSupplyLT > 0)
            validOrder = (enoughA * positiveA * enoughB * positiveB * positiveTotalLT).reveal()

            surplusA = (amtA * poolB > amtB * poolA)
            nonSurplusA = 1 - surplusA
            changeA = validOrder * (surplusA * amtB * poolA / poolB + nonSurplusA * amtA)
            changeB = validOrder * (surplusA * amtB + nonSurplusA * amtA * poolB / poolA)
            changeLT = changeA * totalSupplyLT / poolA

            balanceA -= changeA
            balanceB -= changeB
            balanceLT += changeLT
            poolA += changeA
            poolB += changeB
            totalSupplyLT += changeLT

            balanceA = balanceA.v
            balanceB = balanceB.v
            balanceLT = balanceLT.v
            poolA = poolA.v
            poolB = poolB.v
            totalSupplyLT = totalSupplyLT.v
            mpcOutput(balanceA, balanceB, balanceLT, poolA, poolB, totalSupplyLT)

            print('after addLiquidity', balanceA, balanceB, balanceLT, poolA, poolB, totalSupplyLT)

            writeDB(f'balance_{tokenA}_{user}', balanceA.to_bytes((balanceA.bit_length() + 7) // 8, 'big'))
            writeDB(f'balance_{tokenB}_{user}', balanceB.to_bytes((balanceB.bit_length() + 7) // 8, 'big'))
            writeDB(f'balance_{tokenA}+{tokenB}_{user}', balanceLT.to_bytes((balanceLT.bit_length() + 7) // 8, 'big'))
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenA}', poolA.to_bytes((poolA.bit_length() + 7) // 8, 'big'))
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenB}', poolB.to_bytes((poolB.bit_length() + 7) // 8, 'big'))
            writeDB(f'total_supply_{tokenA}_{tokenB}', totalSupplyLT.to_bytes((totalSupplyLT.bit_length() + 7) // 8, 'big'))
        }
    }

    function removeLiquidity(address tokenA, address tokenB, $uint amt) public {
        require(tokenA < tokenB);
        address user = msg.sender;

        mpc(address user, address tokenA, address tokenB, $uint amt) {
            balanceA = int.from_bytes(readDB(f'balance_{tokenA}_{user}'), 'big')
            balanceB = int.from_bytes(readDB(f'balance_{tokenB}_{user}'), 'big')
            balanceLT = int.from_bytes(readDB(f'balance_{tokenA}+{tokenB}_{user}'), 'big')
            poolA = int.from_bytes(readDB(f'pool_{tokenA}_{tokenB}_{tokenA}'), 'big')
            poolB = int.from_bytes(readDB(f'pool_{tokenA}_{tokenB}_{tokenB}'), 'big')
            totalSupplyLT = int.from_bytes(readDB(f'total_supply_{tokenA}_{tokenB}'), 'big')

            mpcInput(balanceLT, amt, poolA, poolB, totalSupplyLT, balanceA, balanceB)
            balanceLT = sfix._new(balanceLT)
            amt = sfix._new(amt)
            poolA = sfix._new(poolA)
            poolB = sfix._new(poolB)
            totalSupplyLT = sfix._new(totalSupplyLT)
            balanceA = sfix._new(balanceA)
            balanceB = sfix._new(balanceB)

            enoughLT = balanceLT >= amt
            positiveLT = amt > 0

            validOrder = enoughLT * positiveLT

            changeLT = validOrder * amt
            changeA = changeLT * poolA / totalSupplyLT
            changeB = changeLT * poolB / totalSupplyLT

            poolA -= changeA
            poolB -= changeB
            balanceA += changeA
            balanceB += changeB
            balanceLT -= changeLT
            totalSupplyLT -= changeLT

            zeroTotalLT = (totalSupplyLT == 0).reveal()
            zeroTotalLT = sint(zeroTotalLT)

            print_ln('poolA %s', poolA.reveal())
            print_ln('poolB %s', poolB.reveal())
            print_ln('balanceA %s', balanceA.reveal())
            print_ln('balanceB %s', balanceB.reveal())
            print_ln('balanceLT %s', balanceLT.reveal())
            print_ln('totalSupplyLT %s', totalSupplyLT.reveal())
            print_ln('zeroTotalLT %s', zeroTotalLT.reveal())

            poolA = poolA.v
            poolB = poolB.v
            balanceA = balanceA.v
            balanceB = balanceB.v
            balanceLT = balanceLT.v
            totalSupplyLT = totalSupplyLT.v

            mpcOutput(poolA, poolB, balanceA, balanceB, balanceLT, totalSupplyLT, zeroTotalLT)

            writeDB(f'balance_{tokenA}_{user}', balanceA.to_bytes((balanceA.bit_length() + 7) // 8, 'big'))
            writeDB(f'balance_{tokenB}_{user}', balanceB.to_bytes((balanceB.bit_length() + 7) // 8, 'big'))
            writeDB(f'balance_{tokenA}+{tokenB}_{user}', balanceLT.to_bytes((balanceLT.bit_length() + 7) // 8, 'big'))
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenA}', poolA.to_bytes((poolA.bit_length() + 7) // 8, 'big'))
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenB}', poolB.to_bytes((poolB.bit_length() + 7) // 8, 'big'))
            writeDB(f'total_supply_{tokenA}_{tokenB}', totalSupplyLT.to_bytes((totalSupplyLT.bit_length() + 7) // 8, 'big'))

            if zeroTotalLT == 1:
                price = ''
                set(estimatedPrice, string memory price, address tokenA, address tokenB)
        }
    }

    function trade(address tokenA, address tokenB, $uint amtA, $uint amtB) public {
        require(tokenA < tokenB);
        address user = msg.sender;
        uint tradeSeq = ++tradeCnt;

        mpc(uint tradeSeq, address user, address tokenA, address tokenB, $uint amtA, $uint amtB) {
            balanceA = int.from_bytes(readDB(f'balance_{tokenA}_{user}'), 'big')
            balanceB = int.from_bytes(readDB(f'balance_{tokenB}_{user}'), 'big')
            poolA = int.from_bytes(readDB(f'pool_{tokenA}_{tokenB}_{tokenA}'), 'big')
            poolB = int.from_bytes(readDB(f'pool_{tokenA}_{tokenB}_{tokenB}'), 'big')
            totalPrice = int.from_bytes(readDB(f'totalPrice_{tokenA}_{tokenB}'), 'big')
            totalCnt = int.from_bytes(readDB(f'totalCnt_{tokenA}_{tokenB}'), 'big')

            print('before trade', balanceA, amtA, balanceB, amtB, poolA, poolB, totalPrice, totalCnt)
            mpcInput(balanceA, amtA, balanceB, amtB, poolA, poolB, totalPrice, totalCnt)
            balanceA = sfix._new(balanceA)
            amtA = sfix._new(amtA)
            balanceB = sfix._new(balanceB)
            amtB = sfix._new(amtB)
            poolA = sfix._new(poolA)
            poolB = sfix._new(poolB)
            totalPrice = sfix._new(totalPrice)

            feeRate = 0.003
            batchSize = 2

            validOrder = (amtA * amtB) < 0

            buyA = (amtA > 0)
            totalB = (1 + feeRate) * amtB
            enoughB = (-totalB  <= balanceB)
            actualAmtA = poolA  - poolA * poolB / (poolB  - amtB)
            acceptA = (actualAmtA  >= amtA)
            flagBuyA = validOrder * buyA * enoughB * acceptA

            buyB = 1 - buyA
            totalA = (1 + feeRate) * amtA
            enoughA = (-totalA  <= balanceA)
            actualAmtB = poolB  - poolA * poolB / (poolA  - amtA)
            acceptB = (actualAmtB  >= amtB)
            flagBuyB = validOrder * buyB * enoughA * acceptB

            changeA = flagBuyA * actualAmtA + flagBuyB * totalA
            changeB = flagBuyA * totalB + flagBuyB * actualAmtB

            poolA -= changeA
            poolB -= changeB
            balanceA += changeA
            balanceB += changeB

            print_ln('balanceA %s', balanceA.reveal())
            print_ln('balanceB %s', balanceB.reveal())
            print_ln('poolA %s', poolA.reveal())
            print_ln('poolB %s', poolB.reveal())

            orderSucceed = (flagBuyA + flagBuyB).reveal()
            print_ln('orderSucceed %s', orderSucceed.reveal())

            price = - changeB / (changeA + 1 - orderSucceed)
            print_ln('price %s', price.reveal())
            totalPrice += price
            totalCnt += orderSucceed

            batchPrice = 0
            if_then(totalCnt.reveal() >= batchSize)
            batchPrice = (totalPrice / totalCnt).reveal()
            end_if()
            balanceA = balanceA.v
            balanceB = balanceB.v
            poolA = poolA.v
            poolB = poolB.v
            price = price.v
            totalPrice = totalPrice.v
            batchPrice = sint(batchPrice.v)
            print_ln('batchPrice %s', batchPrice.reveal())
            mpcOutput(balanceA, balanceB, poolA, poolB, price, totalPrice, totalCnt, batchPrice)

            print('after trade', balanceA, balanceB, poolA, poolB, price, totalPrice, totalCnt, batchPrice)
            writeDB(f'balance_{tokenA}_{user}', balanceA.to_bytes((balanceA.bit_length() + 7) // 8, 'big'))
            writeDB(f'balance_{tokenB}_{user}', balanceB.to_bytes((balanceB.bit_length() + 7) // 8, 'big'))
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenA}', poolA.to_bytes((poolA.bit_length() + 7) // 8, 'big'))
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenB}', poolB.to_bytes((poolB.bit_length() + 7) // 8, 'big'))

            if batchPrice > 0:
                batchPrice = str(1. * batchPrice / fp)
                print('batchPrice', batchPrice)
                set(estimatedPrice, string memory batchPrice, address tokenA, address tokenB)
                totalPrice = 0
                totalCnt = 0
            writeDB(f'totalPrice_{tokenA}_{tokenB}', totalPrice.to_bytes((totalPrice.bit_length() + 7) // 8, 'big'))
            writeDB(f'totalCnt_{tokenA}_{tokenB}', totalCnt.to_bytes((totalCnt.bit_length() + 7) // 8, 'big'))

            import time
            returnPriceInterval = 60
            time.sleep(returnPriceInterval)
            writeDB(f'price_{tradeSeq}', price.to_bytes((price.bit_length() + 7) // 8, 'big'))
        }

    }
}
