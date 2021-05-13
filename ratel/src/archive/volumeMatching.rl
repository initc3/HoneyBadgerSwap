int bidCnt

// deposit and withdraw similar to hbswap

mapping { address, address => int, address, sfix } bidList // sort bids in ascending order of bidSeq
mapping { int => sfix } matchedAmt

func submitBid(address tokenA, address tokenB, sfix amt) {
    address user = msg.sender
    int bidSeq = ++bidCnt

    int validBid = (amt != 0).reveal()

    if validBid {
        bidList[tokenA][tokenB].append((bidSeq, user, amt))
    }
}

func volumeMatch(address tokenA, address tokenB, fix price) { // price = amtTokenB / amtTokenA
    sint buySum, sellSum = sfix(0), sfix(0)
    for bidSeq, user, amtB in bidList[tokenA][tokenB] {
        sint amtA = amtB / price

        sint buyB = amtB > 0
        sint enoughA = amtA <= secretBalance[user][tokenA]
        sint buySum += buyB * enoughA * amtB

        sint sellB = 1 - buyB // amtB < 0
        sint enoughB = -amtB <= secretBalance[user][tokenB]
        sint sellSum -= sellB * enoughB * amtB
    }

    sint f = sellSum > buySum
    sint tradeAmt = f * (buySum - sellSum) + sellSum
    sint sellSum, buySum = tradeAmt, tradeAmt

    for bidSeq, user, amtB in bidList[tokenA][tokenB] {
        sint buyB = amtB > 0
        sint z1 = buySum <= 0
        sint z2 = buySum < amtB
        sint actualAmtB = buyB * (1 - z1) * (z2 * (buySum - amtB) + amtB)
        buySum -= actualAmtB
        sint actualAmtA = actualAmtB / price
        secretBalance[user][tokenA] -= actualAmtA
        secretBalance[user][tokenB] += actualAmtB

        sint sellB = 1 - buyB // amtB < 0
        z1 = sellSum <= 0
        z2 = sellSum < -amtB
        actualAmtB = buyB * (1 - z1) * (z2 * (buySum - amtB) + amtB)
        sellSum += actualAmtB
        actualAmtA = actualAmtB / price
        secretBalance[user][tokenA] -= actualAmtA
        secretBalance[user][tokenB] += actualAmtB

        matchedAmt[bidSeq] = actualAmtB
    }

    for bidSeq, _, _ in bidList[tokenA][tokenB] {
        matchedAmt[bidSeq].reveal()
    }

    bidList[tokenA][tokenB].clear()
}