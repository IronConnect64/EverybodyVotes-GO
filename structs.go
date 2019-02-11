package main

type register struct {
	Serial string `json:"serial"`
	MAC    string `json:"mac"`
}

type login struct {
	ID  string `json:"id"`
	MAC string `json:"mac"`
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
}
