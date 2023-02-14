package shared

import (
	"time"
)

type SubscriptionByweekdayEnum string

const (
	SubscriptionByweekdayEnumMonday    SubscriptionByweekdayEnum = "monday"
	SubscriptionByweekdayEnumTuesday   SubscriptionByweekdayEnum = "tuesday"
	SubscriptionByweekdayEnumWednesday SubscriptionByweekdayEnum = "wednesday"
	SubscriptionByweekdayEnumThursday  SubscriptionByweekdayEnum = "thursday"
	SubscriptionByweekdayEnumFriday    SubscriptionByweekdayEnum = "friday"
	SubscriptionByweekdayEnumSaturday  SubscriptionByweekdayEnum = "saturday"
	SubscriptionByweekdayEnumSunday    SubscriptionByweekdayEnum = "sunday"
)

type SubscriptionFrequencyEnum string

const (
	SubscriptionFrequencyEnumDaily   SubscriptionFrequencyEnum = "daily"
	SubscriptionFrequencyEnumWeekly  SubscriptionFrequencyEnum = "weekly"
	SubscriptionFrequencyEnumMonthly SubscriptionFrequencyEnum = "monthly"
	SubscriptionFrequencyEnumYearly  SubscriptionFrequencyEnum = "yearly"
)

type SubscriptionTargetTypeEnum string

const (
	SubscriptionTargetTypeEnumEmail   SubscriptionTargetTypeEnum = "email"
	SubscriptionTargetTypeEnumSlack   SubscriptionTargetTypeEnum = "slack"
	SubscriptionTargetTypeEnumWebhook SubscriptionTargetTypeEnum = "webhook"
)

// SubscriptionInput
// Standard Subscription serializer.
type SubscriptionInput struct {
	Bysetpos      *int64                      `json:"bysetpos,omitempty" form:"name=bysetpos" multipartForm:"name=bysetpos"`
	Byweekday     []SubscriptionByweekdayEnum `json:"byweekday,omitempty" form:"name=byweekday,json" multipartForm:"name=byweekday,json"`
	Count         *int64                      `json:"count,omitempty" form:"name=count" multipartForm:"name=count"`
	Dashboard     *int64                      `json:"dashboard,omitempty" form:"name=dashboard" multipartForm:"name=dashboard"`
	Deleted       *bool                       `json:"deleted,omitempty" form:"name=deleted" multipartForm:"name=deleted"`
	Frequency     SubscriptionFrequencyEnum   `json:"frequency" form:"name=frequency" multipartForm:"name=frequency"`
	Insight       *int64                      `json:"insight,omitempty" form:"name=insight" multipartForm:"name=insight"`
	Interval      *int64                      `json:"interval,omitempty" form:"name=interval" multipartForm:"name=interval"`
	InviteMessage *string                     `json:"invite_message,omitempty" form:"name=invite_message" multipartForm:"name=invite_message"`
	StartDate     time.Time                   `json:"start_date" form:"name=start_date" multipartForm:"name=start_date"`
	TargetType    SubscriptionTargetTypeEnum  `json:"target_type" form:"name=target_type" multipartForm:"name=target_type"`
	TargetValue   string                      `json:"target_value" form:"name=target_value" multipartForm:"name=target_value"`
	Title         *string                     `json:"title,omitempty" form:"name=title" multipartForm:"name=title"`
	UntilDate     *time.Time                  `json:"until_date,omitempty" form:"name=until_date" multipartForm:"name=until_date"`
}

type SubscriptionCreatedBy struct {
	DistinctID *string `json:"distinct_id,omitempty"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	ID         int64   `json:"id"`
	UUID       string  `json:"uuid"`
}

// Subscription
// Standard Subscription serializer.
type Subscription struct {
	Bysetpos         *int64                      `json:"bysetpos,omitempty"`
	Byweekday        []SubscriptionByweekdayEnum `json:"byweekday,omitempty"`
	Count            *int64                      `json:"count,omitempty"`
	CreatedAt        time.Time                   `json:"created_at"`
	CreatedBy        SubscriptionCreatedBy       `json:"created_by"`
	Dashboard        *int64                      `json:"dashboard,omitempty"`
	Deleted          *bool                       `json:"deleted,omitempty"`
	Frequency        SubscriptionFrequencyEnum   `json:"frequency"`
	ID               int64                       `json:"id"`
	Insight          *int64                      `json:"insight,omitempty"`
	Interval         *int64                      `json:"interval,omitempty"`
	InviteMessage    *string                     `json:"invite_message,omitempty"`
	NextDeliveryDate time.Time                   `json:"next_delivery_date"`
	StartDate        time.Time                   `json:"start_date"`
	Summary          string                      `json:"summary"`
	TargetType       SubscriptionTargetTypeEnum  `json:"target_type"`
	TargetValue      string                      `json:"target_value"`
	Title            *string                     `json:"title,omitempty"`
	UntilDate        *time.Time                  `json:"until_date,omitempty"`
}
