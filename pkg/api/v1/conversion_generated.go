package v1

// AUTO-GENERATED FUNCTIONS START HERE
import (
	api "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	resource "github.com/GoogleCloudPlatform/kubernetes/pkg/api/resource"
	v1 "github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1"
	conversion "github.com/GoogleCloudPlatform/kubernetes/pkg/conversion"
	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	apiv1 "github.com/openshift/origin/pkg/authorization/api/v1"
	buildapi "github.com/openshift/origin/pkg/build/api"
	buildapiv1 "github.com/openshift/origin/pkg/build/api/v1"
	deployapi "github.com/openshift/origin/pkg/deploy/api"
	deployapiv1 "github.com/openshift/origin/pkg/deploy/api/v1"
	imageapi "github.com/openshift/origin/pkg/image/api"
	imageapiv1 "github.com/openshift/origin/pkg/image/api/v1"
	oauthapi "github.com/openshift/origin/pkg/oauth/api"
	oauthapiv1 "github.com/openshift/origin/pkg/oauth/api/v1"
	projectapi "github.com/openshift/origin/pkg/project/api"
	projectapiv1 "github.com/openshift/origin/pkg/project/api/v1"
	routeapi "github.com/openshift/origin/pkg/route/api"
	routeapiv1 "github.com/openshift/origin/pkg/route/api/v1"
	sdnapi "github.com/openshift/origin/pkg/sdn/api"
	sdnapiv1 "github.com/openshift/origin/pkg/sdn/api/v1"
	templateapi "github.com/openshift/origin/pkg/template/api"
	templateapiv1 "github.com/openshift/origin/pkg/template/api/v1"
	userapi "github.com/openshift/origin/pkg/user/api"
	userapiv1 "github.com/openshift/origin/pkg/user/api/v1"
	reflect "reflect"
)

func convert_api_EnvVar_To_v1_EnvVar(in *api.EnvVar, out *v1.EnvVar, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.EnvVar))(in)
	}
	out.Name = in.Name
	out.Value = in.Value
	if in.ValueFrom != nil {
		out.ValueFrom = new(v1.EnvVarSource)
		if err := convert_api_EnvVarSource_To_v1_EnvVarSource(in.ValueFrom, out.ValueFrom, s); err != nil {
			return err
		}
	} else {
		out.ValueFrom = nil
	}
	return nil
}

func convert_api_EnvVarSource_To_v1_EnvVarSource(in *api.EnvVarSource, out *v1.EnvVarSource, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.EnvVarSource))(in)
	}
	if in.FieldRef != nil {
		out.FieldRef = new(v1.ObjectFieldSelector)
		if err := convert_api_ObjectFieldSelector_To_v1_ObjectFieldSelector(in.FieldRef, out.FieldRef, s); err != nil {
			return err
		}
	} else {
		out.FieldRef = nil
	}
	return nil
}

func convert_api_ListMeta_To_v1_ListMeta(in *api.ListMeta, out *v1.ListMeta, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.ListMeta))(in)
	}
	out.SelfLink = in.SelfLink
	out.ResourceVersion = in.ResourceVersion
	return nil
}

func convert_api_LocalObjectReference_To_v1_LocalObjectReference(in *api.LocalObjectReference, out *v1.LocalObjectReference, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.LocalObjectReference))(in)
	}
	out.Name = in.Name
	return nil
}

func convert_api_ObjectFieldSelector_To_v1_ObjectFieldSelector(in *api.ObjectFieldSelector, out *v1.ObjectFieldSelector, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.ObjectFieldSelector))(in)
	}
	out.APIVersion = in.APIVersion
	out.FieldPath = in.FieldPath
	return nil
}

func convert_api_ObjectMeta_To_v1_ObjectMeta(in *api.ObjectMeta, out *v1.ObjectMeta, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.ObjectMeta))(in)
	}
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = in.UID
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := s.Convert(&in.CreationTimestamp, &out.CreationTimestamp, 0); err != nil {
		return err
	}
	if in.DeletionTimestamp != nil {
		if err := s.Convert(&in.DeletionTimestamp, &out.DeletionTimestamp, 0); err != nil {
			return err
		}
	} else {
		out.DeletionTimestamp = nil
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for key, val := range in.Annotations {
			out.Annotations[key] = val
		}
	} else {
		out.Annotations = nil
	}
	return nil
}

func convert_api_ObjectReference_To_v1_ObjectReference(in *api.ObjectReference, out *v1.ObjectReference, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.ObjectReference))(in)
	}
	out.Kind = in.Kind
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = in.UID
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.FieldPath = in.FieldPath
	return nil
}

func convert_api_ResourceRequirements_To_v1_ResourceRequirements(in *api.ResourceRequirements, out *v1.ResourceRequirements, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.ResourceRequirements))(in)
	}
	if in.Limits != nil {
		out.Limits = make(v1.ResourceList)
		for key, val := range in.Limits {
			newVal := resource.Quantity{}
			if err := s.Convert(&val, &newVal, 0); err != nil {
				return err
			}
			out.Limits[v1.ResourceName(key)] = newVal
		}
	} else {
		out.Limits = nil
	}
	if in.Requests != nil {
		out.Requests = make(v1.ResourceList)
		for key, val := range in.Requests {
			newVal := resource.Quantity{}
			if err := s.Convert(&val, &newVal, 0); err != nil {
				return err
			}
			out.Requests[v1.ResourceName(key)] = newVal
		}
	} else {
		out.Requests = nil
	}
	return nil
}

func convert_api_TypeMeta_To_v1_TypeMeta(in *api.TypeMeta, out *v1.TypeMeta, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.TypeMeta))(in)
	}
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	return nil
}

func convert_v1_EnvVar_To_api_EnvVar(in *v1.EnvVar, out *api.EnvVar, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.EnvVar))(in)
	}
	out.Name = in.Name
	out.Value = in.Value
	if in.ValueFrom != nil {
		out.ValueFrom = new(api.EnvVarSource)
		if err := convert_v1_EnvVarSource_To_api_EnvVarSource(in.ValueFrom, out.ValueFrom, s); err != nil {
			return err
		}
	} else {
		out.ValueFrom = nil
	}
	return nil
}

func convert_v1_EnvVarSource_To_api_EnvVarSource(in *v1.EnvVarSource, out *api.EnvVarSource, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.EnvVarSource))(in)
	}
	if in.FieldRef != nil {
		out.FieldRef = new(api.ObjectFieldSelector)
		if err := convert_v1_ObjectFieldSelector_To_api_ObjectFieldSelector(in.FieldRef, out.FieldRef, s); err != nil {
			return err
		}
	} else {
		out.FieldRef = nil
	}
	return nil
}

func convert_v1_ListMeta_To_api_ListMeta(in *v1.ListMeta, out *api.ListMeta, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.ListMeta))(in)
	}
	out.SelfLink = in.SelfLink
	out.ResourceVersion = in.ResourceVersion
	return nil
}

func convert_v1_LocalObjectReference_To_api_LocalObjectReference(in *v1.LocalObjectReference, out *api.LocalObjectReference, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.LocalObjectReference))(in)
	}
	out.Name = in.Name
	return nil
}

func convert_v1_ObjectFieldSelector_To_api_ObjectFieldSelector(in *v1.ObjectFieldSelector, out *api.ObjectFieldSelector, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.ObjectFieldSelector))(in)
	}
	out.APIVersion = in.APIVersion
	out.FieldPath = in.FieldPath
	return nil
}

func convert_v1_ObjectMeta_To_api_ObjectMeta(in *v1.ObjectMeta, out *api.ObjectMeta, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.ObjectMeta))(in)
	}
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = in.UID
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := s.Convert(&in.CreationTimestamp, &out.CreationTimestamp, 0); err != nil {
		return err
	}
	if in.DeletionTimestamp != nil {
		if err := s.Convert(&in.DeletionTimestamp, &out.DeletionTimestamp, 0); err != nil {
			return err
		}
	} else {
		out.DeletionTimestamp = nil
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for key, val := range in.Annotations {
			out.Annotations[key] = val
		}
	} else {
		out.Annotations = nil
	}
	return nil
}

func convert_v1_ObjectReference_To_api_ObjectReference(in *v1.ObjectReference, out *api.ObjectReference, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.ObjectReference))(in)
	}
	out.Kind = in.Kind
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = in.UID
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.FieldPath = in.FieldPath
	return nil
}

func convert_v1_ResourceRequirements_To_api_ResourceRequirements(in *v1.ResourceRequirements, out *api.ResourceRequirements, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.ResourceRequirements))(in)
	}
	if in.Limits != nil {
		out.Limits = make(api.ResourceList)
		for key, val := range in.Limits {
			newVal := resource.Quantity{}
			if err := s.Convert(&val, &newVal, 0); err != nil {
				return err
			}
			out.Limits[api.ResourceName(key)] = newVal
		}
	} else {
		out.Limits = nil
	}
	if in.Requests != nil {
		out.Requests = make(api.ResourceList)
		for key, val := range in.Requests {
			newVal := resource.Quantity{}
			if err := s.Convert(&val, &newVal, 0); err != nil {
				return err
			}
			out.Requests[api.ResourceName(key)] = newVal
		}
	} else {
		out.Requests = nil
	}
	return nil
}

func convert_v1_TypeMeta_To_api_TypeMeta(in *v1.TypeMeta, out *api.TypeMeta, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.TypeMeta))(in)
	}
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	return nil
}

func convert_api_ClusterPolicyBindingList_To_v1_ClusterPolicyBindingList(in *authorizationapi.ClusterPolicyBindingList, out *apiv1.ClusterPolicyBindingList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.ClusterPolicyBindingList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]apiv1.ClusterPolicyBinding, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_ClusterPolicyList_To_v1_ClusterPolicyList(in *authorizationapi.ClusterPolicyList, out *apiv1.ClusterPolicyList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.ClusterPolicyList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]apiv1.ClusterPolicy, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_ClusterRole_To_v1_ClusterRole(in *authorizationapi.ClusterRole, out *apiv1.ClusterRole, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.ClusterRole))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if in.Rules != nil {
		out.Rules = make([]apiv1.PolicyRule, len(in.Rules))
		for i := range in.Rules {
			if err := s.Convert(&in.Rules[i], &out.Rules[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func convert_api_ClusterRoleBindingList_To_v1_ClusterRoleBindingList(in *authorizationapi.ClusterRoleBindingList, out *apiv1.ClusterRoleBindingList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.ClusterRoleBindingList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]apiv1.ClusterRoleBinding, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_ClusterRoleList_To_v1_ClusterRoleList(in *authorizationapi.ClusterRoleList, out *apiv1.ClusterRoleList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.ClusterRoleList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]apiv1.ClusterRole, len(in.Items))
		for i := range in.Items {
			if err := convert_api_ClusterRole_To_v1_ClusterRole(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_IsPersonalSubjectAccessReview_To_v1_IsPersonalSubjectAccessReview(in *authorizationapi.IsPersonalSubjectAccessReview, out *apiv1.IsPersonalSubjectAccessReview, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.IsPersonalSubjectAccessReview))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	return nil
}

func convert_api_PolicyBindingList_To_v1_PolicyBindingList(in *authorizationapi.PolicyBindingList, out *apiv1.PolicyBindingList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.PolicyBindingList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]apiv1.PolicyBinding, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_PolicyList_To_v1_PolicyList(in *authorizationapi.PolicyList, out *apiv1.PolicyList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.PolicyList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]apiv1.Policy, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_ResourceAccessReview_To_v1_ResourceAccessReview(in *authorizationapi.ResourceAccessReview, out *apiv1.ResourceAccessReview, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.ResourceAccessReview))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Verb = in.Verb
	out.Resource = in.Resource
	if err := s.Convert(&in.Content, &out.Content, 0); err != nil {
		return err
	}
	out.ResourceName = in.ResourceName
	return nil
}

func convert_api_Role_To_v1_Role(in *authorizationapi.Role, out *apiv1.Role, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.Role))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if in.Rules != nil {
		out.Rules = make([]apiv1.PolicyRule, len(in.Rules))
		for i := range in.Rules {
			if err := s.Convert(&in.Rules[i], &out.Rules[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func convert_api_RoleBindingList_To_v1_RoleBindingList(in *authorizationapi.RoleBindingList, out *apiv1.RoleBindingList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.RoleBindingList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]apiv1.RoleBinding, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_RoleList_To_v1_RoleList(in *authorizationapi.RoleList, out *apiv1.RoleList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.RoleList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]apiv1.Role, len(in.Items))
		for i := range in.Items {
			if err := convert_api_Role_To_v1_Role(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_SubjectAccessReviewResponse_To_v1_SubjectAccessReviewResponse(in *authorizationapi.SubjectAccessReviewResponse, out *apiv1.SubjectAccessReviewResponse, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*authorizationapi.SubjectAccessReviewResponse))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	return nil
}

func convert_v1_ClusterPolicyBindingList_To_api_ClusterPolicyBindingList(in *apiv1.ClusterPolicyBindingList, out *authorizationapi.ClusterPolicyBindingList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.ClusterPolicyBindingList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]authorizationapi.ClusterPolicyBinding, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_ClusterPolicyList_To_api_ClusterPolicyList(in *apiv1.ClusterPolicyList, out *authorizationapi.ClusterPolicyList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.ClusterPolicyList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]authorizationapi.ClusterPolicy, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_ClusterRole_To_api_ClusterRole(in *apiv1.ClusterRole, out *authorizationapi.ClusterRole, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.ClusterRole))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if in.Rules != nil {
		out.Rules = make([]authorizationapi.PolicyRule, len(in.Rules))
		for i := range in.Rules {
			if err := s.Convert(&in.Rules[i], &out.Rules[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func convert_v1_ClusterRoleBindingList_To_api_ClusterRoleBindingList(in *apiv1.ClusterRoleBindingList, out *authorizationapi.ClusterRoleBindingList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.ClusterRoleBindingList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]authorizationapi.ClusterRoleBinding, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_ClusterRoleList_To_api_ClusterRoleList(in *apiv1.ClusterRoleList, out *authorizationapi.ClusterRoleList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.ClusterRoleList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]authorizationapi.ClusterRole, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_ClusterRole_To_api_ClusterRole(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_IsPersonalSubjectAccessReview_To_api_IsPersonalSubjectAccessReview(in *apiv1.IsPersonalSubjectAccessReview, out *authorizationapi.IsPersonalSubjectAccessReview, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.IsPersonalSubjectAccessReview))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_PolicyBindingList_To_api_PolicyBindingList(in *apiv1.PolicyBindingList, out *authorizationapi.PolicyBindingList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.PolicyBindingList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]authorizationapi.PolicyBinding, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_PolicyList_To_api_PolicyList(in *apiv1.PolicyList, out *authorizationapi.PolicyList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.PolicyList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]authorizationapi.Policy, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_ResourceAccessReview_To_api_ResourceAccessReview(in *apiv1.ResourceAccessReview, out *authorizationapi.ResourceAccessReview, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.ResourceAccessReview))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Verb = in.Verb
	out.Resource = in.Resource
	if err := s.Convert(&in.Content, &out.Content, 0); err != nil {
		return err
	}
	out.ResourceName = in.ResourceName
	return nil
}

func convert_v1_Role_To_api_Role(in *apiv1.Role, out *authorizationapi.Role, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.Role))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if in.Rules != nil {
		out.Rules = make([]authorizationapi.PolicyRule, len(in.Rules))
		for i := range in.Rules {
			if err := s.Convert(&in.Rules[i], &out.Rules[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func convert_v1_RoleBindingList_To_api_RoleBindingList(in *apiv1.RoleBindingList, out *authorizationapi.RoleBindingList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.RoleBindingList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]authorizationapi.RoleBinding, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_RoleList_To_api_RoleList(in *apiv1.RoleList, out *authorizationapi.RoleList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.RoleList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]authorizationapi.Role, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_Role_To_api_Role(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_SubjectAccessReviewResponse_To_api_SubjectAccessReviewResponse(in *apiv1.SubjectAccessReviewResponse, out *authorizationapi.SubjectAccessReviewResponse, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*apiv1.SubjectAccessReviewResponse))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	return nil
}

func convert_api_Build_To_v1_Build(in *buildapi.Build, out *buildapiv1.Build, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.Build))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := convert_api_BuildSpec_To_v1_BuildSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := convert_api_BuildStatus_To_v1_BuildStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func convert_api_BuildConfig_To_v1_BuildConfig(in *buildapi.BuildConfig, out *buildapiv1.BuildConfig, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildConfig))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := convert_api_BuildConfigSpec_To_v1_BuildConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := convert_api_BuildConfigStatus_To_v1_BuildConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func convert_api_BuildConfigList_To_v1_BuildConfigList(in *buildapi.BuildConfigList, out *buildapiv1.BuildConfigList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildConfigList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]buildapiv1.BuildConfig, len(in.Items))
		for i := range in.Items {
			if err := convert_api_BuildConfig_To_v1_BuildConfig(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_BuildConfigSpec_To_v1_BuildConfigSpec(in *buildapi.BuildConfigSpec, out *buildapiv1.BuildConfigSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildConfigSpec))(in)
	}
	if in.Triggers != nil {
		out.Triggers = make([]buildapiv1.BuildTriggerPolicy, len(in.Triggers))
		for i := range in.Triggers {
			if err := convert_api_BuildTriggerPolicy_To_v1_BuildTriggerPolicy(&in.Triggers[i], &out.Triggers[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	if err := convert_api_BuildSpec_To_v1_BuildSpec(&in.BuildSpec, &out.BuildSpec, s); err != nil {
		return err
	}
	return nil
}

func convert_api_BuildConfigStatus_To_v1_BuildConfigStatus(in *buildapi.BuildConfigStatus, out *buildapiv1.BuildConfigStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildConfigStatus))(in)
	}
	out.LastVersion = in.LastVersion
	return nil
}

func convert_api_BuildList_To_v1_BuildList(in *buildapi.BuildList, out *buildapiv1.BuildList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]buildapiv1.Build, len(in.Items))
		for i := range in.Items {
			if err := convert_api_Build_To_v1_Build(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_BuildLog_To_v1_BuildLog(in *buildapi.BuildLog, out *buildapiv1.BuildLog, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildLog))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	return nil
}

func convert_api_BuildLogOptions_To_v1_BuildLogOptions(in *buildapi.BuildLogOptions, out *buildapiv1.BuildLogOptions, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildLogOptions))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Follow = in.Follow
	out.NoWait = in.NoWait
	return nil
}

func convert_api_BuildOutput_To_v1_BuildOutput(in *buildapi.BuildOutput, out *buildapiv1.BuildOutput, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildOutput))(in)
	}
	if in.To != nil {
		out.To = new(v1.ObjectReference)
		if err := convert_api_ObjectReference_To_v1_ObjectReference(in.To, out.To, s); err != nil {
			return err
		}
	} else {
		out.To = nil
	}
	if in.PushSecret != nil {
		out.PushSecret = new(v1.LocalObjectReference)
		if err := convert_api_LocalObjectReference_To_v1_LocalObjectReference(in.PushSecret, out.PushSecret, s); err != nil {
			return err
		}
	} else {
		out.PushSecret = nil
	}
	return nil
}

func convert_api_BuildRequest_To_v1_BuildRequest(in *buildapi.BuildRequest, out *buildapiv1.BuildRequest, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildRequest))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if in.Revision != nil {
		out.Revision = new(buildapiv1.SourceRevision)
		if err := convert_api_SourceRevision_To_v1_SourceRevision(in.Revision, out.Revision, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if in.TriggeredByImage != nil {
		out.TriggeredByImage = new(v1.ObjectReference)
		if err := convert_api_ObjectReference_To_v1_ObjectReference(in.TriggeredByImage, out.TriggeredByImage, s); err != nil {
			return err
		}
	} else {
		out.TriggeredByImage = nil
	}
	return nil
}

func convert_api_BuildSource_To_v1_BuildSource(in *buildapi.BuildSource, out *buildapiv1.BuildSource, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildSource))(in)
	}
	out.Type = buildapiv1.BuildSourceType(in.Type)
	if in.Git != nil {
		out.Git = new(buildapiv1.GitBuildSource)
		if err := convert_api_GitBuildSource_To_v1_GitBuildSource(in.Git, out.Git, s); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	out.ContextDir = in.ContextDir
	if in.SourceSecret != nil {
		out.SourceSecret = new(v1.LocalObjectReference)
		if err := convert_api_LocalObjectReference_To_v1_LocalObjectReference(in.SourceSecret, out.SourceSecret, s); err != nil {
			return err
		}
	} else {
		out.SourceSecret = nil
	}
	return nil
}

func convert_api_BuildSpec_To_v1_BuildSpec(in *buildapi.BuildSpec, out *buildapiv1.BuildSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildSpec))(in)
	}
	out.ServiceAccount = in.ServiceAccount
	if err := convert_api_BuildSource_To_v1_BuildSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	if in.Revision != nil {
		out.Revision = new(buildapiv1.SourceRevision)
		if err := convert_api_SourceRevision_To_v1_SourceRevision(in.Revision, out.Revision, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if err := convert_api_BuildStrategy_To_v1_BuildStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	if err := convert_api_BuildOutput_To_v1_BuildOutput(&in.Output, &out.Output, s); err != nil {
		return err
	}
	if err := convert_api_ResourceRequirements_To_v1_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	return nil
}

func convert_api_BuildStatus_To_v1_BuildStatus(in *buildapi.BuildStatus, out *buildapiv1.BuildStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildStatus))(in)
	}
	out.Phase = buildapiv1.BuildPhase(in.Phase)
	out.Cancelled = in.Cancelled
	out.Message = in.Message
	if in.StartTimestamp != nil {
		if err := s.Convert(&in.StartTimestamp, &out.StartTimestamp, 0); err != nil {
			return err
		}
	} else {
		out.StartTimestamp = nil
	}
	if in.CompletionTimestamp != nil {
		if err := s.Convert(&in.CompletionTimestamp, &out.CompletionTimestamp, 0); err != nil {
			return err
		}
	} else {
		out.CompletionTimestamp = nil
	}
	out.Duration = in.Duration
	if in.Config != nil {
		out.Config = new(v1.ObjectReference)
		if err := convert_api_ObjectReference_To_v1_ObjectReference(in.Config, out.Config, s); err != nil {
			return err
		}
	} else {
		out.Config = nil
	}
	return nil
}

func convert_api_BuildStrategy_To_v1_BuildStrategy(in *buildapi.BuildStrategy, out *buildapiv1.BuildStrategy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildStrategy))(in)
	}
	out.Type = buildapiv1.BuildStrategyType(in.Type)
	if in.DockerStrategy != nil {
		out.DockerStrategy = new(buildapiv1.DockerBuildStrategy)
		if err := convert_api_DockerBuildStrategy_To_v1_DockerBuildStrategy(in.DockerStrategy, out.DockerStrategy, s); err != nil {
			return err
		}
	} else {
		out.DockerStrategy = nil
	}
	if in.SourceStrategy != nil {
		out.SourceStrategy = new(buildapiv1.SourceBuildStrategy)
		if err := convert_api_SourceBuildStrategy_To_v1_SourceBuildStrategy(in.SourceStrategy, out.SourceStrategy, s); err != nil {
			return err
		}
	} else {
		out.SourceStrategy = nil
	}
	if in.CustomStrategy != nil {
		out.CustomStrategy = new(buildapiv1.CustomBuildStrategy)
		if err := convert_api_CustomBuildStrategy_To_v1_CustomBuildStrategy(in.CustomStrategy, out.CustomStrategy, s); err != nil {
			return err
		}
	} else {
		out.CustomStrategy = nil
	}
	return nil
}

func convert_api_BuildTriggerPolicy_To_v1_BuildTriggerPolicy(in *buildapi.BuildTriggerPolicy, out *buildapiv1.BuildTriggerPolicy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.BuildTriggerPolicy))(in)
	}
	out.Type = buildapiv1.BuildTriggerType(in.Type)
	if in.GitHubWebHook != nil {
		out.GitHubWebHook = new(buildapiv1.WebHookTrigger)
		if err := convert_api_WebHookTrigger_To_v1_WebHookTrigger(in.GitHubWebHook, out.GitHubWebHook, s); err != nil {
			return err
		}
	} else {
		out.GitHubWebHook = nil
	}
	if in.GenericWebHook != nil {
		out.GenericWebHook = new(buildapiv1.WebHookTrigger)
		if err := convert_api_WebHookTrigger_To_v1_WebHookTrigger(in.GenericWebHook, out.GenericWebHook, s); err != nil {
			return err
		}
	} else {
		out.GenericWebHook = nil
	}
	if in.ImageChange != nil {
		out.ImageChange = new(buildapiv1.ImageChangeTrigger)
		if err := convert_api_ImageChangeTrigger_To_v1_ImageChangeTrigger(in.ImageChange, out.ImageChange, s); err != nil {
			return err
		}
	} else {
		out.ImageChange = nil
	}
	return nil
}

func convert_api_CustomBuildStrategy_To_v1_CustomBuildStrategy(in *buildapi.CustomBuildStrategy, out *buildapiv1.CustomBuildStrategy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.CustomBuildStrategy))(in)
	}
	if err := convert_api_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.PullSecret != nil {
		out.PullSecret = new(v1.LocalObjectReference)
		if err := convert_api_LocalObjectReference_To_v1_LocalObjectReference(in.PullSecret, out.PullSecret, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		out.Env = make([]v1.EnvVar, len(in.Env))
		for i := range in.Env {
			if err := convert_api_EnvVar_To_v1_EnvVar(&in.Env[i], &out.Env[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ExposeDockerSocket = in.ExposeDockerSocket
	return nil
}

func convert_api_DockerBuildStrategy_To_v1_DockerBuildStrategy(in *buildapi.DockerBuildStrategy, out *buildapiv1.DockerBuildStrategy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.DockerBuildStrategy))(in)
	}
	if in.From != nil {
		out.From = new(v1.ObjectReference)
		if err := convert_api_ObjectReference_To_v1_ObjectReference(in.From, out.From, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	if in.PullSecret != nil {
		out.PullSecret = new(v1.LocalObjectReference)
		if err := convert_api_LocalObjectReference_To_v1_LocalObjectReference(in.PullSecret, out.PullSecret, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	out.NoCache = in.NoCache
	if in.Env != nil {
		out.Env = make([]v1.EnvVar, len(in.Env))
		for i := range in.Env {
			if err := convert_api_EnvVar_To_v1_EnvVar(&in.Env[i], &out.Env[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ForcePull = in.ForcePull
	return nil
}

func convert_api_GitBuildSource_To_v1_GitBuildSource(in *buildapi.GitBuildSource, out *buildapiv1.GitBuildSource, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.GitBuildSource))(in)
	}
	out.URI = in.URI
	out.Ref = in.Ref
	out.HTTPProxy = in.HTTPProxy
	out.HTTPSProxy = in.HTTPSProxy
	return nil
}

func convert_api_GitSourceRevision_To_v1_GitSourceRevision(in *buildapi.GitSourceRevision, out *buildapiv1.GitSourceRevision, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.GitSourceRevision))(in)
	}
	out.Commit = in.Commit
	if err := convert_api_SourceControlUser_To_v1_SourceControlUser(&in.Author, &out.Author, s); err != nil {
		return err
	}
	if err := convert_api_SourceControlUser_To_v1_SourceControlUser(&in.Committer, &out.Committer, s); err != nil {
		return err
	}
	out.Message = in.Message
	return nil
}

func convert_api_ImageChangeTrigger_To_v1_ImageChangeTrigger(in *buildapi.ImageChangeTrigger, out *buildapiv1.ImageChangeTrigger, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.ImageChangeTrigger))(in)
	}
	out.LastTriggeredImageID = in.LastTriggeredImageID
	return nil
}

func convert_api_SourceBuildStrategy_To_v1_SourceBuildStrategy(in *buildapi.SourceBuildStrategy, out *buildapiv1.SourceBuildStrategy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.SourceBuildStrategy))(in)
	}
	if err := convert_api_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.PullSecret != nil {
		out.PullSecret = new(v1.LocalObjectReference)
		if err := convert_api_LocalObjectReference_To_v1_LocalObjectReference(in.PullSecret, out.PullSecret, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		out.Env = make([]v1.EnvVar, len(in.Env))
		for i := range in.Env {
			if err := convert_api_EnvVar_To_v1_EnvVar(&in.Env[i], &out.Env[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.Scripts = in.Scripts
	out.Incremental = in.Incremental
	out.ForcePull = in.ForcePull
	return nil
}

func convert_api_SourceControlUser_To_v1_SourceControlUser(in *buildapi.SourceControlUser, out *buildapiv1.SourceControlUser, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.SourceControlUser))(in)
	}
	out.Name = in.Name
	out.Email = in.Email
	return nil
}

func convert_api_SourceRevision_To_v1_SourceRevision(in *buildapi.SourceRevision, out *buildapiv1.SourceRevision, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.SourceRevision))(in)
	}
	out.Type = buildapiv1.BuildSourceType(in.Type)
	if in.Git != nil {
		out.Git = new(buildapiv1.GitSourceRevision)
		if err := convert_api_GitSourceRevision_To_v1_GitSourceRevision(in.Git, out.Git, s); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	return nil
}

func convert_api_WebHookTrigger_To_v1_WebHookTrigger(in *buildapi.WebHookTrigger, out *buildapiv1.WebHookTrigger, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapi.WebHookTrigger))(in)
	}
	out.Secret = in.Secret
	return nil
}

func convert_v1_Build_To_api_Build(in *buildapiv1.Build, out *buildapi.Build, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.Build))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := convert_v1_BuildSpec_To_api_BuildSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := convert_v1_BuildStatus_To_api_BuildStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_BuildConfig_To_api_BuildConfig(in *buildapiv1.BuildConfig, out *buildapi.BuildConfig, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildConfig))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := convert_v1_BuildConfigSpec_To_api_BuildConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := convert_v1_BuildConfigStatus_To_api_BuildConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_BuildConfigList_To_api_BuildConfigList(in *buildapiv1.BuildConfigList, out *buildapi.BuildConfigList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildConfigList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]buildapi.BuildConfig, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_BuildConfig_To_api_BuildConfig(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_BuildConfigSpec_To_api_BuildConfigSpec(in *buildapiv1.BuildConfigSpec, out *buildapi.BuildConfigSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildConfigSpec))(in)
	}
	if in.Triggers != nil {
		out.Triggers = make([]buildapi.BuildTriggerPolicy, len(in.Triggers))
		for i := range in.Triggers {
			if err := convert_v1_BuildTriggerPolicy_To_api_BuildTriggerPolicy(&in.Triggers[i], &out.Triggers[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	if err := convert_v1_BuildSpec_To_api_BuildSpec(&in.BuildSpec, &out.BuildSpec, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_BuildConfigStatus_To_api_BuildConfigStatus(in *buildapiv1.BuildConfigStatus, out *buildapi.BuildConfigStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildConfigStatus))(in)
	}
	out.LastVersion = in.LastVersion
	return nil
}

func convert_v1_BuildList_To_api_BuildList(in *buildapiv1.BuildList, out *buildapi.BuildList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]buildapi.Build, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_Build_To_api_Build(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_BuildLog_To_api_BuildLog(in *buildapiv1.BuildLog, out *buildapi.BuildLog, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildLog))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_BuildLogOptions_To_api_BuildLogOptions(in *buildapiv1.BuildLogOptions, out *buildapi.BuildLogOptions, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildLogOptions))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Follow = in.Follow
	out.NoWait = in.NoWait
	return nil
}

func convert_v1_BuildOutput_To_api_BuildOutput(in *buildapiv1.BuildOutput, out *buildapi.BuildOutput, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildOutput))(in)
	}
	if in.To != nil {
		out.To = new(api.ObjectReference)
		if err := convert_v1_ObjectReference_To_api_ObjectReference(in.To, out.To, s); err != nil {
			return err
		}
	} else {
		out.To = nil
	}
	if in.PushSecret != nil {
		out.PushSecret = new(api.LocalObjectReference)
		if err := convert_v1_LocalObjectReference_To_api_LocalObjectReference(in.PushSecret, out.PushSecret, s); err != nil {
			return err
		}
	} else {
		out.PushSecret = nil
	}
	return nil
}

func convert_v1_BuildRequest_To_api_BuildRequest(in *buildapiv1.BuildRequest, out *buildapi.BuildRequest, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildRequest))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if in.Revision != nil {
		out.Revision = new(buildapi.SourceRevision)
		if err := convert_v1_SourceRevision_To_api_SourceRevision(in.Revision, out.Revision, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if in.TriggeredByImage != nil {
		out.TriggeredByImage = new(api.ObjectReference)
		if err := convert_v1_ObjectReference_To_api_ObjectReference(in.TriggeredByImage, out.TriggeredByImage, s); err != nil {
			return err
		}
	} else {
		out.TriggeredByImage = nil
	}
	return nil
}

func convert_v1_BuildSource_To_api_BuildSource(in *buildapiv1.BuildSource, out *buildapi.BuildSource, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildSource))(in)
	}
	out.Type = buildapi.BuildSourceType(in.Type)
	if in.Git != nil {
		out.Git = new(buildapi.GitBuildSource)
		if err := convert_v1_GitBuildSource_To_api_GitBuildSource(in.Git, out.Git, s); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	out.ContextDir = in.ContextDir
	if in.SourceSecret != nil {
		out.SourceSecret = new(api.LocalObjectReference)
		if err := convert_v1_LocalObjectReference_To_api_LocalObjectReference(in.SourceSecret, out.SourceSecret, s); err != nil {
			return err
		}
	} else {
		out.SourceSecret = nil
	}
	return nil
}

func convert_v1_BuildSpec_To_api_BuildSpec(in *buildapiv1.BuildSpec, out *buildapi.BuildSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildSpec))(in)
	}
	out.ServiceAccount = in.ServiceAccount
	if err := convert_v1_BuildSource_To_api_BuildSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	if in.Revision != nil {
		out.Revision = new(buildapi.SourceRevision)
		if err := convert_v1_SourceRevision_To_api_SourceRevision(in.Revision, out.Revision, s); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if err := convert_v1_BuildStrategy_To_api_BuildStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	if err := convert_v1_BuildOutput_To_api_BuildOutput(&in.Output, &out.Output, s); err != nil {
		return err
	}
	if err := convert_v1_ResourceRequirements_To_api_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_BuildStatus_To_api_BuildStatus(in *buildapiv1.BuildStatus, out *buildapi.BuildStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildStatus))(in)
	}
	out.Phase = buildapi.BuildPhase(in.Phase)
	out.Cancelled = in.Cancelled
	out.Message = in.Message
	if in.StartTimestamp != nil {
		if err := s.Convert(&in.StartTimestamp, &out.StartTimestamp, 0); err != nil {
			return err
		}
	} else {
		out.StartTimestamp = nil
	}
	if in.CompletionTimestamp != nil {
		if err := s.Convert(&in.CompletionTimestamp, &out.CompletionTimestamp, 0); err != nil {
			return err
		}
	} else {
		out.CompletionTimestamp = nil
	}
	out.Duration = in.Duration
	if in.Config != nil {
		out.Config = new(api.ObjectReference)
		if err := convert_v1_ObjectReference_To_api_ObjectReference(in.Config, out.Config, s); err != nil {
			return err
		}
	} else {
		out.Config = nil
	}
	return nil
}

func convert_v1_BuildStrategy_To_api_BuildStrategy(in *buildapiv1.BuildStrategy, out *buildapi.BuildStrategy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildStrategy))(in)
	}
	out.Type = buildapi.BuildStrategyType(in.Type)
	if in.DockerStrategy != nil {
		out.DockerStrategy = new(buildapi.DockerBuildStrategy)
		if err := convert_v1_DockerBuildStrategy_To_api_DockerBuildStrategy(in.DockerStrategy, out.DockerStrategy, s); err != nil {
			return err
		}
	} else {
		out.DockerStrategy = nil
	}
	if in.SourceStrategy != nil {
		out.SourceStrategy = new(buildapi.SourceBuildStrategy)
		if err := convert_v1_SourceBuildStrategy_To_api_SourceBuildStrategy(in.SourceStrategy, out.SourceStrategy, s); err != nil {
			return err
		}
	} else {
		out.SourceStrategy = nil
	}
	if in.CustomStrategy != nil {
		out.CustomStrategy = new(buildapi.CustomBuildStrategy)
		if err := convert_v1_CustomBuildStrategy_To_api_CustomBuildStrategy(in.CustomStrategy, out.CustomStrategy, s); err != nil {
			return err
		}
	} else {
		out.CustomStrategy = nil
	}
	return nil
}

func convert_v1_BuildTriggerPolicy_To_api_BuildTriggerPolicy(in *buildapiv1.BuildTriggerPolicy, out *buildapi.BuildTriggerPolicy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.BuildTriggerPolicy))(in)
	}
	out.Type = buildapi.BuildTriggerType(in.Type)
	if in.GitHubWebHook != nil {
		out.GitHubWebHook = new(buildapi.WebHookTrigger)
		if err := convert_v1_WebHookTrigger_To_api_WebHookTrigger(in.GitHubWebHook, out.GitHubWebHook, s); err != nil {
			return err
		}
	} else {
		out.GitHubWebHook = nil
	}
	if in.GenericWebHook != nil {
		out.GenericWebHook = new(buildapi.WebHookTrigger)
		if err := convert_v1_WebHookTrigger_To_api_WebHookTrigger(in.GenericWebHook, out.GenericWebHook, s); err != nil {
			return err
		}
	} else {
		out.GenericWebHook = nil
	}
	if in.ImageChange != nil {
		out.ImageChange = new(buildapi.ImageChangeTrigger)
		if err := convert_v1_ImageChangeTrigger_To_api_ImageChangeTrigger(in.ImageChange, out.ImageChange, s); err != nil {
			return err
		}
	} else {
		out.ImageChange = nil
	}
	return nil
}

func convert_v1_CustomBuildStrategy_To_api_CustomBuildStrategy(in *buildapiv1.CustomBuildStrategy, out *buildapi.CustomBuildStrategy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.CustomBuildStrategy))(in)
	}
	if err := convert_v1_ObjectReference_To_api_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.PullSecret != nil {
		out.PullSecret = new(api.LocalObjectReference)
		if err := convert_v1_LocalObjectReference_To_api_LocalObjectReference(in.PullSecret, out.PullSecret, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		out.Env = make([]api.EnvVar, len(in.Env))
		for i := range in.Env {
			if err := convert_v1_EnvVar_To_api_EnvVar(&in.Env[i], &out.Env[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ExposeDockerSocket = in.ExposeDockerSocket
	return nil
}

func convert_v1_DockerBuildStrategy_To_api_DockerBuildStrategy(in *buildapiv1.DockerBuildStrategy, out *buildapi.DockerBuildStrategy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.DockerBuildStrategy))(in)
	}
	if in.From != nil {
		out.From = new(api.ObjectReference)
		if err := convert_v1_ObjectReference_To_api_ObjectReference(in.From, out.From, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	if in.PullSecret != nil {
		out.PullSecret = new(api.LocalObjectReference)
		if err := convert_v1_LocalObjectReference_To_api_LocalObjectReference(in.PullSecret, out.PullSecret, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	out.NoCache = in.NoCache
	if in.Env != nil {
		out.Env = make([]api.EnvVar, len(in.Env))
		for i := range in.Env {
			if err := convert_v1_EnvVar_To_api_EnvVar(&in.Env[i], &out.Env[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ForcePull = in.ForcePull
	return nil
}

func convert_v1_GitBuildSource_To_api_GitBuildSource(in *buildapiv1.GitBuildSource, out *buildapi.GitBuildSource, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.GitBuildSource))(in)
	}
	out.URI = in.URI
	out.Ref = in.Ref
	out.HTTPProxy = in.HTTPProxy
	out.HTTPSProxy = in.HTTPSProxy
	return nil
}

func convert_v1_GitSourceRevision_To_api_GitSourceRevision(in *buildapiv1.GitSourceRevision, out *buildapi.GitSourceRevision, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.GitSourceRevision))(in)
	}
	out.Commit = in.Commit
	if err := convert_v1_SourceControlUser_To_api_SourceControlUser(&in.Author, &out.Author, s); err != nil {
		return err
	}
	if err := convert_v1_SourceControlUser_To_api_SourceControlUser(&in.Committer, &out.Committer, s); err != nil {
		return err
	}
	out.Message = in.Message
	return nil
}

func convert_v1_ImageChangeTrigger_To_api_ImageChangeTrigger(in *buildapiv1.ImageChangeTrigger, out *buildapi.ImageChangeTrigger, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.ImageChangeTrigger))(in)
	}
	out.LastTriggeredImageID = in.LastTriggeredImageID
	return nil
}

func convert_v1_SourceBuildStrategy_To_api_SourceBuildStrategy(in *buildapiv1.SourceBuildStrategy, out *buildapi.SourceBuildStrategy, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.SourceBuildStrategy))(in)
	}
	if err := convert_v1_ObjectReference_To_api_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.PullSecret != nil {
		out.PullSecret = new(api.LocalObjectReference)
		if err := convert_v1_LocalObjectReference_To_api_LocalObjectReference(in.PullSecret, out.PullSecret, s); err != nil {
			return err
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		out.Env = make([]api.EnvVar, len(in.Env))
		for i := range in.Env {
			if err := convert_v1_EnvVar_To_api_EnvVar(&in.Env[i], &out.Env[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.Scripts = in.Scripts
	out.Incremental = in.Incremental
	out.ForcePull = in.ForcePull
	return nil
}

func convert_v1_SourceControlUser_To_api_SourceControlUser(in *buildapiv1.SourceControlUser, out *buildapi.SourceControlUser, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.SourceControlUser))(in)
	}
	out.Name = in.Name
	out.Email = in.Email
	return nil
}

func convert_v1_SourceRevision_To_api_SourceRevision(in *buildapiv1.SourceRevision, out *buildapi.SourceRevision, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.SourceRevision))(in)
	}
	out.Type = buildapi.BuildSourceType(in.Type)
	if in.Git != nil {
		out.Git = new(buildapi.GitSourceRevision)
		if err := convert_v1_GitSourceRevision_To_api_GitSourceRevision(in.Git, out.Git, s); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	return nil
}

func convert_v1_WebHookTrigger_To_api_WebHookTrigger(in *buildapiv1.WebHookTrigger, out *buildapi.WebHookTrigger, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*buildapiv1.WebHookTrigger))(in)
	}
	out.Secret = in.Secret
	return nil
}

func convert_api_DeploymentConfigList_To_v1_DeploymentConfigList(in *deployapi.DeploymentConfigList, out *deployapiv1.DeploymentConfigList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*deployapi.DeploymentConfigList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]deployapiv1.DeploymentConfig, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_DeploymentConfigRollback_To_v1_DeploymentConfigRollback(in *deployapi.DeploymentConfigRollback, out *deployapiv1.DeploymentConfigRollback, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*deployapi.DeploymentConfigRollback))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func convert_api_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(in *deployapi.DeploymentConfigRollbackSpec, out *deployapiv1.DeploymentConfigRollbackSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*deployapi.DeploymentConfigRollbackSpec))(in)
	}
	if err := convert_api_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.IncludeTriggers = in.IncludeTriggers
	out.IncludeTemplate = in.IncludeTemplate
	out.IncludeReplicationMeta = in.IncludeReplicationMeta
	out.IncludeStrategy = in.IncludeStrategy
	return nil
}

func convert_v1_DeploymentConfigList_To_api_DeploymentConfigList(in *deployapiv1.DeploymentConfigList, out *deployapi.DeploymentConfigList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*deployapiv1.DeploymentConfigList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]deployapi.DeploymentConfig, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_DeploymentConfigRollback_To_api_DeploymentConfigRollback(in *deployapiv1.DeploymentConfigRollback, out *deployapi.DeploymentConfigRollback, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*deployapiv1.DeploymentConfigRollback))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_DeploymentConfigRollbackSpec_To_api_DeploymentConfigRollbackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_DeploymentConfigRollbackSpec_To_api_DeploymentConfigRollbackSpec(in *deployapiv1.DeploymentConfigRollbackSpec, out *deployapi.DeploymentConfigRollbackSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*deployapiv1.DeploymentConfigRollbackSpec))(in)
	}
	if err := convert_v1_ObjectReference_To_api_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.IncludeTriggers = in.IncludeTriggers
	out.IncludeTemplate = in.IncludeTemplate
	out.IncludeReplicationMeta = in.IncludeReplicationMeta
	out.IncludeStrategy = in.IncludeStrategy
	return nil
}

func convert_api_ImageList_To_v1_ImageList(in *imageapi.ImageList, out *imageapiv1.ImageList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapi.ImageList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]imageapiv1.Image, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_ImageStream_To_v1_ImageStream(in *imageapi.ImageStream, out *imageapiv1.ImageStream, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapi.ImageStream))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := s.Convert(&in.Spec, &out.Spec, 0); err != nil {
		return err
	}
	if err := s.Convert(&in.Status, &out.Status, 0); err != nil {
		return err
	}
	return nil
}

func convert_api_ImageStreamImage_To_v1_ImageStreamImage(in *imageapi.ImageStreamImage, out *imageapiv1.ImageStreamImage, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapi.ImageStreamImage))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := s.Convert(&in.Image, &out.Image, 0); err != nil {
		return err
	}
	return nil
}

func convert_api_ImageStreamList_To_v1_ImageStreamList(in *imageapi.ImageStreamList, out *imageapiv1.ImageStreamList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapi.ImageStreamList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]imageapiv1.ImageStream, len(in.Items))
		for i := range in.Items {
			if err := convert_api_ImageStream_To_v1_ImageStream(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_ImageStreamTag_To_v1_ImageStreamTag(in *imageapi.ImageStreamTag, out *imageapiv1.ImageStreamTag, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapi.ImageStreamTag))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := s.Convert(&in.Image, &out.Image, 0); err != nil {
		return err
	}
	return nil
}

func convert_v1_ImageList_To_api_ImageList(in *imageapiv1.ImageList, out *imageapi.ImageList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapiv1.ImageList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]imageapi.Image, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_ImageStream_To_api_ImageStream(in *imageapiv1.ImageStream, out *imageapi.ImageStream, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapiv1.ImageStream))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := s.Convert(&in.Spec, &out.Spec, 0); err != nil {
		return err
	}
	if err := s.Convert(&in.Status, &out.Status, 0); err != nil {
		return err
	}
	return nil
}

func convert_v1_ImageStreamImage_To_api_ImageStreamImage(in *imageapiv1.ImageStreamImage, out *imageapi.ImageStreamImage, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapiv1.ImageStreamImage))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := s.Convert(&in.Image, &out.Image, 0); err != nil {
		return err
	}
	return nil
}

func convert_v1_ImageStreamList_To_api_ImageStreamList(in *imageapiv1.ImageStreamList, out *imageapi.ImageStreamList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapiv1.ImageStreamList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]imageapi.ImageStream, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_ImageStream_To_api_ImageStream(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_ImageStreamTag_To_api_ImageStreamTag(in *imageapiv1.ImageStreamTag, out *imageapi.ImageStreamTag, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*imageapiv1.ImageStreamTag))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := s.Convert(&in.Image, &out.Image, 0); err != nil {
		return err
	}
	return nil
}

func convert_api_OAuthAccessToken_To_v1_OAuthAccessToken(in *oauthapi.OAuthAccessToken, out *oauthapiv1.OAuthAccessToken, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapi.OAuthAccessToken))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	return nil
}

func convert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in *oauthapi.OAuthAccessTokenList, out *oauthapiv1.OAuthAccessTokenList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapi.OAuthAccessTokenList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]oauthapiv1.OAuthAccessToken, len(in.Items))
		for i := range in.Items {
			if err := convert_api_OAuthAccessToken_To_v1_OAuthAccessToken(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in *oauthapi.OAuthAuthorizeToken, out *oauthapiv1.OAuthAuthorizeToken, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapi.OAuthAuthorizeToken))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	return nil
}

func convert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in *oauthapi.OAuthAuthorizeTokenList, out *oauthapiv1.OAuthAuthorizeTokenList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapi.OAuthAuthorizeTokenList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]oauthapiv1.OAuthAuthorizeToken, len(in.Items))
		for i := range in.Items {
			if err := convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_OAuthClient_To_v1_OAuthClient(in *oauthapi.OAuthClient, out *oauthapiv1.OAuthClient, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapi.OAuthClient))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Secret = in.Secret
	out.RespondWithChallenges = in.RespondWithChallenges
	if in.RedirectURIs != nil {
		out.RedirectURIs = make([]string, len(in.RedirectURIs))
		for i := range in.RedirectURIs {
			out.RedirectURIs[i] = in.RedirectURIs[i]
		}
	} else {
		out.RedirectURIs = nil
	}
	return nil
}

func convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in *oauthapi.OAuthClientAuthorization, out *oauthapiv1.OAuthClientAuthorization, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapi.OAuthClientAuthorization))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	return nil
}

func convert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in *oauthapi.OAuthClientAuthorizationList, out *oauthapiv1.OAuthClientAuthorizationList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapi.OAuthClientAuthorizationList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]oauthapiv1.OAuthClientAuthorization, len(in.Items))
		for i := range in.Items {
			if err := convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_OAuthClientList_To_v1_OAuthClientList(in *oauthapi.OAuthClientList, out *oauthapiv1.OAuthClientList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapi.OAuthClientList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]oauthapiv1.OAuthClient, len(in.Items))
		for i := range in.Items {
			if err := convert_api_OAuthClient_To_v1_OAuthClient(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_OAuthAccessToken_To_api_OAuthAccessToken(in *oauthapiv1.OAuthAccessToken, out *oauthapi.OAuthAccessToken, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapiv1.OAuthAccessToken))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	return nil
}

func convert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList(in *oauthapiv1.OAuthAccessTokenList, out *oauthapi.OAuthAccessTokenList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapiv1.OAuthAccessTokenList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]oauthapi.OAuthAccessToken, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_OAuthAccessToken_To_api_OAuthAccessToken(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(in *oauthapiv1.OAuthAuthorizeToken, out *oauthapi.OAuthAuthorizeToken, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapiv1.OAuthAuthorizeToken))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	return nil
}

func convert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList(in *oauthapiv1.OAuthAuthorizeTokenList, out *oauthapi.OAuthAuthorizeTokenList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapiv1.OAuthAuthorizeTokenList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]oauthapi.OAuthAuthorizeToken, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_OAuthClient_To_api_OAuthClient(in *oauthapiv1.OAuthClient, out *oauthapi.OAuthClient, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapiv1.OAuthClient))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Secret = in.Secret
	out.RespondWithChallenges = in.RespondWithChallenges
	if in.RedirectURIs != nil {
		out.RedirectURIs = make([]string, len(in.RedirectURIs))
		for i := range in.RedirectURIs {
			out.RedirectURIs[i] = in.RedirectURIs[i]
		}
	} else {
		out.RedirectURIs = nil
	}
	return nil
}

func convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(in *oauthapiv1.OAuthClientAuthorization, out *oauthapi.OAuthClientAuthorization, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapiv1.OAuthClientAuthorization))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	return nil
}

func convert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList(in *oauthapiv1.OAuthClientAuthorizationList, out *oauthapi.OAuthClientAuthorizationList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapiv1.OAuthClientAuthorizationList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]oauthapi.OAuthClientAuthorization, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_OAuthClientList_To_api_OAuthClientList(in *oauthapiv1.OAuthClientList, out *oauthapi.OAuthClientList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauthapiv1.OAuthClientList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]oauthapi.OAuthClient, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_OAuthClient_To_api_OAuthClient(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_Project_To_v1_Project(in *projectapi.Project, out *projectapiv1.Project, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapi.Project))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := convert_api_ProjectSpec_To_v1_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := convert_api_ProjectStatus_To_v1_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func convert_api_ProjectList_To_v1_ProjectList(in *projectapi.ProjectList, out *projectapiv1.ProjectList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapi.ProjectList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]projectapiv1.Project, len(in.Items))
		for i := range in.Items {
			if err := convert_api_Project_To_v1_Project(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_ProjectRequest_To_v1_ProjectRequest(in *projectapi.ProjectRequest, out *projectapiv1.ProjectRequest, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapi.ProjectRequest))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

func convert_api_ProjectSpec_To_v1_ProjectSpec(in *projectapi.ProjectSpec, out *projectapiv1.ProjectSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapi.ProjectSpec))(in)
	}
	if in.Finalizers != nil {
		out.Finalizers = make([]v1.FinalizerName, len(in.Finalizers))
		for i := range in.Finalizers {
			out.Finalizers[i] = v1.FinalizerName(in.Finalizers[i])
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func convert_api_ProjectStatus_To_v1_ProjectStatus(in *projectapi.ProjectStatus, out *projectapiv1.ProjectStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapi.ProjectStatus))(in)
	}
	out.Phase = v1.NamespacePhase(in.Phase)
	return nil
}

func convert_v1_Project_To_api_Project(in *projectapiv1.Project, out *projectapi.Project, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapiv1.Project))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ProjectSpec_To_api_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := convert_v1_ProjectStatus_To_api_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_ProjectList_To_api_ProjectList(in *projectapiv1.ProjectList, out *projectapi.ProjectList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapiv1.ProjectList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]projectapi.Project, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_Project_To_api_Project(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_ProjectRequest_To_api_ProjectRequest(in *projectapiv1.ProjectRequest, out *projectapi.ProjectRequest, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapiv1.ProjectRequest))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

func convert_v1_ProjectSpec_To_api_ProjectSpec(in *projectapiv1.ProjectSpec, out *projectapi.ProjectSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapiv1.ProjectSpec))(in)
	}
	if in.Finalizers != nil {
		out.Finalizers = make([]api.FinalizerName, len(in.Finalizers))
		for i := range in.Finalizers {
			out.Finalizers[i] = api.FinalizerName(in.Finalizers[i])
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func convert_v1_ProjectStatus_To_api_ProjectStatus(in *projectapiv1.ProjectStatus, out *projectapi.ProjectStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*projectapiv1.ProjectStatus))(in)
	}
	out.Phase = api.NamespacePhase(in.Phase)
	return nil
}

func convert_api_RouteList_To_v1_RouteList(in *routeapi.RouteList, out *routeapiv1.RouteList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*routeapi.RouteList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]routeapiv1.Route, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_RouteList_To_api_RouteList(in *routeapiv1.RouteList, out *routeapi.RouteList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*routeapiv1.RouteList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]routeapi.Route, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_ClusterNetwork_To_v1_ClusterNetwork(in *sdnapi.ClusterNetwork, out *sdnapiv1.ClusterNetwork, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*sdnapi.ClusterNetwork))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Network = in.Network
	out.HostSubnetLength = in.HostSubnetLength
	return nil
}

func convert_api_ClusterNetworkList_To_v1_ClusterNetworkList(in *sdnapi.ClusterNetworkList, out *sdnapiv1.ClusterNetworkList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*sdnapi.ClusterNetworkList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]sdnapiv1.ClusterNetwork, len(in.Items))
		for i := range in.Items {
			if err := convert_api_ClusterNetwork_To_v1_ClusterNetwork(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_HostSubnet_To_v1_HostSubnet(in *sdnapi.HostSubnet, out *sdnapiv1.HostSubnet, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*sdnapi.HostSubnet))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Host = in.Host
	out.HostIP = in.HostIP
	out.Subnet = in.Subnet
	return nil
}

func convert_api_HostSubnetList_To_v1_HostSubnetList(in *sdnapi.HostSubnetList, out *sdnapiv1.HostSubnetList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*sdnapi.HostSubnetList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]sdnapiv1.HostSubnet, len(in.Items))
		for i := range in.Items {
			if err := convert_api_HostSubnet_To_v1_HostSubnet(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_ClusterNetwork_To_api_ClusterNetwork(in *sdnapiv1.ClusterNetwork, out *sdnapi.ClusterNetwork, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*sdnapiv1.ClusterNetwork))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Network = in.Network
	out.HostSubnetLength = in.HostSubnetLength
	return nil
}

func convert_v1_ClusterNetworkList_To_api_ClusterNetworkList(in *sdnapiv1.ClusterNetworkList, out *sdnapi.ClusterNetworkList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*sdnapiv1.ClusterNetworkList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]sdnapi.ClusterNetwork, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_ClusterNetwork_To_api_ClusterNetwork(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_HostSubnet_To_api_HostSubnet(in *sdnapiv1.HostSubnet, out *sdnapi.HostSubnet, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*sdnapiv1.HostSubnet))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Host = in.Host
	out.HostIP = in.HostIP
	out.Subnet = in.Subnet
	return nil
}

func convert_v1_HostSubnetList_To_api_HostSubnetList(in *sdnapiv1.HostSubnetList, out *sdnapi.HostSubnetList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*sdnapiv1.HostSubnetList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]sdnapi.HostSubnet, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_HostSubnet_To_api_HostSubnet(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_Parameter_To_v1_Parameter(in *templateapi.Parameter, out *templateapiv1.Parameter, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*templateapi.Parameter))(in)
	}
	out.Name = in.Name
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	return nil
}

func convert_api_TemplateList_To_v1_TemplateList(in *templateapi.TemplateList, out *templateapiv1.TemplateList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*templateapi.TemplateList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]templateapiv1.Template, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_Parameter_To_api_Parameter(in *templateapiv1.Parameter, out *templateapi.Parameter, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*templateapiv1.Parameter))(in)
	}
	out.Name = in.Name
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	return nil
}

func convert_v1_TemplateList_To_api_TemplateList(in *templateapiv1.TemplateList, out *templateapi.TemplateList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*templateapiv1.TemplateList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]templateapi.Template, len(in.Items))
		for i := range in.Items {
			if err := s.Convert(&in.Items[i], &out.Items[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_Identity_To_v1_Identity(in *userapi.Identity, out *userapiv1.Identity, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapi.Identity))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ProviderName = in.ProviderName
	out.ProviderUserName = in.ProviderUserName
	if err := convert_api_ObjectReference_To_v1_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	if in.Extra != nil {
		out.Extra = make(map[string]string)
		for key, val := range in.Extra {
			out.Extra[key] = val
		}
	} else {
		out.Extra = nil
	}
	return nil
}

func convert_api_IdentityList_To_v1_IdentityList(in *userapi.IdentityList, out *userapiv1.IdentityList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapi.IdentityList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]userapiv1.Identity, len(in.Items))
		for i := range in.Items {
			if err := convert_api_Identity_To_v1_Identity(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_api_User_To_v1_User(in *userapi.User, out *userapiv1.User, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapi.User))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.FullName = in.FullName
	if in.Identities != nil {
		out.Identities = make([]string, len(in.Identities))
		for i := range in.Identities {
			out.Identities[i] = in.Identities[i]
		}
	} else {
		out.Identities = nil
	}
	if in.Groups != nil {
		out.Groups = make([]string, len(in.Groups))
		for i := range in.Groups {
			out.Groups[i] = in.Groups[i]
		}
	} else {
		out.Groups = nil
	}
	return nil
}

func convert_api_UserIdentityMapping_To_v1_UserIdentityMapping(in *userapi.UserIdentityMapping, out *userapiv1.UserIdentityMapping, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapi.UserIdentityMapping))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := convert_api_ObjectReference_To_v1_ObjectReference(&in.Identity, &out.Identity, s); err != nil {
		return err
	}
	if err := convert_api_ObjectReference_To_v1_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	return nil
}

func convert_api_UserList_To_v1_UserList(in *userapi.UserList, out *userapiv1.UserList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapi.UserList))(in)
	}
	if err := convert_api_TypeMeta_To_v1_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_api_ListMeta_To_v1_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]userapiv1.User, len(in.Items))
		for i := range in.Items {
			if err := convert_api_User_To_v1_User(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_Identity_To_api_Identity(in *userapiv1.Identity, out *userapi.Identity, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapiv1.Identity))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ProviderName = in.ProviderName
	out.ProviderUserName = in.ProviderUserName
	if err := convert_v1_ObjectReference_To_api_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	if in.Extra != nil {
		out.Extra = make(map[string]string)
		for key, val := range in.Extra {
			out.Extra[key] = val
		}
	} else {
		out.Extra = nil
	}
	return nil
}

func convert_v1_IdentityList_To_api_IdentityList(in *userapiv1.IdentityList, out *userapi.IdentityList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapiv1.IdentityList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]userapi.Identity, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_Identity_To_api_Identity(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func convert_v1_User_To_api_User(in *userapiv1.User, out *userapi.User, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapiv1.User))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.FullName = in.FullName
	if in.Identities != nil {
		out.Identities = make([]string, len(in.Identities))
		for i := range in.Identities {
			out.Identities[i] = in.Identities[i]
		}
	} else {
		out.Identities = nil
	}
	if in.Groups != nil {
		out.Groups = make([]string, len(in.Groups))
		for i := range in.Groups {
			out.Groups[i] = in.Groups[i]
		}
	} else {
		out.Groups = nil
	}
	return nil
}

func convert_v1_UserIdentityMapping_To_api_UserIdentityMapping(in *userapiv1.UserIdentityMapping, out *userapi.UserIdentityMapping, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapiv1.UserIdentityMapping))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectReference_To_api_ObjectReference(&in.Identity, &out.Identity, s); err != nil {
		return err
	}
	if err := convert_v1_ObjectReference_To_api_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	return nil
}

func convert_v1_UserList_To_api_UserList(in *userapiv1.UserList, out *userapi.UserList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*userapiv1.UserList))(in)
	}
	if err := convert_v1_TypeMeta_To_api_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := convert_v1_ListMeta_To_api_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]userapi.User, len(in.Items))
		for i := range in.Items {
			if err := convert_v1_User_To_api_User(&in.Items[i], &out.Items[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func init() {
	err := api.Scheme.AddGeneratedConversionFuncs(
		convert_api_BuildConfigList_To_v1_BuildConfigList,
		convert_api_BuildConfigSpec_To_v1_BuildConfigSpec,
		convert_api_BuildConfigStatus_To_v1_BuildConfigStatus,
		convert_api_BuildConfig_To_v1_BuildConfig,
		convert_api_BuildList_To_v1_BuildList,
		convert_api_BuildLogOptions_To_v1_BuildLogOptions,
		convert_api_BuildLog_To_v1_BuildLog,
		convert_api_BuildOutput_To_v1_BuildOutput,
		convert_api_BuildRequest_To_v1_BuildRequest,
		convert_api_BuildSource_To_v1_BuildSource,
		convert_api_BuildSpec_To_v1_BuildSpec,
		convert_api_BuildStatus_To_v1_BuildStatus,
		convert_api_BuildStrategy_To_v1_BuildStrategy,
		convert_api_BuildTriggerPolicy_To_v1_BuildTriggerPolicy,
		convert_api_Build_To_v1_Build,
		convert_api_ClusterNetworkList_To_v1_ClusterNetworkList,
		convert_api_ClusterNetwork_To_v1_ClusterNetwork,
		convert_api_ClusterPolicyBindingList_To_v1_ClusterPolicyBindingList,
		convert_api_ClusterPolicyList_To_v1_ClusterPolicyList,
		convert_api_ClusterRoleBindingList_To_v1_ClusterRoleBindingList,
		convert_api_ClusterRoleList_To_v1_ClusterRoleList,
		convert_api_ClusterRole_To_v1_ClusterRole,
		convert_api_CustomBuildStrategy_To_v1_CustomBuildStrategy,
		convert_api_DeploymentConfigList_To_v1_DeploymentConfigList,
		convert_api_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec,
		convert_api_DeploymentConfigRollback_To_v1_DeploymentConfigRollback,
		convert_api_DockerBuildStrategy_To_v1_DockerBuildStrategy,
		convert_api_EnvVarSource_To_v1_EnvVarSource,
		convert_api_EnvVar_To_v1_EnvVar,
		convert_api_GitBuildSource_To_v1_GitBuildSource,
		convert_api_GitSourceRevision_To_v1_GitSourceRevision,
		convert_api_HostSubnetList_To_v1_HostSubnetList,
		convert_api_HostSubnet_To_v1_HostSubnet,
		convert_api_IdentityList_To_v1_IdentityList,
		convert_api_Identity_To_v1_Identity,
		convert_api_ImageChangeTrigger_To_v1_ImageChangeTrigger,
		convert_api_ImageList_To_v1_ImageList,
		convert_api_ImageStreamImage_To_v1_ImageStreamImage,
		convert_api_ImageStreamList_To_v1_ImageStreamList,
		convert_api_ImageStreamTag_To_v1_ImageStreamTag,
		convert_api_ImageStream_To_v1_ImageStream,
		convert_api_IsPersonalSubjectAccessReview_To_v1_IsPersonalSubjectAccessReview,
		convert_api_ListMeta_To_v1_ListMeta,
		convert_api_LocalObjectReference_To_v1_LocalObjectReference,
		convert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList,
		convert_api_OAuthAccessToken_To_v1_OAuthAccessToken,
		convert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList,
		convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken,
		convert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList,
		convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization,
		convert_api_OAuthClientList_To_v1_OAuthClientList,
		convert_api_OAuthClient_To_v1_OAuthClient,
		convert_api_ObjectFieldSelector_To_v1_ObjectFieldSelector,
		convert_api_ObjectMeta_To_v1_ObjectMeta,
		convert_api_ObjectReference_To_v1_ObjectReference,
		convert_api_Parameter_To_v1_Parameter,
		convert_api_PolicyBindingList_To_v1_PolicyBindingList,
		convert_api_PolicyList_To_v1_PolicyList,
		convert_api_ProjectList_To_v1_ProjectList,
		convert_api_ProjectRequest_To_v1_ProjectRequest,
		convert_api_ProjectSpec_To_v1_ProjectSpec,
		convert_api_ProjectStatus_To_v1_ProjectStatus,
		convert_api_Project_To_v1_Project,
		convert_api_ResourceAccessReview_To_v1_ResourceAccessReview,
		convert_api_ResourceRequirements_To_v1_ResourceRequirements,
		convert_api_RoleBindingList_To_v1_RoleBindingList,
		convert_api_RoleList_To_v1_RoleList,
		convert_api_Role_To_v1_Role,
		convert_api_RouteList_To_v1_RouteList,
		convert_api_SourceBuildStrategy_To_v1_SourceBuildStrategy,
		convert_api_SourceControlUser_To_v1_SourceControlUser,
		convert_api_SourceRevision_To_v1_SourceRevision,
		convert_api_SubjectAccessReviewResponse_To_v1_SubjectAccessReviewResponse,
		convert_api_TemplateList_To_v1_TemplateList,
		convert_api_TypeMeta_To_v1_TypeMeta,
		convert_api_UserIdentityMapping_To_v1_UserIdentityMapping,
		convert_api_UserList_To_v1_UserList,
		convert_api_User_To_v1_User,
		convert_api_WebHookTrigger_To_v1_WebHookTrigger,
		convert_v1_BuildConfigList_To_api_BuildConfigList,
		convert_v1_BuildConfigSpec_To_api_BuildConfigSpec,
		convert_v1_BuildConfigStatus_To_api_BuildConfigStatus,
		convert_v1_BuildConfig_To_api_BuildConfig,
		convert_v1_BuildList_To_api_BuildList,
		convert_v1_BuildLogOptions_To_api_BuildLogOptions,
		convert_v1_BuildLog_To_api_BuildLog,
		convert_v1_BuildOutput_To_api_BuildOutput,
		convert_v1_BuildRequest_To_api_BuildRequest,
		convert_v1_BuildSource_To_api_BuildSource,
		convert_v1_BuildSpec_To_api_BuildSpec,
		convert_v1_BuildStatus_To_api_BuildStatus,
		convert_v1_BuildStrategy_To_api_BuildStrategy,
		convert_v1_BuildTriggerPolicy_To_api_BuildTriggerPolicy,
		convert_v1_Build_To_api_Build,
		convert_v1_ClusterNetworkList_To_api_ClusterNetworkList,
		convert_v1_ClusterNetwork_To_api_ClusterNetwork,
		convert_v1_ClusterPolicyBindingList_To_api_ClusterPolicyBindingList,
		convert_v1_ClusterPolicyList_To_api_ClusterPolicyList,
		convert_v1_ClusterRoleBindingList_To_api_ClusterRoleBindingList,
		convert_v1_ClusterRoleList_To_api_ClusterRoleList,
		convert_v1_ClusterRole_To_api_ClusterRole,
		convert_v1_CustomBuildStrategy_To_api_CustomBuildStrategy,
		convert_v1_DeploymentConfigList_To_api_DeploymentConfigList,
		convert_v1_DeploymentConfigRollbackSpec_To_api_DeploymentConfigRollbackSpec,
		convert_v1_DeploymentConfigRollback_To_api_DeploymentConfigRollback,
		convert_v1_DockerBuildStrategy_To_api_DockerBuildStrategy,
		convert_v1_EnvVarSource_To_api_EnvVarSource,
		convert_v1_EnvVar_To_api_EnvVar,
		convert_v1_GitBuildSource_To_api_GitBuildSource,
		convert_v1_GitSourceRevision_To_api_GitSourceRevision,
		convert_v1_HostSubnetList_To_api_HostSubnetList,
		convert_v1_HostSubnet_To_api_HostSubnet,
		convert_v1_IdentityList_To_api_IdentityList,
		convert_v1_Identity_To_api_Identity,
		convert_v1_ImageChangeTrigger_To_api_ImageChangeTrigger,
		convert_v1_ImageList_To_api_ImageList,
		convert_v1_ImageStreamImage_To_api_ImageStreamImage,
		convert_v1_ImageStreamList_To_api_ImageStreamList,
		convert_v1_ImageStreamTag_To_api_ImageStreamTag,
		convert_v1_ImageStream_To_api_ImageStream,
		convert_v1_IsPersonalSubjectAccessReview_To_api_IsPersonalSubjectAccessReview,
		convert_v1_ListMeta_To_api_ListMeta,
		convert_v1_LocalObjectReference_To_api_LocalObjectReference,
		convert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList,
		convert_v1_OAuthAccessToken_To_api_OAuthAccessToken,
		convert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList,
		convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken,
		convert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList,
		convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization,
		convert_v1_OAuthClientList_To_api_OAuthClientList,
		convert_v1_OAuthClient_To_api_OAuthClient,
		convert_v1_ObjectFieldSelector_To_api_ObjectFieldSelector,
		convert_v1_ObjectMeta_To_api_ObjectMeta,
		convert_v1_ObjectReference_To_api_ObjectReference,
		convert_v1_Parameter_To_api_Parameter,
		convert_v1_PolicyBindingList_To_api_PolicyBindingList,
		convert_v1_PolicyList_To_api_PolicyList,
		convert_v1_ProjectList_To_api_ProjectList,
		convert_v1_ProjectRequest_To_api_ProjectRequest,
		convert_v1_ProjectSpec_To_api_ProjectSpec,
		convert_v1_ProjectStatus_To_api_ProjectStatus,
		convert_v1_Project_To_api_Project,
		convert_v1_ResourceAccessReview_To_api_ResourceAccessReview,
		convert_v1_ResourceRequirements_To_api_ResourceRequirements,
		convert_v1_RoleBindingList_To_api_RoleBindingList,
		convert_v1_RoleList_To_api_RoleList,
		convert_v1_Role_To_api_Role,
		convert_v1_RouteList_To_api_RouteList,
		convert_v1_SourceBuildStrategy_To_api_SourceBuildStrategy,
		convert_v1_SourceControlUser_To_api_SourceControlUser,
		convert_v1_SourceRevision_To_api_SourceRevision,
		convert_v1_SubjectAccessReviewResponse_To_api_SubjectAccessReviewResponse,
		convert_v1_TemplateList_To_api_TemplateList,
		convert_v1_TypeMeta_To_api_TypeMeta,
		convert_v1_UserIdentityMapping_To_api_UserIdentityMapping,
		convert_v1_UserList_To_api_UserList,
		convert_v1_User_To_api_User,
		convert_v1_WebHookTrigger_To_api_WebHookTrigger,
	)
	if err != nil {
		// If one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

// AUTO-GENERATED FUNCTIONS END HERE
