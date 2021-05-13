int gameCnt
mapping { int => Game } gameBoard
int[] activeGameList
int[] completedGameList

struct Game {
    address player1, player2
    sint value1, value2
    string status // active, ready, completed
    address winner
}

func createGame(sint value1) {
    gameBoard[++gameCnt] = Game{
        player1 : msg.sender
        value1 : value1
        status: active
    }
    activeGamelist.add(gameCnt)
}

func joinGame(int gameId, sint value2) {
    require(gameBoard[gameId].status == active)
    gameBoard[gameId].player2 = msg.sender
    gameBoard[gameId].value2 = value2
    gameBoard[gameId].status = ready
    activeGamelist.remove(gameCnt)
}

func startRecon(int gameId) {
    Game game = gameBoard[gameId]
    require(game.status == ready)

    int dif = (game.value1 - game.value2).reveal()

    submit(gameBoard[gameId].result = (dif == 1 || dif == -2) ? game.player1 : game.player2)

    gameBoard[gameId].status = completed

    completedGameList.add(gameId)
}

