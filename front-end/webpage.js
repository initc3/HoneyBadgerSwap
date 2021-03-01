const p = BigInt("52435875175126190479447740508185965837690552500527637822603658699938581184513")
const n = 4
const t = 1
const fp = 1 << 16
const eth = "0x0000000000000000000000000000000000000000"
const checkPointInterval = 20 * 1000

const hbswapAddr = "0xe4d40ec72bf5da61a872af12011c7cadd9c49793"
const token1 = "0x63e7f20503256ddcfec64872aadb785d5a290cbb"
const token2 = "0x403b0f962566ffb960d0de98875dc09603aa67e9"

const basePort = 58080
const feeRate = 0.003
const fixPointDigit = 8

// **** Internal functions ****

function isETH(token) {
    return token == tokenList.get('eth')
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
}

function fromFloat(f) {
    return BigInt(Math.round(parseFloat(f) * fp))
}

function toFloat(i) {
    return parseFloat(i) / fp
}

function getElement(st, idx) {
    return st.slice(2 + idx * 64, 2 + (idx + 1) * 64)
}

function getInt(st, idx) {
    return parseInt('0x' + getElement(st, idx))
}

function extended_gcd(a, b) {
    let s = BigInt(0)
    let old_s = BigInt(1)
    let t = BigInt(1)
    let old_t = BigInt(0)
    let r = b
    let old_r = a

    while (r !== BigInt(0)) {
        quotient = ~~(old_r / r);
        [old_r, r] = [r, old_r - quotient * r];
        [old_s, s] = [s, old_s - quotient * s];
        [old_t, t] = [t, old_t - quotient * t];
    }

    return [old_r, old_s, old_t]
}

function mod_reduce(x, p) {
    let r = x % p
    return r >= 0 ? r : r + p
}

function modular_inverse(x, p) {
    let gcd, s, t;
    [gcd, s, t] = extended_gcd(x, p)
    return gcd > 0 ? s : -s
}

function interpolate(n, t, r, p) {
    if (r.length !== n) {
        return false
    }
    let f0 = BigInt(0)
    let f
    for (let i = 0; i <= t; i++) {
        f = BigInt(1)
        for (let j = 0; j <= t; j++) {
            if (i !== j) {
                f *= mod_reduce((BigInt(0) - BigInt(j + 1)) * modular_inverse(BigInt(i - j), p), p)
            }
        }
        f0 += mod_reduce(r[i] * f, p)
    }
    return mod_reduce(f0, p)
}

// ********

// **** Fetch secret value from mpc servers ****

async function getSecretBalance(token, user, prefix='') {
    let shares = []
    for (let i = 0; i < 4; i++) {
        url = "http://localhost:" + (basePort + i) + "/balance/" + token + ',' + user
        console.log(url)
        const share = (await (await fetch(url, {mode: 'cors'})).json()).balance
        $("#" + prefix + i).text(share)
        shares.push(BigInt(share))
    }
    return toFloat(interpolate(n, t, shares, p))
}

async function getTradePrice(tradeSeq) {
    let shares = []
    for (let i = 0; i < 4; i++) {
        url = "http://localhost:" + (basePort + i) + "/price/" + tradeSeq
        console.log(url)
        let share = (await (await fetch(url, {mode: 'cors'})).json()).price
        $("#tradePrice" + i).text(share)
        shares.push(BigInt(share))
    }
    return toFloat(interpolate(n, t, shares, p))
}

async function getInputmasks(srv, idxes) {
    url = "http://localhost:" + (basePort + srv) + "/inputmasks/" + idxes
    const shares = (await (await fetch(url, {mode: 'cors'})).json()).inputmask_shares.split(',')
    return [BigInt(shares[0]), BigInt(shares[1])]
}

async function getServerLog(srv, lines) {
    console.log('get log')
    url = "http://localhost:" + (basePort + srv) + "/log/" + lines
    const log = (await (await fetch(url, {mode: 'cors'})).json()).log
    $("#log").text(log)
}

// ********

// **** Access values on blockchain ****

async function getPersonalBalance(token, user) {
    console.log('getPersonalBalance')
    if (isETH(token)) {
        return await web3.eth.getBalance(user)
    } else {
        return await (contractList.get(token)).methods.balanceOf(user).call()
    }
}

async function getContractBalance(token, user) {
    return await hbswapContract.methods.balances(token, user).call()
}

// ********

// **** Global functions ****

async function trade() {
    $("#idxes").text(' ')
    $("#masks").text(' ')
    $("#seq").text(' ')
    $("#seq").text(' ')
    $("#price").text(' ')
    $("#tradePrice0").text(' ')
    $("#tradePrice1").text(' ')
    $("#tradePrice2").text(' ')
    $("#tradePrice3").text(' ')

    const tokenFrom = tokenList.get($( "#tradeFromToken option:selected" ).text())
    const tokenTo = tokenList.get($( "#tradeToToken option:selected" ).text())

    if (tokenFrom == tokenTo) {
        return
    }

    // Step 1: Claim inputmasks
    let tx = await hbswapContract.methods.tradePrep().send({from: user})
    data = tx['events']['TradePrep']['raw']['data']
    const idxA = getInt(data, 1)
    const idxB = getInt(data, 2)
    const idxes = idxA + ',' + idxB
    $("#idxes").text(idxes)

    // Step 2: Fetch inputmasks from servers
    let maskAShares = []
    let maskBShares = []
    for (let i = 0; i < 4; i++) {
        const shares = await getInputmasks(i, idxes)
        maskAShares.push(shares[0])
        maskBShares.push(shares[1])
    }

    // Step 3: Reconstruct inputmasks
    const maskA = interpolate(n, t, maskAShares, p)
    const maskB = interpolate(n, t, maskBShares, p)
    $("#masks").text(maskA + ',' + maskB)

    // Step 4: Publish masked inputs
    console.log($("#minRecv").text())
    let tokenA, tokenB, valueA, valueB
    if (tokenFrom < tokenTo) {
        tokenA = tokenFrom
        tokenB = tokenTo
        valueA = -fromFloat($("#amtTradeFrom").val())
        valueB = fromFloat($("#minRecv").text())
    } else {
        tokenA = tokenTo
        tokenB = tokenFrom
        valueA = fromFloat($("#minRecv").text())
        valueB = -fromFloat($("#amtTradeFrom").val())
    }
    const maskedValueA = valueA + maskA
    const maskedValueB = valueB + maskB
    tx = await hbswapContract.methods.trade(tokenA, tokenB, idxA, idxB, maskedValueA, maskedValueB).send({from: user})


    // Step 5: Get price of current trade
    data = tx['events']['Trade']['raw']['data']
    const tradeSeq = getInt(data, 0)
    $("#seq").text(tradeSeq)
    const price = await getTradePrice(tradeSeq)
    $("#price").text(price)
}

async function deposit() {
    const token = tokenList.get($( "#depositToken option:selected" ).text())
    const amt = $("#deposit").val()

    // Display balances before public deposit
    const prevSecretBalance = await getSecretBalance(token, user, 'depositUpdate')
    $("#depositStatus").text('public depositing...')
    $("#personalBalance").text(await getPersonalBalance(token, user))
    $("#contractBalance").text(await getContractBalance(token, user))
    $("#secretBalance").text(prevSecretBalance)
    $("#balance").text(prevSecretBalance)

    // Public deposit
    if (isETH(token)) {
        await hbswapContract.methods.deposit(token, amt).send({from: user, value: amt})
    } else {
        // Approve before token transfer
        await contractList.get(token).methods.approve(hbswapAddr, amt).send({from: user})

        await hbswapContract.methods.deposit(token, amt).send({from: user})
    }

    // Display balances after public deposit
    $("#depositStatus").text('secret depositing...')
    $("#personalBalance").text(await getPersonalBalance(token, user))
    $("#contractBalance").text(await getContractBalance(token, user))

    // Secret deposit
    await hbswapContract.methods.secretDeposit(token, amt).send({from: user})

    while (true) {
        if (prevSecretBalance < await getSecretBalance(token, user, 'depositUpdate')) {
            break
        }
        await sleep(5000)
    }

    // Display balances after secret deposit
    const curSecretBalance = await getSecretBalance(token, user, 'depositUpdate')
    $("#depositStatus").text('done')
    $("#contractBalance").text(await getContractBalance(token, user))
    $("#secretBalance").text(curSecretBalance)
    $("#balance").text(curSecretBalance)
}

async function withdraw() {
    const token = tokenList.get($( "#depositToken option:selected" ).text())
    const amt = $("#deposit").val()

    // Display balances before secret withdraw
    const prevContractBalance = await getContractBalance(token, user)
    const prevSecretBalance = await getSecretBalance(token, user, 'depositUpdate')
    $("#depositStatus").text('secret withdrawing...')
    $("#personalBalance").text(await getPersonalBalance(token, user))
    $("#contractBalance").text(prevContractBalance)
    $("#secretBalance").text(prevSecretBalance)
    $("#balance").text(prevSecretBalance)

    // Secret withdraw
    await hbswapContract.methods.secretWithdraw(token, amt).send({from: user})
    while (true) {
        if (prevContractBalance < await getContractBalance(token, user, 'depositUpdate')) {
            break
        }
        await sleep(5000)
    }

    // Display balances after secret withdraw
    const curSecretBalance = await getSecretBalance(token, user, 'depositUpdate')
    $("#depositStatus").text('public withdrawing...')
    $("#contractBalance").text(await getContractBalance(token, user))
    $("#secretBalance").text(curSecretBalance)
    $("#balance").text(curSecretBalance)

    // Public withdraw
    await hbswapContract.methods.withdraw(token, amt).send({from: user})

    // Display balances after public withdraw
    $("#depositStatus").text('done')
    $("#personalBalance").text(await getPersonalBalance(token, user))
    $("#contractBalance").text(await getContractBalance(token, user))
}

async function initPool() {
    const tokenA = tokenList.get($( "#poolTokenA option:selected" ).text())
    const tokenB = tokenList.get($( "#poolTokenB option:selected" ).text())
    const amtA = $("#amtPoolTokenA").val()
    const amtB = $("#amtPoolTokenB").val()

    const prevSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)

    await hbswapContract.methods.initPool(tokenA, tokenB, amtA, amtB).send({from: user})
    let curSecretLiquidityTokenBalance
    while (true) {
        curSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)
        if (prevSecretLiquidityTokenBalance != curSecretLiquidityTokenBalance) {
            break
        }
        await sleep(5000)
    }

    updatePoolPair()
}

async function addLiquidity() {
    const tokenA = tokenList.get($( "#poolTokenA option:selected" ).text())
    const tokenB = tokenList.get($( "#poolTokenB option:selected" ).text())
    const amtA = $("#amtPoolTokenA").val()
    const amtB = $("#amtPoolTokenB").val()

    const prevSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)

    await hbswapContract.methods.addLiquidity(tokenA, tokenB, amtA, amtB).send({from: user})

    let curSecretLiquidityTokenBalance
    while (true) {
        curSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)
        if (prevSecretLiquidityTokenBalance != curSecretLiquidityTokenBalance) {
            break
        }
        await sleep(5000)
    }

    updatePoolPair()
}

async function removeLiquidity() {
    const tokenA = tokenList.get($( "#poolTokenA option:selected" ).text())
    const tokenB = tokenList.get($( "#poolTokenB option:selected" ).text())
    const amt = $("#amtPoolLiquidityToken").val()

    const prevSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)

    await hbswapContract.methods.removeLiquidity(tokenA, tokenB, amt).send({from: user})

    let curSecretLiquidityTokenBalance
    while (true) {
        curSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)
        if (prevSecretLiquidityTokenBalance != curSecretLiquidityTokenBalance) {
            break
        }
        await sleep(5000)
    }

    updatePoolPair()
}

async function updateTradePair() {
    $("#tradeInfo").text(" ")
    $("#estPriceTrade").text(" ")
    $("#balanceTradeFrom").text(" ")
    $("#balanceTradeTo").text(" ")
    $("#idxes").text(' ')
    $("#masks").text(' ')
    $("#seq").text(' ')
    $("#seq").text(' ')
    $("#price").text(' ')
    $("#tradePrice0").text(' ')
    $("#tradePrice1").text(' ')
    $("#tradePrice2").text(' ')
    $("#tradePrice3").text(' ')

    const fromToken = tokenList.get($( "#tradeFromToken option:selected" ).text())
    const toToken = tokenList.get($( "#tradeToToken option:selected" ).text())

    if (fromToken == toToken) {
        $("#tradeInfo").text("Error: invalid token pair!")
        return
    }

    let price
    if (fromToken < toToken) {
        price = await hbswapContract.methods.prices(fromToken, toToken).call()
        $("#estPriceTrade").text(price)
    } else {
        price = await hbswapContract.methods.prices(toToken, fromToken).call()
        $("#estPriceTrade").text(price)
        price = 1. / parseFloat(price)
    }

    $("#balanceTradeFrom").text(await getSecretBalance(fromToken, user))
    $("#balanceTradeTo").text(await getSecretBalance(toToken, user))
}

async function updateAmtTradeFrom() {
    const _fromToken = $( "#tradeFromToken option:selected").text()
    const fromToken = tokenList.get(_fromToken)
    const toToken = tokenList.get($( "#tradeToToken option:selected" ).text())

    if (fromToken == toToken) {
        $("#tradeInfo").text("Error: invalid token pair!")
        return
    }

    const amtTradeFrom = $( "#amtTradeFrom" ).val()
    const slippage = $( "#slippage option:selected" ).val()

    let price
    if (fromToken < toToken) {
        price = await hbswapContract.methods.prices(fromToken, toToken).call()
    } else {
        price = await hbswapContract.methods.prices(toToken, fromToken).call()
        price = 1. / parseFloat(price)
    }

    const amtTradeTo = amtTradeFrom * price
    const minReceived = amtTradeTo * (1 - slippage)
    const fee = amtTradeFrom * feeRate

    $("#amtTradeTo").val(amtTradeTo.toFixed(fixPointDigit))
    $("#minRecv").text(minReceived.toFixed(fixPointDigit))
    $("#fee").text(fee.toFixed(fixPointDigit) + ' ' + _fromToken)
}

async function updateAmtTradeTo() {
    const _fromToken = $( "#tradeFromToken option:selected").text()
    const fromToken = tokenList.get(_fromToken)
    const toToken = tokenList.get($( "#tradeToToken option:selected" ).text())

    if (fromToken == toToken) {
        $("#tradeInfo").text("Error: invalid token pair!")
        return
    }

    const amtTradeTo = $( "#amtTradeTo" ).val()
    const slippage = $( "#slippage option:selected" ).val()

    let price
    if (fromToken < toToken) {
        price = await hbswapContract.methods.prices(fromToken, toToken).call()
    } else {
        price = await hbswapContract.methods.prices(toToken, fromToken).call()
        price = 1. / parseFloat(price)
    }

    const amtTradeFrom = amtTradeTo / price
    const minReceived = amtTradeTo * (1 - slippage)
    const fee = amtTradeFrom * feeRate

    $("#amtTradeFrom").val(amtTradeFrom.toFixed(fixPointDigit))
    $("#minRecv").text(minReceived.toFixed(fixPointDigit))
    $("#fee").text(fee.toFixed(fixPointDigit) + ' ' + _fromToken)
}

async function updatePoolPair() {
    // Get pool pair
    const tokenA = tokenList.get($( "#poolTokenA option:selected" ).text())
    const tokenB = tokenList.get($( "#poolTokenB option:selected" ).text())

    if (tokenA >= tokenB) {
        $("#poolInfo").text("Error: invalid token pair!")
        $("#estPricePool").text(' ')
        $("#balancePoolTokenA").text(' ')
        $("#balancePoolTokenB").text(' ')
        $("#balancePoolLiquidityToken").text(' ')
        return
    }

    // Update estimated price
    let price = await hbswapContract.methods.prices(tokenA, tokenB).call()
    if (price == '') {
        price = 'Pool not initiated!'
    }

    $("#poolInfo").text(" ")
    $("#estPricePool").text(price)
    $("#balancePoolTokenA").text(await getSecretBalance(tokenA, user))
    $("#balancePoolTokenB").text(await getSecretBalance(tokenB, user))
    $("#balancePoolLiquidityToken").text(await getSecretBalance(tokenA + '+' + tokenB, user))
}

async function updateDepositToken() {
    const token = tokenList.get($( "#depositToken option:selected" ).text())

    $("#balance").text(await getSecretBalance(token, user))
    $("#depositStatus").text(' ')
    $("#personalBalance").text(' ')
    $("#contractBalance").text(' ')
    $("#secretBalance").text(' ')
    $("#depositUpdate0").text(' ')
    $("#depositUpdate1").text(' ')
    $("#depositUpdate2").text(' ')
    $("#depositUpdate3").text(' ')
}

// ********

async function init() {
    window.web3 = new Web3(ethereum)

    window.user = (await ethereum.request({ method: 'eth_requestAccounts'}))[0]
    $("#user").text(user)

    const hbswapABI = JSON.parse($("#hbswapABI").text())
    window.hbswapContract = new web3.eth.Contract(hbswapABI, hbswapAddr)

    const tokenABI = JSON.parse($("#tokenABI").text())
    window.tokenList = new Map()
    tokenList.set('eth', eth)
    tokenList.set('token1', token1)
    tokenList.set('token2', token2)
    window.contractList = new Map()
    for (let [k, v] of tokenList) {
        contractList.set(v, new web3.eth.Contract(tokenABI, v))
    }

    await updateTradePair()
    await updateDepositToken()
    await updatePoolPair()

    while (true) {
        await getServerLog(0, 100)
        await sleep(checkPointInterval)
    }
}

init()