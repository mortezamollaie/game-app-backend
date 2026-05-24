package entity

type Role uint8

const (
	TypicalUserRole = iota + 1
	SuperAdminRole
	StaffRole
	CopyWriterRole
	AccountantRole
)

func (r Role) String() string {
	switch r {
	case TypicalUserRole:
		return "TypicalUserRole"
	case SuperAdminRole:
		return "SuperAdminRole"
	case StaffRole:
		return "StaffRole"
	case CopyWriterRole:
		return "CopyWriterRole"
	case AccountantRole:
		return "AccountantRole"
	default:
		return "Unknown"
	}
}
