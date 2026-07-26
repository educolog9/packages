package functions

import (
	"testing"
	"time"

	"github.com/educolog9/packages/enums"
	"github.com/golang-jwt/jwt/v5"
)

const testSecret = "test-secret"

// signLikeUserService mints a token the same way the user service's SignToken
// does: MapClaims, HS256, and `exp` as a bare Unix timestamp. The point is to
// prove that tokens issued by the v3 code path still validate here after the
// move to v5 RegisteredClaims — the two serialise identically on the wire.
func signLikeUserService(t *testing.T, method jwt.SigningMethod, secret string, exp time.Time) string {
	t.Helper()
	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"id":             "user-123",
		"name":           "Ada",
		"lastName":       "Lovelace",
		"profilePicture": "https://example.test/ada.png",
		"roles":          []string{string(enums.Admin)},
		"isBlocked":      false,
		"isConfirmed":    true,
		"organization":   "org-1",
		"group":          "group-1",
		"email":          "ada@example.test",
		"version":        1,
		"exp":            exp.Unix(),
	})
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("signing test token: %v", err)
	}
	return signed
}

func TestValidateTokenAcceptsTokenFromUserService(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)

	claims, err := ValidateToken(signLikeUserService(t, jwt.SigningMethodHS256, testSecret, time.Now().Add(time.Hour)))
	if err != nil {
		t.Fatalf("expected the token to validate, got %v", err)
	}

	if claims.ID != "user-123" {
		t.Errorf("ID = %q, want %q", claims.ID, "user-123")
	}
	if claims.Email != "ada@example.test" {
		t.Errorf("Email = %q, want %q", claims.Email, "ada@example.test")
	}
	if claims.OrganizationID != "org-1" {
		t.Errorf("OrganizationID = %q, want %q", claims.OrganizationID, "org-1")
	}
	if !claims.IsConfirmed || claims.IsBlocked {
		t.Errorf("IsConfirmed = %v, IsBlocked = %v; want true, false", claims.IsConfirmed, claims.IsBlocked)
	}
	if !claims.IsAdmin() {
		t.Error("IsAdmin() = false, want true")
	}
	// exp arrives as a bare Unix number and has to land in RegisteredClaims.
	if claims.ExpiresAt == nil {
		t.Fatal("ExpiresAt is nil; the exp claim did not decode")
	}
	if !claims.ExpiresAt.After(time.Now()) {
		t.Errorf("ExpiresAt = %v, want a future time", claims.ExpiresAt)
	}
}

func TestValidateTokenRejectsExpiredToken(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)

	if _, err := ValidateToken(signLikeUserService(t, jwt.SigningMethodHS256, testSecret, time.Now().Add(-time.Hour))); err == nil {
		t.Fatal("expired token was accepted")
	}
}

// This is the case WithValidMethods actually changes. The keyfunc hands back an
// HMAC secret regardless of what the token asks for, so without the option any
// HMAC variant is honoured and the token below validates. It is not by itself
// an exploit — the attacker still needs the secret — but it is the mechanism
// behind algorithm-confusion attacks, and it becomes one the moment the keyfunc
// returns an asymmetric public key. Removing the option makes this test fail.
func TestValidateTokenRejectsUnexpectedSigningMethod(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)

	if _, err := ValidateToken(signLikeUserService(t, jwt.SigningMethodHS512, testSecret, time.Now().Add(time.Hour))); err == nil {
		t.Fatal("token signed with HS512 was accepted; only HS256 is issued")
	}
}

// v5 rejects alg=none on its own — the keyfunc returns []byte and
// SigningMethodNone only accepts the UnsafeAllowNoneSignatureType sentinel, so
// this passes with or without WithValidMethods. Kept as a regression guard.
func TestValidateTokenRejectsNoneAlgorithm(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)

	unsigned := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"id":    "attacker",
		"roles": []string{string(enums.Admin)},
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	token, err := unsigned.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		t.Fatalf("building unsigned token: %v", err)
	}

	if _, err := ValidateToken(token); err == nil {
		t.Fatal("token with alg=none was accepted")
	}
}

func TestValidateTokenRejectsWrongSecret(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)

	if _, err := ValidateToken(signLikeUserService(t, jwt.SigningMethodHS256, "not-the-secret", time.Now().Add(time.Hour))); err == nil {
		t.Fatal("token signed with the wrong secret was accepted")
	}
}

func TestValidateTokenRejectsGarbage(t *testing.T) {
	t.Setenv("JWT_SECRET", testSecret)

	for _, tokenString := range []string{"", "not.a.token", "a.b.c"} {
		if _, err := ValidateToken(tokenString); err == nil {
			t.Errorf("token string %q was accepted", tokenString)
		}
	}
}
