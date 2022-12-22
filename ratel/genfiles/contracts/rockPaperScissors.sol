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


    uint public N;
    uint public T;
    mapping (address => bool) public isServer;

    uint public opCnt;
    mapping (uint => string) public opEvent;
    mapping (uint => bytes) public opContent;


    constructor(address[] memory servers, uint threshold) public {
        N = servers.length;
        T = threshold;
        require(T <= N / 2);
        for (uint i = 0; i < servers.length; i++) {
            isServer[servers[i]] = true;
        }
    }


    modifier onlyServer() {
        require(isServer[msg.sender], "not an authorized party");
        _;
    }


    event InputMask(uint[] inpusMaskIndexes);
    uint public inputMaskCnt;
    mapping (uint => address) public inputMaskOwner;
    function reserveInput(uint num) public {
        address user = msg.sender;
        uint[] memory inputMaskIndexes = new uint[](num);
        for (uint i = 0; i < num; i++) {
            uint inputMaskIdx = inputMaskCnt++;
            inputMaskOwner[inputMaskIdx] = user;
            inputMaskIndexes[i] = inputMaskIdx;
        }
        emit InputMask(inputMaskIndexes);
    }


    event Finalization(uint seq);
    mapping (uint => uint) public voteFinalization;
    mapping (uint => mapping (address => bool)) public votedFinalization;
    mapping (uint => bool) public finalized;
    function finalize(uint seq) public onlyServer {
        address server = msg.sender;
        require(!votedFinalization[seq][server] && !finalized[seq]);
        votedFinalization[seq][server] = true;
        voteFinalization[seq] += 1;
        if (voteFinalization[seq] > T) {
            finalized[seq] = true;
            emit Finalization(seq);
        }
    }


    event GenInputMask(uint inputMaskCnt, uint inputMaskVersion);
    uint public inputMaskVersion;
    function genInputMask(uint _inputMaskCnt) public {
        //TODO: agree on request
        if (_inputMaskCnt == 0) _inputMaskCnt = inputMaskCnt;
        emit GenInputMask(_inputMaskCnt, ++inputMaskVersion);
    }


    function createGame(uint idxValue1, uint maskedValue1, string memory zkpstmt0) public {
        require(inputMaskOwner[idxValue1] == msg.sender);

        address player1 = msg.sender;
        uint gameId = ++gameCnt;

        uint seqCreateGame = opCnt++;
        opEvent[seqCreateGame] = "CreateGame";
        opContent[seqCreateGame] = abi.encode(gameId, player1, idxValue1, maskedValue1, zkpstmt0);

        emit CreateGame(seqCreateGame, gameId, player1, idxValue1, maskedValue1, zkpstmt0);
    }


    function joinGame(uint gameId, uint idxValue2, uint maskedValue2, string memory zkpstmt0, string memory zkpstmt1) public {
        require(inputMaskOwner[idxValue2] == msg.sender);

        require(status[gameId] == 1);
        address player2 = msg.sender;

        uint seqJoinGame = opCnt++;
        opEvent[seqJoinGame] = "JoinGame";
        opContent[seqJoinGame] = abi.encode(gameId, player2, idxValue2, maskedValue2, zkpstmt0, zkpstmt1);

        emit JoinGame(seqJoinGame, gameId, player2, idxValue2, maskedValue2, zkpstmt0, zkpstmt1);
    }


    function startRecon(uint gameId) public { // 1 < 2; 2 < 3; 3 < 1;

        require(status[gameId] == 2);
        status[gameId]++;

        uint seqStartRecon = opCnt++;
        opEvent[seqStartRecon] = "StartRecon";
        opContent[seqStartRecon] = abi.encode(gameId);

        emit StartRecon(seqStartRecon, gameId);
    }

    event CreateGame(uint seqCreateGame, uint gameId, address player1, uint idxValue1, uint maskedValue1, string zkpstmt0);
    event JoinGame(uint seqJoinGame, uint gameId, address player2, uint idxValue2, uint maskedValue2, string zkpstmt0, string zkpstmt1);
    event StartRecon(uint seqStartRecon, uint gameId);


    mapping(uint => mapping(address => uint)) public statusValue;
    mapping(uint => mapping(uint => uint)) public statusCount;
    function statusSet(uint curStatus, uint gameId) public onlyServer {
        address server = msg.sender;
        if (statusValue[gameId][server] != 0) {
            statusCount[gameId][statusValue[gameId][server]]--;
        }
        statusValue[gameId][server] = curStatus;
        statusCount[gameId][statusValue[gameId][server]]++;
        if (statusCount[gameId][curStatus] > T) {
            status[gameId] = curStatus;
        }
    }


    mapping(uint => mapping(address => string)) public winnersValue;
    mapping(uint => mapping(string => uint)) public winnersCount;
    function winnersSet(string memory winner, uint gameId) public onlyServer {
        address server = msg.sender;
        winnersValue[gameId][server] = winner;
        winnersCount[gameId][winnersValue[gameId][server]]++;
        if (winnersCount[gameId][winner] > T) {
            winners[gameId] = winner;
        }
    }


}
