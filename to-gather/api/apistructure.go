package api

type LoginPostData struct {
	Username      string      `json:"username"`
	StdNum        string      `json:"std_num"`
}

type LoginReturnData struct {
	Token         string      `json:"token"`
	UID           int         `json:"uid"`
	StdNum        string      `json:"std_num"`
}

type ActivityPostData struct {
	Year          int         `json:"year"`
	Month         int         `json:"month"`
	Day           int         `json:"day"`
	Time          string      `json:"time"`
	Location      string      `json:"location"`
	Event         string      `json:"event"`
	Question      string      `json:"question"`
	Tel           string      `json:"tel"`
	QQ            string      `json:"qq"`
}

type ActivityPostReturnData struct {
	ActivityID    int         `json:"activity_id"`
}

type ActivityStatu struct {
	Pickable      bool        `json:"pickable"`
	Close         bool        `json:"close"`
}

type ActivityGetReturnData struct {
	Date          string      `json:"date"`
	Time          string      `json:"time"`
	Location      string      `json:"location"`
	Event         string      `json:"event"`
	QQ            string      `json:"qq"`
	Tel           string      `json:"tel"`
	Question      string      `json:"question"`
	PosterName    string      `json:"poster_name"`
	PosterNum     string      `json:"poster_num"`
	Status        ActivityStatu `json:"status"`
}

type ActivityPickData struct {
	Answer        string      `json:"answer"`
}