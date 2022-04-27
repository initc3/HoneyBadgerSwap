import asyncio


def infinite_precision_calc(balanceA, amtA, balanceB, amtB, poolA, poolB):
    from decimal import Decimal

    feeRate = Decimal(0.003)

    poolProduct = poolA * poolB

    totalA = (1 + feeRate) * amtA
    totalB = (1 + feeRate) * amtB

    actualAmtA = poolA - poolProduct / (poolB - amtB)
    actualAmtB = poolB - poolProduct / (poolA - amtA)

    buyA = amtA > 0
    acceptA = actualAmtA >= amtA
    acceptB = actualAmtB >= amtB
    buyB = 1 - buyA

    flagBuyA = buyA * acceptA
    flagBuyB = buyB * acceptB

    changeA = flagBuyA * actualAmtA + flagBuyB * totalA
    changeB = flagBuyA * totalB + flagBuyB * actualAmtB

    poolA -= changeA
    poolB -= changeB
    balanceA += changeA
    balanceB += changeB

    # print(changeA, changeB)
    # print(balanceA, balanceB, poolA, poolB)

    return balanceA, balanceB, poolA, poolB


def infinite_precision(balanceA, balanceB, poolA, poolB, amtA, amtB, repetition):
    from decimal import Decimal

    balanceA = Decimal(balanceA)
    balanceB = Decimal(balanceB)
    poolA = Decimal(poolA)
    poolB = Decimal(poolB)
    amtA = Decimal(amtA)
    amtB = Decimal(amtB)

    for i in range(repetition):
        balanceA, balanceB, poolA, poolB = infinite_precision_calc(balanceA, amtA, balanceB, amtB, poolA, poolB)
        balanceA, balanceB, poolA, poolB = infinite_precision_calc(balanceA, amtB, balanceB, amtA, poolA, poolB)

    print(f'infinite_precision {balanceA, balanceB, poolA, poolB}')


async def mp_spdz_calc(server_id, balanceA, amtA, balanceB, amtB, poolA, poolB, repetition):
    from ratel.src.python.utils import location_sharefile, mpc_port, int_to_hex, sz, hex_to_int, prog, blsPrime

    players = 3
    threshold = 1
    port = mpc_port
    totalCnt = 0

    for i in range(repetition * 2):
        if server_id == 0:
            print(i)
        file = location_sharefile(server_id, port)
        with open(file, "wb") as f:
            f.write(
                int_to_hex(balanceA)
                + int_to_hex(amtA)
                + int_to_hex(balanceB)
                + int_to_hex(amtB)
                + int_to_hex(poolA)
                + int_to_hex(poolB)
                + int_to_hex(totalCnt)
            )

        cmd = f'{prog} -N {players} -T {threshold} -p {server_id} -pn {port} -P {blsPrime} hbswapTrade1'
        proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)
        stdout, stderr = await proc.communicate()
        if stdout:
            print(f'[stdout]\n{stdout.decode()}')

        input_arg_num = 7
        with open(file, "rb") as f:
            f.seek(input_arg_num * sz)
            balanceA = hex_to_int(f.read(sz))
            balanceB = hex_to_int(f.read(sz))
            poolA = hex_to_int(f.read(sz))
            poolB = hex_to_int(f.read(sz))

        await asyncio.sleep(0.1)

        amtA, amtB = amtB, amtA



async def mp_spdz_fixed_point(balanceA, balanceB, poolA, poolB, amtA, amtB, repetition):
    from ratel.src.python.utils import fp

    players = 3

    balanceA *= fp
    balanceB *= fp
    poolA *= fp
    poolB *= fp
    amtA *= fp
    amtB *= fp

    tasks = []
    for server_id in range(players):
        tasks.append(mp_spdz_calc(server_id, balanceA, amtA, balanceB, amtB, poolA, poolB, repetition))
    await asyncio.gather(*tasks)


if __name__ == '__main__':
    repetition = 1
    balanceA, balanceB = 9000, 9000
    poolA, poolB = 1000, 1000
    amtA = 0.5
    amtB = -1

    asyncio.run(mp_spdz_fixed_point(balanceA, balanceB, poolA, poolB, amtA, amtB, repetition))

    infinite_precision(balanceA, balanceB, poolA, poolB, amtA, amtB, repetition)
