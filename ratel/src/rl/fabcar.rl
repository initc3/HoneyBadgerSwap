pragma solidity ^0.5.0;

contract fabcar {

    uint public nextTruckId;

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
            timeLoad *= fp
            timeUnload *= fp

            mpcInput(timeLoad, timeUnload)
            timeLoad = sfix._new(timeLoad)
            timeUnload = sfix._new(timeUnload)
            print_ln('timeLoad %s', timeLoad.reveal())
            print_ln('timeUnload %s', timeUnload.reveal())

            validShipment = (timeLoad <= timeUnload).reveal()
            validShipment = sint(validShipment)

            mpcOutput(validShipment)

            print(validShipment)
            if validShipment == 1:
                truckRegistry = readDB(f'truckRegistry_{truckId}')
                try:
                    import ast
                    truckRegistry = truckRegistry.decode(encoding='utf-8')
                    truckRegistry = list(ast.literal_eval(truckRegistry))
                except:
                    truckRegistry = []

                truckRegistry.append((timeLoad, timeUnload))
                print(truckRegistry)

                truckRegistry = str(truckRegistry)
                truckRegistry = bytes(truckRegistry, encoding='utf-8')
                writeDB(f'truckRegistry_{truckId}', truckRegistry)
        }
    }

    function queryPositions(uint truckId, $uint tL, $uint tR) public {
        mpc(uint truckId, $uint tL, $uint tR) {
            tL *= fp
            tR *= fp

            truckRegistry = readDB(f'truckRegistry_{truckId}')
            try:
                import ast
                truckRegistry = truckRegistry.decode(encoding='utf-8')
                truckRegistry = list(ast.literal_eval(truckRegistry))
            except:
                truckRegistry = []

            print('truckRegistry', truckRegistry)
            positions = []
            for i, (timeLoad, timeUnload) in enumerate(truckRegistry):
                print(i, (timeLoad, timeUnload))
                mpcInput(tL, tR, timeLoad, timeUnload)
                tL = sfix._new(tL)
                tR = sfix._new(tR)
                timeLoad = sfix._new(timeLoad)
                timeUnload = sfix._new(timeUnload)

                inRange = ((timeLoad <= tR) * (timeUnload >= tL)).reveal()
                inRange = sint(inRange)

                mpcOutput(inRange)

                print('inRange', inRange)
                if inRange == 1:
                    positions.append(i)
                print(positions)
        }
    }

    function queryNumber(uint truckId, $uint tL, $uint tR) public {
        mpc(uint truckId, $uint tL, $uint tR) {
            tL *= fp
            tR *= fp

            truckRegistry = readDB(f'truckRegistry_{truckId}')
            try:
                import ast
                truckRegistry = truckRegistry.decode(encoding='utf-8')
                truckRegistry = list(ast.literal_eval(truckRegistry))
            except:
                truckRegistry = []

            print('truckRegistry', truckRegistry)
            cnt = 0
            for i, (timeLoad, timeUnload) in enumerate(truckRegistry):
                print(i, (timeLoad, timeUnload))
                mpcInput(tL, tR, timeLoad, timeUnload, cnt)
                tL = sfix._new(tL)
                tR = sfix._new(tR)
                timeLoad = sfix._new(timeLoad)
                timeUnload = sfix._new(timeUnload)

                inRange = (timeLoad <= tR) * (timeUnload >= tL)
                print_ln('inRange %s', inRange.reveal())
                cnt += inRange
                print_ln('cnt %s', cnt.reveal())

                mpcOutput(cnt)

            mpcInput(cnt)

            cnt = sint(cnt.reveal())

            mpcOutput(cnt)
            print(cnt)
        }
    }

    function queryFirst(uint truckId, $uint tL, $uint tR) public {
        mpc(uint truckId, $uint tL, $uint tR) {
            print('yes')
            tL *= fp
            tR *= fp

            truckRegistry = readDB(f'truckRegistry_{truckId}')
            try:
                import ast
                truckRegistry = truckRegistry.decode(encoding='utf-8')
                truckRegistry = list(ast.literal_eval(truckRegistry))
            except:
                truckRegistry = []

            print('truckRegistry', truckRegistry)
            a, b = 0, 0
            for i, (timeLoad, timeUnload) in enumerate(truckRegistry):
                mpcInput(tL, tR, timeLoad, timeUnload, a, b)
                tL = sfix._new(tL)
                tR = sfix._new(tR)
                timeLoad = sfix._new(timeLoad)
                timeUnload = sfix._new(timeUnload)
                a = sfix._new(a)
                b = sfix._new(b)
                print_ln('a %s', a.reveal())
                print_ln('b %s', b.reveal())

                inRange = (timeLoad <= tR) * (timeUnload >= tL)
                firstMatched = (a == 0) * inRange
                print_ln('inRange %s', inRange.reveal())
                print_ln('firstMatched %s', firstMatched.reveal())
                a += firstMatched * timeLoad
                b += firstMatched * timeUnload
                print_ln('a %s', a.reveal())
                print_ln('b %s', b.reveal())

                a = a.v
                b = b.v
                mpcOutput(a, b)

            mpcInput(a, b)
            a = sfix._new(a)
            b = sfix._new(b)

            a = sint(a.reveal().v)
            b = sint(b.reveal().v)
            mpcOutput(a, b)
            a //= fp
            b //= fp
            print(a, b)
        }
    }
}
