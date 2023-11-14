package repos

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCommitsItemCheckRunsGetResponse 
type ItemItemCommitsItemCheckRunsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The check_runs property
    check_runs []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckRunable
    // The total_count property
    total_count *int32
}
// NewItemItemCommitsItemCheckRunsGetResponse instantiates a new ItemItemCommitsItemCheckRunsGetResponse and sets the default values.
func NewItemItemCommitsItemCheckRunsGetResponse()(*ItemItemCommitsItemCheckRunsGetResponse) {
    m := &ItemItemCommitsItemCheckRunsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCommitsItemCheckRunsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCommitsItemCheckRunsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCommitsItemCheckRunsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCommitsItemCheckRunsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckRuns gets the check_runs property value. The check_runs property
func (m *ItemItemCommitsItemCheckRunsGetResponse) GetCheckRuns()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckRunable) {
    return m.check_runs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCommitsItemCheckRunsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["check_runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCheckRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckRunable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckRunable)
                }
            }
            m.SetCheckRuns(res)
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
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemItemCommitsItemCheckRunsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemCommitsItemCheckRunsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCheckRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCheckRuns()))
        for i, v := range m.GetCheckRuns() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("check_runs", cast)
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
func (m *ItemItemCommitsItemCheckRunsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckRuns sets the check_runs property value. The check_runs property
func (m *ItemItemCommitsItemCheckRunsGetResponse) SetCheckRuns(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckRunable)() {
    m.check_runs = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemCommitsItemCheckRunsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemItemCommitsItemCheckRunsGetResponseable 
type ItemItemCommitsItemCheckRunsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckRuns()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckRunable)
    GetTotalCount()(*int32)
    SetCheckRuns(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CheckRunable)()
    SetTotalCount(value *int32)()
}