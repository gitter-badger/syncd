// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
	reqApi "github.com/dreamans/syncd/router/route/api"
    "github.com/dreamans/syncd"
    "github.com/dreamans/syncd/router/user"
    "github.com/dreamans/syncd/router/server"
    "github.com/dreamans/syncd/router/project"
	"github.com/dreamans/syncd/router/deploy"
	"github.com/dreamans/syncd/router/middleware"
)

func RegisterRoute() {
    api := syncd.App.Gin.Group("/api", middleware.ApiPriv())
    {
        api.POST(reqApi.LOGIN, user.Login)
        api.GET(reqApi.LOGIN_STATUS, user.LoginStatus)

        api.POST(reqApi.SERVER_GROUP_ADD, server.GroupAdd)
        api.GET(reqApi.SERVER_GROUP_LIST, server.GroupList)
        api.POST(reqApi.SERVER_GROUP_DELETE, server.GroupDelete)
        api.GET(reqApi.SERVER_GROUP_DETAIL, server.GroupDetail)
        api.POST(reqApi.SERVER_GROUP_UPDATE, server.GroupUpdate)
        api.POST(reqApi.SERVER_ADD, server.ServerAdd)
        api.POST(reqApi.SERVER_UPDATE, server.ServerUpdate)
        api.GET(reqApi.SERVER_LIST, server.ServerList)
        api.POST(reqApi.SERVER_DELETE, server.ServerDelete)
        api.GET(reqApi.SERVER_DETAIL, server.ServerDetail)

        api.POST(reqApi.USER_ROLE_ADD, user.RoleAdd)
        api.POST(reqApi.USER_ROLE_UPDATE, user.RoleUpdate)
        api.GET(reqApi.USER_ROLE_LIST, user.RoleList)
        api.GET(reqApi.USER_ROLE_DETAIL, user.RoleDetail)
        api.POST(reqApi.USER_ROLE_DELETE, user.RoleDelete)
        api.POST(reqApi.USER_ADD, user.UserAdd)
        api.POST(reqApi.USER_UPDATE, user.UserUpdate)
        api.GET(reqApi.USER_LIST, user.UserList)
        api.GET(reqApi.USER_EXISTS, user.UserExists)
        api.GET(reqApi.USER_DETAIL, user.UserDetail)
        api.POST(reqApi.USER_DELETE, user.UserDelete)

        api.POST(reqApi.PROJECT_SPACE_ADD, project.SpaceAdd)
        api.POST(reqApi.PROJECT_SPACE_UPDATE, project.SpaceUpdate)
        api.GET(reqApi.PROJECT_SPACE_LIST, project.SpaceList)
        api.GET(reqApi.PROJECT_SPACE_DETAIL, project.SpaceDetail)
        api.POST(reqApi.PROJECT_SPACE_DELETE, project.SpaceDelete)
        api.GET(reqApi.PROJECT_MEMBER_SEARCH, project.MemberSearch)
        api.POST(reqApi.PROJECT_MEMBER_ADD, project.MemberAdd)
        api.GET(reqApi.PROJECT_MEMBER_LIST, project.MemberList)
        api.POST(reqApi.PROJECT_MEMBER_REMOVE, project.MemberRemove)
        api.POST(reqApi.PROJECT_ADD, project.ProjectAdd)
        api.POST(reqApi.PROJECT_UPDATE, project.ProjectUpdate)
        api.GET(reqApi.PROJECT_LIST, project.ProjectList)
        api.POST(reqApi.PROJECT_SWITCHSTATUS, project.ProjectSwitchStatus)
        api.GET(reqApi.PROJECT_DETAIL, project.ProjectDetail)
        api.POST(reqApi.PROJECT_DELETE, project.ProjectDelete)
        api.POST(reqApi.PROJECT_BUILDSCRIPT, project.ProjectBuildScript)

        api.GET(reqApi.DEPLOY_APPLY_PROJECT_DETAIL, deploy.ApplyProjectDetail)
		api.POST(reqApi.DEPLOY_APPLY_SUBMIT, deploy.ApplySubmit)
		api.GET(reqApi.DEPLOY_APPLY_PROJECT_ALL, deploy.ApplyProjectAll)
		
    }
}