package commands

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/v1alpha2"
	"github.com/oam-dev/kubevela/apis/types"
	"github.com/oam-dev/kubevela/pkg/commands/util"
	"github.com/oam-dev/kubevela/pkg/serverlib"
)

func TestUp(t *testing.T) {
	client := fake.NewFakeClientWithScheme(scheme.Scheme)
	ioStream := util.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	env := types.EnvMeta{
		Name:      "up",
		Namespace: "env-up",
		Issuer:    "up",
	}
	o := serverlib.AppfileOptions{
		Kubecli: client,
		IO:      ioStream,
		Env:     &env,
	}
	app := &v1alpha2.Application{}
	app.Name = "app-up"
	msg := o.Info(app)
	assert.Contains(t, msg, "App has been deployed")
	assert.Contains(t, msg, fmt.Sprintf("App status: vela status %s", app.Name))
}

func TestNewUpCommandPersistentPreRunE(t *testing.T) {
	io := util.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	fakeC := types.Args{}
	cmd := NewUpCommand(fakeC, io)
	assert.Nil(t, cmd.PersistentPreRunE(new(cobra.Command), []string{}))
}
