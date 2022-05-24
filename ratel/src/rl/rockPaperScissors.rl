pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract rockPaperScissors {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public gameCnt;

    mapping (uint => uint) public status; // active-1, ready-2, completed-3

    mapping (uint => string) public winners;


    constructor() public {}


    function createGame($#uint value1) public {
        address player1 = msg.sender;
        uint gameId = ++gameCnt;

        mpc(uint gameId, address player1, $#uint value1) {
            mpcInput(sint value1)

            valid = ((value1.greater_equal(1, bit_length=bit_length)) * (value1.less_equal(3, bit_length=bit_length))).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                game = {
                    'player1': player1,
                    'value1': value1,
                }
                print('**** game', game)
                writeDB(f'gameBoard_{gameId}', game, dict)

                curStatus = 1
                set(status, uint curStatus, uint gameId)
        }
    }


    function joinGame(uint gameId, $uint value2) public {
        require(status[gameId] == 1);
        address player2 = msg.sender;

        mpc(uint gameId, address player2, $uint value2) {
            game = readDB(f'gameBoard_{gameId}', dict)

            mpcInput(sint value2)

            valid = ((value2.greater_equal(1, bit_length=bit_length)) * (value2.less_equal(3, bit_length=bit_length))).reveal()

            mpcOutput(cint valid)

            print('**** valid', valid)
            if valid == 1:
                game['player2'] = player2
                game['value2'] = value2

                print('**** game', game)

                writeDB(f'gameBoard_{gameId}', game, dict)

                curStatus = 2
                set(status, uint curStatus, uint gameId)
        }
    }


    function startRecon(uint gameId) public { // 1 < 2; 2 < 3; 3 < 1;
        require(status[gameId] == 2);
        status[gameId]++;

        mpc(uint gameId) {
            game = readDB(f'gameBoard_{gameId}', dict)

            value1 = game['value1']
            value2 = game['value2']

            mpcInput(sint value1, sint value2)
            print_ln('**** value1 %s', value1.reveal())
            print_ln('**** value2 %s', value2.reveal())

            result = (value1 - value2).reveal()

            print_ln('**** result %s', result)
            mpcOutput(cint result)

            if result > 2:
                result -= prime
            print('****', result)
            if result == 0:
                print('**** tie')
                winner = 'tie'
            elif result == 1 or result == -2:
                print('**** winner-player1')
                winner = 'player1'
            else:
                print('**** winner-player2')
                winner = 'player2'

            set(winners, string memory winner, uint gameId)
        }
    }
}
