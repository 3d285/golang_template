#!/usr/bin/env bash
# Demonstrates AES-128-CBC interop with Go.
# Uses hex key/iv for clarity (no salt).

set -euo pipefail

KEY_HEX=31323334353637383930313233343536     # "1234567890123456"
IV_HEX=31323334353637383930313233343536      # "1234567890123456"

PLAINTEXT='{"msg":"支付成功","amt":99.99,"code":200}'

echo "[OpenSSL] Encrypting..."
CIPHERTEXT_BASE64=$(printf "%s" "$PLAINTEXT" | openssl enc -aes-128-cbc -K "$KEY_HEX" -iv "$IV_HEX" -nosalt -base64)
echo "CIPHERTEXT_BASE64=$CIPHERTEXT_BASE64"

echo "[OpenSSL] Decrypting (roundtrip check)..."
printf "%s" "$CIPHERTEXT_BASE64" | openssl base64 -d |   openssl enc -aes-128-cbc -d -K "$KEY_HEX" -iv "$IV_HEX" -nosalt

echo
echo "Now verify the same in Go tests: go test ./pkg/crypto/aesx -run CBC"
