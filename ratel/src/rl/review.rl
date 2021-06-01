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

    function initSession(uint timeRegistration, uint timeReview, uint _numReviewer, address[] memory _reviewerAddrs) public {
        require(_numReviewer > 0 && _numReviewer == _reviewerAddrs.length);
        require(timeRegistration > 0 && timeReview > 0);

        uint sessionId = ++sessionCnt;

        dueRegistration[sessionId] = block.number + timeRegistration;
        dueReview[sessionId] = dueRegistration[sessionId] + timeReview;
        numReviewer[sessionId] = _numReviewer;
        reviewerAddrs[sessionId] = _reviewerAddrs;

        emit NewSession(sessionId);
    }

    function registerPaper(uint sessionId) public {
        require(block.number < dueRegistration[sessionId]);
        numPaper[sessionId]++;
    }

    function assignReviewer(uint sessionId, uint reviewersPerPaper, uint[] memory reviewers) public {
        require(block.number >= dueRegistration[sessionId] && block.number < dueReview[sessionId]);
        require(!reviewerAssigned[sessionId]);
        require(numPaper[sessionId] * reviewersPerPaper == reviewers.length);

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
        require(block.number < dueReview[sessionId]);
        require(reviewerAssigned[sessionId]);

        address reviewer = msg.sender;
        require(isReviewer[sessionId][paperId][reviewer]);

        require(!reviewed[sessionId][paperId][reviewer]);
        reviewed[sessionId][paperId][reviewer] = true;

        mpc(uint sessionId, uint paperId, $uint score) {
            print('*** id', sessionId, paperId)

            totalScore = int.from_bytes(readDB(f'totalScore_{sessionId}_{paperId}'), 'big')
            totalCnt = int.from_bytes(readDB(f'totalCnt{sessionId}_{paperId}'), 'big')

            print('*** total', totalScore, totalCnt)

            score *= fp
            mpcInput(score, totalScore, totalCnt)
            score = sfix._new(score)
            totalScore = sfix._new(totalScore)
            totalCnt = sfix._new(totalCnt)

            print_ln('score %s', score.reveal())
            print_ln('totalScore %s', totalScore.reveal())
            print_ln('totalCnt %s', totalCnt.reveal())

            maxScore = 5

            validScore = (score >= 0) * (score <= maxScore)
            if_then(validScore.reveal())
            totalScore += score
            totalCnt += 1
            end_if()

            print_ln('totalScore %s', totalScore.reveal())
            print_ln('totalCnt %s', totalCnt.reveal())

            totalScore = totalScore.v
            totalCnt = totalCnt.v
            mpcOutput(totalScore, totalCnt)

            writeDB(f'totalScore_{sessionId}_{paperId}', totalScore.to_bytes((totalScore.bit_length() + 7) // 8, 'big'))
            writeDB(f'totalCnt{sessionId}_{paperId}', totalCnt.to_bytes((totalCnt.bit_length() + 7) // 8, 'big'))
        }
    }

    function calcResult(uint sessionId, $uint line) public {
        require(block.number >= dueReview[sessionId]);

        uint paperNum = numPaper[sessionId];

        mpc(uint sessionId, uint paperNum, $uint line) {
            for paperId in range(paperNum):
                totalScore = int.from_bytes(readDB(f'totalScore_{sessionId}_{paperId}'), 'big')
                totalCnt = int.from_bytes(readDB(f'totalCnt{sessionId}_{paperId}'), 'big')

                print('line', line)

                mpcInput(totalScore, totalCnt, line)
                totalScore = sfix._new(totalScore)
                totalCnt = sfix._new(totalCnt)
                line = sfix._new(line)

                print_ln('line %s', line.reveal())
                print_ln('totalScore %s', totalScore.reveal())
                print_ln('totalCnt %s', totalCnt.reveal())

                accept = (totalScore / totalCnt) >= line
                accept = sint(accept.reveal())

                mpcOutput(accept)

                print('!!!', paperId, accept)
        }
    }
}
