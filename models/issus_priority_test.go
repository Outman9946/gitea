// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateIssuePriority(t *testing.T) {
	assert.NoError(t, PrepareTestDatabase())

	issue := AssertExistsAndLoadBean(t, &Issue{ID: 1}).(*Issue)
	repo := AssertExistsAndLoadBean(t, &Repository{ID: issue.RepoID}).(*Repository)
	doer := AssertExistsAndLoadBean(t, &User{ID: repo.OwnerID}).(*User)

	issue.Priority = PriorityPinned

	err := UpdateIssuePriority(issue, doer)
	assert.NoError(t, err)

	issue = AssertExistsAndLoadBean(t, &Issue{ID: 1}).(*Issue)
	assert.EqualValues(t, PriorityPinned, issue.Priority)

	issue.Priority = -1
	err = UpdateIssuePriority(issue, doer)
	assert.Error(t, err)
	assert.EqualValues(
		t, err, ErrIssueInvalidPriority{ID: issue.ID, RepoID: issue.Repo.ID, DesiredPriority: issue.Priority})
}
