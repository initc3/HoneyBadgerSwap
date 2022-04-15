pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract hbswap {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint constant public Decimals = 10**15;
    uint constant public Fp = 2**16;

    mapping (address => mapping (address => uint)) public publicBalance;

    mapping (address => mapping (address => string)) public estimatedPrice;
    mapping (address => string) public estimatedPriceValue;
    mapping (string => uint) public estimatedPriceCount;

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
            #print('**** begin deposit', seqSecretDeposit)

            secretBalance = readDB(f'balance_{token}_{user}', int)

            mpcInput(sfix secretBalance, sfix amt)

            secretBalance += amt
            print_ln("**** secretBalance %s", secretBalance.reveal())

            mpcOutput(sfix secretBalance)

            #print('**** secretBalance', token, secretBalance)
            writeDB(f'balance_{token}_{user}', secretBalance, int)

            #print('**** end deposit', seqSecretDeposit)
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

        mpc(address user, address token, uint amt) {
            balance = readDB(f'balance_{token}_{user}', int)

            mpcInput(sfix balance, sfix amt)

            enough = (balance >= amt).reveal()

            if_then(enough)
            balance -= amt
            end_if()

            mpcOutput(sfix balance, cint enough)

            print('****', balance, enough)
            if enough == 1:
                add(publicBalance, uint seqSecretWithdraw, uint amt, address token, address user)
                writeDB(f'balance_{token}_{user}', balance, int)
        }
    }

    function initPool(address tokenA, address tokenB, uint amtA, uint amtB) public {
        require(tokenA < tokenB && amtA > 0 && amtB > 0);
        address user = msg.sender;

        mpc(address user, address tokenA, address tokenB, uint amtA, uint amtB) {
            balanceA = readDB(f'balance_{tokenA}_{user}', int)
            balanceB = readDB(f'balance_{tokenB}_{user}', int)
            balanceLT = readDB(f'balance_{tokenA}+{tokenB}_{user}', int)
            poolA = readDB(f'pool_{tokenA}_{tokenB}_{tokenA}', int)
            poolB = readDB(f'pool_{tokenA}_{tokenB}_{tokenB}', int)
            totalSupplyLT = readDB(f'total_supply_{tokenA}_{tokenB}', int)

            import math
            amtLT = math.floor(math.sqrt(amtA * amtB))

            #print('**** before initPool', balanceA, amtA, balanceB, amtB, totalSupplyLT, balanceLT, poolA, poolB, amtLT)
            mpcInput(sfix balanceA, sfix amtA, sfix balanceB, sfix amtB, sfix totalSupplyLT, sfix balanceLT, sfix poolA, sfix poolB, sfix amtLT)

            enoughA = balanceA >= amtA
            enoughB = balanceB >= amtB
            zeroTotalLT = totalSupplyLT == 0
            validOrder = (enoughA * enoughB * zeroTotalLT).reveal()
            print_ln("**** balanceA %s", balanceA.reveal())
            print_ln("**** balanceB %s", balanceB.reveal())
            print_ln("**** poolA %s", poolA.reveal())
            print_ln("**** poolB %s", poolB.reveal())

            if_then(validOrder)
            balanceA -= amtA
            balanceB -= amtB
            balanceLT += amtLT
            poolA += amtA
            poolB += amtB
            totalSupplyLT += amtLT
            end_if()
            print_ln("**** balanceA %s", balanceA.reveal())
            print_ln("**** balanceB %s", balanceB.reveal())
            print_ln("**** balanceLT %s", balanceLT.reveal())
            print_ln("**** poolA %s", poolA.reveal())
            print_ln("**** poolB %s", poolB.reveal())
            print_ln("**** totalSupplyLT %s", totalSupplyLT.reveal())

            mpcOutput(cint validOrder, sfix balanceA, sfix balanceB, sfix balanceLT, sfix poolA, sfix poolB, sfix totalSupplyLT)

            #print('**** after initPool', balanceA, balanceB, balanceLT, poolA, poolB, totalSupplyLT)

            writeDB(f'balance_{tokenA}_{user}', balanceA, int)
            writeDB(f'balance_{tokenB}_{user}', balanceB, int)
            writeDB(f'balance_{tokenA}+{tokenB}_{user}', balanceLT, int)
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenA}', poolA, int)
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenB}', poolB, int)
            writeDB(f'total_supply_{tokenA}_{tokenB}', totalSupplyLT, int)

            if validOrder == 1:
                initPrice = str(1. * amtB / amtA)
                print('**** initPrice', initPrice, tokenA, tokenB)
                set(estimatedPrice, string memory initPrice, address tokenA, address tokenB)

        }
    }

    function addLiquidity(address tokenA, address tokenB, $uint amtA, $uint amtB) public {
        require(tokenA < tokenB);
        address user = msg.sender;

        mpc(address user, address tokenA, address tokenB, $uint amtA, $uint amtB) {
            balanceA = readDB(f'balance_{tokenA}_{user}', int)
            balanceB = readDB(f'balance_{tokenB}_{user}', int)
            balanceLT = readDB(f'balance_{tokenA}+{tokenB}_{user}', int)
            poolA = readDB(f'pool_{tokenA}_{tokenB}_{tokenA}', int)
            poolB = readDB(f'pool_{tokenA}_{tokenB}_{tokenB}', int)
            totalSupplyLT = readDB(f'total_supply_{tokenA}_{tokenB}', int)

            mpcInput(sfix balanceA, sfix amtA, sfix balanceB, sfix amtB, sfix totalSupplyLT, sfix balanceLT, sfix poolA, sfix poolB)
            #print_ln('**** poolA %s', poolA.reveal())
            #print_ln('**** poolB %s', poolB.reveal())
            #print_ln('**** balanceA %s', balanceA.reveal())
            #print_ln('**** balanceB %s', balanceB.reveal())
            #print_ln('**** balanceLT %s', balanceLT.reveal())
            #print_ln('**** totalSupplyLT %s', totalSupplyLT.reveal())
            #print_ln('**** amtA %s', amtA.reveal())
            #print_ln('**** amtB %s', amtB.reveal())

            enoughA = balanceA >= amtA
            positiveA = amtA > 0
            enoughB = balanceB >= amtB
            positiveB = amtB > 0
            positiveTotalLT = totalSupplyLT > 0
            validOrder = enoughA * positiveA * enoughB * positiveB * positiveTotalLT
            #print_ln('**** enoughA %s', enoughA.reveal())
            #print_ln('**** positiveA %s', positiveA.reveal())
            #print_ln('**** enoughB %s', enoughB.reveal())
            #print_ln('**** positiveB %s', positiveB.reveal())
            #print_ln('**** positiveTotalLT %s', positiveTotalLT.reveal())
            #print_ln('**** validOrder %s', validOrder.reveal())

            surplusA = (amtA * poolB) > (amtB * poolA)
            nonSurplusA = 1 - surplusA
            changeA = validOrder * (surplusA * amtB * poolA / poolB + nonSurplusA * amtA)
            changeB = validOrder * (surplusA * amtB + nonSurplusA * amtA * poolB / poolA)
            changeLT = changeA * totalSupplyLT / poolA
            #print_ln('**** surplusA %s', surplusA.reveal())
            #print_ln('**** nonSurplusA %s', nonSurplusA.reveal())
            #print_ln('**** changeA %s', changeA.reveal())
            #print_ln('**** changeB %s', changeB.reveal())
            #print_ln('**** changeLT %s', changeLT.reveal())

            balanceA -= changeA
            balanceB -= changeB
            balanceLT += changeLT
            poolA += changeA
            poolB += changeB
            totalSupplyLT += changeLT

            #print_ln('**** poolA %s', poolA.reveal())
            #print_ln('**** poolB %s', poolB.reveal())
            #print_ln('**** balanceA %s', balanceA.reveal())
            #print_ln('**** balanceB %s', balanceB.reveal())
            #print_ln('**** balanceLT %s', balanceLT.reveal())
            #print_ln('**** totalSupplyLT %s', totalSupplyLT.reveal())
            mpcOutput(sfix balanceA, sfix balanceB, sfix balanceLT, sfix poolA, sfix poolB, sfix totalSupplyLT)

            writeDB(f'balance_{tokenA}_{user}', balanceA, int)
            writeDB(f'balance_{tokenB}_{user}', balanceB, int)
            writeDB(f'balance_{tokenA}+{tokenB}_{user}', balanceLT, int)
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenA}', poolA, int)
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenB}', poolB, int)
            writeDB(f'total_supply_{tokenA}_{tokenB}', totalSupplyLT, int)
        }
    }

    function removeLiquidity(address tokenA, address tokenB, $uint amt) public {
        require(tokenA < tokenB);
        address user = msg.sender;

        mpc(address user, address tokenA, address tokenB, $uint amt) {
            balanceA = readDB(f'balance_{tokenA}_{user}', int)
            balanceB = readDB(f'balance_{tokenB}_{user}', int)
            balanceLT = readDB(f'balance_{tokenA}+{tokenB}_{user}', int)
            poolA = readDB(f'pool_{tokenA}_{tokenB}_{tokenA}', int)
            poolB = readDB(f'pool_{tokenA}_{tokenB}_{tokenB}', int)
            totalSupplyLT = readDB(f'total_supply_{tokenA}_{tokenB}', int)

            mpcInput(sfix balanceLT, sfix amt, sfix poolA, sfix poolB, sfix totalSupplyLT, sfix balanceA, sfix balanceB)

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

            print_ln('**** poolA %s', poolA.reveal())
            print_ln('**** poolB %s', poolB.reveal())
            print_ln('**** balanceA %s', balanceA.reveal())
            print_ln('**** balanceB %s', balanceB.reveal())
            print_ln('**** balanceLT %s', balanceLT.reveal())
            print_ln('**** totalSupplyLT %s', totalSupplyLT.reveal())
            print_ln('**** zeroTotalLT %s', zeroTotalLT)

            mpcOutput(sfix poolA, sfix poolB, sfix balanceA, sfix balanceB, sfix balanceLT, sfix totalSupplyLT, cint zeroTotalLT)

            writeDB(f'balance_{tokenA}_{user}', balanceA, int)
            writeDB(f'balance_{tokenB}_{user}', balanceB, int)
            writeDB(f'balance_{tokenA}+{tokenB}_{user}', balanceLT, int)
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenA}', poolA, int)
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenB}', poolB, int)
            writeDB(f'total_supply_{tokenA}_{tokenB}', totalSupplyLT, int)

            print('**** zeroTotalLT', zeroTotalLT)
            if zeroTotalLT == 1:
                price = ''
                set(estimatedPrice, string memory price, address tokenA, address tokenB)
        }
    }

    function trade(address tokenA, address tokenB, $uint amtA, $uint amtB) public {
        require(tokenA < tokenB);
        address user = msg.sender;

        mpc(address user, address tokenA, address tokenB, $uint amtA, $uint amtB) {
            times = []

            import time
            times.append(time.perf_counter())

            balanceA = readDB(f'balance_{tokenA}_{user}', int)
            balanceB = readDB(f'balance_{tokenB}_{user}', int)
            poolA = readDB(f'pool_{tokenA}_{tokenB}_{tokenA}', int)
            poolB = readDB(f'pool_{tokenA}_{tokenB}_{tokenB}', int)
            totalPrice = readDB(f'totalPrice_{tokenA}_{tokenB}', int)
            totalCnt = readDB(f'totalCnt_{tokenA}_{tokenB}', int)

            times.append(time.perf_counter())

            print(f'**** seqTrade {seqTrade} start')

            mpcInput(sfix balanceA, sfix amtA, sfix balanceB, sfix amtB, sfix poolA, sfix poolB, sfix totalPrice, sint totalCnt)

            print_ln('**** balanceA %s', balanceA.reveal())
            print_ln('**** balanceB %s', balanceB.reveal())
            print_ln('**** poolA %s', poolA.reveal())
            print_ln('**** poolB %s', poolB.reveal())

            feeRate = 0.003
            batchSize = 1000

            validOrder = (amtA * amtB) < 0

            poolProduct = poolA * poolB

            buyA = amtA > 0
            totalB = (1 + feeRate) * amtB
            enoughB = (-totalB) <= balanceB
            actualAmtA = poolA - poolProduct / (poolB - amtB)
            acceptA = actualAmtA >= amtA
            flagBuyA = validOrder * buyA * enoughB * acceptA

            buyB = 1 - buyA
            totalA = (1 + feeRate) * amtA
            enoughA = (-totalA) <= balanceA
            actualAmtB = poolB - poolProduct / (poolA - amtA)
            acceptB = actualAmtB >= amtB
            flagBuyB = validOrder * buyB * enoughA * acceptB

            changeA = flagBuyA * actualAmtA + flagBuyB * totalA
            changeB = flagBuyA * totalB + flagBuyB * actualAmtB

            poolA -= changeA
            poolB -= changeB
            balanceA += changeA
            balanceB += changeB

            print_ln('**** balanceA %s', balanceA.reveal())
            print_ln('**** balanceB %s', balanceB.reveal())
            print_ln('**** poolA %s', poolA.reveal())
            print_ln('**** poolB %s', poolB.reveal())

            orderSucceed = flagBuyA + flagBuyB

            price = - changeB / (changeA + 1 - orderSucceed)
            totalPrice += price
            totalCnt += orderSucceed

            batchPrice = ((totalCnt >= batchSize).reveal()).if_else((totalPrice / totalCnt).reveal(), cfix.from_int(0))

            mpcOutput(sfix balanceA, sfix balanceB, sfix poolA, sfix poolB, sfix price, sfix totalPrice, sint totalCnt, cfix batchPrice)

            times.append(time.perf_counter())

            writeDB(f'balance_{tokenA}_{user}', balanceA, int)
            writeDB(f'balance_{tokenB}_{user}', balanceB, int)
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenA}', poolA, int)
            writeDB(f'pool_{tokenA}_{tokenB}_{tokenB}', poolB, int)

            if batchPrice > 0:
                batchPrice = str(1. * batchPrice / fp)
                print('**** batchPrice', batchPrice)
                set(estimatedPrice, string memory batchPrice, address tokenA, address tokenB)
                totalPrice = 0
                totalCnt = 0

            writeDB(f'totalPrice_{tokenA}_{tokenB}', totalPrice, int)
            writeDB(f'totalCnt_{tokenA}_{tokenB}', totalCnt, int)
            writeDB(f'price_{seqTrade}', price, int)

            times.append(time.perf_counter())

            print(f'**** seqTrade {seqTrade} finish')

            with open(f'ratel/benchmark/data/latency_{server.serverID}.csv', 'a') as f:
                for op, t in enumerate(times):
                    f.write(f'trade\t'
                            f'seq\t{seqTrade}\t'
                            f'op\t{op + 1}\t'
                            f'cur_time\t{t}\n')
        }
    }
}
