package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitCommit_tree 
type GitCommit_tree struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // SHA for the commit
    sha *string
    // The url property
    url *string
}
// NewGitCommit_tree instantiates a new gitCommit_tree and sets the default values.
func NewGitCommit_tree()(*GitCommit_tree) {
    m := &GitCommit_tree{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGitCommit_treeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGitCommit_treeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitCommit_tree(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GitCommit_tree) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GitCommit_tree) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetSha gets the sha property value. SHA for the commit
func (m *GitCommit_tree) GetSha()(*string) {
    return m.sha
}
// GetUrl gets the url property value. The url property
func (m *GitCommit_tree) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *GitCommit_tree) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *GitCommit_tree) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSha sets the sha property value. SHA for the commit
func (m *GitCommit_tree) SetSha(value *string)() {
    m.sha = value
}
// SetUrl sets the url property value. The url property
func (m *GitCommit_tree) SetUrl(value *string)() {
    m.url = value
}
// GitCommit_treeable 
type GitCommit_treeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSha()(*string)
    GetUrl()(*string)
    SetSha(value *string)()
    SetUrl(value *string)()
}