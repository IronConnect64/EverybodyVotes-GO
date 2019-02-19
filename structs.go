package main

type register struct {
	Serial string `json:"serial"`
	MAC    string `json:"mac"`
	Token  string `json:"token"`
}

type vote struct {
	ID     string `json:"id"`
	PollID int    `json:"poll"`
	Choice string `json:"choice"`
}

type poll struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	A        string `json:"a"`
	B        string `json:"b"`
	Latest   bool   `json:"latest"`
}
