package auth

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/confluentinc/ccloud-sdk-go"
	sdkMock "github.com/confluentinc/ccloud-sdk-go/mock"
	orgv1 "github.com/confluentinc/ccloudapis/org/v1"

	"github.com/confluentinc/mds-sdk-go"
	mdsMock "github.com/confluentinc/mds-sdk-go/mock"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/config"
	"github.com/confluentinc/cli/internal/pkg/log"
	cliMock "github.com/confluentinc/cli/mock"
)

func TestCredentialsOverride(t *testing.T) {
	req := require.New(t)
	currentEmail := os.Getenv("XX_CCLOUD_EMAIL")
	currentPassword := os.Getenv("XX_CCLOUD_PASSWORD")

	os.Setenv("XX_CCLOUD_EMAIL", "test-email")
	os.Setenv("XX_CCLOUD_PASSWORD", "test-password")

	prompt := prompt("cody@confluent.io", "iambatman")
	auth := &sdkMock.Auth{
		LoginFunc: func(ctx context.Context, idToken string, username string, password string) (string, error) {
			return "y0ur.jwt.T0kEn", nil
		},
		UserFunc: func(ctx context.Context) (*orgv1.GetUserReply, error) {
			return &orgv1.GetUserReply{
				User: &orgv1.User{
					Id:        23,
					Email:     "test-email",
					FirstName: "Cody",
				},
				Accounts: []*orgv1.Account{{Id: "a-595", Name: "Default"}},
			}, nil
		},
	}
	user := &sdkMock.User{
		CheckEmailFunc: func(ctx context.Context, user *orgv1.User) (*orgv1.User, error) {
			return &orgv1.User{
				Email: "test-email",
			}, nil
		},
	}
	cmds, cfg := newAuthCommand(prompt, auth, user, "ccloud", req)

	output, err := pcmd.ExecuteCommand(cmds.Commands[0])
	req.NoError(err)
	req.Contains(output, "Logged in as test-email")

	req.Equal("y0ur.jwt.T0kEn", cfg.AuthToken)
	req.Equal(&orgv1.User{Id: 23, Email: "test-email", FirstName: "Cody"}, cfg.Auth.User)

	os.Setenv("XX_CCLOUD_EMAIL", currentEmail)
	os.Setenv("XX_CCLOUD_PASSWORD", currentPassword)
}

func TestLoginSuccess(t *testing.T) {
	req := require.New(t)

	prompt := prompt("cody@confluent.io", "iambatman")
	auth := &sdkMock.Auth{
		LoginFunc: func(ctx context.Context, idToken string, username string, password string) (string, error) {
			return "y0ur.jwt.T0kEn", nil
		},
		UserFunc: func(ctx context.Context) (*orgv1.GetUserReply, error) {
			return &orgv1.GetUserReply{
				User: &orgv1.User{
					Id:        23,
					Email:     "cody@confluent.io",
					FirstName: "Cody",
				},
				Accounts: []*orgv1.Account{{Id: "a-595", Name: "Default"}},
			}, nil
		},
	}
	user := &sdkMock.User{
		CheckEmailFunc: func(ctx context.Context, user *orgv1.User) (*orgv1.User, error) {
			return &orgv1.User{
				Email: "test-email",
			}, nil
		},
	}

	suite := []struct {
		cliName string
		args    []string
	}{
		{
			cliName: "ccloud",
		},
		{
			cliName: "confluent",
			args: []string{
				"--url=http://localhost:8090",
			},
		},
	}

	for _, s := range suite {
		// Login to the CLI control plane
		cmds, cfg := newAuthCommand(prompt, auth, user, s.cliName, req)

		output, err := pcmd.ExecuteCommand(cmds.Commands[0], s.args...)
		req.NoError(err)
		req.Contains(output, "Logged in as cody@confluent.io")

		req.Equal("y0ur.jwt.T0kEn", cfg.AuthToken)
		contextName := fmt.Sprintf("login-cody@confluent.io-%s", cfg.AuthURL)
		req.Contains(cfg.Platforms, cfg.AuthURL)
		req.Equal(cfg.AuthURL, cfg.Platforms[cfg.AuthURL].Server)
		req.Contains(cfg.Credentials, "username-cody@confluent.io")
		req.Equal("cody@confluent.io", cfg.Credentials["username-cody@confluent.io"].Username)
		req.Contains(cfg.Contexts, contextName)
		req.Equal(cfg.AuthURL, cfg.Contexts[contextName].Platform)
		req.Equal("username-cody@confluent.io", cfg.Contexts[contextName].Credential)
		if s.cliName == "ccloud" {
			// MDS doesn't set some things like cfg.Auth.User since e.g. an MDS user != an orgv1 (ccloud) User
			req.Equal(&orgv1.User{Id: 23, Email: "cody@confluent.io", FirstName: "Cody"}, cfg.Auth.User)
		} else {
			req.Equal("http://localhost:8090", cfg.AuthURL)
		}
	}
}

func TestLoginFail(t *testing.T) {
	req := require.New(t)

	prompt := prompt("cody@confluent.io", "iamrobin")
	auth := &sdkMock.Auth{
		LoginFunc: func(ctx context.Context, idToken string, username string, password string) (string, error) {
			return "", &ccloud.InvalidLoginError{}
		},
	}
	user := &sdkMock.User{
		CheckEmailFunc: func(ctx context.Context, user *orgv1.User) (*orgv1.User, error) {
			return &orgv1.User{
				Email: "test-email",
			}, nil
		},
	}
	cmds, _ := newAuthCommand(prompt, auth, user, "ccloud", req)

	_, err := pcmd.ExecuteCommand(cmds.Commands[0])
	req.Contains(err.Error(), "You have entered an incorrect username or password.")
}

func TestURLRequiredWithMDS(t *testing.T) {
	req := require.New(t)

	prompt := prompt("cody@confluent.io", "iamrobin")
	auth := &sdkMock.Auth{
		LoginFunc: func(ctx context.Context, idToken string, username string, password string) (string, error) {
			return "", &ccloud.InvalidLoginError{}
		},
	}
	cmds, _ := newAuthCommand(prompt, auth, nil, "confluent", req)

	_, err := pcmd.ExecuteCommand(cmds.Commands[0])
	req.Contains(err.Error(), "required flag(s) \"url\" not set")
}

func TestLogout(t *testing.T) {
	req := require.New(t)

	prompt := prompt("cody@confluent.io", "iamrobin")
	auth := &sdkMock.Auth{}
	cmds, cfg := newAuthCommand(prompt, auth, nil, "ccloud", req)

	cfg.AuthToken = "some.token.here"
	cfg.Auth = &config.AuthConfig{User: &orgv1.User{Id: 23}}
	req.NoError(cfg.Save())

	output, err := pcmd.ExecuteCommand(cmds.Commands[1])
	req.NoError(err)
	req.Contains(output, "You are now logged out")

	cfg = config.New()
	cfg.CLIName = "ccloud"
	req.NoError(cfg.Load())
	req.Empty(cfg.AuthToken)
	req.Empty(cfg.Auth)
}

func Test_credentials_NoSpacesAroundEmail_ShouldSupportSpacesAtBeginOrEnd(t *testing.T) {
	req := require.New(t)

	prompt := prompt(" cody@confluent.io ", " iamrobin ")
	auth := &sdkMock.Auth{}
	cmds, _ := newAuthCommand(prompt, auth, nil, "ccloud", req)

	user, pass, err := cmds.credentials(cmds.Commands[0], "Email", nil)
	req.NoError(err)
	req.Equal("cody@confluent.io", user)
	req.Equal(" iamrobin ", pass)
}


func Test_SelfSignedCerts(t *testing.T) {
	req := require.New(t)
	mdsConfig := mds.NewConfiguration()
	mdsClient := mds.NewAPIClient(mdsConfig)
	cfg := config.New()
	cfg.Logger = log.New()
	cfg.CLIName = "confluent"
	prompt := prompt("cody@confluent.io", "iambatman")
	cmds := newCommands(&cliMock.Commander{}, cfg, log.New(), mdsClient, prompt, nil, nil, cliMock.NewDummyAnalyticsMock())

	for _, c := range cmds.Commands {
		c.PersistentFlags().CountP("verbose", "v", "Increase output verbosity")
	}

	// Create a test certificate to be read in by the command
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1234),
		Subject: pkix.Name{ Organization:  []string{"testorg"} },
	}
	priv, err := rsa.GenerateKey(rand.Reader, 512)
	req.NoError(err, "Couldn't generate private key")
	ca_b, err := x509.CreateCertificate(rand.Reader, ca, ca, &priv.PublicKey, priv)
	req.NoError(err, "Couldn't generate certificate from private key")
	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ca_b})
	cmds.certReader = bytes.NewReader(pemBytes)

	cert, err := x509.ParseCertificate(ca_b)
	req.NoError(err, "Couldn't reparse certificate")
	expectedSubject := cert.RawSubject
	mdsClient.TokensAuthenticationApi = &mdsMock.TokensAuthenticationApi{
		GetTokenFunc: func(ctx context.Context, xSPECIALRYANHEADER string) (mds.AuthenticationResponse, *http.Response, error) {
			req.NotEqual(http.DefaultClient, mdsClient)
			transport, ok := mdsClient.GetConfig().HTTPClient.Transport.(*http.Transport)
			req.True(ok)
			req.NotEqual(http.DefaultTransport, transport)
			found := false
			for _, actualSubject := range transport.TLSClientConfig.RootCAs.Subjects() {
				if bytes.Equal(expectedSubject, actualSubject) {
					found = true
					break
				}
			}
			req.True(found, "Certificate not found in client.")
			return mds.AuthenticationResponse{
				AuthToken: "y0ur.jwt.T0kEn",
				TokenType: "JWT",
				ExpiresIn: 100,
			}, nil, nil
		},
	}
	_, err = pcmd.ExecuteCommand(cmds.Commands[0], "--url=http://localhost:8090", "--ca-cert-path=testcert.pem")
	req.NoError(err)
}

func prompt(username, password string) *cliMock.Prompt {
	return &cliMock.Prompt{
		ReadStringFunc: func(delim byte) (string, error) {
			return "cody@confluent.io", nil
		},
		ReadPasswordFunc: func() ([]byte, error) {
			return []byte(" iamrobin "), nil
		},
	}
}

func newAuthCommand(prompt pcmd.Prompt, auth *sdkMock.Auth, user *sdkMock.User, cliName string, req *require.Assertions) (*commands, *config.Config) {
	var mockAnonHTTPClientFactory = func(baseURL string, logger *log.Logger) *ccloud.Client {
		req.Equal("https://confluent.cloud", baseURL)
		return &ccloud.Client{Auth: auth, User: user}
	}
	var mockJwtHTTPClientFactory = func(ctx context.Context, jwt, baseURL string, logger *log.Logger) *ccloud.Client {
		return &ccloud.Client{Auth: auth, User: user}
	}
	cfg := config.New()
	cfg.Logger = log.New()
	cfg.CLIName = cliName
	var mdsClient *mds.APIClient
	if cliName == "confluent" {
		mdsConfig := mds.NewConfiguration()
		mdsClient = mds.NewAPIClient(mdsConfig)
		mdsClient.TokensAuthenticationApi = &mdsMock.TokensAuthenticationApi{
			GetTokenFunc: func(ctx context.Context, xSPECIALRYANHEADER string) (mds.AuthenticationResponse, *http.Response, error) {
				return mds.AuthenticationResponse{
					AuthToken: "y0ur.jwt.T0kEn",
					TokenType: "JWT",
					ExpiresIn: 100,
				}, nil, nil
			},
		}
	}
	commands := newCommands(&cliMock.Commander{}, cfg, log.New(), mdsClient, prompt, mockAnonHTTPClientFactory, mockJwtHTTPClientFactory, cliMock.NewDummyAnalyticsMock())
	for _, c := range commands.Commands {
		c.PersistentFlags().CountP("verbose", "v", "Increase output verbosity")
	}
	return commands, cfg
}
