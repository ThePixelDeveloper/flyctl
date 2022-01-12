package volumes

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/superfly/flyctl/internal/cli/internal/command"
	"github.com/superfly/flyctl/internal/cli/internal/config"
	"github.com/superfly/flyctl/internal/cli/internal/flag"
	"github.com/superfly/flyctl/internal/cli/internal/render"
	"github.com/superfly/flyctl/internal/client"
	"github.com/superfly/flyctl/pkg/iostreams"
)

func newShow() *cobra.Command {
	const (
		long = `how details of an app's volume. Requires the volume's ID
number to operate. This can be found through the volumes list command`

		short = "Show details of an app's volume"
	)

	cmd := command.New("show <id>", short, long, runShow)
	cmd.Args = cobra.ExactArgs(1)

	flag.Add(
		cmd,
	)

	return cmd
}

func runShow(ctx context.Context) (err error) {
	cfg := config.FromContext(ctx)
	client := client.FromContext(ctx).API()

	volumeID := flag.FirstArg(ctx)

	volume, err := client.GetVolume(ctx, volumeID)
	if err != nil {
		return
	}

	out := iostreams.FromContext(ctx).Out

	if cfg.JSONOutput {
		_ = render.JSON(out, volume)
		return
	}

	fmt.Fprintf(out, "%10s: %s\n", "ID", volume.ID)
	fmt.Fprintf(out, "%10s: %s\n", "Name", volume.Name)
	fmt.Fprintf(out, "%10s: %s\n", "App", volume.App.Name)
	fmt.Fprintf(out, "%10s: %s\n", "Region", volume.Region)
	fmt.Fprintf(out, "%10s: %s\n", "Zone", volume.Host.ID)
	fmt.Fprintf(out, "%10s: %d\n", "Size GB", volume.SizeGb)
	fmt.Fprintf(out, "%10s: %t\n", "Encrypted", volume.Encrypted)
	fmt.Fprintf(out, "%10s: %s\n", "Created at", volume.CreatedAt.Format(time.RFC822))

	return
}
