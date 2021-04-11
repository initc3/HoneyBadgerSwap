const p = BigInt('52435875175126190479447740508185965837690552500527637822603658699938581184513')
const n = 4
const t = 1
const fp = 1 << 16

const hbswapAddr = '0x82ac888f567365362ea290f89368b7885227fc7e'
// const token1 = '0x63e7f20503256ddcfec64872aadb785d5a290cbb'
// const token2 = '0x403b0f962566ffb960d0de98875dc09603aa67e9'
const ethAddr = '0x0000000000000000000000000000000000000000'
const hbsAddr = "0x78160ee9e55fd81626f98d059c84d21d8b71bfda"
const daiAddr = "0x4f96fe3b7a6cf9725f59d353f723c1bdb64ca6aa"

const decimals = 1 // TODO: 10**18
const checkPointInterval = 20 * 1000
const feeRate = 0.003
const displayPrecision = 4
// const host = 'https://www.ratelang.org'
// const basePort = 8080
const host = 'http://localhost'
const basePort = 58080

// **** Internal functions ****

function isETH(token) {
    return token == tokenList.get('eth')
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
}

function floatToFix(f) {
    return BigInt(Math.round(f * fp))
}

function fixToFloat(i) {
    return parseFloat(i) / fp
}

function transferValue(x) {
    return x * decimals
}

function getElement(st, idx) {
    return st.slice(2 + idx * 64, 2 + (idx + 1) * 64)
}

function getInt(st, idx) {
    return parseInt('0x' + getElement(st, idx))
}

async function getInputMaskIndexes(num) {
    let tx = await hbswapContract.methods.reserveInput(num).send({from: user})
    data = tx['events']['InputMask']['raw']['data']
    var indexes = []
    for (let i = 0; i < num; i++) {
        indexes.push(getInt(data, i + 2))
    }
    return indexes
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

function interpolate(t, r, p) {
    let f0 = BigInt(0)
    for (let i = 0; i < t; i++) {
        let f = BigInt(1)
        for (let j = 0; j < t; j++) {
            if (i !== j) {
                f *= mod_reduce((BigInt(0) - BigInt(j + 1)) * modular_inverse(BigInt(i - j), p), p)
            }
        }
        f0 += mod_reduce(r[i] * f, p)
    }
    return mod_reduce(f0, p)
}

function reconstruct(n, t, shares, p) {
    if (shares.length !== n) {
        return false
    }
    value = interpolate(t + 1, shares, p)
    for (let i = t + 2; i < n + 1; i++) {
        _value = interpolate(i, shares, p)
        if (_value != value) {
            return NaN
        }
    }
    return value
}

// **** Fetch secret value from mpc servers ****

async function getSecretBalance(token, user, prefix='') {
    let shares = []
    for (let i = 0; i < n; i++) {
        url = host + ':' + (basePort + i) + '/balance/' + token + ',' + user
        console.log(url)
        const share = (await (await fetch(url, {mode: 'cors'})).json()).balance
        $('#' + prefix + i).text(share)
        shares.push(BigInt(share))
    }
    return fixToFloat(reconstruct(n, t, shares, p))
}

async function getTradePrice(tradeSeq) {
    let shares = []
    for (let i = 0; i < n; i++) {
        url = host + ':' + (basePort + i) + '/price/' + tradeSeq
        console.log(url)
        let share = (await (await fetch(url, {mode: 'cors'})).json()).price
        if (share == '') {
            return ''
        }
        $('#tradePrice' + i).text(share)
        shares.push(BigInt(share))
    }
    return fixToFloat(reconstruct(n, t, shares, p))
}

async function getInputmasks(num, idxes) {
    //Fetch inputmask shares from servers
    let shares = []
    for (let i = 0; i < num; i++) shares.push([])
    for (let srv = 0; srv < n; srv++) {
        url = host + ':' + (basePort + srv) + '/inputmasks/' + idxes
        console.log(url)
        const tmp = (await (await fetch(url, {mode: 'cors'})).json()).inputmask_shares.split(',')
        for (let i = 0; i < num; i++) {
            shares[i].push(BigInt(tmp[i]))
        }
    }

    //Reconstruct inputmasks
    let masks = []
    for (let i = 0; i < num; i++) {
        masks.push(reconstruct(n, t, shares[i], p))
    }

    return masks
}

async function getServerLog(srv, lines) {
    url = host + ':' + (basePort + srv) + '/log/' + lines
    const log = (await (await fetch(url, {mode: 'cors'})).json()).log
    $('#log').text(log)
}

// **** Access values on blockchain ****

async function getPersonalBalance(token, user) {
    let personalBalance
    if (isETH(token)) {
        personalBalance =  await web3.eth.getBalance(user)
    } else {
        personalBalance = await (contractList.get(token)).methods.balanceOf(user).call()
    }
    return parseFloat(personalBalance) / decimals
}

async function getPublicBalance(token, user) {
    return fixToFloat(await hbswapContract.methods.publicBalance(token, user).call())
}

// **** Global functions ****

async function trade() {
    if (!updateTradePair()) {
        return
    }

    const tokenFrom = tokenList.get($( '#tradeFromToken option:selected' ).text())
    const tokenTo = tokenList.get($( '#tradeToToken option:selected' ).text())

    // Check trading amount is valid
    const amtTradeFrom = parseFloat($( '#amtTradeFrom' ).val())
    const amtTradeTo = parseFloat($( '#amtTradeTo' ).val())

    const fee = amtTradeFrom * feeRate
    const balanceFrom = await getSecretBalance(tokenFrom, user)
    if (isNaN(amtTradeFrom) || amtTradeFrom <= 0 || amtTradeFrom + fee > balanceFrom) {
        $('#tradeInfo').text("Error: invalid amount for tokenFrom!")
        $('#tradeInfo').show()
        return
    }
    if (isNaN(amtTradeTo) || amtTradeTo <= 0) {
        $('#tradeInfo').text("Error: invalid amount for tokenTo!")
        $('#tradeInfo').show()
        return
    }

    $('#tradeStatus').text('Getting inputmasks...')
    $('#tradeStatusNeutral').show()
    const idxes = await getInputMaskIndexes(2)
    $('#tradeIdxes').text(idxes)

    const masks = await getInputmasks(2, idxes)
    $('#tradeMasks').text(masks)

    // Step 4: Publish masked inputs
    const slippage = $( '#slippage' ).val()
    const minReceived = amtTradeTo * (1 - slippage)

    let tokenA, tokenB, valueA, valueB
    if (tokenFrom < tokenTo) {
        tokenA = tokenFrom
        tokenB = tokenTo
        valueA = -floatToFix(amtTradeFrom)
        valueB = floatToFix(minReceived)
    } else {
        tokenA = tokenTo
        tokenB = tokenFrom
        valueA = floatToFix(minReceived)
        valueB = -floatToFix(amtTradeFrom)
    }
    const maskedValueA = valueA + masks[0]
    const maskedValueB = valueB + masks[1]
    $('#tradeStatus').text('Submitting trade order...')
    tx = await hbswapContract.methods.trade(tokenA, tokenB, idxes[0], idxes[1], maskedValueA, maskedValueB).send({from: user})

    // Step 5: Get price of current trade
    data = tx['events']['Trade']['raw']['data']
    const tradeSeq = getInt(data, 0)
    $('#seq').text(tradeSeq)
    while (true) {
        const price = await getTradePrice(tradeSeq)
        if (!(typeof price === 'string')) {
            $('#tradeStatusNeutral').hide()
            $('#price').text(price.toFixed(displayPrecision))
            if (price == 0) {
                $('#tradeStatus').text('Trade failed')
                $('#tradeStatusFail').show()
            } else {
                $('#tradeStatus').text('Trade succeed')
                $('#tradeStatusSucceed').show()
            }
            break
        }
        await sleep(5000)
    }

    $('#balanceTradeFrom').text((await getSecretBalance(tokenFrom, user)).toFixed(displayPrecision))
    $('#balanceTradeTo').text((await getSecretBalance(tokenTo, user)).toFixed(displayPrecision))
}

async function deposit() {
    await updateDepositToken()

    const token = tokenList.get($( '#depositToken option:selected' ).text())
    const amt = parseFloat($('#depositAmt').val()) // float

    // Check amt is valid
    const prevPersonalBalance = await getPersonalBalance(token, user)
    if (isNaN(amt) || amt <= 0 || amt > prevPersonalBalance) {
        $('#depositInfo').text("Error: invalid deposit amount!")
        $('#depositInfo').show()
        return
    }

    // Display balances before public deposit
    const prevPublicBalance = await getPublicBalance(token, user)
    const prevSecretBalance = await getSecretBalance(token, user, 'depositUpdate')
    $('#depositBalance').text(prevSecretBalance.toFixed(displayPrecision))
    $('#personalBalance').text(prevPersonalBalance.toFixed(displayPrecision))
    $('#contractBalance').text(prevPublicBalance.toFixed(displayPrecision))
    $('#secretBalance').text(prevSecretBalance.toFixed(displayPrecision))

    // Public deposit
    const fixAmt = floatToFix(amt)
    const transferAmt = transferValue(amt)
    if (isETH(token)) {
        $('#depositStatus').text('public depositing...')
        await hbswapContract.methods.publicDeposit(token, fixAmt).send({from: user, value: transferAmt})
    } else {
        // Approve before token transfer
        $('#depositStatus').text('approving tokens...')
        await contractList.get(token).methods.approve(hbswapAddr, transferAmt).send({from: user})

        $('#depositStatus').text('public depositing...')
        await hbswapContract.methods.publicDeposit(token, fixAmt).send({from: user})
    }

    // Display balances after public deposit
    $('#depositStatus').text('secret depositing...')
    $('#personalBalance').text((await getPersonalBalance(token, user)).toFixed(displayPrecision))
    $('#contractBalance').text((await getPublicBalance(token, user)).toFixed(displayPrecision))

    // Secret deposit
    await hbswapContract.methods.secretDeposit(token, fixAmt).send({from: user})

    let curSecretBalance
    while (true) {
        curSecretBalance = await getSecretBalance(token, user, 'depositUpdate')
        if (!isNaN(curSecretBalance) && prevSecretBalance != curSecretBalance) {
            break
        }
        await sleep(5000)
    }

    // Display balances after secret deposit
    $('#depositBalance').text(curSecretBalance.toFixed(displayPrecision))
    $('#depositStatus').text('done')
    $('#contractBalance').text((await getPublicBalance(token, user)).toFixed(displayPrecision))
    $('#secretBalance').text(curSecretBalance.toFixed(displayPrecision))
}

async function withdraw() {
    await updateDepositToken()

    const token = tokenList.get($( '#depositToken option:selected' ).text())
    const amt = parseFloat($('#depositAmt').val()) // float

    // Check amount is valid
    let prevSecretBalance = await getSecretBalance(token, user)
    if (isNaN(amt) || amt <= 0 || amt > prevSecretBalance) {
        $('#withdrawInfo').text("Error: invalid withdraw amount!")
        $('#withdrawInfo').show()
        return
    }

    // Display balances before secret withdraw
    const prevPersonalBalance = await getPersonalBalance(token, user)
    const prevContractBalance = await getPublicBalance(token, user)
    prevSecretBalance = await getSecretBalance(token, user, 'depositUpdate')
    $('#depositBalance').text(prevSecretBalance.toFixed(displayPrecision))
    $('#depositStatus').text('secret withdrawing...')
    $('#personalBalance').text(prevPersonalBalance.toFixed(displayPrecision))
    $('#contractBalance').text(prevContractBalance.toFixed(displayPrecision))
    $('#secretBalance').text(prevSecretBalance.toFixed(displayPrecision))

    // Secret withdraw
    const fixAmt = floatToFix(amt)
    await hbswapContract.methods.secretWithdraw(token, fixAmt).send({from: user})
    // Wait for secret balance change
    let curSecretBalance
    while (true) {
        curSecretBalance = await getSecretBalance(token, user, 'depositUpdate')
        if (!isNaN(curSecretBalance) && prevSecretBalance != curSecretBalance) {
            break
        }
        await sleep(5000)
    }
    $('#depositBalance').text(curSecretBalance.toFixed(displayPrecision))
    $('#secretBalance').text(curSecretBalance.toFixed(displayPrecision))
    // Wait for tx confirmation
    let curContractBalance
    while (true) {
        curContractBalance = await getPublicBalance(token, user)
        if (prevContractBalance != curContractBalance) {
            break
        }
        await sleep(5000)
    }
    // Display balances after secret withdraw
    $('#depositStatus').text('public withdrawing...')
    $('#contractBalance').text(curContractBalance.toFixed(displayPrecision))

    // Public withdraw
    await hbswapContract.methods.publicWithdraw(token, fixAmt).send({from: user})

    // Display balances after public withdraw
    $('#depositStatus').text('done')
    $('#personalBalance').text((await getPersonalBalance(token, user)).toFixed(displayPrecision))
    $('#contractBalance').text((await getPublicBalance(token, user)).toFixed(displayPrecision))
}

async function initPool() {
    if (!await updatePoolPair()) return

    const tokenA = tokenList.get($( '#poolTokenA option:selected' ).text())
    const tokenB = tokenList.get($( '#poolTokenB option:selected' ).text())
    const amtA = parseFloat($('#amtPoolTokenA').val())
    const amtB = parseFloat($('#amtPoolTokenB').val())

    const price = await hbswapContract.methods.prices(tokenA, tokenB).call()
    if (price != '') {
        $('#addInfo').text("Error: pool already initiated!")
        $('#addInfo').show()
        return
    }

    const balanceA = await getSecretBalance(tokenA, user)
    if (isNaN(amtA) || amtA <= 0 || amtA > balanceA) {
        $('#addInfo').text("Error: invalid amount for tokenA!")
        $('#addInfo').show()
        return
    }

    const balanceB = await getSecretBalance(tokenB, user)
    if (isNaN(amtB) || amtB <= 0 || amtB > balanceB) {
        $('#addInfo').text("Error: invalid amount for tokenB!")
        $('#addInfo').show()
        return
    }

    $('#poolStatus').text("Sending transaction...")

    await hbswapContract.methods.initPool(tokenA, tokenB, floatToFix(amtA), floatToFix(amtB)).send({from: user})

    while (true) {
        const price = await hbswapContract.methods.prices(tokenA, tokenB).call()
        if (price != '') {
            $('#estPricePool').text(price)
            break
        }
        await sleep(5000)
    }

    updatePoolPair()
    $('#poolStatus').text("Done")
}

async function addLiquidity() {
    if (!await updatePoolPair()) return

    const tokenA = tokenList.get($( '#poolTokenA option:selected' ).text())
    const tokenB = tokenList.get($( '#poolTokenB option:selected' ).text())
    const amtA = parseFloat($('#amtPoolTokenA').val())
    const amtB = parseFloat($('#amtPoolTokenB').val())

    const price = await hbswapContract.methods.prices(tokenA, tokenB).call()
    if (price == '') {
        $('#addInfo').text("Error: pool not initiated!")
        $('#addInfo').show()
        return
    }

    const balanceA = await getSecretBalance(tokenA, user)
    if (isNaN(amtA) || amtA <= 0 || amtA > balanceA) {
        $('#addInfo').text("Error: invalid amount for tokenA!")
        $('#addInfo').show()
        return
    }

    const balanceB = await getSecretBalance(tokenB, user)
    if (isNaN(amtB) || amtB <= 0 || amtB > balanceB) {
        $('#addInfo').text("Error: invalid amount for tokenB!")
        $('#addInfo').show()
        return
    }

    $('#poolStatus').text("Getting input masks...")

    const prevSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)

    const idxes = await getInputMaskIndexes(2)
    $('#poolIdxes').text(idxes)

    const masks = await getInputmasks(2, idxes)
    $('#poolMasks').text(masks)

    $('#poolStatus').text("Sending transaction...")
    const maskedAmtA = floatToFix(amtA) + masks[0]
    const maskedAmtB = floatToFix(amtB) + masks[1]
    await hbswapContract.methods.addLiquidity(tokenA, tokenB, idxes[0], idxes[1], maskedAmtA, maskedAmtB).send({from: user})

    let curSecretLiquidityTokenBalance
    while (true) {
        curSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)
        if (!isNaN(curSecretLiquidityTokenBalance) && prevSecretLiquidityTokenBalance != curSecretLiquidityTokenBalance) {
            break
        }
        await sleep(5000)
    }

    updatePoolPair()
    $('#poolStatus').text("Done")
}

async function removeLiquidity() {
    if (!await updatePoolPair()) return

    const tokenA = tokenList.get($( '#poolTokenA option:selected' ).text())
    const tokenB = tokenList.get($( '#poolTokenB option:selected' ).text())
    const amt = parseFloat($('#amtPoolLiquidityToken').val())

    const price = await hbswapContract.methods.prices(tokenA, tokenB).call()
    if (price == '') {
        $('#removeInfo').text("Error: pool not initiated!")
        $('#removeInfo').show()
        return
    }

    const prevSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)
    if (isNaN(amt) || amt <= 0 || amt > prevSecretLiquidityTokenBalance) {
        $('#removeInfo').text("Error: invalid amount for liquidity token!")
        $('#removeInfo').show()
        return
    }

    $('#poolStatus').text("Getting input masks...")

    const idxes = await getInputMaskIndexes(1)
    $('#poolIdxes').text(idxes)

    const masks = await getInputmasks(1, idxes)
    $('#poolMasks').text(masks)

    $('#poolStatus').text("Sending transaction...")

    const maskedAmt = floatToFix(amt) + masks[0]
    await hbswapContract.methods.removeLiquidity(tokenA, tokenB, idxes[0], maskedAmt).send({from: user})

    let curSecretLiquidityTokenBalance
    while (true) {
        curSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user)
        if (!isNaN(curSecretLiquidityTokenBalance) && prevSecretLiquidityTokenBalance != curSecretLiquidityTokenBalance) {
            break
        }
        await sleep(5000)
    }

    updatePoolPair()
    $('#poolStatus').text("Done")
}

async function updateTradePair() {
    $('#balanceTradeFrom').empty()
    $('#balanceTradeTo').empty()
    $('#estPriceTrade').empty()
    $('#minRecv').empty()
    $('#fee').empty()
    $('#tradeInfo').empty()
    $('#tradeInfo').hide()
    $('#tradeStatus').empty()
    $('#tradeIdxes').empty()
    $('#tradeMasks').empty()
    $('#seq').empty()
    $('#price').empty()
    $('#tradePrice0').empty()
    $('#tradePrice1').empty()
    $('#tradePrice2').empty()
    $('#tradePrice3').empty()
    $('#tradeStatusNeutral').hide()
    $('#tradeStatusFail').hide()
    $('#tradeStatusSucceed').hide()

    const fromToken = tokenList.get($( '#tradeFromToken option:selected' ).text())
    const toToken = tokenList.get($( '#tradeToToken option:selected' ).text())

    if (fromToken == toToken) {
        $('#tradeInfo').text('Error: invalid token pair!')
        $('#tradeInfo').show()
        return false
    }

    $('#balanceTradeFrom').text((await getSecretBalance(fromToken, user)).toFixed(displayPrecision))
    $('#balanceTradeTo').text((await getSecretBalance(toToken, user)).toFixed(displayPrecision))

    let price
    if (fromToken < toToken) {
        price = parseFloat(await hbswapContract.methods.prices(fromToken, toToken).call())
    } else {
        price = parseFloat(await hbswapContract.methods.prices(toToken, fromToken).call())
    }
    if (price == '') {
        $('#tradeInfo').text('Error: pool not initiated!')
        $('#tradeInfo').show()
        return false
    }
    $('#estPriceTrade').text(price.toFixed(displayPrecision))

    updateAmtTradeFrom()

    return true
}

async function updateAmtTradeFrom() {
    const _fromToken = $( '#tradeFromToken option:selected').text()
    const fromToken = tokenList.get(_fromToken)
    const toToken = tokenList.get($( '#tradeToToken option:selected' ).text())

    if (fromToken == toToken) {
        $('#tradeInfo').text('Error: invalid token pair!')
        $('#tradeInfo').show()
        return
    }

    const amtTradeFrom = $( '#amtTradeFrom' ).val()
    const slippage = $( '#slippage option:selected' ).val()

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

    $('#amtTradeTo').val(amtTradeTo.toFixed(displayPrecision))
    $('#minRecv').text(minReceived.toFixed(displayPrecision))
    $('#fee').text(fee.toFixed(displayPrecision) + ' ' + _fromToken)

}

async function updateAmtTradeTo() {
    const _fromToken = $( '#tradeFromToken option:selected').text()
    const fromToken = tokenList.get(_fromToken)
    const toToken = tokenList.get($( '#tradeToToken option:selected' ).text())

    if (fromToken == toToken) {
        $('#tradeInfo').text('Error: invalid token pair!')
        $('#tradeInfo').show()
        return
    }

    const amtTradeTo = $( '#amtTradeTo' ).val()
    const slippage = $( '#slippage option:selected' ).val()

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

    $('#amtTradeFrom').val(amtTradeFrom.toFixed(displayPrecision))
    $('#minRecv').text(minReceived.toFixed(displayPrecision))
    $('#fee').text(fee.toFixed(displayPrecision) + ' ' + _fromToken)
}

async function updatePoolPair() {
    // Get pool pair
    const tokenA = tokenList.get($( '#poolTokenA option:selected' ).text())
    const tokenB = tokenList.get($( '#poolTokenB option:selected' ).text())

    $('#balancePoolTokenA').empty()
    $('#balancePoolTokenB').empty()
    $('#estPricePool').empty()
    $('#addInfo').empty()
    $('#addInfo').hide()
    $('#removeInfo').empty()
    $('#removeInfo').hide()
    $('#balancePoolLiquidityToken').empty()
    $('#poolStatus').empty()
    $('#poolIdxes').empty()
    $('#poolMasks').empty()

    if (tokenA >= tokenB) {
        $('#addInfo').text('Error: invalid token pair!')
        $('#addInfo').show()
        $('#removeInfo').text('Error: invalid token pair!')
        $('#removeInfo').show()
        return false
    }

    // Update estimated price
    let price = await hbswapContract.methods.prices(tokenA, tokenB).call()
    if (price == '') {
        price = 'Pool not initiated!'
    }

    $('#balancePoolTokenA').text((await getSecretBalance(tokenA, user)).toFixed(displayPrecision))
    $('#balancePoolTokenB').text((await getSecretBalance(tokenB, user)).toFixed(displayPrecision))
    $('#estPricePool').text(price)
    $('#balancePoolLiquidityToken').text((await getSecretBalance(tokenA + '+' + tokenB, user)).toFixed(displayPrecision))

    return true
}

async function updateDepositToken() {
    const token = tokenList.get($( '#depositToken option:selected' ).text())
    $('#depositBalance').text((await getSecretBalance(token, user)).toFixed(displayPrecision))
    $('#depositInfo').empty()
    $('#depositInfo').hide()
    $('#withdrawInfo').empty()
    $('#withdrawInfo').hide()
    $('#depositStatus').empty()
    $('#personalBalance').empty()
    $('#contractBalance').empty()
    $('#secretBalance').empty()
    $('#depositUpdate0').empty()
    $('#depositUpdate1').empty()
    $('#depositUpdate2').empty()
    $('#depositUpdate3').empty()
}

// **** Initialization ****

async function init() {
    window.web3 = new Web3(ethereum)

    window.user = (await ethereum.request({ method: 'eth_requestAccounts'}))[0]
    $('#user').text(user)

    const hbswapABI = JSON.parse($('#hbswapABI').text())
    window.hbswapContract = new web3.eth.Contract(hbswapABI, hbswapAddr)

    const tokenABI = JSON.parse($('#tokenABI').text())
    window.tokenList = new Map()
    tokenList.set('ETH', ethAddr)
    tokenList.set('HBS', hbsAddr)
    tokenList.set('DAI', daiAddr)
    // tokenList.set('token1', token1)
    // tokenList.set('token2', token2)
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