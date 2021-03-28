pragma solidity ^0.5.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract HbSwap {
    using SafeMath for uint;
    using SafeERC20 for IERC20;

    event InputMask(uint[] inpusMaskIndexes);
    event SecretDeposit(address token, address user, uint amt);
    event SecretWithdraw(uint seq, address token, address user, uint amt);
    event InitPool(address user, address tokenA, address tokenB, uint amtA, uint amtB);
    event AddLiquidity(address user, address tokenA, address tokenB, uint idxA, uint idxB, uint maskedAmtA, uint maskedAmtB);
    event RemoveLiquidity(address user, address tokenA, address tokenB, uint idx, uint maskedAmt);
    event Trade(uint tradeSeq, address user, address tokenA, address tokenB, uint idxA, uint idxB, uint maskedAmtA, uint maskedAmtB);

    struct SecretWithdrawIntention {
        address token;
        address user;
        uint amt;
    }

    // TODO: support for tokens with different decimals
    uint constant public Decimals = 1; //TODO: 10**18;
    uint constant public Fp = 2**16;

    uint public threshold;
    mapping (address => bool) public servers;

    mapping (address => mapping(address => uint)) public publicBalance;

    uint public secretWithdrawCnt;
    mapping (uint => SecretWithdrawIntention) public secretWithdrawMap;
    mapping (uint => mapping (address => bool)) public consentRecord;
    mapping (uint => uint) consentCounter;
    mapping (uint => bool) expired;

    uint public inputmaskCnt;
    mapping (uint => address) public inputMaskOwner;

    uint public tradeCnt;

    mapping (address => mapping (address => string)) public prices;
    mapping (address => mapping (address => uint)) public lastUpdateSeq;
    mapping (address => mapping (address => mapping (uint => mapping (string => uint)))) proposalCnt;
    mapping (address => mapping (address => mapping (uint => mapping (address => bool)))) propose;

    constructor(address[] memory _servers, uint _threshold) public {
        for (uint i = 0; i < _servers.length; i++) {
            servers[_servers[i]] = true;
        }
        threshold = _threshold;
    }

    function publicDeposit(address _token, uint _amt) payable public {
        address user = msg.sender;
        require(_amt > 0, "invalid amount");

        if (_token == address(0x0)) {
            require(msg.value * Fp == _amt * Decimals, "amounts not match" );
        } else {
            IERC20 token = IERC20(_token);
            token.safeTransferFrom(user, address(this), _amt / Fp * Decimals);
        }
        publicBalance[_token][user] += _amt;
    }

    function secretDeposit(address _token, uint _amt) public {
        address user = msg.sender;

        require(_amt > 0 && publicBalance[_token][user] >= _amt, "invalid amount");
        publicBalance[_token][user] -= _amt;

        emit SecretDeposit(_token, user, _amt);
    }

    function secretWithdraw(address _token, uint _amt) public {
        address user = msg.sender;
        require(_amt > 0, "invalid amount");

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
            publicBalance[secretWithdrawIntention.token][secretWithdrawIntention.user] += secretWithdrawIntention.amt;
            expired[_seq] = true;
        }
    }

    function publicWithdraw(address _token, uint _amt) public {
        address payable user = msg.sender;
        require(_amt > 0 && publicBalance[_token][user] >= _amt, "invalid amount");

        if (_token == address(0x0)) {
            user.transfer(_amt / Fp * Decimals);
        } else {
            IERC20 token = IERC20(_token);
            token.safeTransfer(user, _amt / Fp * Decimals);
        }
        publicBalance[_token][user] -= _amt;
    }

    function reserveInput(uint _num) public {
        address user = msg.sender;
        uint[] memory inputMaskIndexes = new uint[](_num);

        for (uint i = 0; i < _num; i++) {
            uint inputMaskIdx = inputmaskCnt++;
            inputMaskOwner[inputMaskIdx] = user;
            inputMaskIndexes[i] = inputMaskIdx;
        }

        emit InputMask(inputMaskIndexes);
    }

    function initPool(address _tokenA, address _tokenB, uint _amtA, uint _amtB) public {
        require(_tokenA < _tokenB, "invalid trading pair");
        require(_amtA > 0 && _amtB > 0, "invalid amount");

        address user = msg.sender;

        emit InitPool(msg.sender, _tokenA, _tokenB, _amtA, _amtB);
    }

    function addLiquidity(address _tokenA, address _tokenB, uint _idxA, uint _idxB, uint _maskedAmtA, uint _maskedAmtB) public {
        require(_tokenA < _tokenB, "invalid trading pair");

        address user = msg.sender;
        require(inputMaskOwner[_idxA] == user, "unauthorized inputmask");
        require(inputMaskOwner[_idxB] == user, "unauthorized inputmask");

        emit AddLiquidity(user, _tokenA, _tokenB, _idxA, _idxB, _maskedAmtA, _maskedAmtB);
    }

    function removeLiquidity(address _tokenA, address _tokenB, uint _idx, uint _maskedAmt) public {
        require(_tokenA < _tokenB, "invalid trading pair");

        address user = msg.sender;
        require(inputMaskOwner[_idx] == user, "unauthorized inputmask");

        emit RemoveLiquidity(user, _tokenA, _tokenB, _idx, _maskedAmt);
    }

    function trade(address _tokenA, address _tokenB, uint _idxA, uint _idxB, uint _maskedA, uint _maskedB) public {
        require(_tokenA < _tokenB, "invalid trading pair");

        address user = msg.sender;
        require(inputMaskOwner[_idxA] == user, "unauthorized inputmask");
        require(inputMaskOwner[_idxB] == user, "unauthorized inputmask");

        tradeCnt += 1;
        emit Trade(tradeCnt, user, _tokenA, _tokenB, _idxA, _idxB, _maskedA, _maskedB);
    }

    function updatePrice(address _tokenA, address _tokenB, uint _checkpointSeq, string memory _price) public {
        address server = msg.sender;
        require(servers[server], "not a valid server");

        if (!propose[_tokenA][_tokenB][_checkpointSeq][server] || _checkpointSeq == 0) {
            propose[_tokenA][_tokenB][_checkpointSeq][server] = true;
            proposalCnt[_tokenA][_tokenB][_checkpointSeq][_price] += 1;

            if (proposalCnt[_tokenA][_tokenB][_checkpointSeq][_price] > threshold && _checkpointSeq >= lastUpdateSeq[_tokenA][_tokenB]) {
                prices[_tokenA][_tokenB] = _price;
                lastUpdateSeq[_tokenA][_tokenB] = _checkpointSeq;
            }
        }
    }

    //TODO: for test only, remove before deployment
    function resetPrice(address _tokenA, address _tokenB, address[] memory _servers) public {
        prices[_tokenA][_tokenB] = "";
        lastUpdateSeq[_tokenA][_tokenB] = 0;
        for (uint i = 0; i < _servers.length; i++) {
            propose[_tokenA][_tokenB][0][_servers[i]] = false;
        }
        proposalCnt[_tokenA][_tokenB][0]["0.0"] = 0;
    }

    //TODO: for test only, remove before deployment
    function resetBalance(address _token, address _user) public {
        publicBalance[_token][_user] = 0;
    }
}
