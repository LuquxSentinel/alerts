package types

type Location struct {
	Longitute string	`json:"longitute"`
	Latitude string	`json:"latitude"`
}

type PublishLocationInput struct {
	Location Location `json:"location"`
	ChannelID string `json:"channel_id"`
}