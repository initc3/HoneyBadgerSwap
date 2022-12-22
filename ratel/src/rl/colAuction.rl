pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract colAuction{
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    address constant public ETH_addr = 0x0000000000000000000000000000000000000000;
    uint public colAuctionCnt;

    mapping (uint => uint) public biddersCnt;

    mapping (uint => uint) public curPriceList;
    mapping (uint => uint) public floorPriceList;
    mapping (uint => uint) public startPriceList;
    mapping (uint => uint) public totalAmtList;
    mapping (uint => address) public tokenAddrList;
    mapping (uint => address) public appAddrList;
    mapping (uint => address) public creatorAddrList;
    mapping (uint => uint) public debtList;
    
    mapping (uint => uint) public checkTime;

    mapping (uint => uint) public status; // closed-1 created-2 submitted --- bidders_num+2 
   

    constructor() public {}

    function createAuction(uint StartPrice, uint FloorPrice, uint totalAmt, uint debt, address token, address appAddr, address creator_addr) public{
        uint colAuctionId = ++colAuctionCnt;
        curPriceList[colAuctionId] = StartPrice;
        floorPriceList[colAuctionId] = FloorPrice;
        startPriceList[colAuctionId] = StartPrice;
        totalAmtList[colAuctionId] = totalAmt;

        checkTime[colAuctionId] = block.number;

        status[colAuctionId] = 2;

        tokenAddrList[colAuctionId] = token;
        appAddrList[colAuctionId] = appAddr;
        creatorAddrList[colAuctionId] = creator_addr;
        debtList[colAuctionId] = debt;
    }

    function scheduleCheck(uint colAuctionId) public {
        uint lastTime = checkTime[colAuctionId];
        uint curTime = block.number;
        require(lastTime + 10 < curTime);
        checkTime[colAuctionId] = block.number;

        uint curPrice = curPriceList[colAuctionId]*(20000-curTime+lastTime)/20000;
        curPriceList[colAuctionId] = curPrice;

        uint FloorPrice = floorPriceList[colAuctionId];
        uint totalAmt = totalAmtList[colAuctionId];

        address token_addr = tokenAddrList[colAuctionId];
        address appAddr = appAddrList[colAuctionId];
        address creatorAddr = creatorAddrList[colAuctionId];

        uint n = biddersCnt[colAuctionId];

        uint debt = debtList[colAuctionId];

        mpc(uint colAuctionId, uint n, uint curPrice, uint FloorPrice, uint totalAmt, uint debt, address token_addr, address appAddr, address creatorAddr, address ETH_addr){

            import time
            times = []
            times.append(time.perf_counter())

            cur_token_creator_balance = readDB(f'balanceBoard_{token_addr}_{creatorAddr}',int)
            cur_token_app_balance = readDB(f'balanceBoard_{token_addr}_{appAddr}',int)
            cur_eth_creator_balance = readDB(f'balanceBoard_{ETH_addr}_{creatorAddr}',int)

            import time
            add_benchmark_res_info = ''

            if curPrice < FloorPrice:
    
                times.append(time.perf_counter())

                for i in range(n):
                    pricei,Pi,Amti,recover_debti = await runCheckFail(server, token_addr, i, colAuctionId)
                    await runCheckFailUpdate(server, token_addr, i, colAuctionId,pricei,Pi,Amti,recover_debti)

                times.append(time.perf_counter())

                print(colAuctionId,'Auction failed!!!!!!!!!')

                add_benchmark_res_info = 'auctionFailed\t colAuctionId\t' + str(colAuctionId) +'\t'


            else:

                times.append(time.perf_counter())

                sum_amt = 0

                for i in range(n):
                    sum_amt = await runCheckAuction(server, i, colAuctionId,sum_amt,curPrice)

    
                times.append(time.perf_counter())

                mpcInput(sint sum_amt,sint curPrice, sint debt)
                
                cur_debt = sum_amt*curPrice
                v1 = (debt <= cur_debt)
                aucDone = v1.reveal()

                print_ln('cur_debt curPrice debt v1 %s %s %s %s',cur_debt.reveal(),curPrice.reveal(),debt.reveal(),v1.reveal())

                mpcOutput(cint aucDone)


                if aucDone == 1:
                    curAmt = totalAmt
                    app_token_amt = 0

                    times.append(time.perf_counter())

                    for i in range(n):
                        pricei, Pi, Amti,recover_debti = await runCheckSuccess(server, i, colAuctionId)
                        curAmt, app_token_amt = await runCheckSuccessUpdate(server, i, colAuctionId, token_addr, ETH_addr, curPrice, curAmt, app_token_amt,pricei,Pi,Amti,recover_debti)

                    cur_recover_debt = curPrice * totalAmt
                    cur_token_creator_balance = (cur_token_creator_balance + cur_recover_debt)%prime
                    cur_token_app_balance = (cur_token_app_balance - app_token_amt ) %prime

                    times.append(time.perf_counter())

                    add_benchmark_res_info = 'auctionSuccess\t colAuctionId\t' + str(colAuctionId) +'\t'

                    print(colAuctionId,'Auction success!!!!!!!!!')



            writeDB(f'balanceBoard_{ETH_addr}_{creatorAddr}',cur_eth_creator_balance,int)
            writeDB(f'balanceBoard_{token_addr}_{creatorAddr}',cur_token_creator_balance,int)
            writeDB(f'balanceBoard_{token_addr}_{appAddr}',cur_token_app_balance,int)

            if add_benchmark_res_info != '':
                cur_time = time.strftime("%D %H:%M:%S",time.localtime())
                with open(f'ratel/benchmark/data/latency.csv', 'a') as f:
                    f.write(f'{add_benchmark_res_info}\t'
                            f'cur_time\t{cur_time}\n')

                with open(f'ratel/benchmark/data/latency_{server.serverID}.csv', 'a') as f:
                    for op, t in enumerate(times):
                        f.write(f'auction end\t'
                                f'op\t{op + 1}\t'
                                f'cur_time\t{t}\n')


        }
    }

    pureMpc checkAuction(server, i, colAuctionId,sum_amt,curPrice) {
        bids = readDB(f'bidsBoard_{colAuctionId}_{i+1}', dict)

        Xi = bids['price']
        Pi = bids['address']
        Amti = bids['amt']

        mpcInput(sint Xi, sint curPrice, sint Amti, sint sum_amt)

        v1 = (curPrice <= Xi)
        new_sum_amt = sum_amt + v1*Amti

        mpcOutput(sint new_sum_amt)

        return new_sum_amt
    }

    pureMpc checkFail(server, token_addr, i, colAuctionId) {
        bids = readDB(f'bidsBoard_{colAuctionId}_{i+1}', dict)
    
        pricei = bids['price']
        Pi = bids['address']
        Amti = bids['amt']
        recover_debti = bids['recover_debt']

        return pricei,Pi,Amti,recover_debti
    } 

    pureMpc checkFailUpdate(server, token_addr, i, colAuctionId,pricei,Pi,Amti,recover_debti) {
        cur_token_balance = readDB(f'balanceBoard_{token_addr}_{Pi}',int)

        cur_token_balance = cur_token_balance + recover_debti

        writeDB(f'balanceBoard_{token_addr}_{Pi}',cur_token_balance,int)
    }

    pureMpc checkSuccess(server, i, colAuctionId) {
        bids = readDB(f'bidsBoard_{colAuctionId}_{i+1}', dict)

        pricei = bids['price']
        Pi = bids['address']
        Amti = bids['amt']
        recover_debti = bids['recover_debt']

        return pricei, Pi, Amti,recover_debti
    }

    pureMpc checkSuccessUpdate(server, i, colAuctionId, token_addr, ETH_addr, curPrice, curAmt, app_token_amt,pricei,Pi,Amti,recover_debti){
        
        cur_eth_balance = readDB(f'balanceBoard_{ETH_addr}_{Pi}',int)
        cur_token_balance = readDB(f'balanceBoard_{token_addr}_{Pi}',int)

        mpcInput(sint cur_eth_balance,sint cur_token_balance,sint pricei,sint curPrice,sint curAmt,sint Amti,sint recover_debti,sint app_token_amt)
        
        v1 = (curAmt >= Amti)

        realAmt1 = v1*Amti
        realAmt2 = (1-v1)*curAmt
        realAmt = realAmt1 + realAmt2

        print_ln('v1,realAmt %s %s',v1.reveal(),realAmt.reveal())

        cur_eth_balance = cur_eth_balance + realAmt

        origin_recover_debt = recover_debti
        actual_recover_debt = curPrice*realAmt

        print_ln('origin_recover_debt actual_recover_debt:%s %s',origin_recover_debt.reveal(),actual_recover_debt.reveal())

        cur_token_balance = cur_token_balance + origin_recover_debt - actual_recover_debt

        curAmt = curAmt - realAmt
        
        app_token_amt = app_token_amt + actual_recover_debt

        print_ln('curAmt,cur_eth_balance,cur_token_balance,app_token_amt:%s %s %s %s',curAmt.reveal(),cur_eth_balance.reveal(),cur_token_balance.reveal(),app_token_amt.reveal())
        
        mpcOutput(sint curAmt,sint cur_eth_balance,sint cur_token_balance,sint app_token_amt)

        writeDB(f'balanceBoard_{ETH_addr}_{Pi}',cur_eth_balance,int)
        writeDB(f'balanceBoard_{token_addr}_{Pi}',cur_token_balance,int)
    
        return curAmt, app_token_amt
    }

    function initClient(address token_addr, address user_addr) public{
        mpc(address user_addr,address token_addr){
            init_balance = 10000000000
            writeDB(f'balanceBoard_{token_addr}_{user_addr}',init_balance,int)
        }
    }

    function submitBids(uint colAuctionId, $uint price, $uint Amt,address bidder_addr) public {
        address P = bidder_addr;

        uint bidders_id = biddersCnt[colAuctionId]+1;
        biddersCnt[colAuctionId] = bidders_id;

        address token_addr = tokenAddrList[colAuctionId];
        address appAddr = appAddrList[colAuctionId];

        mpc(uint colAuctionId, uint bidders_id, $uint price, address P, $uint Amt, address token_addr, address appAddr){
            times = []

            import time
            times.append(time.perf_counter())
            start_time = time.strftime("%D %H:%M:%S",time.localtime())

            cur_token_balance = readDB(f'balanceBoard_{token_addr}_{P}',int)
            cur_app_balance = readDB(f'balanceBoard_{token_addr}_{appAddr}',int)

            times.append(time.perf_counter())

            mpcInput(sint cur_token_balance,sint cur_app_balance,sint price,sint Amt)

            recover_debt = price*Amt
            cur_token_balance = cur_token_balance - recover_debt
            cur_app_balance = cur_app_balance + recover_debt

            mpcOutput(sint cur_token_balance,sint cur_app_balance,sint recover_debt)

            times.append(time.perf_counter())

            bid = {
                'price': price,
                'amt': Amt,
                'address': P,
                'recover_debt':recover_debt,
            }

            writeDB(f'bidsBoard_{colAuctionId}_{bidders_id}',bid,dict)
            writeDB(f'balanceBoard_{token_addr}_{P}',cur_token_balance,int)
            writeDB(f'balanceBoard_{token_addr}_{appAddr}',cur_app_balance,int)

            times.append(time.perf_counter())

            with open(f'ratel/benchmark/data/latency.csv', 'a') as f:
                f.write(f'submit_bid\t'
                        f'colAuctionId\t{colAuctionId}\t'
                        f'bidders_id\t{bidders_id}\t'
                        f'start_time\t{start_time}\n')
        
            with open(f'ratel/benchmark/data/latency_{server.serverID}.csv', 'a') as f:
                for op, t in enumerate(times):
                    f.write(f'submit_bid\t'
                            f'seq\t{seqSubmitBids}\t'
                            f'op\t{op + 1}\t'
                            f'cur_time\t{t}\n')
        }
    }


}
