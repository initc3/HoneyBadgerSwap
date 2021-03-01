//////////////////
// common codes //
//////////////////
struct MaskedInput{
    uint inputMaskIdx
    uint maskedValue
}

//////////////////
// global codes //
//////////////////
uint usedInputMasks
mapping (uint => address) inputMaskOwner
uint tradeCnt

event Trade(uint tradeSeq, address user, address tokenA, address tokenB, MaskedInput amtA, MaskedInput amtB)

func reserveInput(uint _num) public {
    for (uint i = 0; i < _num; i++) {
        uint inputMaskIdx = usedInputMasks
        inputMaskOnwer[inputMaskIdx] = msg.sender
        usedInputMasks++
    }
}

func trade(address _tokenA, address _tokenB, MaskedInput _amtA, MaskedInput _amtB) public {
    require(_tokenA < _tokenB)
    address user = msg.sender
    require(_checkOwnership(_amtA, user))
    require(_checkOwnership(_amtB, user))
    emit Trade(tradeCnt, user, _tokenA, _tokenB, _amtA, _amtB)
    tradeCnt++
}

func _checkOwnership(MaskedInput input, address user) internal return bool {
    return inputMaskOwner[input.inputMaskIdx] == user
}

//////////////////
// local codes //
//////////////////
uint totalInputMasks

func genInputMask(uint _num) {
    uint[_num] newInputMasks = _genInputMask(_num)
    for inputMask in newInputMasks) {
        uint inputMaskIdx = totalInputMasks
        totalInputMasks++
        dbPut(f'inputmask_{inputMaskIdx}', inputMask)
    }
}

// example: key = 'inputmask_7'
func httpRequest(_key) {
    return dbGet(_key)
}

func trade(uint _tradeSeq, address _user, address _tokenA, address _tokenB, MaskedInput _amtA, MaskedInput _amtB) {
    sint poolA = dbGet(f'pool_{_tokenA}_{_tokenB}_{_tokenA}')
    sint poolB = dbGet(f'pool_{_tokenA}_{_tokenB}_{_tokenB}')
    sint balanceA = dbGet(f'balance_{_tokenA}_{_user}')
    sint balanceB = dbGet(f'balance_{_tokenB}_{_user}')
    sint amtA = _amtA.maskedValue - dbGet(f'inputmask_{_amtA.inputMaskIdx}')
    sint amtB = _amtB.maskedValue - dbGet(f'inputmask_{_amtB.inputMaskIdx}')

    sint validOrder = (amtA * amtB) < 0

    sint buyA = (amtA > 0)
    sint enoughB = (-amtB <= balanceB)
    sint actualAmtA = poolA - poolA * poolB / (poolB - amtB)
    sint acceptA = (actualAmtA >= amtA)
    sint flagBuyA = validOrder * buyA * enougthB * accecptA

    sint buyB = 1 - buyA
    sint enoughA = (-amtA <= balanceA)
    sint actualAmtB = poolB - poolA * poolB / (poolA - amtA)
    sint acceptB = (actualAmtB >= amtB)
    sint flatBuyB = validOrder * buyB * enoughA * acceptB

    sint changeA = flagBuyA * actualAmtA + flagBuyB * amtA
    sint changeB = flagBuyA * amtB + flagBuyB * actualAmtB
    poolA -= changeA
    poolB -= changeB
    balanceA += changeA
    balanceB += changeB

    dbPut(f'pool_{_tokenA}_{_tokenB}_{_tokenA}', poolA)
    dbPut(f'pool_{_tokenA}_{_tokenB}_{_tokenB}', poolB)
    dbPut(f'balance_{_tokenA}_{_user}', balanceA)
    dbPut(f'balance_{_tokenB}_{_user}', balanceB)
}

