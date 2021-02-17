const p = BigInt("52435875175126190479447740508185965837690552500527637822603658699938581184513");
const n = 4;
const t = 1;
const fp = 1 << 16;
const eth = "0x0000000000000000000000000000000000000000";

endpoint = 'https://kovan.infura.io/v3/6a82d2519efb4d748c02552e02e369c1';
var web3 = new Web3();
web3.setProvider(new web3.providers.HttpProvider(endpoint));

const hbswapAddr = "0x77527DB365eC8dE2296D33464224E045bD7882C8";
const token1 = "0x63e7F20503256DdCFEC64872aAdb785d5A290CBb";
const token2 = "0x403B0F962566Ffb960d0dE98875dc09603Aa67e9";

const hbswapABI = JSON.parse($("#hbswapABI").text());
const hbswapContract = new web3.eth.Contract(hbswapABI, hbswapAddr);

const tokenABI = JSON.parse($("#tokenABI").text());
var tokenList = new Map();
tokenList.set('eth', eth);
tokenList.set('token1', token1);
tokenList.set('token2', token2);
var contractList = new Map();
for (let [k, v] of tokenList) {
    contractList.set(v, new web3.eth.Contract(tokenABI, v))
}

/* let account = web3.eth.accounts.decrypt({"address":"ef860fb0634474ae0e6cec1a8e0dbe7c70a280a5","crypto":{"cipher":"aes-128-ctr","ciphertext":"ef78844ca575e03eab239f3a97cf5f2393326b3190a103e8ecfa4c8f0fa087a8","cipherparams":{"iv":"a56aa4c42ffc07a13708b2fe61a98e3d"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"6bb662badf3affd44acf50aec280074ee487ad86f30967ef69e64582f2d9074f"},"mac":"1e684483d72e49e4dfa5b9b4caba7a04ba9091f6512bacbd30a0e6976424f807"},"id":"f47efed8-2874-4801-85af-4a90af7f9c2b","version":3}, ""); */
const privateKey = '0xa2452f41d937aa23f2d6be28e953293827f15a6277104c857fdd3373f766745a'
const user = web3.eth.accounts.privateKeyToAccount(privateKey).address;
$("#user").text(user);

// **** Internal functions ****

function isETH(token) {
    return token == tokenList.get('eth')
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function fromFloat(f) {
    return BigInt(Math.round(parseFloat(f) * fp));
}

function toFloat(i) {
    return parseFloat(i) / fp;
}

async function sendTx(tx) {
    const signedTx = await web3.eth.accounts.signTransaction(tx, privateKey);
    const receipt = await web3.eth.sendSignedTransaction(
        signedTx.rawTransaction
    );
    return receipt;
}

function getElement(st, idx) {
    return st.slice(2 + idx * 64, 2 + (idx + 1) * 64)
}

function getInt(st, idx) {
    return parseInt('0x' + getElement(st, idx))
}

function extended_gcd(a, b) {
    let s = BigInt(0);
    let old_s = BigInt(1);
    let t = BigInt(1);
    let old_t = BigInt(0);
    let r = b;
    let old_r = a;

    while (r !== BigInt(0)) {
        quotient = ~~(old_r / r);
        [old_r, r] = [r, old_r - quotient * r];
        [old_s, s] = [s, old_s - quotient * s];
        [old_t, t] = [t, old_t - quotient * t];
    }

    return [old_r, old_s, old_t];
}

function mod_reduce(x, p) {
    let r = x % p;
    return r >= 0 ? r : r + p;
}

function modular_inverse(x, p) {
    let gcd, s, t;
    [gcd, s, t] = extended_gcd(x, p);
    return gcd > 0 ? s : -s;
}

function interpolate(n, t, r, p) {
    if (r.length !== n) {
        return false
    }
    let f0 = BigInt(0);
    let f;
    for (let i = 0; i <= t; i++) {
        f = BigInt(1);
        for (let j = 0; j <= t; j++) {
            if (i !== j) {
                f *= mod_reduce((BigInt(0) - BigInt(j + 1)) * modular_inverse(BigInt(i - j), p), p);
            }
        }
        f0 += mod_reduce(r[i] * f, p);
    }
    return mod_reduce(f0, p);
}

// ********

// **** Fetch secret value from mpc servers ****

async function getSecretBalance(token, user) {
    let shares = [];
    for (let i = 0; i < 4; i++) {
        url = "http://localhost:" + (8080 + i) + "/balance/" + token + ',' + user;
        console.log(url);
        const share = (await (await fetch(url, {mode: 'cors'})).json()).balance;
        shares.push(BigInt(share));
    }
    return toFloat(interpolate(n, t, shares, p))
}

async function getTradePrice(tradeSeq) {
    let shares = [];
    for (let i = 0; i < 4; i++) {
        url = "http://localhost:" + (8080 + i) + "/price/" + tradeSeq;
        console.log(url);
        let share = (await (await fetch(url, {mode: 'cors'})).json()).price;
        shares.push(BigInt(share));
    }
    return toFloat(interpolate(n, t, shares, p))
}

async function getInputmasks(srv, idxes) {
    url = "http://localhost:" + (8080 + srv) + "/inputmasks/" + idxes;
    const shares = (await (await fetch(url, {mode: 'cors'})).json()).inputmask_shares.split(',');
    return [BigInt(shares[0]), BigInt(shares[1])];
}

// ********

// **** Access values on blockchain ****

async function getPersonalBalance(token, user) {
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
    $("#idxes").text(' ');
    $("#masks").text(' ');
    $("#seq").text(' ');
    $("#seq").text(' ');
    $("#price").text(' ');

    const tokenA = tokenList.get($( "#tradeTokenA option:selected" ).text());
    const tokenB = tokenList.get($( "#tradeTokenB option:selected" ).text());

    if (tokenA >= tokenB) {
        return
    }

    // Step 1: Claim inputmasks
    let tx = {
        to: hbswapAddr,
        gas: 210000,
        data: hbswapContract.methods.tradePrep().encodeABI()
    };
    let receipt = await sendTx(tx);
    data = receipt.logs[0].data;
    const idxA = getInt(data, 1);
    const idxB = getInt(data, 2);
    const idxes = idxA + ',' + idxB;
    $("#idxes").text(idxes);

    // Step 2: Fetch inputmasks from servers
    let maskAShares = [];
    let maskBShares = [];
    for (let i = 0; i < 4; i++) {
        const shares = await getInputmasks(i, idxes);
        maskAShares.push(shares[0]);
        maskBShares.push(shares[1]);
    }

    // Step 3: Reconstruct inputmasks
    const maskA = interpolate(n, t, maskAShares, p);
    const maskB = interpolate(n, t, maskBShares, p);
    $("#masks").text(maskA + ',' + maskB);

    // Step 4: Publish masked inputs
    const maskedValueA = fromFloat($("#amtTradeTokenA").val()) + maskA;
    const maskedValueB = fromFloat($("#amtTradeTokenB").val()) + maskB;
    tx = {
        to: hbswapAddr,
        gas: 210000,
        data: hbswapContract.methods.trade(tokenA, tokenB, idxA, idxB, maskedValueA, maskedValueB).encodeABI()
    };
    receipt = await sendTx(tx);

    // Step 5: Get price of current trade
    data = receipt.logs[0].data;
    const tradeSeq = getInt(data, 0);
    $("#seq").text(tradeSeq);
    const price = await getTradePrice(tradeSeq);
    $("#price").text(price);
}

async function deposit() {
    const token = tokenList.get($( "#depositToken option:selected" ).text());
    const amt = $("#deposit").val();

    // Display balances before withdraw
    const prevPersonalBalance = await getPersonalBalance(token, user);
    const prevContractBalance = await getContractBalance(token, user);
    const prevSecretBalance = await getSecretBalance(token, user);
    $("#balance").text(prevSecretBalance);
    $("#depositStatus").text('Pending');
    $("#prevPersonalBalance").text(prevPersonalBalance);
    $("#prevContractBalance").text(prevContractBalance);
    $("#prevSecretBalance").text(prevSecretBalance);
    $("#curPersonalBalance").text(' ');
    $("#curContractBalance").text(' ');
    $("#curSecretBalance").text(' ');

    // Public deposit
    let tx;
    if (isETH(token)) {
        // Transfer ethers to hbswap contract
        tx = {
            to: hbswapAddr,
            gas: 210000,
            value: amt, //web3.utils.toWei(amt, 'ether'),
            data: hbswapContract.methods.deposit(token, amt).encodeABI()
        }
    } else {
        // Approve before token transfer
        tx = {
            to: token,
            gas: 210000,
            data: contractList.get(token).methods.approve(hbswapAddr, amt).encodeABI()
        };
        await sendTx(tx);

        // Transfer tokens to hbswap contract
        tx = {
            to: hbswapAddr,
            gas: 210000,
            data: hbswapContract.methods.deposit(token, amt).encodeABI()
        }
    }
    await sendTx(tx);

    // Secret deposit
    tx = {
        to: hbswapAddr,
        gas: 210000,
        data: hbswapContract.methods.secretDeposit(token, amt).encodeABI()
    };
    await sendTx(tx);
    while (true) {
        if (prevSecretBalance < await getSecretBalance(token, user)) {
            break
        }
        await sleep(5000)
    }

    // Display balances after withdraw
    const curPersonalBalance = await getPersonalBalance(token, user);
    const curContractBalance = await getContractBalance(token, user);
    const curSecretBalance = await getSecretBalance(token, user);
    $("#balance").text(curSecretBalance);
    $("#depositStatus").text('Done');
    $("#curPersonalBalance").text(curPersonalBalance);
    $("#curContractBalance").text(curContractBalance);
    $("#curSecretBalance").text(curSecretBalance);
}

async function withdraw() {
    const token = tokenList.get($( "#depositToken option:selected" ).text());
    const amt = $("#deposit").val();

    // Display balances before withdraw
    const prevPersonalBalance = await getPersonalBalance(token, user);
    const prevContractBalance = await getContractBalance(token, user);
    const prevSecretBalance = await getSecretBalance(token, user);
    $("#balance").text(prevSecretBalance);
    $("#depositStatus").text('Pending');
    $("#prevPersonalBalance").text(prevPersonalBalance);
    $("#prevContractBalance").text(prevContractBalance);
    $("#prevSecretBalance").text(prevSecretBalance);
    $("#curPersonalBalance").text(' ');
    $("#curContractBalance").text(' ');
    $("#curSecretBalance").text(' ');

    // Secret withdraw
    let tx = {
        to: hbswapAddr,
        gas: 210000,
        data: hbswapContract.methods.secretWithdraw(token, amt).encodeABI()
    };
    await sendTx(tx);
    while (true) {
        if (prevContractBalance < await getContractBalance(token, user)) {
            break
        }
        await sleep(5000)
    }

    // Public withdraw
    tx = {
        to: hbswapAddr,
        gas: 210000,
        data: hbswapContract.methods.withdraw(token, amt).encodeABI()
    };
    await sendTx(tx);

    // Display balances after withdraw
    const curPersonalBalance = await getPersonalBalance(token, user);
    const curContractBalance = await getContractBalance(token, user);
    const curSecretBalance = await getSecretBalance(token, user);
    $("#balance").text(curSecretBalance);
    $("#depositStatus").text('Done');
    $("#curPersonalBalance").text(curPersonalBalance);
    $("#curContractBalance").text(curContractBalance);
    $("#curSecretBalance").text(curSecretBalance);
}

async function initPool() {
    const tokenA = tokenList.get($( "#poolTokenA option:selected" ).text());
    const tokenB = tokenList.get($( "#poolTokenB option:selected" ).text());
    const amtA = $("#amtPoolTokenA").val();
    const amtB = $("#amtPoolTokenB").val();
    console.log(tokenA);
    console.log(tokenB);
    console.log(amtA);
    console.log(amtB);

    tx = {
        to: hbswapAddr,
        gas: 210000,
        data: hbswapContract.methods.initPool(tokenA, tokenB, amtA, amtB).encodeABI()
    };
    await sendTx(tx)

    const prevSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user);
    let curSecretLiquidityTokenBalance;
    while (true) {
        curSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user);
        if (prevSecretLiquidityTokenBalance != curSecretLiquidityTokenBalance) {
            break
        }
        await sleep(5000)
    }

    updatePoolPair()
}

async function addLiquidity() {
    const tokenA = tokenList.get($( "#poolTokenA option:selected" ).text());
    const tokenB = tokenList.get($( "#poolTokenB option:selected" ).text());
    const amtA = $("#amtPoolTokenA").val();
    const amtB = $("#amtPoolTokenB").val();

    tx = {
        to: hbswapAddr,
        gas: 210000,
        data: hbswapContract.methods.addLiquidity(tokenA, tokenB, amtA, amtB).encodeABI()
    };
    await sendTx(tx)

    const prevSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user);
    let curSecretLiquidityTokenBalance;
    while (true) {
        curSecretLiquidityTokenBalance = await getSecretBalance(tokenA + '+' + tokenB, user);
        if (prevSecretLiquidityTokenBalance != curSecretLiquidityTokenBalance) {
            break
        }
        await sleep(5000)
    }

    updatePoolPair()
}

async function removeLiquidity() {
    const tokenA = tokenList.get($( "#poolTokenA option:selected" ).text());
    const tokenB = tokenList.get($( "#poolTokenB option:selected" ).text());
    const amt = $("#amtPoolLiquidityToken").val();

    tx = {
        to: hbswapAddr,
        gas: 210000,
        data: hbswapContract.methods.removeLiquidity(tokenA, tokenB, amt).encodeABI()
    };
    await sendTx(tx);

    updatePoolPair()
}

async function updateTradePair() {
    // Get trade pair
    const tokenA = tokenList.get($( "#tradeTokenA option:selected" ).text());
    const tokenB = tokenList.get($( "#tradeTokenB option:selected" ).text());

    if (tokenA >= tokenB) {
        $("#tradeInfo").text("Error: invalid token pair!");
        return
    }
    $("#tradeInfo").text(" ");

    // Update estimated price
    let price = await hbswapContract.methods.prices(tokenA, tokenB).call();
    if (price == '') {
        price = 'not available'
    }
    $("#estPriceTrade").text(price);
}

async function updatePoolPair() {
    // Get pool pair
    const tokenA = tokenList.get($( "#poolTokenA option:selected" ).text());
    const tokenB = tokenList.get($( "#poolTokenB option:selected" ).text());

    if (tokenA >= tokenB) {
        $("#poolInfo").text("Error: invalid token pair!");
        $("#estPricePool").text(' ');
        $("#balancePoolTokenA").text(' ');
        $("#balancePoolTokenB").text(' ');
        $("#balancePoolLiquidityToken").text(' ');
        return
    }

    // Update estimated price
    let price = await hbswapContract.methods.prices(tokenA, tokenB).call();
    if (price == '') {
        price = 'Pool not initiated!'
    }

    $("#poolInfo").text(" ");
    $("#estPricePool").text(price);
    $("#balancePoolTokenA").text(await getSecretBalance(tokenA, user));
    $("#balancePoolTokenB").text(await getSecretBalance(tokenB, user));
    $("#balancePoolLiquidityToken").text(await getSecretBalance(tokenA + '+' + tokenB, user));
}

async function updateDepositToken() {
    const token = tokenList.get($( "#depositToken option:selected" ).text());

    $("#balance").text(await getSecretBalance(token, user));
    $("#depositStatus").text(' ');
    $("#prevPersonalBalance").text(' ');
    $("#prevContractBalance").text(' ');
    $("#prevSecretBalance").text(' ');
    $("#curPersonalBalance").text(' ');
    $("#curContractBalance").text(' ');
    $("#curSecretBalance").text(' ');
}

// ********

async function init() {
    await updateTradePair();
    await updateDepositToken();
    await updatePoolPair();
}

init();