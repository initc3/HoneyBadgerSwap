pragma solidity ^0.6.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract HbSwap {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    event InitPool(address tokenA, address tokenB, uint amtA, uint amtB);
    event TradePrep(address user, uint idxA, uint idxB);
    event Trade(uint tradeSeq, address user, address tokenA, address tokenB, uint idxA, uint idxB, uint maskedA, uint maskedB);
    event SecretDeposit(address token, address user, uint amt);
    event SecretWithdraw(uint seq, address token, address user, uint amt);

    mapping (address => mapping (address => bool)) public tradingPairs;
    mapping (address => mapping (address => mapping (address => uint))) public liquidityToken;
    mapping (address => mapping (address => string)) public prices;
    mapping (address => mapping (address => uint)) public updateTimes;

    uint public threshold;
    uint public inputmaskCnt;
    mapping (address => mapping(address => uint)) public balances;

    uint public serverNum;
    mapping (address => bool) public servers;

    uint public tradeCnt;

    struct SecretWithdrawIntention {
        address token;
        address user;
        uint amt;
    }

    uint public secretWithdrawCnt;
    mapping (uint => SecretWithdrawIntention) public secretWithdrawMap;
    mapping (uint => mapping (address => bool)) public consentRecord;
    mapping (uint => uint) consentCounter;
    mapping (uint => bool) expired;

    constructor(address[] memory _servers, uint _threshold) public {
        for (uint i = 0; i < _servers.length; i++) {
            servers[_servers[i]] = true;
        }
        threshold = _threshold;
    }

    function initPool(address _tokenA, address _tokenB, uint _amtA, uint _amtB) payable public {
        require(_tokenA < _tokenB, "invalid trading pair");
        require(!tradingPairs[_tokenA][_tokenB], "pool already initiated");

        address user = msg.sender;

        if (_tokenA != address(0x0)) {
            IERC20 tokenA = IERC20(_tokenA);
            tokenA.safeTransferFrom(user, address(this), _amtA);
        }
        if (_tokenB != address(0x0)) {
            IERC20 tokenB = IERC20(_tokenB);
            tokenB.safeTransferFrom(user, address(this), _amtB);
        }

        tradingPairs[_tokenA][_tokenB] = true;
        liquidityToken[_tokenA][_tokenB][user] = _sqrt(_amtA * _amtB);

        emit InitPool(_tokenA, _tokenB, _amtA, _amtB);
    }

    function deposit(address _token, uint _amt) payable public {
        address user = msg.sender;

        if (_token == address(0x0)) {
            balances[_token][user] += msg.value;
        } else {
            IERC20 token = IERC20(_token);
            token.safeTransferFrom(user, address(this), _amt);
            balances[_token][user] += _amt;
        }
    }

    function withdraw(address _token, uint _amt) public {
        address payable user = msg.sender;

        if (_token == address(0x0)) {
            balances[_token][user] -= _amt;
            user.transfer(_amt);
        } else {
            IERC20 token = IERC20(_token);
            token.safeTransfer(user, _amt);
            balances[_token][user] -= _amt;
        }
    }

    function secretDeposit(address _token, uint _amt) public {
        address user = msg.sender;

        require(balances[_token][user] >= _amt, "not enough balance");
        balances[_token][user] -= _amt;

        emit SecretDeposit(_token, user, _amt);
    }

    function secretWithdraw(address _token, uint _amt) public {
        address user = msg.sender;

        secretWithdrawCnt += 1;
        secretWithdrawMap[secretWithdrawCnt] = SecretWithdrawIntention(_token, user, _amt);

        emit SecretWithdraw(secretWithdrawCnt, _token, user, _amt);
    }

    function consent(uint _seq) public {
        address server = msg.sender;

        require(servers[server], "not a valid server");
        require(!consentRecord[_seq][server], "already consent");

        consentRecord[_seq][server] = true;
        consentCounter[_seq] += 1;
        if (consentCounter[_seq] > threshold && !expired[_seq]) {
            SecretWithdrawIntention memory secretWithdrawIntention = secretWithdrawMap[_seq];
            balances[secretWithdrawIntention.token][secretWithdrawIntention.user] += secretWithdrawIntention.amt;
            expired[_seq] = true;
        }
    }

    function tradePrep() public {
        emit TradePrep(msg.sender, inputmaskCnt, inputmaskCnt + 1);
        inputmaskCnt += 2;
    }

    function trade(address _tokenA, address _tokenB, uint _idxA, uint _idxB, uint _maskedA, uint _maskedB) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        require(tradingPairs[_tokenA][_tokenB], "pool not exist");
        tradeCnt += 1;
        emit Trade(tradeCnt, msg.sender, _tokenA, _tokenB, _idxA, _idxB, _maskedA, _maskedB);
    }

    function reset() public {
        inputmaskCnt = 0;
    }

    function updatePrice(address _tokenA, address _tokenB, string calldata _price) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        require(tradingPairs[_tokenA][_tokenB], "pool not exist");
        prices[_tokenA][_tokenB] = _price;
        updateTimes[_tokenA][_tokenB] = block.number;
    }

    // babylonian method (https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Babylonian_method)
    function _sqrt(uint y) internal pure returns (uint z) {
        if (y > 3) {
            z = y;
            uint x = y / 2 + 1;
            while (x < z) {
                z = x;
                x = (y / x + x) / 2;
            }
        } else if (y != 0) {
            z = 1;
        }
    }
}
