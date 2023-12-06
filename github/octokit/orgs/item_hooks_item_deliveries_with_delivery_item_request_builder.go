package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\hooks\{hook_id}\deliveries\{delivery_id}
type ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attempts the attempts property
func (m *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) Attempts()(*ItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    return NewItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal instantiates a new WithDelivery_ItemRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    m := &ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}", pathParameters),
    }
    return m
}
// NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder instantiates a new WithDelivery_ItemRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a delivery for a webhook configured in an organization.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/webhooks#get-a-webhook-delivery-for-an-organization-webhook
func (m *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.HookDeliveryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateHookDeliveryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.HookDeliveryable), nil
}
// ToGetRequestInformation returns a delivery for a webhook configured in an organization.
func (m *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    return NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
