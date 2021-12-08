package model

type Stats struct {
	ID                int64 `json:"id"`
	UserID            int64 `json:"user_id"`
	BannedBefore      bool  `json:"banned_before"`
	UsersMet          uint  `json:"users_met"`
	MessagesSent      uint  `json:"messages_sent"`
	AverageMessageLen uint  `json:"average_message_length"`
	LinksInMessages   uint  `json:"links_in_messages"`
}

type UserMessage struct {
	Sender string `json:"sender_id"`
	Body   string `json:"message"`
}
