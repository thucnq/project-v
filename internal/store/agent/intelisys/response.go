package intelisysclient

import (
	"workflow/internal/model/agent"
)

type SearchUserResponse struct {
	Data  []*UserInfo `json:"data"`
	Count int         `json:"count"`
	Links Links       `json:"links"`
}

// Convert SearchUserResponse to agent.UserInfo
func (o SearchUserResponse) ToAgent() *agent.SearchUserResponse {
	obj := &agent.SearchUserResponse{
		Count: o.Count,
		Links: agent.Links{
			Next: o.Links.Next,
			Prev: o.Links.Prev,
		},
	}
	// Get UserIDs
	for _, v := range o.Data {
		obj.UserIDs = append(obj.UserIDs, v.ID)
	}
	return obj
}

type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type UserInfo struct {
	ID                 int64              `json:"id"`
	Lang               string             `json:"lang"`
	DisplayName        string             `json:"display_name"`
	Avatar             string             `json:"avatar"`
	AvatarThumbPattern string             `json:"avatar_thumb_pattern"`
	IdentifierCode     string             `json:"identifier_code,omitempty"`
	Email              string             `json:"email,omitempty"`
	PhoneNumber        string             `json:"phone_number,omitempty"`
	Title              string             `json:"title,omitempty"`
	Department         string             `json:"department,omitempty"`
	CompanyName        string             `json:"company_name,omitempty"`
	EmployeeCode       string             `json:"employee_code,omitempty"`
	DepartmentID       string             `json:"department_id,omitempty"`
	DepartmentIDs      interface{}        `json:"department_ids,omitempty"`
	Departments        interface{}        `json:"departments,omitempty"`
	ListDepartments    []DepartmentInUser `json:"list_departments,omitempty"`
	Role               int                `json:"role,omitempty"`
	RoleID             string             `json:"role_id,omitempty"`
	State              int                `json:"state,omitempty"`
	Status             int                `json:"status,omitempty"`
	Supervisor         bool               `json:"supervisor,omitempty"`
	IdentityCode       string             `json:"identity_code,omitempty"`
}

type DepartmentInUser struct {
	TreeID        string   `json:"tree_id"`
	TreeName      string   `json:"tree_name"`
	DepartmentIDs []string `json:"department_ids"`
	Departments   []string `json:"departments"`
	DepartmentID  string   `json:"department_id"`
	Department    string   `json:"department"`
	RoleID        string   `json:"role_id"`
	Title         string   `json:"title"`
}

// Convert UserInfo to agent.UserInfo
func (o UserInfo) ToAgent() *agent.UserInfo {
	obj := &agent.UserInfo{}
	obj.Id = o.ID
	obj.Lang = o.Lang
	obj.DisplayName = o.DisplayName
	obj.Avatar = o.Avatar
	obj.AvatarThumbPattern = o.AvatarThumbPattern
	obj.IdentifierCode = o.IdentifierCode
	obj.Email = o.Email
	obj.PhoneNumber = o.PhoneNumber
	obj.Title = o.Title
	obj.Department = o.Department
	obj.CompanyName = o.CompanyName
	obj.EmployeeCode = o.EmployeeCode
	obj.DepartmentID = o.DepartmentID
	if v, ok := o.DepartmentIDs.([]string); ok {
		obj.DepartmentIDs = v
	}
	if v, ok := o.Departments.([]string); ok {
		obj.Departments = v
	}
	for _, v := range o.ListDepartments {
		obj.ListDepartments = append(obj.ListDepartments, agent.DepartmentInUser{
			TreeID:        v.TreeID,
			TreeName:      v.TreeName,
			DepartmentIDs: v.DepartmentIDs,
			Departments:   v.Departments,
			DepartmentID:  v.DepartmentID,
			Department:    v.Department,
			RoleID:        v.RoleID,
			Title:         v.Title,
		})
	}
	obj.Role = o.Role
	obj.RoleID = o.RoleID
	obj.State = o.State
	obj.Status = o.Status
	obj.Supervisor = o.Supervisor
	obj.IdentityCode = o.IdentityCode
	return obj
}
