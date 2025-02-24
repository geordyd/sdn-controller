package traffic

type TrafficReceived struct {
	Packet Traffic
}

type TrafficAllowed struct {
	Packet Traffic
}

type TrafficBlocked struct {
	Packet Traffic
}
