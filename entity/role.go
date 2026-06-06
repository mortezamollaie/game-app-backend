package entity

type Role uint8

const (
	TypicalUserRole Role = iota + 1
	SuperAdminRole  Role = iota + 2
	StaffRole       Role = iota + 3
	CopyWriterRole  Role = iota + 4
	AccountantRole  Role = iota + 5
)

const (
	TypicalUserStr = "typicalUser"
	SuperAdminStr  = "superAdmin"
	StaffStr       = "staff"
	CopyWriterStr  = "copyWriter"
	AccountantStr  = "accountant"
)

func (r Role) String() string {
	switch r {
	case TypicalUserRole:
		return TypicalUserStr
	case SuperAdminRole:
		return SuperAdminStr
	case StaffRole:
		return StaffStr
	case CopyWriterRole:
		return CopyWriterStr
	case AccountantRole:
		return AccountantStr
	default:
		return TypicalUserStr
	}
}

func MapToRoleEntity(roleStr string) Role {
	switch roleStr {
	case TypicalUserStr:
		return TypicalUserRole
	case SuperAdminStr:
		return SuperAdminRole
	case StaffStr:
		return StaffRole
	case CopyWriterStr:
		return CopyWriterRole
	case AccountantStr:
		return AccountantRole
	default:
		return TypicalUserRole
	}
}
