// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package smtp_test

import (
	"code.gitea.io/gitea/models/login"
	"code.gitea.io/gitea/services/auth"
	"code.gitea.io/gitea/services/auth/source/smtp"
)

// This test file exists to assert that our Source exposes the interfaces that we expect
// It tightly binds the interfaces and implementation without breaking go import cycles

type sourceInterface interface {
	auth.PasswordAuthenticator
	login.Config
	login.SkipVerifiable
	login.HasTLSer
	login.UseTLSer
	login.SourceSettable
}

var _ (sourceInterface) = &smtp.Source{}
