package realtime

import "fmt"

func (c *Client) StartTyping(roomID string, username string) error {
	_, err := c.ddp.Call("stream-notify-room", fmt.Sprintf("%s/typing", roomID), username, true)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) StopTyping(roomID string, username string) error {
	_, err := c.ddp.Call("stream-notify-room", fmt.Sprintf("%s/typing", roomID), username, false)
	if err != nil {
		return err
	}

	return nil
}
