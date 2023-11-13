package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iea7adb0a6f1f927ec27be9b310739e6399457d01132530bc9be2e75dbde0ec47 "github.com/octokit/go-sdk/repos/item/item/milestones"
)

// ItemItemMilestonesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\milestones
type ItemItemMilestonesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemMilestonesRequestBuilderGetQueryParameters lists milestones for a repository.
type ItemItemMilestonesRequestBuilderGetQueryParameters struct {
    // The direction of the sort. Either `asc` or `desc`.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction of the sort. Either `asc` or `desc`.
    DirectionAsGetDirectionQueryParameterType *iea7adb0a6f1f927ec27be9b310739e6399457d01132530bc9be2e75dbde0ec47.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // What to sort results by. Either `due_on` or `completeness`.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // What to sort results by. Either `due_on` or `completeness`.
    SortAsGetSortQueryParameterType *iea7adb0a6f1f927ec27be9b310739e6399457d01132530bc9be2e75dbde0ec47.GetSortQueryParameterType `uriparametername:"sort"`
    // The state of the milestone. Either `open`, `closed`, or `all`.
    // Deprecated: This property is deprecated, use stateAsGetStateQueryParameterType instead
    State *string `uriparametername:"state"`
    // The state of the milestone. Either `open`, `closed`, or `all`.
    StateAsGetStateQueryParameterType *iea7adb0a6f1f927ec27be9b310739e6399457d01132530bc9be2e75dbde0ec47.GetStateQueryParameterType `uriparametername:"state"`
}
// ItemItemMilestonesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemMilestonesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemMilestonesRequestBuilderGetQueryParameters
}
// ItemItemMilestonesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemMilestonesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMilestone_number gets an item from the go-sdk.repos.item.item.milestones.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemMilestonesRequestBuilder) ByMilestone_number(milestone_number string)(*ItemItemMilestonesWithMilestone_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if milestone_number != "" {
        urlTplParams["milestone_number"] = milestone_number
    }
    return NewItemItemMilestonesWithMilestone_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByMilestone_numberInteger gets an item from the go-sdk.repos.item.item.milestones.item collection
func (m *ItemItemMilestonesRequestBuilder) ByMilestone_numberInteger(milestone_number int32)(*ItemItemMilestonesWithMilestone_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["milestone_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(milestone_number), 10)
    return NewItemItemMilestonesWithMilestone_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemMilestonesRequestBuilderInternal instantiates a new MilestonesRequestBuilder and sets the default values.
func NewItemItemMilestonesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMilestonesRequestBuilder) {
    m := &ItemItemMilestonesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/milestones{?state*,sort*,direction*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemMilestonesRequestBuilder instantiates a new MilestonesRequestBuilder and sets the default values.
func NewItemItemMilestonesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMilestonesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemMilestonesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists milestones for a repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/milestones#list-milestones
func (m *ItemItemMilestonesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemMilestonesRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Milestoneable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateMilestoneFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Milestoneable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Milestoneable)
        }
    }
    return val, nil
}
// Post creates a milestone.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/milestones#create-a-milestone
func (m *ItemItemMilestonesRequestBuilder) Post(ctx context.Context, body ItemItemMilestonesPostRequestBodyable, requestConfiguration *ItemItemMilestonesRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Milestoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateMilestoneFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Milestoneable), nil
}
// ToGetRequestInformation lists milestones for a repository.
func (m *ItemItemMilestonesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemMilestonesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPostRequestInformation creates a milestone.
func (m *ItemItemMilestonesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemMilestonesPostRequestBodyable, requestConfiguration *ItemItemMilestonesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemMilestonesRequestBuilder) WithUrl(rawUrl string)(*ItemItemMilestonesRequestBuilder) {
    return NewItemItemMilestonesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}