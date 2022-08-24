import asyncio

from ratel.src.python.utils import prog, location_sharefile, int_to_hex, prime, sz


async def run(server_id, ):
    file = location_sharefile(server_id, port)
    with open(file, "wb") as f:
        f.write(
            int_to_hex(balanceA[server_id])
            + int_to_hex(amtA[server_id])
            + int_to_hex(balanceB[server_id])
            + int_to_hex(amtB[server_id])
            + int_to_hex(poolA[server_id])
            + int_to_hex(poolB[server_id])
            + int_to_hex(totalCnt[server_id])
            + int_to_hex(seqTrade)
        )

    cmd = f'{prog} -N {players} -T {threshold} -p {server_id} -pn {port} -P {prime} hbswapTrade1'
    proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)
    stdout, stderr = await proc.communicate()
    print(f'[{cmd!r} exited with {proc.returncode}]')
    if stdout:
        print(f'[stdout]\n{stdout.decode()}')
    if stderr:
        print(f'[stderr]\n{stderr.decode()}')

    input_arg_num = 8
    with open(file, "rb") as f:
        f.seek(input_arg_num * sz)


async def main():
    tasks = []
    for server_id in range(players):
        tasks.append(run(server_id))
    await asyncio.gather(*tasks)


if __name__ == '__main__':
    players = 3
    threshold = 1
    port = 5000

    balanceA = [
        4172190578519487839267498352441096556103883083624756335077661355530995201932,
        1107375579706713464561810141839198871350649807869605064153371772775946329137,
        5279566158226201303829308494280295427454532891494361399231033128306351707331
    ]

    amtA = [
        2327823015592411720226077229061973379010719845164411552452406163467949994723,
        4655646031184823440452154458123946758021439690328823104904812326935899956678,
        6983469046777235160678231687185920137032159535493234657357218490403849918633
    ]

    balanceB = [
        4647230090283695356433416699130758431335691550479986272813380108058633586587,
        2057454603235128498893646835218522621814266741580064939624809277831223098317,
        6704684693518823855327063534349281053149958292060051212438189385889266861036
    ]

    amtB = [
        5610134210538746484052348216162776701095058724886962827965992506882328386231,
        3983262843745230754131509869282559161333001090394018049930034075479202587009,
        2356391476951715024210671522402341621570943455901073271894075644076076787787
    ]

    poolA = [
        3064814998812774374705688210601897684753233275755151270924289582755114409057,
        6129629997625548749411376421203795369506466551510302541848579165510163281852,
        1957439419106060910143878068762698813402583467885546206770917809979757903658
    ]

    poolB = [
        2589775487048566857539769863912235809521424808899921333188570830227476024402,
        5179550974097133715079539727824471619042849617799842666377141660454886512672,
        532320883813438358646123028693713187707158067319856393563761552396842749953
    ]

    totalCnt = [
        394863230693209454757376375114956126393026772355578047483291218111669132352,
        789726461386418909514752750229912252786053544711156094966582436223338264702,
        1184589692079628364272129125344868379179080317066734142449873654335007397052
    ]

    seqTrade = 12

    asyncio.run(main())
