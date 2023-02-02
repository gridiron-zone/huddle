package types

// DONTCOVER

import (
	subspacestypes "github.com/gridiron-zone/huddle/x/subspaces/types"
)

var (
	// PermissionReportContent allows users to report contents
	PermissionReportContent = subspacestypes.RegisterPermission("report content")

	// PermissionDeleteOwnReports allows users to delete existing reports made by their own
	PermissionDeleteOwnReports = subspacestypes.RegisterPermission("delete own reports")

	// PermissionManageReports allows users to manage other users reports
	PermissionManageReports = subspacestypes.RegisterPermission("manage reports")

	// PermissionManageReasons allows users to manage a subspace reasons for reporting
	PermissionManageReasons = subspacestypes.RegisterPermission("manage reasons")
)
