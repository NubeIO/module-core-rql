package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetTicket(ticketUUID string) (*model.Ticket, error) {
	url := fmt.Sprintf("/api/tickets/%s", ticketUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Ticket{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Ticket)
	return out, nil
}

func (inst *Client) GetTickets() ([]model.Ticket, error) {
	url := "/api/tickets?with_comments=true&with_teams=true&with_members=true"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.Ticket{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Ticket
	out = *resp.Result().(*[]model.Ticket)
	return out, nil
}

func (inst *Client) AddTicket(body *model.Ticket) (*model.Ticket, error) {
	url := "/api/tickets"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Ticket{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Ticket)
	return out, nil
}

func (inst *Client) DeleteTicket(ticketUUID string) (bool, error) {
	url := fmt.Sprintf("/api/tickets/%s", ticketUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		Delete(url))
	if err != nil {
		return false, err
	}
	return resp.String() == "true", nil
}

func (inst *Client) UpdateTicket(ticketUUID string, body *model.Ticket) (*model.Ticket, error) {
	url := fmt.Sprintf("/api/tickets/%s", ticketUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(body).
		SetResult(&model.Ticket{}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Ticket)
	return out, nil
}

type TicketUpdateStatusBody struct {
	Status string `json:"status"`
}

func (inst *Client) UpdateTicketStatus(ticketUUID string, ticketStatus string) (*bool, error) {
	url := fmt.Sprintf("/api/tickets/%s/status", ticketUUID)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(TicketUpdateStatusBody{
			Status: ticketStatus,
		}).
		SetResult(false).
		Put(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*bool)
	return out, nil
}

type TicketUpdatePriorityBody struct {
	Priority string `json:"priority"`
}

func (inst *Client) UpdateTicketPriority(ticketUUID string, ticketPriority string) (*bool, error) {
	url := fmt.Sprintf("/api/tickets/%s/priority", ticketUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(TicketUpdatePriorityBody{
			Priority: ticketPriority,
		}).
		SetResult(false).
		Put(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*bool)
	return out, nil
}

func (inst *Client) UpdateTicketTeams(ticketUUID string, teams []string) ([]*model.TicketTeam, error) {
	url := fmt.Sprintf("/api/tickets/%s/teams", ticketUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(teams).
		SetResult(&[]*model.TicketTeam{}).
		Put(url))
	if err != nil {
		return nil, err
	}

	out := resp.Result().(*[]*model.TicketTeam)
	return *out, nil
}

func (inst *Client) UpdateTicketMembers(ticketUUID string, members []string) ([]*model.TicketMember, error) {
	url := fmt.Sprintf("/api/tickets/%s/members", ticketUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(members).
		SetResult(&[]*model.TicketMember{}).
		Put(url))
	if err != nil {
		return nil, err
	}

	out := resp.Result().(*[]*model.TicketMember)
	return *out, nil
}

type TicketAddCommentBody struct {
	TicketUUID string `json:"ticket_uuid"`
	Content    string `json:"content"`
}

func (inst *Client) AddTicketComment(ticketUUID string, content string) (*model.Ticket, error) {
	url := "/api/tickets/comments"
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetBody(TicketAddCommentBody{
			TicketUUID: ticketUUID,
			Content:    content,
		}).
		SetResult(&model.Ticket{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.Ticket)
	return out, nil
}
