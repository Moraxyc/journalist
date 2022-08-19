// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/item"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Categories holds the value of the "categories" field.
	Categories string `json:"categories,omitempty"`
	// CrawledTitle holds the value of the "crawled_title" field.
	CrawledTitle string `json:"crawled_title,omitempty"`
	// CrawledAuthor holds the value of the "crawled_author" field.
	CrawledAuthor string `json:"crawled_author,omitempty"`
	// CrawledExcerpt holds the value of the "crawled_excerpt" field.
	CrawledExcerpt string `json:"crawled_excerpt,omitempty"`
	// CrawledSiteName holds the value of the "crawled_site_name" field.
	CrawledSiteName string `json:"crawled_site_name,omitempty"`
	// CrawledImage holds the value of the "crawled_image" field.
	CrawledImage string `json:"crawled_image,omitempty"`
	// CrawledContentHTML holds the value of the "crawled_content_html" field.
	CrawledContentHTML string `json:"crawled_content_html,omitempty"`
	// CrawledContentText holds the value of the "crawled_content_text" field.
	CrawledContentText string `json:"crawled_content_text,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges      ItemEdges `json:"edges"`
	feed_items *uuid.UUID
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Feed holds the value of the feed edge.
	Feed *Feed `json:"feed,omitempty"`
	// ReadByUsers holds the value of the read_by_users edge.
	ReadByUsers []*User `json:"read_by_users,omitempty"`
	// Reads holds the value of the reads edge.
	Reads []*Read `json:"reads,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// FeedOrErr returns the Feed value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) FeedOrErr() (*Feed, error) {
	if e.loadedTypes[0] {
		if e.Feed == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: feed.Label}
		}
		return e.Feed, nil
	}
	return nil, &NotLoadedError{edge: "feed"}
}

// ReadByUsersOrErr returns the ReadByUsers value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) ReadByUsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.ReadByUsers, nil
	}
	return nil, &NotLoadedError{edge: "read_by_users"}
}

// ReadsOrErr returns the Reads value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) ReadsOrErr() ([]*Read, error) {
	if e.loadedTypes[2] {
		return e.Reads, nil
	}
	return nil, &NotLoadedError{edge: "reads"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldTitle, item.FieldDescription, item.FieldContent, item.FieldURL, item.FieldAuthor, item.FieldImage, item.FieldCategories, item.FieldCrawledTitle, item.FieldCrawledAuthor, item.FieldCrawledExcerpt, item.FieldCrawledSiteName, item.FieldCrawledImage, item.FieldCrawledContentHTML, item.FieldCrawledContentText:
			values[i] = new(sql.NullString)
		case item.FieldCreatedAt, item.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case item.FieldID:
			values[i] = new(uuid.UUID)
		case item.ForeignKeys[0]: // feed_items
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Item", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case item.FieldTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[j])
			} else if value.Valid {
				i.Title = value.String
			}
		case item.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		case item.FieldContent:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[j])
			} else if value.Valid {
				i.Content = value.String
			}
		case item.FieldURL:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[j])
			} else if value.Valid {
				i.URL = value.String
			}
		case item.FieldAuthor:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[j])
			} else if value.Valid {
				i.Author = value.String
			}
		case item.FieldImage:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[j])
			} else if value.Valid {
				i.Image = value.String
			}
		case item.FieldCategories:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field categories", values[j])
			} else if value.Valid {
				i.Categories = value.String
			}
		case item.FieldCrawledTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawled_title", values[j])
			} else if value.Valid {
				i.CrawledTitle = value.String
			}
		case item.FieldCrawledAuthor:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawled_author", values[j])
			} else if value.Valid {
				i.CrawledAuthor = value.String
			}
		case item.FieldCrawledExcerpt:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawled_excerpt", values[j])
			} else if value.Valid {
				i.CrawledExcerpt = value.String
			}
		case item.FieldCrawledSiteName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawled_site_name", values[j])
			} else if value.Valid {
				i.CrawledSiteName = value.String
			}
		case item.FieldCrawledImage:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawled_image", values[j])
			} else if value.Valid {
				i.CrawledImage = value.String
			}
		case item.FieldCrawledContentHTML:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawled_content_html", values[j])
			} else if value.Valid {
				i.CrawledContentHTML = value.String
			}
		case item.FieldCrawledContentText:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crawled_content_text", values[j])
			} else if value.Valid {
				i.CrawledContentText = value.String
			}
		case item.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case item.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case item.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field feed_items", values[j])
			} else if value.Valid {
				i.feed_items = new(uuid.UUID)
				*i.feed_items = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryFeed queries the "feed" edge of the Item entity.
func (i *Item) QueryFeed() *FeedQuery {
	return (&ItemClient{config: i.config}).QueryFeed(i)
}

// QueryReadByUsers queries the "read_by_users" edge of the Item entity.
func (i *Item) QueryReadByUsers() *UserQuery {
	return (&ItemClient{config: i.config}).QueryReadByUsers(i)
}

// QueryReads queries the "reads" edge of the Item entity.
func (i *Item) QueryReads() *ReadQuery {
	return (&ItemClient{config: i.config}).QueryReads(i)
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("title=")
	builder.WriteString(i.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(i.Content)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(i.URL)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(i.Author)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(i.Image)
	builder.WriteString(", ")
	builder.WriteString("categories=")
	builder.WriteString(i.Categories)
	builder.WriteString(", ")
	builder.WriteString("crawled_title=")
	builder.WriteString(i.CrawledTitle)
	builder.WriteString(", ")
	builder.WriteString("crawled_author=")
	builder.WriteString(i.CrawledAuthor)
	builder.WriteString(", ")
	builder.WriteString("crawled_excerpt=")
	builder.WriteString(i.CrawledExcerpt)
	builder.WriteString(", ")
	builder.WriteString("crawled_site_name=")
	builder.WriteString(i.CrawledSiteName)
	builder.WriteString(", ")
	builder.WriteString("crawled_image=")
	builder.WriteString(i.CrawledImage)
	builder.WriteString(", ")
	builder.WriteString("crawled_content_html=")
	builder.WriteString(i.CrawledContentHTML)
	builder.WriteString(", ")
	builder.WriteString("crawled_content_text=")
	builder.WriteString(i.CrawledContentText)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
