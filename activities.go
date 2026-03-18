package vantagepoint

import (
	"context"
	"fmt"
)

// Activity represents a Vantagepoint activity record. This struct covers core
// business fields from the ~95+ fields returned by the API. Additional fields
// (recurrence detail fields, Outlook integration, custom UDIC fields) exist in
// the API response but are omitted here. Expand from the Postman collection as needed.
//
// Note: VP uses "Y"/"N" strings for boolean fields like CompletionInd, PrivateInd,
// AllDayEventInd. ActivityID is a GUID-style string.
type Activity struct {
	ActivityID         string `json:"ActivityID,omitempty"`
	Type               string `json:"Type,omitempty"`
	TypeName           string `json:"TypeName,omitempty"`
	Subject            string `json:"Subject,omitempty"`
	Notes              string `json:"Notes,omitempty"`
	HasNotes           string `json:"HasNotes,omitempty"`
	Location           string `json:"Location,omitempty"`
	StartDate          string `json:"StartDate,omitempty"`
	EndDate            string `json:"EndDate,omitempty"`
	Duration           int    `json:"Duration,omitempty"`
	Priority           string `json:"Priority,omitempty"`
	PriorityDesc       string `json:"PriorityDesc,omitempty"`
	CompletionInd      string `json:"CompletionInd,omitempty"`
	PrivateInd         string `json:"PrivateInd,omitempty"`
	AllDayEventInd     string `json:"AllDayEventInd,omitempty"`
	ShowTimeAs         string `json:"ShowTimeAs,omitempty"`
	Employee           string `json:"Employee,omitempty"`
	CreateEmployee     string `json:"CreateEmployee,omitempty"`
	Owner              string `json:"Owner,omitempty"`
	OwnerName          string `json:"OwnerName,omitempty"`
	ClientID           string `json:"ClientID,omitempty"`
	ClientName         string `json:"ClientName,omitempty"`
	ContactID          string `json:"ContactID,omitempty"`
	ContactName        string `json:"ContactName,omitempty"`
	Vendor             string `json:"Vendor,omitempty"`
	VendorName         string `json:"VendorName,omitempty"`
	WBS1               string `json:"WBS1,omitempty"`
	WBS2               string `json:"WBS2,omitempty"`
	WBS3               string `json:"WBS3,omitempty"`
	ProjName           string `json:"ProjName,omitempty"`
	CampaignID         string `json:"CampaignID,omitempty"`
	MktName            string `json:"MktName,omitempty"`
	LeadID             string `json:"LeadID,omitempty"`
	RecurrenceInd      string `json:"RecurrenceInd,omitempty"`
	RecurrType         string `json:"RecurrType,omitempty"`
	IsTask             string `json:"isTask,omitempty"`
	IsCalendar         string `json:"isCalendar,omitempty"`
	TaskCompletionDate string `json:"TaskCompletionDate,omitempty"`
	ReminderInd        string `json:"ReminderInd,omitempty"`
	ReminderUnit       int    `json:"ReminderUnit,omitempty"`
	ReminderMinHrDay   string `json:"ReminderMinHrDay,omitempty"`
	ReminderDate       string `json:"ReminderDate,omitempty"`
	CreateUser         string `json:"CreateUser,omitempty"`
	CreateDate         string `json:"CreateDate,omitempty"`
	ModUser            string `json:"ModUser,omitempty"`
	ModDate            string `json:"ModDate,omitempty"`
}

// ActivityContact represents a contact associated with an activity.
type ActivityContact struct {
	ActivityID          string `json:"ActivityID,omitempty"`
	ContactID           string `json:"ContactID,omitempty"`
	Email               string `json:"Email,omitempty"`
	Name                string `json:"Name,omitempty"`
	Phone               string `json:"Phone,omitempty"`
	CellPhone           string `json:"CellPhone,omitempty"`
	PrimaryInd          string `json:"PrimaryInd,omitempty"`
	Title               string `json:"Title,omitempty"`
	Company             string `json:"Company,omitempty"`
	Type                string `json:"Type,omitempty"`
	FirstLast           string `json:"FirstLast,omitempty"`
	LastFirst           string `json:"LastFirst,omitempty"`
	ClientID            string `json:"ClientID,omitempty"`
	ContactHasPhoto     string `json:"ContactHasPhoto,omitempty"`
	ContactPhotoModDate string `json:"ContactPhotoModDate,omitempty"`
}

// ActivityEmployee represents an employee associated with an activity.
type ActivityEmployee struct {
	ActivityID   string `json:"ActivityID,omitempty"`
	Employee     string `json:"Employee,omitempty"`
	Name         string `json:"Name,omitempty"`
	FullName     string `json:"FullName,omitempty"`
	Title        string `json:"Title,omitempty"`
	Owner        string `json:"Owner,omitempty"`
	Email        string `json:"Email,omitempty"`
	WorkPhone    string `json:"WorkPhone,omitempty"`
	WorkPhoneExt string `json:"WorkPhoneExt,omitempty"`
	Org          string `json:"Org,omitempty"`
	OrgName      string `json:"OrgName,omitempty"`
	FirstLast    string `json:"FirstLast,omitempty"`
	LastFirst    string `json:"LastFirst,omitempty"`
}

// ActivityClient represents a client (firm) associated with an activity.
type ActivityClient struct {
	ActivityID string `json:"ActivityID,omitempty"`
	ClientID   string `json:"ClientID,omitempty"`
	Name       string `json:"Name,omitempty"`
	PrimaryInd string `json:"PrimaryInd,omitempty"`
}

// ActivityLink represents a file link associated with an activity.
type ActivityLink struct {
	ActivityID  string `json:"ActivityID,omitempty"`
	LinkID      string `json:"LinkID,omitempty"`
	Description string `json:"Description,omitempty"`
	FilePath    string `json:"FilePath,omitempty"`
	Graphic     string `json:"Graphic,omitempty"`
}

// ListActivities retrieves a list of activities from the Vantagepoint API.
// Pass a *Query to filter, paginate, or sort results; nil retrieves defaults.
func (c *Client) ListActivities(ctx context.Context, q *Query) ([]Activity, error) {
	var activities []Activity
	if err := c.get(ctx, "activity", q, &activities); err != nil {
		return nil, fmt.Errorf("listing activities: %w", err)
	}
	return activities, nil
}

// GetActivity retrieves a single activity by its ActivityID. The Vantagepoint
// API returns an array even for single-resource lookups, so the response is
// decoded as a slice and the first element is returned.
func (c *Client) GetActivity(ctx context.Context, activityID string) (*Activity, error) {
	var activities []Activity
	if err := c.get(ctx, "activity/"+activityID, nil, &activities); err != nil {
		return nil, fmt.Errorf("getting activity %s: %w", activityID, err)
	}
	if len(activities) == 0 {
		return nil, fmt.Errorf("getting activity %s: %w", activityID, ErrNotFound)
	}
	return &activities[0], nil
}

// CreateActivity creates a new activity record in Vantagepoint.
func (c *Client) CreateActivity(ctx context.Context, activity *Activity) (*Activity, error) {
	var result Activity
	if err := c.post(ctx, "activity", activity, &result); err != nil {
		return nil, fmt.Errorf("creating activity: %w", err)
	}
	return &result, nil
}

// UpdateActivity updates an existing activity record identified by ActivityID.
func (c *Client) UpdateActivity(ctx context.Context, activityID string, activity *Activity) (*Activity, error) {
	var result Activity
	if err := c.put(ctx, "activity/"+activityID, activity, &result); err != nil {
		return nil, fmt.Errorf("updating activity %s: %w", activityID, err)
	}
	return &result, nil
}

// DeleteActivity deletes an activity record by its ActivityID.
func (c *Client) DeleteActivity(ctx context.Context, activityID string) error {
	if err := c.delete(ctx, "activity/"+activityID); err != nil {
		return fmt.Errorf("deleting activity %s: %w", activityID, err)
	}
	return nil
}

// ListActivityContacts retrieves contacts associated with an activity.
// Uses GET /activity/{ActivityKey}/contact.
func (c *Client) ListActivityContacts(ctx context.Context, activityKey string, q *Query) ([]ActivityContact, error) {
	var results []ActivityContact
	if err := c.get(ctx, "activity/"+activityKey+"/contact", q, &results); err != nil {
		return nil, fmt.Errorf("listing activity contacts for %s: %w", activityKey, err)
	}
	return results, nil
}

// ListActivityEmployees retrieves employees associated with an activity.
// Uses GET /activity/{ActivityKey}/employee.
func (c *Client) ListActivityEmployees(ctx context.Context, activityKey string, q *Query) ([]ActivityEmployee, error) {
	var results []ActivityEmployee
	if err := c.get(ctx, "activity/"+activityKey+"/employee", q, &results); err != nil {
		return nil, fmt.Errorf("listing activity employees for %s: %w", activityKey, err)
	}
	return results, nil
}

// ListActivityClients retrieves clients (firms) associated with an activity.
// Uses GET /activity/{ActivityKey}/client.
func (c *Client) ListActivityClients(ctx context.Context, activityKey string, q *Query) ([]ActivityClient, error) {
	var results []ActivityClient
	if err := c.get(ctx, "activity/"+activityKey+"/client", q, &results); err != nil {
		return nil, fmt.Errorf("listing activity clients for %s: %w", activityKey, err)
	}
	return results, nil
}

// ListActivityLinks retrieves file links associated with an activity.
// Uses GET /activity/{ActivityKey}/links.
func (c *Client) ListActivityLinks(ctx context.Context, activityKey string, q *Query) ([]ActivityLink, error) {
	var results []ActivityLink
	if err := c.get(ctx, "activity/"+activityKey+"/links", q, &results); err != nil {
		return nil, fmt.Errorf("listing activity links for %s: %w", activityKey, err)
	}
	return results, nil
}
