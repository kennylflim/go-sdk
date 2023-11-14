package orgs

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemRulesetsItemWithRuleset_PutRequestBody 
type ItemRulesetsItemWithRuleset_PutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The actors that can bypass the rules in this ruleset
    bypass_actors []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRulesetBypassActorable
    // Conditions for an organization ruleset. The conditions object should contain both `repository_name` and `ref_name` properties or both `repository_id` and `ref_name` properties.
    conditions i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRulesetConditionsable
    // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
    enforcement *i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleEnforcement
    // The name of the ruleset.
    name *string
    // An array of rules within the ruleset.
    rules []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleable
}
// NewItemRulesetsItemWithRuleset_PutRequestBody instantiates a new ItemRulesetsItemWithRuleset_PutRequestBody and sets the default values.
func NewItemRulesetsItemWithRuleset_PutRequestBody()(*ItemRulesetsItemWithRuleset_PutRequestBody) {
    m := &ItemRulesetsItemWithRuleset_PutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemRulesetsItemWithRuleset_PutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemRulesetsItemWithRuleset_PutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRulesetsItemWithRuleset_PutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassActors gets the bypass_actors property value. The actors that can bypass the rules in this ruleset
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) GetBypassActors()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRulesetBypassActorable) {
    return m.bypass_actors
}
// GetConditions gets the conditions property value. Conditions for an organization ruleset. The conditions object should contain both `repository_name` and `ref_name` properties or both `repository_id` and `ref_name` properties.
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) GetConditions()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRulesetConditionsable) {
    return m.conditions
}
// GetEnforcement gets the enforcement property value. The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) GetEnforcement()(*i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleEnforcement) {
    return m.enforcement
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bypass_actors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRepositoryRulesetBypassActorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRulesetBypassActorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRulesetBypassActorable)
                }
            }
            m.SetBypassActors(res)
        }
        return nil
    }
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrgRulesetConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRulesetConditionsable))
        }
        return nil
    }
    res["enforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ParseRepositoryRuleEnforcement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcement(val.(*i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleEnforcement))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRepositoryRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleable)
                }
            }
            m.SetRules(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the ruleset.
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) GetName()(*string) {
    return m.name
}
// GetRules gets the rules property value. An array of rules within the ruleset.
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) GetRules()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleable) {
    return m.rules
}
// Serialize serializes information the current object
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBypassActors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBypassActors()))
        for i, v := range m.GetBypassActors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("bypass_actors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcement() != nil {
        cast := (*m.GetEnforcement()).String()
        err := writer.WriteStringValue("enforcement", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
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
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassActors sets the bypass_actors property value. The actors that can bypass the rules in this ruleset
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) SetBypassActors(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRulesetBypassActorable)() {
    m.bypass_actors = value
}
// SetConditions sets the conditions property value. Conditions for an organization ruleset. The conditions object should contain both `repository_name` and `ref_name` properties or both `repository_id` and `ref_name` properties.
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) SetConditions(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRulesetConditionsable)() {
    m.conditions = value
}
// SetEnforcement sets the enforcement property value. The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) SetEnforcement(value *i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleEnforcement)() {
    m.enforcement = value
}
// SetName sets the name property value. The name of the ruleset.
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) SetName(value *string)() {
    m.name = value
}
// SetRules sets the rules property value. An array of rules within the ruleset.
func (m *ItemRulesetsItemWithRuleset_PutRequestBody) SetRules(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleable)() {
    m.rules = value
}
// ItemRulesetsItemWithRuleset_PutRequestBodyable 
type ItemRulesetsItemWithRuleset_PutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassActors()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRulesetBypassActorable)
    GetConditions()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRulesetConditionsable)
    GetEnforcement()(*i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleEnforcement)
    GetName()(*string)
    GetRules()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleable)
    SetBypassActors(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRulesetBypassActorable)()
    SetConditions(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrgRulesetConditionsable)()
    SetEnforcement(value *i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleEnforcement)()
    SetName(value *string)()
    SetRules(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleable)()
}