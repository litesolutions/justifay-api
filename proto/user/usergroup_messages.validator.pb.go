// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user/usergroup_messages.proto

package user

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *UserGroupRequest) Validate() error {
	return nil
}
func (this *UserGroupMembershipRequest) Validate() error {
	return nil
}
func (this *UserGroup) Validate() error {
	for _, item := range this.Followers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Followers", err)
			}
		}
	}
	for _, item := range this.Members {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Members", err)
			}
		}
	}
	for _, item := range this.MemberOfGroups {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MemberOfGroups", err)
			}
		}
	}
	for _, item := range this.Links {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
			}
		}
	}
	for _, item := range this.Tags {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Tags", err)
			}
		}
	}
	if this.Address != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Address); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Address", err)
		}
	}
	if this.Privacy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Privacy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Privacy", err)
		}
	}
	for _, item := range this.RecommendedArtists {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RecommendedArtists", err)
			}
		}
	}
	return nil
}
func (this *UserGroupCreateRequest) Validate() error {
	return nil
}
func (this *UserGroupUpdateRequest) Validate() error {
	return nil
}
func (this *UserGroupPrivateResponse) Validate() error {
	return nil
}
func (this *UserGroupListResponse) Validate() error {
	for _, item := range this.Usergroup {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Usergroup", err)
			}
		}
	}
	return nil
}
func (this *UserGroupPublicResponse) Validate() error {
	for _, item := range this.Links {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
			}
		}
	}
	return nil
}
func (this *UserGroupRecommended) Validate() error {
	return nil
}
func (this *UserGroupMembers) Validate() error {
	return nil
}
func (this *Group) Validate() error {
	return nil
}
func (this *GroupTypes) Validate() error {
	for _, item := range this.Types {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Types", err)
			}
		}
	}
	return nil
}
func (this *Link) Validate() error {
	return nil
}
func (this *Privacy) Validate() error {
	return nil
}
func (this *GroupedUserGroups) Validate() error {
	for _, item := range this.Groups {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Groups", err)
			}
		}
	}
	return nil
}