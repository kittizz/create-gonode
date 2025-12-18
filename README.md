# create-gonode

Build Node.js modules with Go. Write your logic in Go, compile to JavaScript, and publish as npm package with full TypeScript support.

## Features

- 🚀 **Go to JavaScript** - Compile Go code to JavaScript using GopherJS
- 📦 **npm Ready** - Publish as standard npm package
- 🔷 **TypeScript Support** - Auto-generated type definitions
- 🌐 **Multi-platform** - Build for Node.js and Browser
- ⚡ **Modern Tooling** - Powered by tsdown (Rolldown-based bundler)
- 🔧 **npm Scripts** - Simple build automation with npm/pnpm

## How It Works

```
┌─────────────┐     ┌───────────┐     ┌─────────────┐     ┌──────────────┐
│   Go Code   │ ──▶ │ GopherJS  │ ──▶ │ JavaScript  │ ──▶ │  npm Package │
│  (*.go)     │     │ Compiler  │     │   (*.js)    │     │  (.js + .d.ts)│
└─────────────┘     └───────────┘     └─────────────┘     └──────────────┘
```

## Quick Start

### Prerequisites
- Go 1.19+ (for GopherJS compatibility)
- Node.js 18+
- pnpm (recommended)

```bash
pnpm run setup
```

### Installation

```bash
# Clone this template
git clone https://github.com/kittizz/create-gonode.git my-go-module
cd my-go-module

# Setup toolchain
pnpm run setup
```

### Project Structure

```
create-gonode/
├── pkg
│   ├── (go packages)
├── src/
│   ├── go/           # Go source code
│   └── node/         # Node exports (GopherJS)
│       ├── exports.go    # Node.js exports (GopherJS)
│       ├── lib.d.ts      # TypeScript type definitions
│       ├── lib.js        # Node.js build output
│       └── index.ts      # TypeScript entry point
├── dist/
│   ├── node/             # Node.js build output
├── go.mod
├── package.json
├── build.js
└── tsdown.config.ts
```

## Usage

### 1. Write Your Go Code

```go
// mylib.go
package mylib

func Add(a, b int) int {
    return a + b
}

func Greet(name string) string {
    return "Hello, " + name + "!"
}
```

### 2. Export to JavaScript

```go
// src/node/exports.go
//go:build js

package main

import (
    "github.com/user/mylib"
    "github.com/gopherjs/gopherjs/js"
)

func main() {
    js.Module.Get("exports").Set("add", mylib.Add)
    js.Module.Get("exports").Set("greet", mylib.Greet)
}
```

### 3. Create TypeScript Entry

```typescript
// src/node/index.ts
import * as lib from "./lib.js";

export const add: (a: number, b: number) => number = lib.add;
export const greet: (name: string) => string = lib.greet;
```

### 4. Build

```bash
pnpm run build
```

### 5. Use in Node.js

```typescript
import { add, greet } from "my-go-module";

console.log(add(2, 3));        // 5
console.log(greet("World"));   // Hello, World!
```

## Available Scripts

```bash
pnpm run setup       # Install Go toolchain and dependencies
pnpm run build       # Build all (node + ts)
pnpm run build:node  # Build Node.js library with GopherJS
pnpm run build:ts    # Build TypeScript library
pnpm run clean       # Clean dist folder
pnpm run test        # Run tests
```

## Configuration

### package.json exports

```json
{
  "exports": {
    "./node": {
      "types": "./dist/node/index.d.ts",
      "import": "./dist/node/index.mjs"
    },
    "./browser": {
      "types": "./dist/browser/index.d.ts",
      "import": "./dist/browser/index.js"
    }
  }
}
```

### tsdown.config.ts

```typescript
import { defineConfig } from "tsdown";

export default defineConfig([
  {
    entry: "src/node/index.ts",
    format: "esm",
    outDir: "dist/node",
    platform: "node",
    dts: true,
  },
  {
    entry: "src/browser/index.ts",
    format: "esm",
    outDir: "dist/browser",
    platform: "browser",
    dts: true,
  },
]);
```

## Why Go + Node.js?

| Use Case | Benefit |
|----------|---------|
| **CPU-intensive tasks** | Go's performance in JavaScript |
| **Existing Go libraries** | Reuse Go code in Node.js |
| **Cryptography** | Share encryption logic across platforms |
| **Complex algorithms** | Write once, run everywhere |
| **Code obfuscation** | Compiled JS is harder to reverse-engineer |

## Limitations

- GopherJS supports Go 1.19 (not the latest Go version)
- Some Go packages may not be compatible with GopherJS
- Generated JavaScript is larger than hand-written code
- No goroutine support (JavaScript is single-threaded)

## Tech Stack

- [GopherJS](https://github.com/gopherjs/gopherjs) - Go to JavaScript compiler
- [tsdown](https://github.com/nicolo-ribaudo/tsdown) - TypeScript bundler (Rolldown-based)

## License

MIT

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
