package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type GroupsResponse struct {
	Status
	models.Pagination
	Groups []models.Group `json:"groups"`
}

type GroupResponse struct {
	Status
	Group models.Group `json:"group"`
}

type GroupMembersResponse = ChannelMembersResponse

type GroupMessagesResponse = ChannelMessagesResponse

// CreateGroup Creates a new private group, optionally including specified users. The group creator is always included.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/create
func (c *Client) CreateGroup(group *models.CreateGroupRequest) (*models.Group, error) {
	body, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	response := new(GroupResponse)

	err = c.Post("groups.create", bytes.NewBuffer(body), response)
	if err != nil {
		return nil, err
	}

	return &response.Group, err
}

// DeleteGroup remove a private channel.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/delete
func (c *Client) DeleteGroup(group *models.Group) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, group.ID)
	return c.Post("groups.delete", bytes.NewBufferString(body), new(GroupResponse))
}

// GetGroupInfo retrieves the information about the private group, only if you're part of the group.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/info
func (c *Client) GetGroupInfo(group *models.Group) (*models.Group, error) {
	response := new(GroupResponse)
	switch {
	case group.Name != "" && group.ID == "":
		if err := c.Get("groups.info", url.Values{"roomName": []string{group.Name}}, response); err != nil {
			return nil, err
		}
	default:
		if err := c.Get("groups.info", url.Values{"roomId": []string{group.ID}}, response); err != nil {
			return nil, err
		}
	}

	return &response.Group, nil
}

// InviteGroup Adds a user to the private group.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/invite
func (c *Client) InviteGroup(group *models.InviteGroupRequest) (*models.Group, error) {
	body, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	response := new(GroupResponse)

	err = c.Post("groups.invite", bytes.NewBuffer(body), response)
	if err != nil {
		return nil, err
	}

	return &response.Group, err
}

// Removes a user from the private group.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/kick
func (c *Client) KickGroup(group *models.InviteGroupRequest) (*models.Group, error) {
	body, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	response := new(GroupResponse)

	err = c.Post("groups.kick", bytes.NewBuffer(body), response)
	if err != nil {
		return nil, err
	}

	return &response.Group, err
}

// LeaveGroup remove a private channel.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/leave
func (c *Client) LeaveGroup(group *models.Group) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, group.ID)
	return c.Post("groups.leave", bytes.NewBufferString(body), new(GroupResponse))
}

// ListGroup remove a private channel.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/list
func (c *Client) ListGroup() ([]models.Group, error) {
	response := new(GroupsResponse)
	err := c.Get("groups.list", nil, response)
	return response.Groups, err
}

// MembersGroup remove a private channel.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/list
func (c *Client) MembersGroup(group *models.Group) ([]models.User, error) {
	response := new(GroupMembersResponse)
	err := c.Get("groups.members", url.Values{"roomId": []string{group.ID}}, response)
	return response.Members, err
}

// SetAnnouncementGroup remove a private channel.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/setannouncement
func (c *Client) SetAnnouncementGroup(groupID, announcement string) error {
	var body = fmt.Sprintf(`{ "roomId": "%s", "announcement": "%s" }`, groupID, announcement)
	return c.Post("groups.setAnnouncement", bytes.NewBufferString(body), new(GroupResponse))
}

// AddOwnerGroup gives the role of owner for a user in the current group.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/addowner
func (c *Client) AddOwnerGroup(group *models.InviteGroupRequest) (*models.Group, error) {
	body, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	response := new(GroupResponse)

	err = c.Post("groups.addOwner", bytes.NewBuffer(body), response)
	if err != nil {
		return nil, err
	}

	return &response.Group, err
}

// RemoveOwnerGroup removes the role of owner from a user in the current Group.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/removeowner
func (c *Client) RemoveOwnerGroup(group *models.InviteGroupRequest) (*models.Group, error) {
	body, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	response := new(GroupResponse)

	err = c.Post("groups.removeOwner", bytes.NewBuffer(body), response)
	if err != nil {
		return nil, err
	}

	return &response.Group, err
}

// HistoryGroup retrieves the messages from a private group, only if you're part of the group.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/history
func (c *Client) HistoryGroup(group *models.Group) ([]models.Message, error) {
	response := new(GroupMessagesResponse)
	err := c.Get("groups.history", url.Values{"roomId": []string{group.ID}}, response)
	return response.Messages, err
}

// MessagesGroup Lists all of the specific group messages on the server. It supports the Offset, Count, and Sort Query Parameters along with Query and Fields Query Parameters.
//
// https://docs.rocket.chat/api/rest-api/methods/groups/messages
func (c *Client) MessagesGroup(group *models.Group) ([]models.Message, error) {
	response := new(GroupMessagesResponse)
	err := c.Get("groups.messages", url.Values{"roomId": []string{group.ID}}, response)
	return response.Messages, err
}
