pragma solidity ^0.6.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract HbSwap {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    event TradePrep(address user, uint idxA, uint idxB);
    event Trade(address user, address tokenA, address tokenB, uint idxA, uint idxB, uint maskedA, uint maskedB);
    event SecretDeposit(address token, address user, uint amt);
    event SecretWithdraw(address token, address user, uint amt);

    uint public inputmaskCnt;
    mapping (address => mapping(address => uint)) public balances;

    constructor() public {}

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

        balances[_token][user] += _amt;

        emit SecretWithdraw(_token, user, _amt);
    }

    function tradePrep() public {
        emit TradePrep(msg.sender, inputmaskCnt, inputmaskCnt + 1);
        inputmaskCnt += 2;
    }

    function trade(address _tokenA, address _tokenB, uint _idxA, uint _idxB, uint _maskedA, uint _maskedB) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        emit Trade(msg.sender, _tokenA, _tokenB, _idxA, _idxB, _maskedA, _maskedB);
    }
}
