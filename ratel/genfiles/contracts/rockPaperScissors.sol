pragma solidity ^0.5.0;
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
contract rockPaperScissors {
    using SafeMath for uint;
    using SafeERC20 for IERC20;
    uint public gameCnt;
    mapping (uint => uint) public status; // active-1, ready-2, completed-3
    mapping (address => uint) public statusValue;
    mapping (uint => uint) public statusCount;
    mapping (uint => string) public winners;
    mapping (address => string) public winnersValue;
    mapping (string => uint) public winnersCount;
    // TODO:
    event CreateGame(uint seqCreateGame, uint gameId, address player1, uint idxValue1, uint maskedValue1, uint idxBlinding, uint maskedBlinding, uint8[] proof, uint8[] commitment);
    event JoinGame(uint seqJoinGame, uint gameId, address player2, uint idxValue2, uint maskedValue2);
    event StartRecon(uint seqStartRecon, uint gameId);

    event InputMask(uint[] inpusMaskIndexes);
    uint public inputMaskCnt;
    mapping (uint => address) public inputMaskOwner;

    event NewServer(address server);
    uint public N;
    uint public T;
    mapping (address => bool) public isServer;
    mapping (address => uint) public votes;
    mapping (address => mapping (address => bool)) voted;

    uint public opCnt;
    mapping (uint => string) public opEvent;
    mapping (uint => bytes) public opContent;

    event GenInputMask(uint inputMaskCnt, uint committeeChangeCnt);
    uint public isInputMaskReady;
    uint public committeeChangeCnt;
    mapping (address => uint) public numCommittee;

    mapping (uint => uint) public voteFinalization;
    mapping (uint => mapping (address => bool)) public votedFinalization;
    mapping (uint => uint) public finalizedTime;
    constructor(address[] memory servers, uint threshold) public {
        N = servers.length;
        isInputMaskReady = N;
        require(T <= N / 2);
        T = threshold;
        for (uint i = 0; i < servers.length; i++) {
            isServer[servers[i]] = true;
            votes[servers[i]] = N;
            for (uint j = 0; j < servers.length; j++) {
                voted[servers[i]][servers[j]] = true;
            }
        }
    }

    function registerServer() public {
        address server = msg.sender;
        require(!isServer[server]);
        emit NewServer(server);
    }

    function addServer(address s) public {
        address server = msg.sender;
        require(isServer[server]);
        require(!voted[server][s]);
        voted[server][s] = true;
        votes[s] += 1;
        if (votes[s] > T && !isServer[s]) {
            isServer[s] = true;
            N += 1;
            isInputMaskReady = 0;
            committeeChangeCnt += 1;
            emit GenInputMask(inputMaskCnt, committeeChangeCnt);
        }
    }

    function setReady(uint num) public {
        address server = msg.sender;
        require(isServer[server]);
        if (numCommittee[server] < num) {
            numCommittee[server] = num;
            isInputMaskReady += 1;
        }
    }

    function resetThreshold() public {
    }

    function removeServer(address s) public {
        address server = msg.sender;
        require(isServer[server]);
        require(voted[server][s]);
        voted[server][s] = false;
        votes[s] -= 1;
        if (votes[s] <= T) {
            isServer[s] = false;
            N -= 1;
        }
    }

    function reserveInput(uint num) public {
        require(isInputMaskReady > T);
        address user = msg.sender;
        uint[] memory inputMaskIndexes = new uint[](num);
        for (uint i = 0; i < num; i++) {
            uint inputMaskIdx = inputMaskCnt++;
            inputMaskOwner[inputMaskIdx] = user;
           inputMaskIndexes[i] = inputMaskIdx;
        }
        emit InputMask(inputMaskIndexes);
    }

    function finalize(uint seq) public {
        address server = msg.sender;
        require(isServer[server] && !votedFinalization[seq][server]);
        votedFinalization[seq][server] = true;
        voteFinalization[seq] += 1;
        if (voteFinalization[seq] > T && finalizedTime[seq] == 0) {
            finalizedTime[seq] = block.number;
        }
    }

    function winnersSet(string memory winner, uint gameId) public {
        address server = msg.sender;
        require(isServer[server]);
        if (bytes(winnersValue[server]).length > 0) {
            winnersCount[winnersValue[server]]--;
        }
        winnersValue[server] = winner;
        winnersCount[winnersValue[server]]++;
        if (winnersCount[winner] > T) {
            winners[gameId] = winner;
        }
    }

    function statusSet(uint curStatus, uint gameId) public {
        address server = msg.sender;
        require(isServer[server]);
        if (statusValue[server] != 0) {
            statusCount[statusValue[server]]--;
        }
        statusValue[server] = curStatus;
        statusCount[statusValue[server]]++;
        if (statusCount[curStatus] > T) {
            status[gameId] = curStatus;
        }
    }

    // TODO: 
    function createGame(uint idxValue1, uint maskedValue1, uint idxBlinding, uint maskedBlinding, uint8[] memory proof, uint8[] memory commitment) public {
        require(inputMaskOwner[idxValue1] == msg.sender);
	require(inputMaskOwner[idxBlinding] == msg.sender);

        address player1 = msg.sender;
        uint gameId = ++gameCnt;

        uint seqCreateGame = opCnt++;
        opEvent[seqCreateGame] = "CreateGame";
        opContent[seqCreateGame] = abi.encode(gameId, player1, idxValue1, maskedValue1, idxBlinding, maskedBlinding);
        emit CreateGame(seqCreateGame, gameId, player1, idxValue1, maskedValue1, idxBlinding, maskedBlinding, proof, commitment);
    }

    function joinGame(uint gameId, uint idxValue2, uint maskedValue2) public {
        require(inputMaskOwner[idxValue2] == msg.sender);

        require(status[gameId] == 1);
        address player2 = msg.sender;

        uint seqJoinGame = opCnt++;
        opEvent[seqJoinGame] = "JoinGame";
        opContent[seqJoinGame] = abi.encode(gameId, player2, idxValue2, maskedValue2);
        emit JoinGame(seqJoinGame, gameId, player2, idxValue2, maskedValue2);
    }

    function startRecon(uint gameId) public { // 1 < 2; 2 < 3; 3 < 1;
        require(status[gameId] == 2);
        status[gameId]++;

        uint seqStartRecon = opCnt++;
        opEvent[seqStartRecon] = "StartRecon";
        opContent[seqStartRecon] = abi.encode(gameId);
        emit StartRecon(seqStartRecon, gameId);
    }

}
