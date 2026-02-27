# BLS

# BLS Signatures and Ethereum 2.0

Implementation of the BLS signature scheme using the [CIRCL library](https://github.com/cloudflare/circl) by Cloudflare, written in Go. This project covers key generation, signing, verification, and signature aggregation — as used in the Ethereum 2.0 Proof-of-Stake consensus mechanism.

---

## Project Structure

```
.
├── go.mod
├── go.sum
├── README.md
├── keygen.go        # Key generation
├── sign.go          # Signing
├── verify.go        # Verification
├── aggregate.go     # Signature aggregation
├── bls_test.go      # Test cases
└── report/          # Written report (PDF)
```

---

## Compilation and Installation

### Prerequisites

- [Go](https://go.dev/dl/) 1.21 or later
- Internet connection (to fetch the CIRCL dependency on first run)

Verify your Go installation:
```bash
go version
```

### Clone the Repository

```bash
git clone https://github.com/armancher/BLS
cd BLS
```

### Install Dependencies

CIRCL is managed as a Go module. To download it:
```bash
go mod download
```

This fetches `github.com/cloudflare/circl` automatically based on `go.mod` and `go.sum` — no manual cloning of CIRCL is needed.

### Build

To compile and verify everything builds correctly:
```bash
go build ./...
```

---

## Running the Test Cases

Run all tests with:
```bash
go test ./... -v
```

The `-v` flag prints each test name and its result. You should see output like:

```
--- PASS: TestKeyGen (0.00s)
--- PASS: TestSignAndVerify (0.00s)
--- PASS: TestAggregation (0.00s)
PASS
```

A `PASS` result for all tests confirms the implementation is correct. Any `FAIL` indicates an issue with that specific function.

To run a single test:
```bash
go test -v -run TestAggregation
```

---

## Running Benchmarks

To benchmark the performance of key generation, signing, verification, and aggregation:
```bash
go test -bench=. -benchmem
```

Output will show each benchmark's name, number of iterations, time per operation, and memory usage. Example:
```
BenchmarkSign-8       5000    300000 ns/op    1024 B/op    12 allocs/op
```

Higher iterations and lower `ns/op` indicate better performance.

---

## Dependencies

| Library | Version | Purpose |
|---|---|---|
| [cloudflare/circl](https://github.com/cloudflare/circl) | v1.6.3 | BLS signature primitives and pairing-based elliptic curves |

---

## Background

BLS (Boneh–Lynn–Shacham) signatures rely on bilinear pairings on elliptic curves. Their key property is that multiple signatures over the same or different messages can be **aggregated** into a single compact signature, verified in one pairing check. Ethereum 2.0 uses BLS12-381 curve-based BLS signatures for validator attestations, dramatically reducing the on-chain data required to record thousands of validator votes per slot.

---

## AI Usage

Parts of this project (scaffolding, README) were developed with assistance from Claude (Anthropic). All cryptographic logic and analysis was reviewed and understood by the authors. AI usage is declared in accordance with DTU guidelines.



##  group 7
Arman Cheraghvnadi s252657
Simone Panella s253125
Riccardo Lussana s253032