package traffic

import "github.com/google/uuid"

type Traffic struct {
	ID              uuid.UUID
	SourceIP        string
	DestinationPort int
}
