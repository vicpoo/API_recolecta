package notificacion_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vicpoo/API_recolecta/config"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
	"github.com/vicpoo/API_recolecta/src/notificacion/infrastructure"
)

func TestFCMClient_SendNotification_Integration(t *testing.T) {
	// This is an integration test and requires ADC or a credentials file.
	// Set the FCM_CREDENTIALS_FILE or GOOGLE_APPLICATION_CREDENTIALS environment variable to run.
	if os.Getenv("FCM_CREDENTIALS_FILE") == "" && os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		t.Skip("Skipping FCM integration test: Set FCM_CREDENTIALS_FILE or GOOGLE_APPLICATION_CREDENTIALS to run.")
	}

	cfg, err := config.LoadConfig()
	require.NoError(t, err, "Failed to load config")

	fcmClient, err := infrastructure.NewFCMClient(cfg.FCMCredentialsFile)
	require.NoError(t, err, "Failed to create FCM client. Check your credentials.")
	require.NotNil(t, fcmClient, "FCM client should not be nil")

	// Use a fake token to verify connectivity with FCM.
	// Sending to an invalid token proves the request reached FCM servers for validation.
	fakeToken := "f1234567890_this_is_an_invalid_token_and_should_fail"
	userTokens := map[string]string{
		"test-user-id": fakeToken,
	}

	notification := &domain.PushNotification{
		Title: "Integration Test Title",
		Body:  "Integration Test Body",
		Type:  "test",
	}

	results, err := fcmClient.Send(context.Background(), userTokens, notification)

	// A top-level error indicates a problem with the request itself (e.g., auth failure, network issue).
	// Individual token errors are handled within the 'results' map.
	require.NoError(t, err, "fcmClient.Send should not return a top-level error")

	// We expect a result for our test user.
	require.Contains(t, results, "test-user-id", "Results map should contain an entry for the test user")

	result := results["test-user-id"]

	// The send should fail because the token is invalid. This is the expected outcome.
	assert.False(t, result.Success, "Sending to a fake token should not be successful")
	assert.NotEmpty(t, result.Error, "There should be an error message for the failed send")

	// Check for common error messages associated with invalid tokens.
	// This confirms that the error came from FCM and not our own system.
	// Error messages from FCM can vary, so we check for substrings.
	errorLower := strings.ToLower(result.Error)
	isInvalidArgument := strings.Contains(errorLower, "invalid-argument")
	isNotRegistered := strings.Contains(errorLower, "registration-token-not-registered")
	isNotFound := strings.Contains(errorLower, "was not found")
	isNotValidFCMToken := strings.Contains(errorLower, "not a valid fcm registration token")

	assert.True(t, isInvalidArgument || isNotRegistered || isNotFound || isNotValidFCMToken, "Error message should indicate an invalid or unregistered token, but got: %s", result.Error)
}
