int nextTruckId
mapping { int => sint, sint } truckRegistry

func createTruck() {
    truckId = nextTruckId++
}

func recordShipment(int truckId, sint timeLoad, sint timeUnload) {
    sint validShipment = (timeLoad <= timeUnload).reveal()
    if validShipment == 1 {
        truckRegistry.append((timeLoad, timeUnload))
    }
}

func queryPositions(int truckId, sint tL, sint tR) {
    positions = []
    for i, (timeLoad, timeUnload) in enumerate(truckRegistry[truckId]) {
        if (timeLoad <= tR).reveal() && (timeUnload >= tL).reveal() {
            positions.append(i)
        }
    }
}

func queryNumber(int truckId, sint tL, sint tR) {
    sint cnt = 0
    for timeLoad, timeUnload in truckRegistry[truckId] {
        cnt += (timeLoad <= tR) * (timeUnload >= tL)
    }
    cnt.reveal()
}

func queryFirst(int truckId, sint tL, sint tR) {
    record = sint(0), sint(0)
    for timeLoad, timeUnload in truckRegistry[truckId] {
        matched = (timeLoad <= tR) * (timeUnload >= tL)
        firstMatched = (record.first == 0) * matched
        record = firstMatched * timeLoad, firstMatched * timeUnload
    }
    record.reveal()
}