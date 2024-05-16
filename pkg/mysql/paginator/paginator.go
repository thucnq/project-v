package paginator

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"

	"github.com/iancoleman/strcase"

	stringutils "project-v/pkg/strings"
)

const (
	defaultLimit = 10
	defaultOrder = DESC
)

func New() *Paginator {
	return &Paginator{}
}

// Credit: https://github.com/pilagod/gorm-cursor-paginator
type Paginator struct {
	cursor    Cursor
	next      Cursor
	keys      []string
	tableKeys []string
	limit     int
	order     Order
}

func (p *Paginator) SetCursor(value string) {
	decode, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return
	}

	token := string(decode)
	if strings.HasPrefix(token, "next-") {
		p.cursor.NextPageToken = normalizeToken(string(decode))
	}

	if strings.HasPrefix(token, "prev-") {
		p.cursor.PrevPageToken = normalizeToken(string(decode))
	}
}

// SetAfterCursor sets paging after cursor
func (p *Paginator) SetAfterCursor(afterCursor string) {
	decode, err := base64.StdEncoding.DecodeString(afterCursor)
	if err != nil {
		return
	}

	p.cursor.NextPageToken = normalizeToken(string(decode))
}

// SetBeforeCursor sets paging before cursor
func (p *Paginator) SetBeforeCursor(beforeCursor string) {
	decode, err := base64.StdEncoding.DecodeString(beforeCursor)
	if err != nil {
		return
	}

	p.cursor.PrevPageToken = normalizeToken(string(decode))
}

func normalizeToken(token string) *string {
	if strings.HasPrefix(token, "prev-") {
		return stringutils.String(strings.TrimPrefix(token, "prev-"))
	}

	if strings.HasPrefix(token, "next-") {
		return stringutils.String(strings.TrimPrefix(token, "next-"))
	}

	return stringutils.String(token)
}

// SetKeys sets paging keys
func (p *Paginator) SetKeys(keys ...string) {
	p.keys = append(p.keys, keys...)
}

// SetLimit sets paging limit
func (p *Paginator) SetLimit(limit int) {
	p.limit = limit
}

// SetOrder sets paging order
func (p *Paginator) SetOrder(order Order) {
	p.order = order
}

// GetNextCursor returns cursor for next pagination
func (p *Paginator) GetNextCursor() Cursor {
	cursor := p.next
	if cursor.NextPageToken != nil {
		nextToken := base64.StdEncoding.EncodeToString([]byte("next-" + stringutils.StringValue(cursor.NextPageToken)))
		cursor.NextPageToken = stringutils.String(nextToken)
	}
	if cursor.PrevPageToken != nil {
		prevToken := base64.StdEncoding.EncodeToString([]byte("prev-" + stringutils.StringValue(cursor.PrevPageToken)))
		cursor.PrevPageToken = stringutils.String(prevToken)
	}
	return cursor
}

// Paginate paginates data
func (p *Paginator) Paginate(stmt *gorm.DB, out interface{}) *gorm.DB {
	p.initOptions()
	p.initTableKeys(stmt, out)
	result := p.appendPagingQuery(stmt, out).Find(out)
	// out must be a pointer or gorm will panic above
	elements := reflect.ValueOf(out).Elem()
	if elements.Kind() == reflect.Slice && elements.Len() > 0 {
		p.postProcess(out)
	}
	return result
}

/* private */

func (p *Paginator) initOptions() {
	if len(p.keys) == 0 {
		p.keys = append(p.keys, "CreatedAt", "ID")
	}
	if p.limit == 0 {
		p.limit = defaultLimit
	}
	if p.order == "" {
		p.order = defaultOrder
	}
}

func (p *Paginator) initTableKeys(db *gorm.DB, out interface{}) {
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(out)
	table := stmt.Schema.Table
	for _, key := range p.keys {
		p.tableKeys = append(
			p.tableKeys, fmt.Sprintf("%s.%s", table, strcase.ToSnake(key)),
		)
	}
}

func (p *Paginator) appendPagingQuery(stmt *gorm.DB, out interface{}) *gorm.DB {
	decoder, _ := NewCursorDecoder(out, p.keys...)
	var fields []interface{}
	if p.hasAfterCursor() {
		fields = decoder.Decode(*p.cursor.NextPageToken)
	} else if p.hasBeforeCursor() {
		fields = decoder.Decode(*p.cursor.PrevPageToken)
	}
	if len(fields) > 0 {

		stmt = stmt.Where(
			p.getCursorquery(),
			p.getCursorQueryArgs(fields)...,
		)
	}
	stmt = stmt.Limit(p.limit + 1)
	stmt = stmt.Order(p.getOrder())
	return stmt
}

func (p *Paginator) hasAfterCursor() bool {
	return p.cursor.NextPageToken != nil
}

func (p *Paginator) hasBeforeCursor() bool {
	return !p.hasAfterCursor() && p.cursor.PrevPageToken != nil
}

func (p *Paginator) getCursorquery() string {
	qs := make([]string, len(p.tableKeys))
	op := p.getOperator()
	composite := ""
	for i, sqlKey := range p.tableKeys {
		qs[i] = fmt.Sprintf("%s%s %s ?", composite, sqlKey, op)
		composite = fmt.Sprintf("%s%s = ? AND ", composite, sqlKey)
	}
	return strings.Join(qs, " OR ")
}

func (p *Paginator) getCursorQueryArgs(fields []interface{}) (args []interface{}) {
	for i := 1; i <= len(fields); i++ {
		args = append(args, fields[:i]...)
	}
	return
}

func (p *Paginator) getOperator() string {
	if (p.hasAfterCursor() && p.order == ASC) ||
		(p.hasBeforeCursor() && p.order == DESC) {
		return ">"
	}
	return "<"
}

func (p *Paginator) getOrder() string {
	order := p.order
	if p.hasBeforeCursor() {
		order = flip(p.order)
	}
	orders := make([]string, len(p.tableKeys))
	for index, sqlKey := range p.tableKeys {
		orders[index] = fmt.Sprintf("%s %s", sqlKey, order)
	}
	return strings.Join(orders, ", ")
}

func (p *Paginator) postProcess(out interface{}) {
	elements := reflect.ValueOf(out).Elem()
	hasMore := elements.Len() > p.limit
	if hasMore {
		elements.Set(elements.Slice(0, elements.Len()-1))
	}
	if p.hasBeforeCursor() {
		elements.Set(reverse(elements))
	}
	encoder := NewCursorEncoder(p.keys...)
	if p.hasBeforeCursor() || hasMore {
		cursor := encoder.Encode(elements.Index(elements.Len() - 1))
		p.next.NextPageToken = &cursor
	}
	if p.hasAfterCursor() || (hasMore && p.hasBeforeCursor()) {
		cursor := encoder.Encode(elements.Index(0))
		p.next.PrevPageToken = &cursor
	}
	return
}

func reverse(v reflect.Value) reflect.Value {
	result := reflect.MakeSlice(v.Type(), 0, v.Cap())
	for i := v.Len() - 1; i >= 0; i-- {
		result = reflect.Append(result, v.Index(i))
	}
	return result
}
