package testclient

import (
	"testing"

	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/errors"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client/testclient"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
)

func TestNewClient(t *testing.T) {
	o := testclient.NewObjects(kapi.Scheme, kapi.Scheme)
	if err := testclient.AddObjectsFromPath("../../../test/integration/fixtures/test-deployment-config.json", o, kapi.Scheme); err != nil {
		t.Fatal(err)
	}
	oc, _ := NewFixtureClients(o)
	list, err := oc.DeploymentConfigs("test").List(labels.Everything(), fields.Everything())
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Items) != 1 {
		t.Fatalf("unexpected list %#v", list)
	}

	// same result
	list, err = oc.DeploymentConfigs("test").List(labels.Everything(), fields.Everything())
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Items) != 1 {
		t.Fatalf("unexpected list %#v", list)
	}
	t.Logf("list: %#v", list)
}

func TestErrors(t *testing.T) {
	o := testclient.NewObjects(kapi.Scheme, kapi.Scheme)
	o.Add(&kapi.List{
		Items: []runtime.Object{
			&(errors.NewNotFound("DeploymentConfigList", "").(*errors.StatusError).ErrStatus),
			&(errors.NewForbidden("DeploymentConfigList", "", nil).(*errors.StatusError).ErrStatus),
		},
	})
	oc, _ := NewFixtureClients(o)
	_, err := oc.DeploymentConfigs("test").List(labels.Everything(), fields.Everything())
	if !errors.IsNotFound(err) {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Logf("error: %#v", err.(*errors.StatusError).Status())
	_, err = oc.DeploymentConfigs("test").List(labels.Everything(), fields.Everything())
	if !errors.IsForbidden(err) {
		t.Fatalf("unexpected error: %v", err)
	}
}
