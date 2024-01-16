#!/usr/bin/env nu

const output_path = "bin"

def build [$architectures: list, $operating_systems: list] {
    $architectures | each { |arch|
        $operating_systems | each { |os|
            $env.GOOS = $os
            $env.GOARCH = $arch

            mut binary_path = $"($output_path)/sbot-($os)-($arch)"

            if $os == "windows" {
                $binary_path = $"($binary_path).exe"
            }

            go build -o $binary_path

            ^echo $"built ($binary_path)"
        }
    }
}

def build-all [] {
    let architectures = ["amd64", "arm64"]
    let operating_systems = ["windows", "linux", "darwin"]

    build $architectures $operating_systems
}

def build-local [] {
    echo "building local binary..."
    go build -o $"($output_path)/sbot"
}

def "main build-all" [] {
    echo "building all binaries..."
    build-all
}

def "main build" [] {
    build-local
}

def main [] {}