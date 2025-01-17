package scan

import (
	"os"
	"strings"

	biutils "github.com/jfrog/build-info-go/utils"
	commandsutils "github.com/jfrog/jfrog-cli-core/v2/artifactory/commands/utils"
	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-core/v2/common/spec"
	corecommondocs "github.com/jfrog/jfrog-cli-core/v2/docs/common"
	coreconfig "github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-cli-core/v2/utils/coreutils"
	"github.com/jfrog/jfrog-cli-core/v2/xray/commands/audit"
	_go "github.com/jfrog/jfrog-cli-core/v2/xray/commands/audit/go"
	"github.com/jfrog/jfrog-cli-core/v2/xray/commands/audit/java"
	"github.com/jfrog/jfrog-cli-core/v2/xray/commands/audit/npm"
	"github.com/jfrog/jfrog-cli-core/v2/xray/commands/audit/python"
	"github.com/jfrog/jfrog-cli-core/v2/xray/commands/scan"
	"github.com/jfrog/jfrog-cli/docs/common"
	auditdocs "github.com/jfrog/jfrog-cli/docs/scan/audit"
	auditgodocs "github.com/jfrog/jfrog-cli/docs/scan/auditgo"
	auditgradledocs "github.com/jfrog/jfrog-cli/docs/scan/auditgradle"
	"github.com/jfrog/jfrog-cli/docs/scan/auditmvn"
	auditnpmdocs "github.com/jfrog/jfrog-cli/docs/scan/auditnpm"
	auditpipdocs "github.com/jfrog/jfrog-cli/docs/scan/auditpip"
	auditpipenvdocs "github.com/jfrog/jfrog-cli/docs/scan/auditpipenv"
	buildscandocs "github.com/jfrog/jfrog-cli/docs/scan/buildscan"
	dockerscandocs "github.com/jfrog/jfrog-cli/docs/scan/dockerscan"
	scandocs "github.com/jfrog/jfrog-cli/docs/scan/scan"
	"github.com/jfrog/jfrog-cli/utils/cliutils"
	"github.com/urfave/cli"

	"github.com/jfrog/jfrog-client-go/utils/errorutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

const auditScanCategory = "Audit & Scan"

func GetCommands() []cli.Command {
	return cliutils.GetSortedCommands(cli.CommandsByName{
		{
			Name:         "audit",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.Audit),
			Aliases:      []string{"audit"},
			Description:  auditdocs.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("audit", auditdocs.GetDescription(), auditdocs.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       AuditCmd,
		},
		{
			Name:         "audit-mvn",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.AuditMvn),
			Aliases:      []string{"am"},
			Description:  auditmvn.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("audit-mvn", auditmvn.GetDescription(), auditmvn.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       AuditMvnCmd,
		},
		{
			Name:         "audit-gradle",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.AuditGradle),
			Aliases:      []string{"ag"},
			Description:  auditgradledocs.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("audit-gradle", auditgradledocs.GetDescription(), auditgradledocs.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       AuditGradleCmd,
		},
		{
			Name:         "audit-npm",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.AuditNpm),
			Aliases:      []string{"an"},
			Description:  auditnpmdocs.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("audit-npm", auditnpmdocs.GetDescription(), auditnpmdocs.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       AuditNpmCmd,
		},
		{
			Name:         "audit-go",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.AuditGo),
			Aliases:      []string{"ago"},
			Description:  auditgodocs.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("audit-go", auditgodocs.GetDescription(), auditgodocs.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       AuditGoCmd,
		},
		{
			Name:         "audit-pip",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.AuditPip),
			Aliases:      []string{"ap"},
			Description:  auditpipdocs.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("audit-pip", auditpipdocs.GetDescription(), auditpipdocs.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       AuditPipCmd,
		},
		{
			Name:         "audit-pipenv",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.AuditPip),
			Aliases:      []string{"ape"},
			Description:  auditpipenvdocs.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("audit-pipenv", auditpipenvdocs.GetDescription(), auditpipenvdocs.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       AuditPipenvCmd,
		},
		{
			Name:         "scan",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.XrScan),
			Aliases:      []string{"s"},
			Description:  scandocs.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("scan", scandocs.GetDescription(), scandocs.Usage),
			UsageText:    scandocs.GetArguments(),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       ScanCmd,
		},
		{
			Name:         "build-scan",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.BuildScan),
			Aliases:      []string{"bs"},
			Description:  buildscandocs.GetDescription(),
			UsageText:    buildscandocs.GetArguments(),
			HelpName:     corecommondocs.CreateUsage("build-scan", buildscandocs.GetDescription(), buildscandocs.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       BuildScan,
		},
		{
			Name:         "docker",
			Category:     auditScanCategory,
			Flags:        cliutils.GetCommandFlags(cliutils.DockerScan),
			Description:  dockerscandocs.GetDescription(),
			HelpName:     corecommondocs.CreateUsage("docker scan", dockerscandocs.GetDescription(), dockerscandocs.Usage),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommondocs.CreateBashCompletionFunc(),
			Action:       DockerCommand,
		},
	})
}

func AuditCmd(c *cli.Context) error {
	wd, err := os.Getwd()
	if errorutils.CheckError(err) != nil {
		return err
	}
	detectedTechnologies, err := coreutils.DetectTechnologies(wd, false, false)
	if err != nil {
		return err
	}
	for tech := range detectedTechnologies {
		log.Info(string(tech) + " detected.")
		switch tech {
		case coreutils.Maven:
			err = AuditMvnCmd(c)
		case coreutils.Gradle:
			err = AuditGradleCmd(c)
		case coreutils.Npm:
			err = AuditNpmCmd(c)
		case coreutils.Go:
			err = AuditGoCmd(c)
		case coreutils.Pip:
			err = AuditPipCmd(c)
		case coreutils.Pipenv:
			err = AuditPipenvCmd(c)
		default:
			log.Info(string(tech), " is currently not supported")
		}
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func AuditMvnCmd(c *cli.Context) error {
	genericAuditCmd, err := createGenericAuditCmd(c)
	if err != nil {
		return err
	}
	xrAuditMvnCmd := java.NewAuditMavenCommand(*genericAuditCmd).SetInsecureTls(c.Bool(cliutils.InsecureTls))
	return commands.Exec(xrAuditMvnCmd)
}

func AuditGradleCmd(c *cli.Context) error {
	genericAuditCmd, err := createGenericAuditCmd(c)
	if err != nil {
		return err
	}
	xrAuditGradleCmd := java.NewAuditGradleCommand(*genericAuditCmd).SetExcludeTestDeps(c.Bool(cliutils.ExcludeTestDeps)).SetUseWrapper(c.Bool(cliutils.UseWrapper))
	return commands.Exec(xrAuditGradleCmd)
}

func AuditNpmCmd(c *cli.Context) error {
	genericAuditCmd, err := createGenericAuditCmd(c)
	if err != nil {
		return err
	}
	var typeRestriction = biutils.All
	switch c.String("dep-type") {
	case "devOnly":
		typeRestriction = biutils.DevOnly
	case "prodOnly":
		typeRestriction = biutils.ProdOnly
	}
	auditNpmCmd := npm.NewAuditNpmCommand(*genericAuditCmd).SetNpmTypeRestriction(typeRestriction)
	return commands.Exec(auditNpmCmd)
}

func AuditGoCmd(c *cli.Context) error {
	genericAuditCmd, err := createGenericAuditCmd(c)
	if err != nil {
		return err
	}
	auditGoCmd := _go.NewAuditGoCommand(*genericAuditCmd)
	return commands.Exec(auditGoCmd)
}

func AuditPipCmd(c *cli.Context) error {
	genericAuditCmd, err := createGenericAuditCmd(c)
	if err != nil {
		return err
	}
	auditPipCmd := python.NewAuditPipCommand(*genericAuditCmd)
	return commands.Exec(auditPipCmd)
}

func AuditPipenvCmd(c *cli.Context) error {
	genericAuditCmd, err := createGenericAuditCmd(c)
	if err != nil {
		return err
	}
	auditPipenvCmd := python.NewAuditPipenvCommand(*genericAuditCmd)
	return commands.Exec(auditPipenvCmd)
}

func createGenericAuditCmd(c *cli.Context) (*audit.AuditCommand, error) {
	auditCmd := audit.NewAuditCommand()
	err := validateXrayContext(c)
	if err != nil {
		return nil, err
	}
	serverDetails, err := createServerDetailsWithConfigOffer(c)
	if err != nil {
		return nil, err
	}
	format, err := commandsutils.GetXrayOutputFormat(c.String("format"))
	if err != nil {
		return nil, err
	}

	auditCmd.SetServerDetails(serverDetails).
		SetOutputFormat(format).
		SetTargetRepoPath(addTrailingSlashToRepoPathIfNeeded(c)).
		SetProject(c.String("project")).
		SetIncludeVulnerabilities(shouldIncludeVulnerabilities(c)).
		SetIncludeLicenses(c.Bool("licenses")).
		SetFail(c.BoolT("fail"))

	if c.String("watches") != "" {
		auditCmd.SetWatches(strings.Split(c.String("watches"), ","))
	}
	return auditCmd, err
}

func ScanCmd(c *cli.Context) error {
	err := validateXrayContext(c)
	if err != nil {
		return err
	}
	serverDetails, err := createServerDetailsWithConfigOffer(c)
	if err != nil {
		return err
	}
	var specFile *spec.SpecFiles
	if c.IsSet("spec") {
		specFile, err = cliutils.GetFileSystemSpec(c)
	} else {
		specFile, err = createDefaultScanSpec(c, addTrailingSlashToRepoPathIfNeeded(c))
	}
	if err != nil {
		return err
	}
	err = spec.ValidateSpec(specFile.Files, false, false)
	if err != nil {
		return err
	}
	threads, err := cliutils.GetThreadsCount(c)
	if err != nil {
		return err
	}
	format, err := commandsutils.GetXrayOutputFormat(c.String("format"))
	if err != nil {
		return err
	}
	cliutils.FixWinPathsForFileSystemSourcedCmds(specFile, c)
	scanCmd := scan.NewScanCommand().SetServerDetails(serverDetails).SetThreads(threads).SetSpec(specFile).SetOutputFormat(format).
		SetProject(c.String("project")).SetIncludeVulnerabilities(shouldIncludeVulnerabilities(c)).
		SetIncludeLicenses(c.Bool("licenses")).SetFail(c.BoolT("fail"))
	if c.String("watches") != "" {
		scanCmd.SetWatches(strings.Split(c.String("watches"), ","))
	}
	return commands.Exec(scanCmd)
}

// Scan published builds with Xray
func BuildScan(c *cli.Context) error {
	if c.NArg() > 2 {
		return cliutils.WrongNumberOfArgumentsHandler(c)
	}
	buildConfiguration := cliutils.CreateBuildConfiguration(c)
	if err := buildConfiguration.ValidateBuildParams(); err != nil {
		return err
	}

	serverDetails, err := createServerDetailsWithConfigOffer(c)
	if err != nil {
		return err
	}
	err = validateXrayContext(c)
	if err != nil {
		return err
	}
	format, err := commandsutils.GetXrayOutputFormat(c.String("format"))
	if err != nil {
		return err
	}
	buildScanCmd := scan.NewBuildScanCommand().SetServerDetails(serverDetails).SetFailBuild(c.BoolT("fail")).SetBuildConfiguration(buildConfiguration).
		SetIncludeVulnerabilities(c.Bool("vuln")).SetOutputFormat(format)
	return commands.Exec(buildScanCmd)
}

func DockerCommand(c *cli.Context) error {
	args := cliutils.ExtractCommand(c)
	cmdName := ""
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			cmdName = arg
			break
		}
	}
	switch cmdName {
	// Commands accepted by 'jf docker'.
	case "scan":
		return dockerScan(c)
	default:
		return errorutils.CheckErrorf("'jf docker %s' command is currently not supported by JFrog CLI", cmdName)
	}
}

func dockerScan(c *cli.Context) error {
	if c.NArg() != 2 {
		return cliutils.WrongNumberOfArgumentsHandler(c)
	}
	serverDetails, err := createServerDetailsWithConfigOffer(c)
	if err != nil {
		return err
	}
	format, err := commandsutils.GetXrayOutputFormat(c.String("format"))
	if err != nil {
		return err
	}
	containerScanCommand := scan.NewDockerScanCommand()
	containerScanCommand.SetServerDetails(serverDetails).SetOutputFormat(format).SetProject(c.String("project")).
		SetIncludeVulnerabilities(shouldIncludeVulnerabilities(c)).SetIncludeLicenses(c.Bool("licenses")).SetFail(c.BoolT("fail"))
	if c.String("watches") != "" {
		containerScanCommand.SetWatches(strings.Split(c.String("watches"), ","))
	}
	containerScanCommand.SetImageTag(c.Args().Get(1))
	return commands.Exec(containerScanCommand)
}

func addTrailingSlashToRepoPathIfNeeded(c *cli.Context) string {
	repoPath := c.String("repo-path")
	if repoPath != "" && !strings.Contains(repoPath, "/") {
		// In case a only repo name was provided (no path) we are adding a trailing slash.
		repoPath += "/"
	}
	return repoPath
}

func createDefaultScanSpec(c *cli.Context, defaultTarget string) (*spec.SpecFiles, error) {
	return spec.NewBuilder().
		Pattern(c.Args().Get(0)).
		Target(defaultTarget).
		Recursive(c.BoolT("recursive")).
		Exclusions(cliutils.GetStringsArrFlagValue(c, "exclusions")).
		Regexp(c.Bool("regexp")).
		Ant(c.Bool("ant")).
		IncludeDirs(c.Bool("include-dirs")).
		BuildSpec(), nil
}

func createServerDetailsWithConfigOffer(c *cli.Context) (*coreconfig.ServerDetails, error) {
	return cliutils.CreateServerDetailsWithConfigOffer(c, true, "xr")
}

func shouldIncludeVulnerabilities(c *cli.Context) bool {
	// If no context was provided by the user, no Violations will be triggered by Xray, so include general vulnerabilities in the command output
	return c.String("watches") == "" && !isProjectProvided(c) && c.String("repo-path") == ""
}

func validateXrayContext(c *cli.Context) error {
	contextFlag := 0
	if c.String("watches") != "" {
		contextFlag++
	}
	if isProjectProvided(c) {
		contextFlag++
	}
	if c.String("repo-path") != "" {
		contextFlag++
	}
	if contextFlag > 1 {
		return errorutils.CheckErrorf("only one of the following flags can be supplied: --watches, --project or --repo-path")
	}
	return nil
}

func isProjectProvided(c *cli.Context) bool {
	return c.String("project") != "" || os.Getenv(coreutils.Project) != ""
}
