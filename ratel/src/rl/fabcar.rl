pragma solidity ^0.5.0;

contract fabcar {

    uint public nextTruckId;
    uint public queryCnt;

    mapping (uint => address) public truckOwner;

    event NewTruck(uint truckId);

    constructor() public {}

    function createTruck() public {
        uint truckId = ++nextTruckId;
        address owner = msg.sender;
        truckOwner[truckId] = owner;
        emit NewTruck(truckId);
    }

    function recordShipment(uint truckId, $uint timeLoad, $uint timeUnload) public {
        require(truckOwner[truckId] != address(0));

        mpc(uint truckId, $uint timeLoad, $uint timeUnload) {
            mpcInput(sint timeLoad, sint timeUnload)
            print_ln('**** timeLoad %s', timeLoad.reveal())
            print_ln('**** timeUnload %s', timeUnload.reveal())

            validShipment = timeLoad.less_equal(timeUnload, bit_length=bit_length).reveal()

            print_ln('**** validShipment %s', validShipment)
            mpcOutput(cint validShipment)

            print('**** validShipment', validShipment)
            if validShipment == 1:
                truckRegistry = readDB(f'truckRegistry_{truckId}', list)

                truckRegistry.append((timeLoad, timeUnload))
                print('**** truckRegistry', truckRegistry)

                writeDB(f'truckRegistry_{truckId}', truckRegistry, list)
        }
    }

    function queryPositions(uint truckId, $uint tL, $uint tR) public {
        uint querySeq = ++queryCnt;

        mpc(uint querySeq, uint truckId, $uint tL, $uint tR) {
            truckRegistry = readDB(f'truckRegistry_{truckId}', list)

            positions = []
            for i, (timeLoad, timeUnload) in enumerate(truckRegistry):
                print('**** i', i)

                mpcInput(sint tL, sint tR, sint timeLoad, sint timeUnload)
                print_ln('**** tL %s', tL.reveal())
                print_ln('**** tR %s', tR.reveal())
                print_ln('**** timeLoad %s', timeLoad.reveal())
                print_ln('**** timeUnload %s', timeUnload.reveal())

                inRange = ((timeLoad.less_equal(tR, bit_length=bit_length)) * (timeUnload.greater_equal(tL, bit_length=bit_length))).reveal()
                print_ln('**** inRange %s', inRange)

                mpcOutput(cint inRange)

                print('**** inRange', inRange)
                if inRange == 1:
                    positions.append(i)
                print('**** positions', positions)

            writeDB(f'query_{querySeq}', positions, list)
        }
    }

    function queryNumber(uint truckId, $uint tL, $uint tR) public {
        uint querySeq = ++queryCnt;

        mpc(uint querySeq, uint truckId, $uint tL, $uint tR) {
            truckRegistry = readDB(f'truckRegistry_{truckId}', list)
            print('**** truckRegistry', truckRegistry)

            cnt = 0
            for i, (timeLoad, timeUnload) in enumerate(truckRegistry):
                print('**** i', i)

                mpcInput(sint tL, sint tR, sint timeLoad, sint timeUnload, sint cnt)
                print_ln('**** tL %s', tL.reveal())
                print_ln('**** tR %s', tR.reveal())
                print_ln('**** timeLoad %s', timeLoad.reveal())
                print_ln('**** timeUnload %s', timeUnload.reveal())

                inRange = (timeLoad.less_equal(tR, bit_length=bit_length)) * (timeUnload.greater_equal(tL, bit_length=bit_length))
                print_ln('**** inRange %s', inRange.reveal())

                cnt += inRange
                print_ln('**** cnt %s', cnt.reveal())

                mpcOutput(sint cnt)

            mpcInput(sint cnt)

            cnt = cnt.reveal()

            mpcOutput(cint cnt)
            print('**** cnt', cnt)

            writeDB(f'query_{querySeq}', cnt, int)
        }
    }

    function queryFirst(uint truckId, $uint tL, $uint tR) public {
        uint querySeq = ++queryCnt;

        mpc(uint querySeq, uint truckId, $uint tL, $uint tR) {
            truckRegistry = readDB(f'truckRegistry_{truckId}', list)
            print('**** truckRegistry', truckRegistry)

            a, b = 0, 0
            for i, (timeLoad, timeUnload) in enumerate(truckRegistry):
                mpcInput(sint tL, sint tR, sint timeLoad, sint timeUnload, sint a, sint b)
                print_ln('**** a %s', a.reveal())
                print_ln('**** b %s', b.reveal())

                inRange = (timeLoad.less_equal(tR, bit_length=bit_length)) * (timeUnload.greater_equal(tL, bit_length=bit_length))
                firstMatched = (a.equal(0, bit_length=bit_length)) * inRange
                print_ln('**** inRange %s', inRange.reveal())
                print_ln('**** firstMatched %s', firstMatched.reveal())
                a += firstMatched * timeLoad
                b += firstMatched * timeUnload

                print_ln('**** a %s', a.reveal())
                print_ln('**** b %s', b.reveal())
                mpcOutput(sint a, sint b)

            mpcInput(sint a, sint b)

            a = a.reveal()
            b = b.reveal()

            mpcOutput(cint a, cint b)

            print('****', 'a', a, 'b', b)

            ans = [a, b]
            writeDB(f'query_{querySeq}', ans, list)
        }
    }
}
