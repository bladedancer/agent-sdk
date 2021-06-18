/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1alpha1"
)

type ReleaseTagMergeFunc func(*v1alpha1.ReleaseTag, *v1alpha1.ReleaseTag) (*v1alpha1.ReleaseTag, error)

// Merge builds a merge option for an update operation
func ReleaseTagMerge(f ReleaseTagMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.ReleaseTag{}, &v1alpha1.ReleaseTag{}

		switch t := prev.(type) {
		case *v1alpha1.ReleaseTag:
			p = t
		case *apiv1.ResourceInstance:
			err := p.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialise prev resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise prev resource, unxexpected resource type: %T", t)
		}

		switch t := new.(type) {
		case *v1alpha1.ReleaseTag:
			n = t
		case *apiv1.ResourceInstance:
			err := n.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialize new resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise new resource, unxexpected resource type: %T", t)
		}

		return f(p, n)
	})
}

// ReleaseTagClient -
type ReleaseTagClient struct {
	client v1.Scoped
}

// UnscopedReleaseTagClient -
type UnscopedReleaseTagClient struct {
	client v1.Unscoped
}

// NewReleaseTagClient -
func NewReleaseTagClient(c v1.Base) (*UnscopedReleaseTagClient, error) {

	client, err := c.ForKind(v1alpha1.ReleaseTagGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedReleaseTagClient{client}, nil

}

// WithScope -
func (c *UnscopedReleaseTagClient) WithScope(scope string) *ReleaseTagClient {
	return &ReleaseTagClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedReleaseTagClient) Get(name string) (*v1alpha1.ReleaseTag, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.ReleaseTag{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedReleaseTagClient) Update(res *v1alpha1.ReleaseTag, opts ...v1.UpdateOption) (*v1alpha1.ReleaseTag, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.ReleaseTag{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *ReleaseTagClient) List(options ...v1.ListOptions) ([]*v1alpha1.ReleaseTag, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.ReleaseTag, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.ReleaseTag{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *ReleaseTagClient) Get(name string) (*v1alpha1.ReleaseTag, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.ReleaseTag{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *ReleaseTagClient) Delete(res *v1alpha1.ReleaseTag) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *ReleaseTagClient) Create(res *v1alpha1.ReleaseTag, opts ...v1.CreateOption) (*v1alpha1.ReleaseTag, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.ReleaseTag{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *ReleaseTagClient) Update(res *v1alpha1.ReleaseTag, opts ...v1.UpdateOption) (*v1alpha1.ReleaseTag, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.ReleaseTag{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
