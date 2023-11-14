package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPublic_membersRequestBuilder builds and executes requests for operations under \orgs\{org}\public_members
type ItemPublic_membersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPublic_membersRequestBuilderGetQueryParameters members of an organization can choose to have their membership publicized or not.
type ItemPublic_membersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemPublic_membersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPublic_membersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPublic_membersRequestBuilderGetQueryParameters
}
// ByUsername gets an item from the octokit.orgs.item.public_members.item collection
func (m *ItemPublic_membersRequestBuilder) ByUsername(username string)(*ItemPublic_membersWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemPublic_membersWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPublic_membersRequestBuilderInternal instantiates a new Public_membersRequestBuilder and sets the default values.
func NewItemPublic_membersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPublic_membersRequestBuilder) {
    m := &ItemPublic_membersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/public_members{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemPublic_membersRequestBuilder instantiates a new Public_membersRequestBuilder and sets the default values.
func NewItemPublic_membersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPublic_membersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPublic_membersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get members of an organization can choose to have their membership publicized or not.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/members#list-public-organization-members
func (m *ItemPublic_membersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPublic_membersRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateSimpleUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SimpleUserable)
        }
    }
    return val, nil
}
// ToGetRequestInformation members of an organization can choose to have their membership publicized or not.
func (m *ItemPublic_membersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPublic_membersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPublic_membersRequestBuilder) WithUrl(rawUrl string)(*ItemPublic_membersRequestBuilder) {
    return NewItemPublic_membersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}