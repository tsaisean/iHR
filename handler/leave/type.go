package leave

import repo "iHR/repositories"

type LeaveHandler struct {
	repo *repo.LeaveRepo
}

func NewLeaveHandler(repo *repo.LeaveRepo) *LeaveHandler {
	return &LeaveHandler{repo: repo}
}
