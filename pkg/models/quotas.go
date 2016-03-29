package models

import (
	"errors"
	"github.com/grafana/grafana/pkg/setting"
	"time"
)

var ErrInvalidQuotaTarget = errors.New("Invalid quota target")

type Quota struct {
	Id      int64
	OrgId   int64
	UserId  int64
	Target  string
	Limit   int64
	Created time.Time
	Updated time.Time
}

type QuotaScope struct {
	Name         string
	Target       string
	DefaultLimit int64
}

type OrgQuotaDTO struct {
	OrgId  int64  `json:"org_id"`
	Target string `json:"target"`
	Limit  int64  `json:"limit"`
	Used   int64  `json:"used"`
}

type GlobalQuotaDTO struct {
	Target string `json:"target"`
	Limit  int64  `json:"limit"`
	Used   int64  `json:"used"`
}

type GetOrgQuotaByTargetQuery struct {
	Target  string
	OrgId   int64
	Default int64
	Result  *OrgQuotaDTO
}

type GetOrgQuotasQuery struct {
	OrgId  int64
	Result []*OrgQuotaDTO
}

type GetGlobalQuotaByTargetQuery struct {
	Target  string
	Default int64
	Result  *GlobalQuotaDTO
}

type UpdateOrgQuotaCmd struct {
	Target string `json:"target"`
	Limit  int64  `json:"limit"`
	OrgId  int64  `json:"-"`
}

func GetQuotaScopes(target string) ([]QuotaScope, error) {
	scopes := make([]QuotaScope, 0)
	switch target {
	case "endpoint":
		scopes = append(scopes,
			QuotaScope{Name: "global", Target: target, DefaultLimit: setting.Quota.Global.Endpoint},
			QuotaScope{Name: "org", Target: target, DefaultLimit: setting.Quota.Org.Endpoint},
		)
		return scopes, nil
	case "collector":
		scopes = append(scopes,
			QuotaScope{Name: "global", Target: target, DefaultLimit: setting.Quota.Global.Collector},
			QuotaScope{Name: "org", Target: target, DefaultLimit: setting.Quota.Org.Collector},
		)
		return scopes, nil
	default:
		return scopes, ErrInvalidQuotaTarget
	}
}
