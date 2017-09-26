package main

import "time"

type Cards []struct {
	ID                string        `json:"id"`
	CheckItemStates   interface{}   `json:"checkItemStates"`
	Closed            bool          `json:"closed"`
	DateLastActivity  time.Time     `json:"dateLastActivity"`
	Desc              string        `json:"desc"`
	DescData          interface{}   `json:"descData"`
	IDBoard           string        `json:"idBoard"`
	IDList            string        `json:"idList"`
	IDMembersVoted    []interface{} `json:"idMembersVoted"`
	IDShort           int           `json:"idShort"`
	IDAttachmentCover interface{}   `json:"idAttachmentCover"`
	Limits            struct {
		Attachments struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"attachments"`
		Checklists struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"checklists"`
		Stickers struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"stickers"`
	} `json:"limits"`
	IDLabels              []string `json:"idLabels"`
	ManualCoverAttachment bool     `json:"manualCoverAttachment"`
	Name                  string   `json:"name"`
	Pos                   float64  `json:"pos"`
	ShortLink             string   `json:"shortLink"`
	Badges                struct {
		Votes             int `json:"votes"`
		AttachmentsByType struct {
			Trello struct {
				Board int `json:"board"`
				Card  int `json:"card"`
			} `json:"trello"`
		} `json:"attachmentsByType"`
		ViewingMemberVoted bool      `json:"viewingMemberVoted"`
		Subscribed         bool      `json:"subscribed"`
		Fogbugz            string    `json:"fogbugz"`
		CheckItems         int       `json:"checkItems"`
		CheckItemsChecked  int       `json:"checkItemsChecked"`
		Comments           int       `json:"comments"`
		Attachments        int       `json:"attachments"`
		Description        bool      `json:"description"`
		Due                time.Time `json:"due"`
		DueComplete        bool      `json:"dueComplete"`
	} `json:"badges"`
	DueComplete  bool          `json:"dueComplete"`
	Due          time.Time     `json:"due"`
	IDChecklists []string      `json:"idChecklists"`
	IDMembers    []interface{} `json:"idMembers"`
	Labels       []struct {
		ID      string      `json:"id"`
		IDBoard string      `json:"idBoard"`
		Name    string      `json:"name"`
		Color   interface{} `json:"color"`
		Uses    int         `json:"uses"`
	} `json:"labels"`
	ShortURL   string `json:"shortUrl"`
	Subscribed bool   `json:"subscribed"`
	URL        string `json:"url"`
}

type CheckList struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IDBoard    string `json:"idBoard"`
	CheckItems []struct {
		State       string      `json:"state"`
		IDChecklist string      `json:"idChecklist"`
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		NameData    interface{} `json:"nameData"`
		Pos         int         `json:"pos"`
	} `json:"checkItems"`
}
