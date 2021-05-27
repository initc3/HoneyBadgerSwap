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
}
