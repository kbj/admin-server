package response

type MenuTreeModel struct {
	ID       *uint            `json:"id"`
	Sequence *int             `json:"sequence"`
	ParentId *uint            `json:"parentId"`
	Path     *string          `json:"path"`
	Icon     *string          `json:"icon"`
	IsHide   *bool            `json:"isHide"`
	Type     *uint8           `json:"type"`
	Level    *uint8           `json:"level"`
	Children *[]MenuTreeModel `json:"children"`
}
