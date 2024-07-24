package models

type ActionADD struct {
	SubjectID string
	ObjectID  string
	Action    string
	Result    bool
	Info      string
}
