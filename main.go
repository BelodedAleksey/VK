package main

import (
	"fmt"

	vkapi "./golang-vk-api"
)

func main() {
	client, err := vkapi.NewVKClient(vkapi.DeviceIPhone, "", "")
	if err != nil {
		fmt.Println("Error VKCLient: ", err)
	}

	// listening sent messages
	/*client.AddLongpollCallback("msgout", func(m *vkapi.LongPollMessage) {
		fmt.Printf("sent message to uid %d\n", m.UserID)
	})

	// starting
	client.ListenLongPollServer()*/

	//Account

	//appWidgets

	//Apps

	//Auth

	//Board

	//Database

	//Docs

	//Other

	//Fave

	//Friends
	_, users, _ := client.FriendsGet(client.Self.UID, 100)
	for _, u := range users {
		fmt.Println(u.LastName, " ", u.FirstName)
	}

	//Gifts

	//Groups

	//leadForms

	//Likes

	//Market

	//Messages
	dialog, _ := client.MessagesGetConversations(10, nil)
	for _, item := range dialog.Items {
		fmt.Println("TEXT: ", item.LastMessage.Text)
	}
	_, messages, err := client.MessagesGetHistory(10, "502201296", nil)
	if err != nil {
		fmt.Println("Error MessagesGetHistory: ", err)
	}
	for _, m := range messages {
		fmt.Println("Message: ", m.Body)
	}

	//Newsfeed

	//Notes

	//Notifications

	//Pages

	//Photos

	//Polls

	//prettyCards

	//Search

	//Stats

	//Status

	//Storage

	//stories

	//streaming

	//Users

	//Utils

	//Video

	//Wall
	fmt.Println("Посты с Подслушано Апатиты")
	wall, err := client.WallGet(-61904449, 5, nil)
	if err != nil {
		fmt.Println("Error WallGet: ", err)
	}
	for _, post := range wall.Posts {
		for _, attach := range post.Attachments {
			fmt.Println("Тип вложения", attach.Type)
			fmt.Println("Фото", attach.Photo.Photo75)
		}
		for _, repost := range post.CopyHistory {
			fmt.Println("Репост: ", repost.Text)
			for _, attach := range repost.Attachments {
				fmt.Println("Тип вложения", attach.Type)
				fmt.Println("Фото", attach.Photo.Photo75)
			}
		}
		fmt.Println("Пост: ", post.Text)
		_, comments, err := client.WallGetComments(5, post.OwnerID, post.ID, nil)
		if err != nil {
			fmt.Println("Error WallGetComments: ", err)
		}
		for _, c := range comments {
			fmt.Println("Коммент: ", c.Text)
		}

	}
	//Widgets

}
