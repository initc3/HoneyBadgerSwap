public import {
    "@openzeppelin/contracts/math/SafeMath.sol"
    "@openzeppelin/contracts/token/ERC20/IERC20.sol"
    "@openzeppelin/contracts/token/ERC20/SafeERC20.sol"
}

public declaration {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint constant public Decimals = 10**15;
    uint constant public Fp = 2**16;

    mapping (address => mapping (address => uint)) public publicBalance;

    event SecretDeposit(address token, address user, uint amt);

    constructor() public {}
}

public func publicDeposit(address token, uint amt) payable {
    address user = msg.sender;
    require(amt > 0);
    if (token == address(0x0)) {
        require(msg.value * Fp == amt * Decimals); // take care: unit conversion
    } else {
        IERC20(token).safeTransferFrom(user, address(this), amt / Fp * Decimals); // take care: unit conversion
    }
    publicBalance[token][user] += amt;
}

public func secretDeposit(address token, uint amt) {
    address user = msg.sender;
    require(amt > 0 && publicBalance[token][user] >= amt);
    publicBalance[token][user] -= amt;

    emit SecretDeposit(token, user, amt);
}

private func secretDeposit() {
    monitorEvent(SecretDeposit, token, user, amt)
    secretBalance = readDB(f'balance_{token}_{user}')
    secretBalance = int.from_bytes(secretBalance, 'big')
    print('!', secretBalance)
    mpcInput(secretBalance, amt)
    secretBalance = sfix._new(secretBalance)
    amt = sfix._new(amt)
    secretBalance += amt
    secretBalance = secretBalance.v
    mpcOutput(secretBalance)
    print(secretBalance)
    secretBalance = secretBalance.to_bytes((secretBalance.bit_length() + 7) // 8, 'big')
    writeDB(f'balance_{token}_{user}', secretBalance)
}