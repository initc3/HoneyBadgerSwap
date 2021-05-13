public declaration {
    mapping (address => mapping (address => uint)) publicBalance;

    event SecretDeposit(address token, address user, uint amt);
}

public func secretDeposit(address token, uint amt) {
    address user = msg.sender;
    require(amt > 0 && publicBalance[token][user] >= amt);
    publicBalance[token][user] -= amt;

    emit SecretDeposit(token, user, amt);
}

private func secretDeposit(string token, string user, uint amt) {
    secretBalance = readDB(f'balance_{token}_{user}')
    mpcInput(secretBalance, amt)
    secretBalance += amt
    mpcOutput(secretBalance)
    writeDB(f'balance_{token}_{user}', secretBalance)
}