package vkapi

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

const (
	ActivityTypeTyping   = "typing"
	ActivityTypeAudioMsg = "audiomessage"
)

//Dialog struct
type Dialog struct {
	Count       int     `json:"count"`
	Items       []*Item `json:"items"`
	UnreadCount int     `json:"unread_count"`
	//Profiles array `json:"profiles"`
	//Groups array `json:"groups"`
}

//Item struct
type Item struct {
	Conversation Conversation `json:"conversation"`
	LastMessage  Message      `json:"last_message"`
}

//Peer struct
type Peer struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	LocalID int    `json:"local_id"`
}

//Conversation struct
type Conversation struct {
	Peer        Peer `json:"peer"`
	InRead      int  `json:"in_read"`
	OutRead     int  `json:"out_read"`
	UnreadCount int  `json:"unread_count"`
	Important   bool `json:"important"`
	Unanswered  bool `json:"unanswered"`
	//PushSettings object `json:"push_settings"
	//CanWrite object `json:"can_write"
	//ChatSettings object `json:"chat_settings"
}

//Message struct
type Message struct {
	MID               int                  `json:"id"`
	Date              int64                `json:"date"`
	PeerID            int                  `json:"peer_id"`
	FromID            int                  `json:"from_id"`
	Text              string               `json:"text"`
	RandomID          int                  `json:"random_id"`
	Ref               string               `json:"ref"`
	RefSource         string               `json:"ref_source"`
	Attachments       []*MessageAttachment `json:"attachments"`
	Important         bool                 `json:"important"`
	Payload           string               `json:"payload"`
	ForwardedMessages []*ForwardedMessage  `json:"fwd_messages"`
	Out               int                  `json:"out"`
	ConversationID    int                  `json:"conversation_message_id"`
	IsHidden          bool                 `json:"is_hidden"`
	Body              string               `json:"body"`
	//Geo object `json:"geo"`
	//Keyboard          object            `json:"keyboard"`
	//ReplyMessage object `json:"reply_message"`
	//Action object `json:"action"`
}

type Push struct {
	Sound         int   `json:"sound"`
	DisabledUntil int64 `json:"disabled_until"`
}

//ForwardedMessage struct
type ForwardedMessage struct {
	UID               int                  `json:"user_id"`
	Date              int64                `json:"date"`
	Body              string               `json:"body"`
	Attachments       []*MessageAttachment `json:"attachments"`
	ForwardedMessages []*ForwardedMessage  `json:"fwd_messages"`
}

//MessageAttachment struct
type MessageAttachment struct {
	Type     string             `json:"type"`
	Audio    *AudioAttachment   `json:"audio"`
	Video    *VideoAttachment   `json:"video"`
	Photo    *PhotoAttachment   `json:"photo"`
	Document *DocAttachment     `json:"doc"`
	Link     *LinkAttachment    `json:"link"`
	Wall     *WallPost          `json:"wall"`
	Sticker  *StickerAttachment `json:"sticker"`
}

type StickerAttachment struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Photo64   string `json:"photo_64"`
	Photo128  string `json:"photo_128"`
	Photo256  string `json:"photo_256"`
	Photo352  string `json:"photo_352"`
	Photo512  string `json:"photo_512"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type HistoryAttachment struct {
	Attachments []HistoryAttachmentItem `json:"items"`
	NextFrom    string                  `json:"next_from"`
}

type HistoryAttachmentItem struct {
	MID        int                `json:"message_id"`
	Attachment *MessageAttachment `json:"attachment"`
}

type AudioAttachment struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Artist    string `json:"artist"`
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
	URL       string `json:"url"`
	Performer string `json:"performer"`
}

type VideoAttachment struct {
	ID            int    `json:"id"`
	OwnerID       int    `json:"owner_id"`
	Title         string `json:"title"`
	Duration      int    `json:"duration"`
	Description   string `json:"description"`
	Date          int64  `json:"date"`
	AddingDate    int64  `json:"adding_date"`
	Views         int    `json:"views"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	Photo130      string `json:"photo130"`
	Photo320      string `json:"photo320"`
	Photo800      string `json:"photo800"`
	FirstFrame320 string `json:"first_frame_320"`
	FirstFrame160 string `json:"first_frame_160"`
	FirstFrame130 string `json:"first_frame_130"`
	FirstFrame800 string `json:"first_frame_800"`
	Player        string `json:"player"`
	CanEdit       int    `json:"can_edit"`
	CanAdd        int    `json:"can_add"`
}

type LinkAttachment struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Target      string `json:"target"`
}

//MessagesAddChatUser func chat_id/user_id/visible_messages_count
func (client *VKClient) MessagesAddChatUser() {

}

//MessagesAllowMessagesFromGroup func group_id/key
func (client *VKClient) MessagesAllowMessagesFromGroup() {

}

//MessagesCreateChat func user_ids/title/group_id
func (client *VKClient) MessagesCreateChat() {

}

//MessagesEdit func peer_id/message/message_id/lat/long
//	attachment/keep_forward_messages/keep_snipets/group_id/dont_parse_links
func (client *VKClient) MessagesEdit() {

}

//MessagesGetConversations func offset/count/filter/extended/start_message_id/fields/group_id/major_sort_id
func (client *VKClient) MessagesGetConversations(count int, params url.Values) (*Dialog, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Add("count", strconv.Itoa(count))
	resp, err := client.MakeRequest("messages.getConversations", params)
	if err != nil {
		return nil, err
	}

	var dialog *Dialog
	json.Unmarshal(resp.Response, &dialog)
	return dialog, nil
}

func (client *VKClient) GetHistoryAttachments(peerID int, mediaType string, count int, params url.Values) (*HistoryAttachment, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Add("count", strconv.Itoa(count))
	params.Add("media_type", mediaType)
	params.Add("peer_id", strconv.Itoa(peerID))

	resp, err := client.MakeRequest("messages.getHistoryAttachments", params)
	if err != nil {
		return nil, err
	}

	var att *HistoryAttachment
	json.Unmarshal(resp.Response, &att)
	return att, nil
}

//MessagesGetHistory func offset/count/user_id/peer_id/start_message_id/rev/extended/fields/group_id
func (client *VKClient) MessagesGetHistory(count int, userID string, params url.Values) (int, []*Message, error) {
	if params == nil {
		params = url.Values{}
	}

	params.Add("user_id", userID)
	params.Add("count", strconv.Itoa(count))

	resp, err := client.MakeRequest("messages.getHistory", params)
	if err != nil {
		return 0, nil, err
	}

	var response *struct {
		Count int        `json:"count"`
		Items []*Message `json:"items"`
	}
	json.Unmarshal(resp.Response, &response)
	return response.Count, response.Items, nil
}

//MessagesGetByID func messages_ids/preview_length/extended/fields/group_id
func (client *VKClient) MessagesGetByID(messageIds []int, params url.Values) (int, []*Message, error) {
	if params == nil {
		params = url.Values{}
	}

	s := ArrayToStr(messageIds)
	params.Add("message_ids", s)

	resp, err := client.MakeRequest("messages.getById", params)
	if err != nil {
		return 0, nil, err
	}

	var response *struct {
		Count int        `json:"count"`
		Items []*Message `json:"items"`
	}
	json.Unmarshal(resp.Response, &response)

	return response.Count, response.Items, nil
}

func (client *VKClient) MessagesSend(user interface{}, message string, params url.Values) (APIResponse, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Add("message", message)

	switch user.(type) {
	case int:
		params.Add("user_id", strconv.Itoa(user.(int)))
	case string:
		params.Add("domain", user.(string))
	}

	resp, err := client.MakeRequest("messages.send", params)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (client *VKClient) MessagesDelete(ids []int, spam int, deleteForAll int) (int, error) {
	params := url.Values{}
	s := ArrayToStr(ids)
	params.Add("message_ids", s)
	params.Add("spam", strconv.Itoa(spam))
	params.Add("delete_for_all", strconv.Itoa(deleteForAll))

	resp, err := client.MakeRequest("messages.delete", params)
	if err != nil {
		return 0, err
	}

	delCount := 0
	var idMap map[string]int
	reader := strings.NewReader(string(resp.Response))
	err = json.NewDecoder(reader).Decode(&idMap)
	if err != nil {
		return 0, err
	}

	for _, v := range idMap {
		if v == 1 {
			delCount++
		}
	}

	return delCount, nil
}

func (client *VKClient) MessagesSetActivity(user int, params url.Values) error {
	if params == nil {
		params = url.Values{}
	}

	params.Add("user_id", strconv.Itoa(user))

	_, err := client.MakeRequest("messages.setActivity", params)
	if err != nil {
		return err
	}

	return nil
}
