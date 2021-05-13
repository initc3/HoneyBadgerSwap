int paperCnt
int sessionCnt
mapping { int => Info } sessionInfo
mapping { int, int, address => bool } reviewed
mapping { int, int => sint } totalScore
mapping { int, int => int } reviewCnt
mapping { int, int => bool } accepted

struct Info {
    int dueRegistration, dueReview
    int numReviewer
    []address reviewerAddrs
    int numPaper
    bool reviewerAssigned
    mapping { int, address => bool } isReviewer
}

func initSession(int dueRegistration, int dueReview, int numReviewer, []address reviewerAddrs) {
    require(numReviewer == reviewerAddrs.length)
    sessionId = ++sessionCnt
    sessionInfo[sessionId] = Info{
        dueRegistration : dueRegistration
        dueReview : dueReview
        numReviewer : numReviewer
        reviewerAddrs : reviewerAddrs
    }
}

func registerPaper(int sessionId) {
    require(block.number < sessionInfo[sessionId].dueRegistration)
    sessionInfo[sessionId].numPaper++
}

func assignReviewer(int sessionId, [][]int reviewers) {
    require(block.number >= sessionInfo[sessionId].dueRegistration)
    require(!sessionInfo[sessionId].reviewerAssigned)

    for i in range(reviewers.length) {
        for j in range(reviewers[i].length) {
            sessionInfo[sessionId].isReviewer[i][reviewerAddrs[reviewers[i][j]]] = true
        }
    }
    sessionInfo[sessionId].reviewerAssigned = true
}

func peerReview(int sessionId, int paperId, sint score) {
    require(block.number < sessionInfo[sessionId].dueReview)
    require(sessionInfo[sessionId].reviewerAssigned)
    address reviewer = msg.sender
    require(sessionInfo[sessionId].isReviewer[paperId][reviewer])
    require(!reviewed[sessionId][paperId][reviewer])

    int validScore = (score >= 0 && score <= maxScore).reveal()
    if validScore == 1 {
        totalScore[sessionId][paperId] += score
        reviewCnt[sessionId][paperId] += 1
        reviewed[sessionId][paperId][reviewer] = true
    }
}

func calcResult(int sessionId, int threshold) {
    require(block.number >= sessionInfo[sessionId].dueReview)

    for paperId in range(sessionInfo[sessionId].numPaper) {
        int accept = (totalScore[sessionId][paperId] / reviewCnt[sessionId][paperId] >= threshold).reveal()
        accepted[sessionId][paperId] = accept
    }
}