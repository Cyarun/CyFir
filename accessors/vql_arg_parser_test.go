package accessors_test

import (
	"context"
	"testing"

	"github.com/Velocidex/ordereddict"
	"github.com/Cyarun/CyFir/accessors"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/constants"
	"github.com/Cyarun/CyFir/file_store/api"
	"github.com/Cyarun/CyFir/file_store/path_specs"
	"github.com/Cyarun/CyFir/json"
	vql_subsystem "github.com/Cyarun/CyFir/vql"
	"github.com/Cyarun/CyFir/vql/acl_managers"
	"github.com/Cyarun/CyFir/vtesting/assert"
	"github.com/Cyarun/CyFir/vtesting/goldie"
	"www.velocidex.com/golang/vfilter"
	"www.velocidex.com/golang/vfilter/arg_parser"

	_ "github.com/Cyarun/CyFir/accessors/file"
)

type testStruct struct {
	Path     *accessors.OSPath `vfilter:"required,field=path"`
	Accessor string            `vfilter:"optional,field=accessor"`
}

type testCases struct {
	name string
	in   vfilter.Any
}

var testcases = []testCases{
	{name: "Simple Path",
		in: []string{
			"Hello", "World",
		}},
	{name: "Path With {",
		in: []string{
			"Hello", "{this is a test}",
		}},
	{name: "FSPathSpec",
		in: path_specs.NewUnsafeFilestorePath("Hello", "World")},
	{name: "FSPathSpec With type",
		in: path_specs.NewUnsafeFilestorePath("Hello", "World").
			SetType(api.PATH_TYPE_FILESTORE_DOWNLOAD_ZIP)},
	{name: "DSPathSpec",
		in: path_specs.NewUnsafeDatastorePath("Hello", "World")},

	{name: "DSPathSpec With Type",
		in: path_specs.NewUnsafeDatastorePath("Hello", "World").
			SetType(api.PATH_TYPE_DATASTORE_PROTO)},

	{name: "OSPath",
		in: accessors.MustNewGenericOSPath("/foo/bar")},

	{name: "PathSpec",
		in: accessors.MustNewGenericOSPath("/foo/bar").PathSpec()},

	{name: "Serialized PathSpec",
		in: `{"Path": "/foo/bar.txt", "Accessor": "zip", "DelegatePath": "/tmp/file.zip", "DelegateAccessor": "file"}`},

	{name: "Multiple parts of mixed type",
		in: []vfilter.Any{accessors.MustNewGenericOSPath("/foo/bar"), "Hello.txt"}},

	// Just join all parts
	{name: "Multiple parts of mixed type",
		in: []vfilter.Any{"/root/home", accessors.MustNewGenericOSPath("/foo/bar"), "Hello.txt"}},

	{name: "Multiple parts of mixed type 2",
		in: []vfilter.Any{"/root/home", `{"Path": "/a/b"}`, "Hello.txt"}},
}

func TestVQLParsing(t *testing.T) {
	config_obj := &config_proto.Config{}

	// To make this test run on Linux and Windows the same we use a
	// neutral accessor.
	device_manager := accessors.GetDefaultDeviceManager(config_obj).Copy()
	device_manager.Register(accessors.DescribeAccessor(
		accessors.NewVirtualFilesystemAccessor(accessors.MustNewLinuxOSPath("")),
		accessors.AccessorDescriptor{
			Name: "virt",
		}))

	ctx := context.Background()
	scope := vql_subsystem.MakeScope().
		AppendVars(ordereddict.NewDict().
			Set(vql_subsystem.ACL_MANAGER_VAR, acl_managers.NullACLManager{}).
			Set(constants.SCOPE_DEVICE_MANAGER, device_manager),
		)

	result := ordereddict.NewDict()

	for _, testcase := range testcases {
		args := ordereddict.NewDict().
			Set("accessor", "virt").
			Set("path", testcase.in)
		arg := &testStruct{}
		err := arg_parser.ExtractArgsWithContext(ctx, scope, args, arg)
		assert.NoError(t, err)

		result.Set(testcase.name, ordereddict.NewDict().
			Set("Components", arg.Path.Components).
			Set("PathSpec", arg.Path.PathSpec()))
	}
	goldie.Assert(t, "TestVQLParsing", json.MustMarshalIndent(result))
}
