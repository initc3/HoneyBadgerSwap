pragma solidity ^0.5.0;


import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";


contract hbswap {
    pureMpc addTotalPrice(server, seqTrade, totalPrice) {
        priceInfo = readDB(f'price_{seqTrade}', list)

        orderSucceed, changeA, changeB = priceInfo

        mpcInput(sfix totalPrice, sint orderSucceed, sfix changeA, sfix changeB)

        price = - changeB / (changeA + 1 - orderSucceed)
        totalPrice += price

        mpcOutput(sfix totalPrice)

        return totalPrice
    }

}
