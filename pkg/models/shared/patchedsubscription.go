package shared

import (
	"time"
)

type PatchedSubscriptionByweekdayEnum string

const (
	PatchedSubscriptionByweekdayEnumMonday    PatchedSubscriptionByweekdayEnum = "monday"
	PatchedSubscriptionByweekdayEnumTuesday   PatchedSubscriptionByweekdayEnum = "tuesday"
	PatchedSubscriptionByweekdayEnumWednesday PatchedSubscriptionByweekdayEnum = "wednesday"
	PatchedSubscriptionByweekdayEnumThursday  PatchedSubscriptionByweekdayEnum = "thursday"
	PatchedSubscriptionByweekdayEnumFriday    PatchedSubscriptionByweekdayEnum = "friday"
	PatchedSubscriptionByweekdayEnumSaturday  PatchedSubscriptionByweekdayEnum = "saturday"
	PatchedSubscriptionByweekdayEnumSunday    PatchedSubscriptionByweekdayEnum = "sunday"
)

type PatchedSubscriptionFrequencyEnum string

const (
	PatchedSubscriptionFrequencyEnumDaily   PatchedSubscriptionFrequencyEnum = "daily"
	PatchedSubscriptionFrequencyEnumWeekly  PatchedSubscriptionFrequencyEnum = "weekly"
	PatchedSubscriptionFrequencyEnumMonthly PatchedSubscriptionFrequencyEnum = "monthly"
	PatchedSubscriptionFrequencyEnumYearly  PatchedSubscriptionFrequencyEnum = "yearly"
)

type PatchedSubscriptionTargetTypeEnum string

const (
	PatchedSubscriptionTargetTypeEnumEmail   PatchedSubscriptionTargetTypeEnum = "email"
	PatchedSubscriptionTargetTypeEnumSlack   PatchedSubscriptionTargetTypeEnum = "slack"
	PatchedSubscriptionTargetTypeEnumWebhook PatchedSubscriptionTargetTypeEnum = "webhook"
)

// PatchedSubscriptionInput
// Standard Subscription serializer.
type PatchedSubscriptionInput struct {
	Bysetpos      *int64                             `json:"bysetpos,omitempty" form:"name=bysetpos" multipartForm:"name=bysetpos"`
	Byweekday     []PatchedSubscriptionByweekdayEnum `json:"byweekday,omitempty" form:"name=byweekday,json" multipartForm:"name=byweekday,json"`
	Count         *int64                             `json:"count,omitempty" form:"name=count" multipartForm:"name=count"`
	Dashboard     *int64                             `json:"dashboard,omitempty" form:"name=dashboard" multipartForm:"name=dashboard"`
	Deleted       *bool                              `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Frequency     *PatchedSubscriptionFrequencyEnum  `json:"frequency,omitempty" form:"name=frequency" multipartForm:"name=frequency"`
	Insight       *int64                             `json:"insight,omitempty" form:"name=insight" multipartForm:"name=insight"`
	Interval      *int64                             `json:"interval,omitempty" form:"name=interval" multipartForm:"name=interval"`
	InviteMessage *string                            `json:"invite_message,omitempty" form:"name=invite_message" multipartForm:"name=invite_message"`
	StartDate     *time.Time                         `json:"start_date,omitempty" form:"name=start_date" multipartForm:"name=start_date"`
	TargetType    *PatchedSubscriptionTargetTypeEnum `json:"target_type,omitempty" form:"name=target_type" multipartForm:"name=target_type"`
	TargetValue   *string                            `json:"target_value,omitempty" form:"name=target_value" multipartForm:"name=target_value"`
	Title         *string                            `json:"title,omitempty" form:"name=title" multipartForm:"name=title"`
	UntilDate     *time.Time                         `json:"until_date,omitempty" form:"name=until_date" multipartForm:"name=until_date"`
}
