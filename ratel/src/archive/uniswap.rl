int returnPriceInterval
int tradeCnt
int batchSize
fix feeRate
mapping { address, address => fix } publicBalance
mapping { address, addrses => string } estimatedPrice

mapping { address, address => sfix } secretBalance
mapping { address, address => sfix, sfix } poolSize
mapping { address, address => sfix } totalSupplyLT
mapping { address, address, address => sfix } secretBalanceLT
mapping { int => sfix} individualTradePrice
mapping { address, address => sfix } totalPrice
mapping { address, address => int } totalCnt

func publicDeposit(address token, fix amt) payable {
  address user = msg.sender
  require(amt > 0)
  if (token == address(0x0)) {
    require(msg.value == amt) // take care: unit conversion
  } else {
    IERC20(token).safeTransferFrom(user, address(this), amt) // take care: unit conversion
  }
  publicBalance[token][user] += amt
}

func secretDeposit(address token, fix amt) {
  address user = msg.sender
  require(amt > 0 && publicBalance[token][user] >= amt)
  publicBalance[token][user] -= amt
  secretBalance[token][user] += amt
}

func secretWithdraw(address token, address user, fix amt) {
  require(secretBalance[token][user] >= amt && amt > 0)
  secretBalance[token][user] -= amt
  submit(publicBalance[token][user] += amt) // "submit" is a special key word to update data on mainchain contract, notice that sidechain consensus is needed to do so
}

func publicWithdraw(address token, fix amt) {
  address user = msg.sender
  require(amt > 0 && publicBalance[token][user] >= amt)
  if (token == address(0x0)) {
    user.transfer(amt) // take care: unit conversion
  } else {
    IERC20(token).safeTransfer(user, amt) // take care: unit conversion
  }
  publicBalance[token][user] -= amt)
}

func trade(address tokenA, address tokenB, sfix amtA, sfix amtB) {
  require(tokenA < tokenB)
  address user = msg.sender
  int tradeSeq = ++tradeCnt

  sint validOrder = (amtA * amtB < 0)

  sint buyA = (amtA > 0)
  sint totalB = (1 + feeRate) * amtB
  sint enoughB = (-totalB  <= secretBalance[tokenB][user])
  sfix poolA, poolB = poolSize[tokenA][tokenB]
  sint actualAmtA = poolA  - poolA * poolB / (poolB  - amtB)
  sint acceptA = (actualAmtA  >= amtA)
  sint flagBuyA = validOrder * buyA * enoughB * acceptA

  sint buyB = 1 - buyA
  sint totalA = (1 + feeRate) * amtA
  sint enoughA = (-totalA  <= secretBalance[tokenA][user])
  sint actualAmtB = poolB  - poolA * poolB / (poolA  - amtA)
  sint acceptB = (actualAmtB  >= amtB)
  sint flatBuyB = validOrder * buyB * enoughA * acceptB

  sfix changeA = flagBuyA * actualAmtA + flagBuyB * totalA
  sfix changeB = flagBuyA * totalB + flagBuyB * actualAmtB

  int orderSucceed = (flagBuyA + flagBuyB).reveal()

  wait(returnPriceInterval) // "wait" is a special key word of local mpc code

  poolSize[tokenA][tokenB] = poolA - changeA, poolB - changeB
  secretBalance[tokenA][user] += changeA
  secretBalance[tokenB][user] += changeB

  sfix price = 0
  if orderSucceed == 1 {
    sfix price = -changeB / changeA
    totalPrice[tokenA][tokenB] += price
    totalCnt[tokenA][tokenB]++
  }
  individualTradePrice[tradeSeq] = price

  if totalCnt[tokenA][tokenB] >= batchSize {
    int batchPrice = (totalPrice[tokenA][tokenB] / totalCnt[tokenA][tokenB]).reveal()
    submit(estimatedPrice[tokenA][tokenB] = batchPrice) // "submit" is a special key word to update data on mainchain contract, notice that sidechain consensus is needed to do so
    totalPrice[tokenA][tokenB] = 0
    totalCnt[tokenA][tokenB] = 0
  }
}

func initPool(address tokenA, address tokenB, fix amtA, fix amtB) {
  require(tokenA < tokenB && amtA > 0 && amtB > 0)
  address user = msg.sender

  sint enoughA = (secretBalance[tokenA][user] >= amtA)
  sint enoughB = (secretBalance[tokenB][user] >= amtB)
  sint zeroTotalLT = (totalSupplyLT[tokenA][tokenB] == 0)
  int validOrder = (enoughA * enoughB * zeroTotalLT).reveal()

  if validOrder == 1 {
    poolSize[tokenA][tokenB] = amtA, amtB
    secretBalance[tokenA][user] -= amtA
    secretBalance[tokenB][user] -= amtB
    fix amtLT = sqrt(amtA * amtB)
    secretBalanceLT[tokenA][tokenB][user] = amtLT
    totalSupplyLT[tokenA][tokenB] = amtLT

    fix initPrice = amtB / amtA
    submit(estimatedPrice[tokenA][tokenB] = initPrice)
  }
}

func addLiquidity(address tokenA, address tokenB, sfix amtA, sfix amtB) {
  require(tokenA < tokenB)
  address user = msg.sender

  // amtA is valid
  sint enoughA = (secretBalance[tokenA][user] >= amtA)
  sint positiveA = (amtA > 0)

  // amtB is valid
  sint enoughB = (secretBalance[tokenB][user] >= amtB)
  sint positiveB = (amtB > 0)

  // pool has been initiated. This must check off-chain. Even though estimated price on contract would be set to "" when pool is empty, the state might not be synced timely.
  sint positiveTotalLT = (totalSupplyLT[tokenA][tokenB] > 0)

  sint validOrder = enoughA * positiveA * enoughB * positiveB * positiveTotalLT

  sint surplusA = amtA * poolB > amtB * poolA
  sint nonSurplusA = 1 - surplusA
  sfix changeA = validOrder * (surplusA * amtB * poolA / poolB + nonSurplusA * amtA)
  sfix changeB = validOrder * (surplusA * amtB + nonSurplusA * amtA * poolB / poolA)
  sfix changeLT = changeA * totalSupplyLT[tokenA][tokenB] / poolA

  poolSize[tokenA][tokenB] += changeA, changeB
  secretBalance[tokenA][user] -= changeA
  secretBalance[tokenB][user] -= changeB
  secretBalanceLT[tokenA][tokenB][user] += changeLT
  totalSupplyLT[tokenA][tokenB] += changeLT
}

func removeLiquidity(address tokenA, address tokenB, sfix amt) {
  require(tokenA < tokenB)
  address user = msg.sender

  // amt is valid
  sint enoughLT = (secretBalanceLT[tokenA][tokenB][user] >= amt)
  sint positiveLT = (amt > 0)

  // totalSupplyLT > 0 because 0 < amt <= balanceLT

  sint validOrder = enoughLT * positiveLT

  sfix poolA, poolB = poolSize[tokenA][tokenB]
  sfix changeLT = validOrder * amt
  sfix changeA = changeLT * poolA / totalSupplyLT
  sfix changeB = changeLT * poolB / totalSupplyLT

  poolSize[tokenA][tokenB] -= changeA, changeB
  secretBalance[tokenA][user] += changeA
  secretBalance[tokenB][user] += changeB
  secretBalanceLT[tokenA][tokenB][user] -= changeLT
  totalSupplyLT[tokenA][tokenB] -= changeLT

  int zeroTotalLT = (totalSupplyLT[tokenA][tokenB] == 0).reveal()
  if zeroTotalLT == 1 {
    submit(estimatedPrice[tokenA][tokenB] = "") // empty string indicates the pool is empty
  }
}