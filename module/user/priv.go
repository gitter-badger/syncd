// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
    "github.com/dreamans/syncd/util/goslice"
    reqApi "github.com/dreamans/syncd/router/route/api"
)

var apiToPrivMap = map[string][]int{}

func init() {
    for priv, apiList := range privToApiMap{
        for _, api := range apiList {
            privMap, _ := apiToPrivMap[api]
            apiToPrivMap[api] = append(privMap, priv)
        }
    }
}

func CheckHavePriv(api string, priv []int) bool {
    privMap, exists := apiToPrivMap[api]
    if !exists {
        return false
    }
    return len(goslice.SliceIntersectInt(privMap, priv)) > 0
}

func PrivIn(privCode int, privList []int) bool {
    return goslice.InSliceInt(privCode, privList)
}

const (
    DEPLOY_APPLY      = 1001 // 填写上线单
    DEPLOY_VIEW_MY    = 1002 // 查看上线单(自己)
    DEPLOY_VIEW_ALL   = 1003 // 查看上线单(全部)
    DEPLOY_AUDIT_MY   = 1004 // 审核上线单(自己)
    DEPLOY_AUDIT_ALL  = 1005 // 审核上线单(全部)
    DEPLOY_DEPLOY_MY  = 1006 // 上线操作(自己)
    DEPLOY_DEPLOY_ALL = 1007 // 上线操作(全部)
    DEPLOY_DROP_MY    = 1008 // 废弃上线单(自己)
    DEPLOY_DROP_ALL   = 1009 // 废弃上线单(全部)
    DEPLOY_EDIT_MY    = 1010 // 编辑上线单(自己)
    DEPLOY_EDIT_ALL   = 1011 // 编辑上线单(全部)

    PROJECT_SPACE_VIEW  = 2001 // 查看空间
    PROJECT_SPACE_NEW   = 2002 // 新增空间
    PROJECT_SPACE_EDIT  = 2003 // 编辑空间
    PROJECT_SPACE_DEL   = 2004 // 删除空间
    PROJECT_USER_VIEW = 2100 // 查看成员
    PROJECT_USER_NEW  = 2101 // 新增成员
    PROJECT_USER_DEL  = 2102 // 删除成员
    PROJECT_VIEW  = 2201 // 查看项目
    PROJECT_NEW   = 2202 // 新增项目
    PROJECT_EDIT  = 2203 // 编辑项目
    PROJECT_DEL   = 2204 // 删除项目
    PROJECT_AUDIT = 2205 // 启用项目
    PROJECT_BUILD = 2206 // 构建设置

    USER_ROLE_VIEW = 3001 // 查看角色
    USER_ROLE_NEW  = 3002 // 新增角色
    USER_ROLE_EDIT = 3003 // 编辑角色
    USER_ROLE_DEL  = 3004 // 删除角色
    USER_VIEW = 3101 // 查看用户
    USER_NEW  = 3102 // 新增用户
    USER_EDIT = 3103 // 编辑用户
    USER_DEL  = 3104 // 删除用户

    SERVER_GROUP_VIEW = 4001 // 查看集群
    SERVER_GROUP_NEW  = 4002 // 新增集群
    SERVER_GROUP_EDIT = 4003 // 编辑集群
    SERVER_GROUP_DEL  = 4004 // 删除集群
    SERVER_VIEW  = 4101 // 查看服务器
    SERVER_NEW   = 4102 // 新增服务器
    SERVER_EDIT  = 4103 // 编辑服务器
    SERVER_DEL   = 4104 // 删除服务器
)

var privToApiMap = map[int][]string{
    SERVER_GROUP_VIEW: []string{
        reqApi.SERVER_GROUP_LIST,
    },
    SERVER_GROUP_NEW: []string{
        reqApi.SERVER_GROUP_ADD,
    },
    SERVER_GROUP_EDIT: []string{
        reqApi.SERVER_GROUP_DETAIL,
        reqApi.SERVER_GROUP_UPDATE,
    },
    SERVER_GROUP_DEL: []string{
        reqApi.SERVER_GROUP_DELETE,
    },
    SERVER_VIEW: []string{
        reqApi.SERVER_LIST,
    },
    SERVER_NEW: []string{
        reqApi.SERVER_GROUP_LIST,
        reqApi.SERVER_ADD,
    },
    SERVER_EDIT: []string{
        reqApi.SERVER_DETAIL,
        reqApi.SERVER_UPDATE,
    },
    SERVER_DEL: []string{
        reqApi.SERVER_DELETE,
    },
    USER_ROLE_VIEW: []string{
        reqApi.USER_ROLE_LIST,
    },
    USER_ROLE_NEW: []string{
        reqApi.USER_ROLE_ADD,
    },
    USER_ROLE_EDIT: []string{
        reqApi.USER_ROLE_DETAIL,
        reqApi.USER_ROLE_UPDATE,
    },
    USER_ROLE_DEL: []string{
        reqApi.USER_ROLE_DELETE,
    },
    USER_VIEW: []string{
        reqApi.USER_LIST,
    },
    USER_NEW: []string{
        reqApi.USER_ADD,
        reqApi.USER_ROLE_LIST,
        reqApi.USER_EXISTS,
    },
    USER_EDIT: []string{
        reqApi.USER_ROLE_LIST,
        reqApi.USER_DETAIL,
        reqApi.USER_UPDATE,
        reqApi.USER_EXISTS,
    },
    USER_DEL: []string{
        reqApi.USER_DELETE,
    },
    PROJECT_SPACE_VIEW: []string{
        reqApi.PROJECT_SPACE_LIST,
    },
    PROJECT_SPACE_NEW: []string{
        reqApi.PROJECT_SPACE_ADD,
    },
    PROJECT_SPACE_EDIT: []string{
        reqApi.PROJECT_SPACE_DETAIL,
        reqApi.PROJECT_SPACE_UPDATE,
    },
    PROJECT_SPACE_DEL: []string{
        reqApi.PROJECT_SPACE_DELETE,
    },
    PROJECT_USER_VIEW: []string{
        reqApi.PROJECT_SPACE_DETAIL,
        reqApi.PROJECT_SPACE_LIST,
        reqApi.PROJECT_MEMBER_LIST,
    },
    PROJECT_USER_NEW: []string{
        reqApi.PROJECT_MEMBER_ADD,
        reqApi.PROJECT_MEMBER_SEARCH,
    },
    PROJECT_USER_DEL: []string{
        reqApi.PROJECT_MEMBER_REMOVE,
    },
    PROJECT_VIEW: []string{
        reqApi.SERVER_GROUP_LIST,
        reqApi.PROJECT_SPACE_LIST,
        reqApi.PROJECT_SPACE_DETAIL,
        reqApi.PROJECT_LIST,
        reqApi.PROJECT_DETAIL,
    },
    PROJECT_NEW: []string{
        reqApi.SERVER_GROUP_LIST,
        reqApi.PROJECT_ADD,
    },
    PROJECT_EDIT: []string{
        reqApi.SERVER_GROUP_LIST,
        reqApi.PROJECT_DETAIL,
        reqApi.PROJECT_UPDATE,
    },
    PROJECT_DEL: []string{
        reqApi.PROJECT_DELETE,
    },
    PROJECT_AUDIT: []string{
        reqApi.PROJECT_SWITCHSTATUS,
    },
    PROJECT_BUILD: []string{
        reqApi.PROJECT_DETAIL,
        reqApi.PROJECT_BUILDSCRIPT,
    },
    DEPLOY_APPLY: []string{
        reqApi.PROJECT_SPACE_LIST,
        reqApi.PROJECT_LIST,
        reqApi.DEPLOY_APPLY_PROJECT_DETAIL,
        reqApi.DEPLOY_APPLY_SUBMIT,
    },
    DEPLOY_VIEW_MY: []string{
        reqApi.DEPLOY_APPLY_PROJECT_ALL,
        //syncd.API_DEPLOY_APPLY_LIST,
        // syncd.API_DEPLOY_APPLY_DETAIL,
        //syncd.API_DEPLOY_APPLY_PROJECT_ALL,
        //syncd.API_DEPLOY_APPLY_LOG,
    },
    DEPLOY_VIEW_ALL: []string{
        reqApi.DEPLOY_APPLY_PROJECT_ALL,
        //syncd.API_DEPLOY_APPLY_LIST,
        //syncd.API_DEPLOY_APPLY_DETAIL,
        //syncd.API_DEPLOY_APPLY_PROJECT_ALL,
        //syncd.API_DEPLOY_APPLY_LOG,
    },
    /*
    DEPLOY_AUDIT_MY: []string{
        syncd.API_DEPLOY_APPLY_AUDIT,
        syncd.API_DEPLOY_APPLY_UNAUDIT,
    },
    DEPLOY_AUDIT_ALL: []string{
        syncd.API_DEPLOY_APPLY_AUDIT,
        syncd.API_DEPLOY_APPLY_UNAUDIT,
    },
    DEPLOY_DROP_MY: []string{
        syncd.API_DEPLOY_APPLY_DISCARD,
    },
    DEPLOY_DROP_ALL: []string{
        syncd.API_DEPLOY_APPLY_DISCARD,
    },
    DEPLOY_EDIT_MY: []string{
        syncd.API_DEPLOY_APPLY_PROJECT_DETAIL,
        syncd.API_DEPLOY_APPLY_TAGLIST,
        syncd.API_DEPLOY_APPLY_COMMITLIST,
        syncd.API_DEPLOY_APPLY_UPDATE,
    },
    DEPLOY_DEPLOY_MY: []string{
        syncd.API_DEPLOY_DEPLOY_START,
        syncd.API_DEPLOY_APPLY_DETAIL,
        syncd.API_DEPLOY_DEPLOY_STATUS,
        syncd.API_DEPLOY_DEPLOY_STOP,
    },
    DEPLOY_DEPLOY_ALL: []string{
        syncd.API_DEPLOY_DEPLOY_START,
        syncd.API_DEPLOY_APPLY_DETAIL,
        syncd.API_DEPLOY_DEPLOY_STATUS,
        syncd.API_DEPLOY_DEPLOY_STOP,
    },
    */
}
