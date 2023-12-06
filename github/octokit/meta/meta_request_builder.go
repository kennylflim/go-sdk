package meta

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// MetaRequestBuilder builds and executes requests for operations under \meta
type MetaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MetaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MetaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMetaRequestBuilderInternal instantiates a new MetaRequestBuilder and sets the default values.
func NewMetaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetaRequestBuilder) {
    m := &MetaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/meta", pathParameters),
    }
    return m
}
// NewMetaRequestBuilder instantiates a new MetaRequestBuilder and sets the default values.
func NewMetaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns meta information about GitHub, including a list of GitHub's IP addresses. For more information, see "[About GitHub's IP addresses](https://docs.github.com/articles/about-github-s-ip-addresses/)."The API's response also includes a list of GitHub's domain names.The values shown in the documentation's response are example values. You must always query the API directly to get the latest values.**Note:** This endpoint returns both IPv4 and IPv6 addresses. However, not all features support IPv6. You should refer to the specific documentation for each feature to determine if IPv6 is supported.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/meta/meta#get-apiname-meta-information
func (m *MetaRequestBuilder) Get(ctx context.Context, requestConfiguration *MetaRequestBuilderGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ApiOverviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateApiOverviewFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ApiOverviewable), nil
}
// ToGetRequestInformation returns meta information about GitHub, including a list of GitHub's IP addresses. For more information, see "[About GitHub's IP addresses](https://docs.github.com/articles/about-github-s-ip-addresses/)."The API's response also includes a list of GitHub's domain names.The values shown in the documentation's response are example values. You must always query the API directly to get the latest values.**Note:** This endpoint returns both IPv4 and IPv6 addresses. However, not all features support IPv6. You should refer to the specific documentation for each feature to determine if IPv6 is supported.
func (m *MetaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MetaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MetaRequestBuilder) WithUrl(rawUrl string)(*MetaRequestBuilder) {
    return NewMetaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
