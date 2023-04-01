package coinbase

type Response[T any] struct {
	Pagination Pagination `json:"pagination"`
	Data       T          `json:"data"`
}

type Pagination struct {
	EndingBefore         interface{} `json:"ending_before"`
	StartingAfter        interface{} `json:"starting_after"`
	PreviousEndingBefore interface{} `json:"previous_ending_before"`
	NextStartingAfter    interface{} `json:"next_starting_after"`
	Limit                int         `json:"limit"`
	Order                string      `json:"order"`
	PreviousUri          interface{} `json:"previous_uri"`
	NextUri              interface{} `json:"next_uri"`
}
