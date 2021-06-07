pragma solidity ^0.5.0;

contract review {
    uint public sessionCnt;

    mapping (uint => uint) public dueRegistration;
    mapping (uint => uint) public dueReview;
    mapping (uint => uint) public numReviewer;
    mapping (uint => address[]) public reviewerAddrs;
    mapping (uint => uint) public numPaper;
    mapping (uint => bool) public reviewerAssigned;
    mapping (uint => mapping (uint => mapping (address => bool))) isReviewer;
    mapping (uint => mapping (uint => mapping (address => bool))) reviewed;

    event NewSession(uint sessionId);

    constructor() public {}

    function initSession(uint timeRegistration, uint _numReviewer, address[] memory _reviewerAddrs) public {
        require(_numReviewer > 0 && _numReviewer == _reviewerAddrs.length);
        require(timeRegistration > 0);

        uint sessionId = ++sessionCnt;

        dueRegistration[sessionId] = block.number + timeRegistration;
        numReviewer[sessionId] = _numReviewer;
        reviewerAddrs[sessionId] = _reviewerAddrs;

        emit NewSession(sessionId);
    }

    function registerPaper(uint sessionId) public {
        require(block.number < dueRegistration[sessionId]);
        numPaper[sessionId]++;
    }

    function assignReviewer(uint sessionId, uint timeReview, uint reviewersPerPaper, uint[] memory reviewers) public {
        require(timeReview > 0);
        require(block.number >= dueRegistration[sessionId] && !reviewerAssigned[sessionId]);
        require(numPaper[sessionId] * reviewersPerPaper == reviewers.length);

        dueReview[sessionId] = block.number + timeReview;
        reviewerAssigned[sessionId] = true;

        uint k = 0;
        for (uint i = 0; i < numPaper[sessionId]; i++) {
            for (uint j = 0; j < reviewersPerPaper; j++) {
                address reviewer = reviewerAddrs[sessionId][reviewers[k]];
                isReviewer[sessionId][i][reviewer] = true;
                k++;
            }
        }
    }

    function peerReview(uint sessionId, uint paperId, $uint score) public {
        require(block.number < dueReview[sessionId] && reviewerAssigned[sessionId]);

        address reviewer = msg.sender;
        require(isReviewer[sessionId][paperId][reviewer]);
        require(!reviewed[sessionId][paperId][reviewer]);

        reviewed[sessionId][paperId][reviewer] = true;

        mpc(uint sessionId, uint paperId, $uint score) {
            print('****', 'sessionId', sessionId, 'paperId', paperId)

            totalScore = readDB(f'totalScore_{sessionId}_{paperId}', int)
            totalCnt = readDB(f'totalCnt{sessionId}_{paperId}', int)

            mpcInput(sint score, sint totalScore, sint totalCnt)

            print_ln('**** score %s', score.reveal())
            print_ln('**** totalScore %s', totalScore.reveal())
            print_ln('**** totalCnt %s', totalCnt.reveal())

            maxScore = 5

            validScore = (score.greater_equal(0, bit_length=bit_length)) * (score.less_equal(maxScore, bit_length=bit_length))
            if_then(validScore.reveal())
            totalScore += score
            totalCnt += 1
            end_if()

            print_ln('**** totalScore %s', totalScore.reveal())
            print_ln('**** totalCnt %s', totalCnt.reveal())

            mpcOutput(sint totalScore, sint totalCnt)

            writeDB(f'totalScore_{sessionId}_{paperId}', totalScore, int)
            writeDB(f'totalCnt{sessionId}_{paperId}', totalCnt, int)
        }
    }

    function calcResult(uint sessionId, $uint line) public {
        require(block.number >= dueReview[sessionId]);

        uint paperNum = numPaper[sessionId];

        mpc(uint sessionId, uint paperNum, $uint line) {
            for paperId in range(paperNum):
                totalScore = readDB(f'totalScore_{sessionId}_{paperId}', int)
                totalCnt = readDB(f'totalCnt{sessionId}_{paperId}', int)

                mpcInput(sint totalScore, sint totalCnt, sfix line)

                print_ln('**** line %s', line.reveal())
                print_ln('**** totalScore %s', totalScore.reveal())
                print_ln('**** totalCnt %s', totalCnt.reveal())

                accept = ((sfix(totalScore) / totalCnt) >= line).reveal()

                mpcOutput(cint accept)

                print('****', 'paperId', paperId, 'accept', accept)

                writeDB(f'result_{sessionId}_{paperId}', accept, int)
        }
    }
}
