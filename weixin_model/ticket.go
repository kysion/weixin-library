package weixin_model

/*
TicketResult 票据返回结果
*/
type TicketResult struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}
