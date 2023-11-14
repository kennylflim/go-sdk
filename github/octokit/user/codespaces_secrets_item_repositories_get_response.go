package user

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesSecretsItemRepositoriesGetResponse 
type CodespacesSecretsItemRepositoriesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repositories property
    repositories []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MinimalRepositoryable
    // The total_count property
    total_count *int32
}
// NewCodespacesSecretsItemRepositoriesGetResponse instantiates a new CodespacesSecretsItemRepositoriesGetResponse and sets the default values.
func NewCodespacesSecretsItemRepositoriesGetResponse()(*CodespacesSecretsItemRepositoriesGetResponse) {
    m := &CodespacesSecretsItemRepositoriesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesSecretsItemRepositoriesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesSecretsItemRepositoriesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesSecretsItemRepositoriesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodespacesSecretsItemRepositoriesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CodespacesSecretsItemRepositoriesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MinimalRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MinimalRepositoryable)
                }
            }
            m.SetRepositories(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetRepositories gets the repositories property value. The repositories property
func (m *CodespacesSecretsItemRepositoriesGetResponse) GetRepositories()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MinimalRepositoryable) {
    return m.repositories
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *CodespacesSecretsItemRepositoriesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *CodespacesSecretsItemRepositoriesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositories()))
        for i, v := range m.GetRepositories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("repositories", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodespacesSecretsItemRepositoriesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositories sets the repositories property value. The repositories property
func (m *CodespacesSecretsItemRepositoriesGetResponse) SetRepositories(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MinimalRepositoryable)() {
    m.repositories = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *CodespacesSecretsItemRepositoriesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// CodespacesSecretsItemRepositoriesGetResponseable 
type CodespacesSecretsItemRepositoriesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepositories()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MinimalRepositoryable)
    GetTotalCount()(*int32)
    SetRepositories(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MinimalRepositoryable)()
    SetTotalCount(value *int32)()
}