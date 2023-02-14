# PostHog Go SDK

<div align="center">
   <p>The PostHog API allows you to perform any action as if you were an authenticated user utilizing the PostHog UI. It is mostly used for getting data out of PostHog, as well as other private actions such as creating a feature flag.</p>
   <a href="https://github.com/speakeasy-sdks/posthog-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/leapml-python-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
   <a href="https://posthog.com/docs/api"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://discord.com/channels/1065392526745403502/1065392527198404670"><img src="https://img.shields.io/static/v1?label=Slack&message=Join&color=7289da&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/posthog-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/posthog-go-sdk"
    "github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/posthog-go-sdk/pkg/models/operations"
)

func main() {
    s := posthog.New()
    
    req := operations.ActionsCountRetrieveRequest{
        PathParams: operations.ActionsCountRetrievePathParams{
            ID: 548814,
            ProjectID: "deserunt",
        },
        QueryParams: operations.ActionsCountRetrieveQueryParams{
            Format: "undefined",
        },
    }
    
    res, err := s.Actions.ActionsCountRetrieve(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Action != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### Actions

* `ActionsCountRetrieve`
* `ActionsCreate`
* `ActionsDestroy` - Hard delete of this model is not allowed. Use a patch API call to set "deleted" to true
* `ActionsList`
* `ActionsPartialUpdate`
* `ActionsPeopleRetrieve`
* `ActionsRetrieve`
* `ActionsUpdate`

### ActivityLog

* `ActivityLogBookmarkActivityNotificationCreate`
* `ActivityLogImportantChangesRetrieve`

### Annotations

* `AnnotationsCreate` - Create, Read, Update and Delete annotations. [See docs](https://posthog.com/docs/user-guides/annotations) for more information on annotations.
* `AnnotationsDestroy` - Hard delete of this model is not allowed. Use a patch API call to set "deleted" to true
* `AnnotationsList` - Create, Read, Update and Delete annotations. [See docs](https://posthog.com/docs/user-guides/annotations) for more information on annotations.
* `AnnotationsPartialUpdate` - Create, Read, Update and Delete annotations. [See docs](https://posthog.com/docs/user-guides/annotations) for more information on annotations.
* `AnnotationsRetrieve` - Create, Read, Update and Delete annotations. [See docs](https://posthog.com/docs/user-guides/annotations) for more information on annotations.
* `AnnotationsUpdate` - Create, Read, Update and Delete annotations. [See docs](https://posthog.com/docs/user-guides/annotations) for more information on annotations.

### AppMetrics

* `AppMetricsErrorDetailsRetrieve`
* `AppMetricsHistoricalExportsRetrieve`
* `AppMetricsHistoricalExportsRetrieve2`
* `AppMetricsRetrieve`

### Cohorts

* `CohortsCreate`
* `CohortsDestroy` - Hard delete of this model is not allowed. Use a patch API call to set "deleted" to true
* `CohortsList`
* `CohortsPartialUpdate`
* `CohortsPersonsRetrieve`
* `CohortsRetrieve`
* `CohortsUpdate`

### DashboardTemplates

* `DashboardTemplatesCreate`
* `DashboardTemplatesRepositoryRetrieve`

### Dashboards

* `DashboardsCreate`
* `DashboardsDestroy` - Hard delete of this model is not allowed. Use a patch API call to set "deleted" to true
* `DashboardsList`
* `DashboardsMoveTilePartialUpdate`
* `DashboardsPartialUpdate`
* `DashboardsRetrieve`
* `DashboardsUpdate`

### Domains

* `DomainsCreate`
* `DomainsDestroy`
* `DomainsList`
* `DomainsPartialUpdate`
* `DomainsRetrieve`
* `DomainsUpdate`
* `DomainsVerifyCreate`

### EventDefinitions

* `EventDefinitionsList`
* `EventDefinitionsPartialUpdate`
* `EventDefinitionsRetrieve`
* `EventDefinitionsUpdate`

### Events

* `EventsRetrieve`
* `EventsValuesRetrieve`

### Experiments

* `ExperimentsCreate`
* `ExperimentsDestroy`
* `ExperimentsList`
* `ExperimentsPartialUpdate`
* `ExperimentsRequiresFlagImplementationRetrieve`
* `ExperimentsResultsRetrieve`
* `ExperimentsRetrieve`
* `ExperimentsSecondaryResultsRetrieve`
* `ExperimentsUpdate`

### Exports

* `ExportsContentRetrieve`
* `ExportsCreate`
* `ExportsRetrieve`

### FeatureFlags

* `FeatureFlagsActivityRetrieve` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsActivityRetrieve2` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsCreate` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsDestroy` - Hard delete of this model is not allowed. Use a patch API call to set "deleted" to true
* `FeatureFlagsEvaluationReasonsRetrieve` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsList` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsLocalEvaluationRetrieve` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsMyFlagsRetrieve` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsPartialUpdate` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsRetrieve` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsRoleAccessCreate`
* `FeatureFlagsRoleAccessDestroy`
* `FeatureFlagsRoleAccessList`
* `FeatureFlagsRoleAccessRetrieve`
* `FeatureFlagsUpdate` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.
* `FeatureFlagsUserBlastRadiusCreate` - Create, read, update and delete feature flags. [See docs](https://posthog.com/docs/user-guides/feature-flags) for more information on feature flags.

If you're looking to use feature flags on your application, you can either use our JavaScript Library or our dedicated endpoint to check if feature flags are enabled for a given user.

### Funnel

* `Funnels`

### Groups

* `GroupsFindRetrieve`
* `GroupsList`
* `GroupsPropertyDefinitionsRetrieve`
* `GroupsPropertyValuesRetrieve`
* `GroupsRelatedRetrieve`

### GroupsTypes

* `GroupsTypesList`
* `GroupsTypesUpdateMetadataPartialUpdate`

### Hooks

* `HooksCreate` - Retrieve, create, update or destroy REST hooks.
* `HooksDestroy` - Retrieve, create, update or destroy REST hooks.
* `HooksList` - Retrieve, create, update or destroy REST hooks.
* `HooksPartialUpdate` - Retrieve, create, update or destroy REST hooks.
* `HooksRetrieve` - Retrieve, create, update or destroy REST hooks.
* `HooksUpdate` - Retrieve, create, update or destroy REST hooks.

### IngestionWarnings

* `IngestionWarningsRetrieve`

### Insights

* `Funnels`
* `Trends`
* `InsightsActivityRetrieve`
* `InsightsActivityRetrieve2`
* `InsightsCancelCreate`
* `InsightsCreate`
* `InsightsDestroy` - Hard delete of this model is not allowed. Use a patch API call to set "deleted" to true
* `InsightsFunnelCorrelationCreate`
* `InsightsFunnelCorrelationRetrieve`
* `InsightsFunnelRetrieve`
* `InsightsList`
* `InsightsMyLastViewedRetrieve` - Returns basic details about the last 5 insights viewed by this user. Most recently viewed first.
* `InsightsPartialUpdate`
* `InsightsPathCreate`
* `InsightsPathRetrieve`
* `InsightsRetentionRetrieve`
* `InsightsRetrieve`
* `InsightsTimingCreate`
* `InsightsTrendRetrieve`
* `InsightsUpdate`
* `InsightsViewedCreate`

### Integrations

* `IntegrationsChannelsRetrieve`
* `IntegrationsCreate`
* `IntegrationsDestroy`
* `IntegrationsList`
* `IntegrationsRetrieve`

### Invites

* `InvitesBulkCreate`
* `InvitesCreate`
* `InvitesDestroy`
* `InvitesList`

### IsGeneratingDemoData

* `IsGeneratingDemoDataRetrieve` - Projects for the current organization.

### Members

* `MembersDestroy`
* `MembersList`
* `MembersPartialUpdate`
* `MembersUpdate`

### Organizations

* `DomainsCreate`
* `DomainsDestroy`
* `DomainsList`
* `DomainsPartialUpdate`
* `DomainsRetrieve`
* `DomainsUpdate`
* `DomainsVerifyCreate`
* `InvitesBulkCreate`
* `InvitesCreate`
* `InvitesDestroy`
* `InvitesList`
* `MembersDestroy`
* `MembersList`
* `MembersPartialUpdate`
* `MembersUpdate`
* `PluginsActivityRetrieve`
* `PluginsCheckForUpdatesRetrieve`
* `PluginsCreate`
* `PluginsDestroy`
* `PluginsList`
* `PluginsPartialUpdate`
* `PluginsRepositoryRetrieve`
* `PluginsRetrieve`
* `PluginsSourceRetrieve`
* `PluginsUpdate`
* `PluginsUpdateSourcePartialUpdate`
* `PluginsUpgradeCreate`
* `ResourceAccessCreate`
* `ResourceAccessDestroy`
* `ResourceAccessList`
* `ResourceAccessPartialUpdate`
* `ResourceAccessRetrieve`
* `ResourceAccessUpdate`
* `RolesCreate`
* `RolesDestroy`
* `RolesList`
* `RolesPartialUpdate`
* `RolesRetrieve`
* `RolesRoleMembershipsCreate`
* `RolesRoleMembershipsDestroy`
* `RolesRoleMembershipsList`
* `RolesUpdate`

### PerformanceEvents

* `PerformanceEventsList`
* `PerformanceEventsRecentPageviewsRetrieve`

### Persons

* `PersonsActivityRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsActivityRetrieve2` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsCohortsRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsDeletePropertyCreate` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsDestroy` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsFunnelCorrelationCreate` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsFunnelCorrelationRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsFunnelCreate` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsFunnelRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsLifecycleRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsPartialUpdate` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsPathCreate` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsPathRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsPropertiesRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsPropertiesTimelineRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsRetentionRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsSplitCreate` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsStickinessRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsTrendsRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsUpdate` - Only for setting properties on the person. "properties" from the request data will be updated via a "$set" event.
This means that only the properties listed will be updated, but other properties won't be removed nor updated.
If you would like to remove a property use the `delete_property` endpoint.
* `PersonsUpdatePropertyCreate` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.
* `PersonsValuesRetrieve` - To create or update persons, use a PostHog library of your choice and [use an identify call](/docs/integrate/identifying-users). This API endpoint is only for reading and deleting.

### PluginConfigs

* `PluginConfigsCreate`
* `PluginConfigsDestroy`
* `PluginConfigsFrontendRetrieve`
* `PluginConfigsJobCreate`
* `PluginConfigsList`
* `PluginConfigsLogsList`
* `PluginConfigsPartialUpdate`
* `PluginConfigsRearrangePartialUpdate`
* `PluginConfigsRetrieve`
* `PluginConfigsUpdate`

### Plugins

* `PluginsActivityRetrieve`
* `PluginsCheckForUpdatesRetrieve`
* `PluginsCreate`
* `PluginsDestroy`
* `PluginsList`
* `PluginsPartialUpdate`
* `PluginsRepositoryRetrieve`
* `PluginsRetrieve`
* `PluginsSourceRetrieve`
* `PluginsUpdate`
* `PluginsUpdateSourcePartialUpdate`
* `PluginsUpgradeCreate`

### Projects

* `Create` - Projects for the current organization.
* `Destroy` - Projects for the current organization.
* `List` - Projects for the current organization.
* `PartialUpdate` - Projects for the current organization.
* `Retrieve` - Projects for the current organization.
* `Update` - Projects for the current organization.

### Prompts

* `PromptsMyPromptsPartialUpdate` - Create, read, update and delete prompt sequences state for a person.
* `PromptsMyPromptsPartialUpdate` - Create, read, update and delete prompt sequences state for a person.

### PropertyDefinitions

* `PropertyDefinitionsList`
* `PropertyDefinitionsPartialUpdate`
* `PropertyDefinitionsRetrieve`
* `PropertyDefinitionsUpdate`

### Query

* `QueryRetrieve`

### ResetToken

* `ResetTokenPartialUpdate` - Projects for the current organization.

### ResourceAccess

* `ResourceAccessCreate`
* `ResourceAccessDestroy`
* `ResourceAccessList`
* `ResourceAccessPartialUpdate`
* `ResourceAccessRetrieve`
* `ResourceAccessUpdate`

### Roles

* `RolesCreate`
* `RolesDestroy`
* `RolesList`
* `RolesPartialUpdate`
* `RolesRetrieve`
* `RolesRoleMembershipsCreate`
* `RolesRoleMembershipsDestroy`
* `RolesRoleMembershipsList`
* `RolesUpdate`

### SessionRecordingPlaylists

* `SessionRecordingPlaylistsCreate`
* `SessionRecordingPlaylistsDestroy` - Hard delete of this model is not allowed. Use a patch API call to set "deleted" to true
* `SessionRecordingPlaylistsList`
* `SessionRecordingPlaylistsPartialUpdate`
* `SessionRecordingPlaylistsRecordingsCreate`
* `SessionRecordingPlaylistsRecordingsDestroy`
* `SessionRecordingPlaylistsRecordingsRetrieve`
* `SessionRecordingPlaylistsRetrieve`
* `SessionRecordingPlaylistsUpdate`

### SessionRecordings

* `SessionRecordingsPropertiesRetrieve`
* `SessionRecordingsRetrieve`
* `SessionRecordingsRetrieve2`
* `SessionRecordingsSnapshotsRetrieve`

### Subscriptions

* `SubscriptionsCreate`
* `SubscriptionsDestroy` - Hard delete of this model is not allowed. Use a patch API call to set "deleted" to true
* `SubscriptionsList`
* `SubscriptionsPartialUpdate`
* `SubscriptionsRetrieve`
* `SubscriptionsUpdate`

### Tags

* `TagsList`

### Trend

* `Trends`

### UploadedMedia

* `UploadedMediaCreate` - 
    When object storage is available this API allows upload of media which can be used, for example, in text cards on dashboards.

    Uploaded media must have a content type beginning with 'image/' and be less than 4MB.
    
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
