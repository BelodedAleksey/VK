package vkapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//Wall struct
type Wall struct {
	Count    int         `json:"count"`
	Posts    []*WallPost `json:"items"`
	Profiles []*User     `json:"profiles"`
	Groups   []*Group    `json:"groups"`
}

//WallPost struct
type WallPost struct {
	ID           int                  `json:"id"`
	FromID       int                  `json:"from_id"`
	OwnerID      int                  `json:"owner_id"`
	ToID         int                  `json:"to_id"`
	Date         int64                `json:"date"`
	MarkedAsAd   int                  `json:"marked_as_ads"`
	IsPinned     int                  `json:"is_pinned"`
	PostType     string               `json:"post_type"`
	CopyPostDate int64                `json:"copy_post_date"`
	CopyPostType string               `json:"copy_post_type"`
	CopyOwnerID  int                  `json:"copy_owner_id"`
	CopyPostID   int                  `json:"copy_post_id"`
	CopyHistory  []*WallPost          `json:"copy_history"`
	CreatedBy    int                  `json:"created_by"`
	Text         string               `json:"text"`
	CanDelete    int                  `json:"can_delete"`
	CanPin       int                  `json:"can_pin"`
	Attachments  []*MessageAttachment `json:"attachments"`
	PostSource   *Source              `json:"post_source"`
	Comments     *Comment             `json:"comments"`
	Likes        *Like                `json:"likes"`
	Reposts      *Repost              `json:"reposts"`
	Online       int                  `json:"online"`
	ReplyCount   int                  `json:"reply_count"`
}

//Comment struct
type Comment struct {
	Count   int `json:"count"`
	CanPost int `json:"can_post"`
}

//PostComment struct
type PostComment struct {
	ID             int                  `json:"id"`
	FromID         int                  `json:"from_id"`
	Date           int64                `json:"date"`
	Text           string               `json:"text"`
	ReplyToUser    int                  `json:"reply_to_user"`
	ReplyToComment int                  `json:"reply_to_comment"`
	Attachments    []*MessageAttachment `json:"attachments"`
	ParentsStack   []*PostComment       `json:"parents_stack"`
	//Thread object `json:"thread"`
}

//Like struct
type Like struct {
	Count      int `json:"count"`
	UserLikes  int `json:"user_likes"`
	CanLike    int `json:"can_like"`
	CanPublish int `json:"can_publish"`
}

//Repost struct
type Repost struct {
	Count        int `json:"count"`
	UserReposted int `json:"user_reposted"`
}

//Source struct
type Source struct {
	Type string `json:"type"`
}

//WallGetComments func owner_id/post_id/need_likes/start_comment_id/offset/
//	count/sort/preview_length/extended/fields/comment_id/thread_items_count
func (client *VKClient) WallGetComments(count int, ownerID int, postID int, params url.Values) (int, []*PostComment, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("count", strconv.Itoa(count))
	params.Set("owner_id", strconv.Itoa(ownerID))
	params.Set("post_id", strconv.Itoa(postID))

	resp, err := client.MakeRequest("wall.getComments", params)
	if err != nil {
		return 0, nil, err
	}
	var comments struct {
		Count    int            `json:"count"`
		Comments []*PostComment `json:"items"`
	}
	json.Unmarshal(resp.Response, &comments)

	return comments.Count, comments.Comments, nil
}

//WallGet func owner_id/domain/offset/count/filter/extended/fields
func (client *VKClient) WallGet(id interface{}, count int, params url.Values) (*Wall, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("count", strconv.Itoa(count))

	switch id.(type) {
	case int:
		params.Set("owner_id", strconv.Itoa(id.(int)))
	case string:
		params.Set("domain", id.(string))
	}

	resp, err := client.MakeRequest("wall.get", params)
	if err != nil {
		return nil, err
	}

	var wall *Wall
	json.Unmarshal(resp.Response, &wall)

	return wall, nil
}

//WallGetByID func posts/extended/copy_history_depth/fileds
func (client *VKClient) WallGetByID(id string, params url.Values) (*Wall, error) {
	if params == nil {
		params = url.Values{}
	}

	params.Set("posts", id)

	resp, err := client.MakeRequest("wall.getById", params)
	if err != nil {
		return nil, err
	}

	wall := &Wall{}
	if params.Get(`extended`) == `1` {
		json.Unmarshal(resp.Response, &wall)
	} else {
		json.Unmarshal(resp.Response, &wall.Posts)
	}
	wall.Count = len(wall.Posts)
	return wall, nil
}

//WallPost func owner_id/friends_only/from_group/message/attachments/services/
//	signed/publish_date/lat/long/place_id/post_id/guid/mark_as_ads/close_comments/
//	mute_notifications
func (client *VKClient) WallPost(ownerID int, message string, params url.Values) (int, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("owner_id", strconv.Itoa(ownerID))
	params.Set("message", message)

	resp, err := client.MakeRequest("wall.post", params)
	if err != nil {
		return 0, err
	}
	m := map[string]int{}
	if err = json.Unmarshal(resp.Response, &m); err != nil {
		return 0, err
	}

	return m["post_id"], nil
}

//WallCreateComment func owner_id/post_id/from_group/message/reply_to_comment/
//	attachments/sticker_id/guid
func (client *VKClient) WallCreateComment(ownerID int, postID int, message string, params url.Values) (int, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("owner_id", strconv.Itoa(ownerID))
	params.Set("post_id", strconv.Itoa(postID))
	params.Set("message", message)

	resp, err := client.MakeRequest("wall.createComment", params)
	if err != nil {
		return 0, err
	}
	m := map[string]int{}
	if err = json.Unmarshal(resp.Response, &m); err != nil {
		return 0, err
	}

	return m["comment_id"], nil

}
