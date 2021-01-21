pragma solidity ^0.6.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract HbSwap {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    event InitPool(address user, address tokenA, address tokenB, uint amtA, uint amtB);
    event AddLiquidity(address user, address tokenA, address tokenB, uint amtA, uint amtB);
    event RemoveLiquidity(address user, address tokenA, address tokenB, uint amt);
    event SecretDeposit(address token, address user, uint amt);
    event TradePrep(address user, uint idxA, uint idxB);
    event Trade(uint tradeSeq, address user, address tokenA, address tokenB, uint idxA, uint idxB, uint maskedA, uint maskedB);
    event SecretWithdraw(uint seq, address token, address user, uint amt);

    struct SecretWithdrawIntention {
        address token;
        address user;
        uint amt;
    }

    uint public threshold;
    mapping (address => bool) public servers;

    uint public inputmaskCnt;
    uint public tradeCnt;
    mapping (address => mapping(address => uint)) public balances;

    uint public secretWithdrawCnt;
    mapping (uint => SecretWithdrawIntention) public secretWithdrawMap;
    mapping (uint => mapping (address => bool)) public consentRecord;
    mapping (uint => uint) consentCounter;
    mapping (uint => bool) expired;

    mapping (address => mapping (address => string)) public prices;
    mapping (address => mapping (address => uint)) public updateTimes;

    constructor(address[] memory _servers, uint _threshold) public {
        for (uint i = 0; i < _servers.length; i++) {
            servers[_servers[i]] = true;
        }
        threshold = _threshold;
    }

    function initPool(address _tokenA, address _tokenB, uint _amtA, uint _amtB) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        emit InitPool(msg.sender, _tokenA, _tokenB, _amtA, _amtB);
    }

    function addLiquidity(address _tokenA, address _tokenB, uint _amtA, uint _amtB) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        emit AddLiquidity(msg.sender, _tokenA, _tokenB, _amtA, _amtB);
    }

    function removeLiquidity(address _tokenA, address _tokenB, uint _amt) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        emit RemoveLiquidity(msg.sender, _tokenA, _tokenB, _amt);
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

    function tradePrep() public {
        emit TradePrep(msg.sender, inputmaskCnt, inputmaskCnt + 1);
        inputmaskCnt += 2;
    }

    function trade(address _tokenA, address _tokenB, uint _idxA, uint _idxB, uint _maskedA, uint _maskedB) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        tradeCnt += 1;
        emit Trade(tradeCnt, msg.sender, _tokenA, _tokenB, _idxA, _idxB, _maskedA, _maskedB);
    }

    function reset() public {
        inputmaskCnt = 0;
    }

    function updatePrice(address _tokenA, address _tokenB, string calldata _price) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        prices[_tokenA][_tokenB] = _price;
        updateTimes[_tokenA][_tokenB] = block.number;
    }

}
