package main

import (
  "crypto/rand"
  "encoding/base64"
)

func GenerateRandomString(s int) (string, error) {
  b, err := generateRandomBytes(s)
  return base64.URLEncoding.EncodeToString(b), err
}

func generateRandomBytes(n int) ([]byte, error) {
  b := make([]byte, n)
  _, err := rand.Read(b)
  // Note that err == nil only if we read len(b) bytes.
  if err != nil {
    return nil, err
  }

  return b, nil
}
