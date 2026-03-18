package vantagepoint

import (
	"context"
	"fmt"
)

// Contact represents a Vantagepoint contact record. This struct covers core
// business fields from the ~119 fields returned by the API. Additional fields
// (firm address block, recurrence/touchpoint fields, Salesforce/QBO integration,
// custom cust* fields, denormalized lookups) exist in the API response but are
// omitted here. Expand from the Postman collection as needed.
//
// Note: VP uses quirky casing for some fields (e.g., CellPhone, not MobilePhone).
// Boolean values are "Y"/"N" strings. ContactStatus values include "A" (active),
// "I" (inactive).
type Contact struct {
	ContactID          string `json:"ContactID,omitempty"`
	ClientID           string `json:"ClientID,omitempty"`
	CLAddress          string `json:"CLAddress,omitempty"`
	Vendor             string `json:"Vendor,omitempty"`
	VEAddress          string `json:"VEAddress,omitempty"`
	Salutation         string `json:"Salutation,omitempty"`
	FirstName          string `json:"FirstName,omitempty"`
	MiddleName         string `json:"MiddleName,omitempty"`
	LastName           string `json:"LastName,omitempty"`
	Suffix             string `json:"Suffix,omitempty"`
	PreferredName      string `json:"PreferredName,omitempty"`
	Title              string `json:"Title,omitempty"`
	ProfessionalSuffix string `json:"ProfessionalSuffix,omitempty"`
	Email              string `json:"Email,omitempty"`
	Phone              string `json:"Phone,omitempty"`
	Fax                string `json:"Fax,omitempty"`
	CellPhone          string `json:"CellPhone,omitempty"`
	HomePhone          string `json:"HomePhone,omitempty"`
	Pager              string `json:"Pager,omitempty"`
	Address1           string `json:"Address1,omitempty"`
	Address2           string `json:"Address2,omitempty"`
	Address3           string `json:"Address3,omitempty"`
	Address4           string `json:"Address4,omitempty"`
	City               string `json:"City,omitempty"`
	State              string `json:"State,omitempty"`
	Zip                string `json:"Zip,omitempty"`
	Country            string `json:"Country,omitempty"`
	Company            string `json:"Company,omitempty"`
	ClientName         string `json:"ClientName,omitempty"`
	ContactStatus      string `json:"ContactStatus,omitempty"`
	CustomCurrencyCode string `json:"CustomCurrencyCode,omitempty"`
	Source             string `json:"Source,omitempty"`
	Owner              string `json:"Owner,omitempty"`
	Rating             string `json:"Rating,omitempty"`
	Market             string `json:"Market,omitempty"`
	Website            string `json:"Website,omitempty"`
	QualifiedStatus    string `json:"QualifiedStatus,omitempty"`
	StatusReason       string `json:"StatusReason,omitempty"`
	StatusDate         string `json:"StatusDate,omitempty"`
	MailingAddress     string `json:"MailingAddress,omitempty"`
	Billing            string `json:"Billing,omitempty"`
	PrimaryInd         string `json:"PrimaryInd,omitempty"`
	Addressee          string `json:"Addressee,omitempty"`
	Memo               string `json:"Memo,omitempty"`
	ProjectDescription string `json:"ProjectDescription,omitempty"`
	HasPhoto           int    `json:"HasPhoto,omitempty"`
	PhotoModDate       string `json:"PhotoModDate,omitempty"`
	CreateUser         string `json:"CreateUser,omitempty"`
	CreateDate         string `json:"CreateDate,omitempty"`
	ModUser            string `json:"ModUser,omitempty"`
	ModDate            string `json:"ModDate,omitempty"`
}

// ContactQualification represents the qualification status of a contact,
// returned by the qualifiedContact endpoint.
type ContactQualification struct {
	ContactID       string `json:"ContactID,omitempty"`
	LastName        string `json:"LastName,omitempty"`
	FirstName       string `json:"FirstName,omitempty"`
	QualifiedStatus string `json:"QualifiedStatus,omitempty"`
	ClientID        string `json:"ClientID,omitempty"`
	ClientName      string `json:"ClientName,omitempty"`
	OpportunityID   string `json:"OpportunityID,omitempty"`
	OpportunityName string `json:"OpportunityName,omitempty"`
	CampaignID      string `json:"CampaignID,omitempty"`
	CampaignName    string `json:"CampaignName,omitempty"`
	StatusDate      string `json:"StatusDate,omitempty"`
}

// ListContacts retrieves a list of contacts from the Vantagepoint API.
// Pass a *Query to filter, paginate, or sort results; nil retrieves defaults.
func (c *Client) ListContacts(ctx context.Context, q *Query) ([]Contact, error) {
	var contacts []Contact
	if err := c.get(ctx, "contact", q, &contacts); err != nil {
		return nil, fmt.Errorf("listing contacts: %w", err)
	}
	return contacts, nil
}

// GetContact retrieves a single contact by its ContactID. The Vantagepoint API
// returns an array even for single-resource lookups, so the response is decoded
// as a slice and the first element is returned.
func (c *Client) GetContact(ctx context.Context, contactID string) (*Contact, error) {
	var contacts []Contact
	if err := c.get(ctx, "contact/"+contactID, nil, &contacts); err != nil {
		return nil, fmt.Errorf("getting contact %s: %w", contactID, err)
	}
	if len(contacts) == 0 {
		return nil, fmt.Errorf("getting contact %s: %w", contactID, ErrNotFound)
	}
	return &contacts[0], nil
}

// CreateContact creates a new contact record in Vantagepoint.
func (c *Client) CreateContact(ctx context.Context, contact *Contact) (*Contact, error) {
	var result Contact
	if err := c.post(ctx, "contact", contact, &result); err != nil {
		return nil, fmt.Errorf("creating contact: %w", err)
	}
	return &result, nil
}

// UpdateContact updates an existing contact record identified by ContactID.
func (c *Client) UpdateContact(ctx context.Context, contactID string, contact *Contact) (*Contact, error) {
	var result Contact
	if err := c.put(ctx, "contact/"+contactID, contact, &result); err != nil {
		return nil, fmt.Errorf("updating contact %s: %w", contactID, err)
	}
	return &result, nil
}

// DeleteContact deletes a contact record by its ContactID.
func (c *Client) DeleteContact(ctx context.Context, contactID string) error {
	if err := c.delete(ctx, "contact/"+contactID); err != nil {
		return fmt.Errorf("deleting contact %s: %w", contactID, err)
	}
	return nil
}

// GetContactQualificationStatus retrieves the qualification status of a contact.
// Uses GET /contact/{ContactID}/qualifiedContact.
func (c *Client) GetContactQualificationStatus(ctx context.Context, contactID string) (*ContactQualification, error) {
	var result ContactQualification
	if err := c.get(ctx, "contact/"+contactID+"/qualifiedContact", nil, &result); err != nil {
		return nil, fmt.Errorf("getting contact qualification status for %s: %w", contactID, err)
	}
	return &result, nil
}

// QualifyContact qualifies a contact (changes status to qualified).
// Uses POST /contact/{ContactID}/qualify.
func (c *Client) QualifyContact(ctx context.Context, contactID string) error {
	if err := c.post(ctx, "contact/"+contactID+"/qualify", nil, nil); err != nil {
		return fmt.Errorf("qualifying contact %s: %w", contactID, err)
	}
	return nil
}

// DisqualifyContact disqualifies a contact.
// Uses POST /contact/{ContactID}/disqualify.
func (c *Client) DisqualifyContact(ctx context.Context, contactID string) error {
	if err := c.post(ctx, "contact/"+contactID+"/disqualify", nil, nil); err != nil {
		return fmt.Errorf("disqualifying contact %s: %w", contactID, err)
	}
	return nil
}

// RevertContactToLead reverts a contact back to lead status.
// Uses POST /contact/{ContactID}/revertToLead.
func (c *Client) RevertContactToLead(ctx context.Context, contactID string) error {
	if err := c.post(ctx, "contact/"+contactID+"/revertToLead", nil, nil); err != nil {
		return fmt.Errorf("reverting contact %s to lead: %w", contactID, err)
	}
	return nil
}
