// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/installpackage/v1alpha1"
	scheme "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InstalledPackagesGetter has a method to return a InstalledPackageInterface.
// A group's client should implement this interface.
type InstalledPackagesGetter interface {
	InstalledPackages(namespace string) InstalledPackageInterface
}

// InstalledPackageInterface has methods to work with InstalledPackage resources.
type InstalledPackageInterface interface {
	Create(*v1alpha1.InstalledPackage) (*v1alpha1.InstalledPackage, error)
	Update(*v1alpha1.InstalledPackage) (*v1alpha1.InstalledPackage, error)
	UpdateStatus(*v1alpha1.InstalledPackage) (*v1alpha1.InstalledPackage, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.InstalledPackage, error)
	List(opts v1.ListOptions) (*v1alpha1.InstalledPackageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstalledPackage, err error)
	InstalledPackageExpansion
}

// installedPackages implements InstalledPackageInterface
type installedPackages struct {
	client rest.Interface
	ns     string
}

// newInstalledPackages returns a InstalledPackages
func newInstalledPackages(c *InstallV1alpha1Client, namespace string) *installedPackages {
	return &installedPackages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the installedPackage, and returns the corresponding installedPackage object, and an error if there is any.
func (c *installedPackages) Get(name string, options v1.GetOptions) (result *v1alpha1.InstalledPackage, err error) {
	result = &v1alpha1.InstalledPackage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("installedpackages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InstalledPackages that match those selectors.
func (c *installedPackages) List(opts v1.ListOptions) (result *v1alpha1.InstalledPackageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.InstalledPackageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("installedpackages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested installedPackages.
func (c *installedPackages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("installedpackages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a installedPackage and creates it.  Returns the server's representation of the installedPackage, and an error, if there is any.
func (c *installedPackages) Create(installedPackage *v1alpha1.InstalledPackage) (result *v1alpha1.InstalledPackage, err error) {
	result = &v1alpha1.InstalledPackage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("installedpackages").
		Body(installedPackage).
		Do().
		Into(result)
	return
}

// Update takes the representation of a installedPackage and updates it. Returns the server's representation of the installedPackage, and an error, if there is any.
func (c *installedPackages) Update(installedPackage *v1alpha1.InstalledPackage) (result *v1alpha1.InstalledPackage, err error) {
	result = &v1alpha1.InstalledPackage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("installedpackages").
		Name(installedPackage.Name).
		Body(installedPackage).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *installedPackages) UpdateStatus(installedPackage *v1alpha1.InstalledPackage) (result *v1alpha1.InstalledPackage, err error) {
	result = &v1alpha1.InstalledPackage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("installedpackages").
		Name(installedPackage.Name).
		SubResource("status").
		Body(installedPackage).
		Do().
		Into(result)
	return
}

// Delete takes name of the installedPackage and deletes it. Returns an error if one occurs.
func (c *installedPackages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("installedpackages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *installedPackages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("installedpackages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched installedPackage.
func (c *installedPackages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstalledPackage, err error) {
	result = &v1alpha1.InstalledPackage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("installedpackages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}