pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract rockPaperScissors {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    uint public gameCnt;

    mapping (uint => uint) public status; // active-1, ready-2, completed-3
    mapping (address => uint) public valueStatus;
    mapping (uint => uint) public countStatus;

    mapping (uint => address) public winners;
    mapping (address => address) public valueWinners;
    mapping (address => uint) public countWinners;

    constructor() public {}

    function createGame($uint value1) public {
        address player1 = msg.sender;
        uint gameId = ++gameCnt;

        mpc(uint gameId, address player1, $uint value1) {
            print(value1, type(value1))

            value1 *= fp
            mpcInput(value1)
            value1 = sfix._new(value1)

            valid = (value1 >= 1) * (value1 <= 3)
            valid = sint(valid.reveal())

            mpcOutput(valid)

            print(valid)
            if valid == 1:
                game = {
                    'player1': player1,
                    'value1': value1 // fp,
                }
                print(game)
                game = str(game)
                print(game)
                game = bytes(game, encoding='utf-8')
                print(game)
                writeDB(f'gameBoard_{gameId}', game)

                curStatus = 1
                set(status, uint curStatus, uint gameId)
        }
    }

    function joinGame(uint gameId, $uint value2) public {
        require(status[gameId] == 1);

        address player2 = msg.sender;

        mpc(uint gameId, address player2, $uint value2) {
            value2 *= fp
            mpcInput(value2)
            value2 = sfix._new(value2)

            valid = (value2 >= 1) * (value2 <= 3)
            valid = sint(valid.reveal())

            mpcOutput(valid)

            print(valid)
            if valid == 1:

                game = readDB(f'gameBoard_{gameId}')
                game = game.decode(encoding='utf-8')
                import ast
                game = dict(ast.literal_eval(game))

                game['player2'] = player2
                game['value2'] = value2 // fp

                print(game)
                game = str(game)
                print(game)
                game = bytes(game, encoding='utf-8')
                print(game)
                writeDB(f'gameBoard_{gameId}', game)

                curStatus = 2
                set(status, uint curStatus, uint gameId)
        }
    }

    function startRecon(uint gameId) public {
        require(status[gameId] == 2);
        status[gameId]++;

        mpc(uint gameId) {
            game = readDB(f'gameBoard_{gameId}')
            game = game.decode(encoding='utf-8')
            game = dict(ast.literal_eval(game))

            value1 = game['value1'] * fp
            value2 = game['value2'] * fp

            mpcInput(value1, value2)
            value1 = sfix._new(value1)
            value2 = sfix._new(value2)

            print_ln('value1 %s', value1.reveal())
            print_ln('value2 %s', value2.reveal())

            result = (value1 - value2).reveal()
            print_ln('result %s', result)
            result = sint(result.v)
            print_ln('result %s', result.reveal())

            mpcOutput(result)
            if result != 0:
                result -= blsPrime
            print(result)
            result //= fp
            print(result)

            if result == 0:
                print('tie')
                winner = 'tie'
            elif result == 1 or result == -2:
                print('player1')
                winner = game['player1']
            else:
                print('player2')
                winner = game['player2']

            set(winners, address winner, uint gameId)
        }
    }
}
