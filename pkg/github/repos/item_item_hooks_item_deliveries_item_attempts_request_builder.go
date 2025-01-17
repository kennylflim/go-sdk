package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\hooks\{hook_id}\deliveries\{delivery_id}\attempts
type ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal instantiates a new AttemptsRequestBuilder and sets the default values.
func NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    m := &ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/hooks/{hook_id}/deliveries/{delivery_id}/attempts", pathParameters),
    }
    return m
}
// NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilder instantiates a new AttemptsRequestBuilder and sets the default values.
func NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post redeliver a webhook delivery for a webhook configured in a repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/webhooks#redeliver-a-delivery-for-a-repository-webhook
func (m *ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemItemHooksItemDeliveriesItemAttemptsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemHooksItemDeliveriesItemAttemptsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemHooksItemDeliveriesItemAttemptsPostResponseable), nil
}
// ToPostRequestInformation redeliver a webhook delivery for a webhook configured in a repository.
func (m *ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) WithUrl(rawUrl string)(*ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    return NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
