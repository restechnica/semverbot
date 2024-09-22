#!/usr/bin/env nu

const output_path = "bin"

def build-all [build_options: list<string>] {
    print "building all binaries..."

    let architectures = ["amd64", "arm64"]
    let operating_systems = ["windows", "linux", "darwin"]

    $architectures | each { |arch|
        $operating_systems | each { |os|
            $env.GOOS = $os
            $env.GOARCH = $arch

            mut binary_path = $"($output_path)/sbot-($os)-($arch)"

            if $os == "windows" {
                $binary_path = $"($binary_path).exe"
            }

            let options =  ["build", "-o",  $binary_path] ++ $build_options

            run-external go ...$options
            print $"built ($binary_path)"
        }
    }
}

def build-local [build_options: list<string>] {
    print "building local binary..."
    let binary_path = $"($output_path)/sbot"
    let options = ["build", "-o", $binary_path] ++ $build_options
    run-external go ...$options
    print $"built ($binary_path)"
}

def get-ldflags [version: string] {
    const go_import_path = "github.com/restechnica/semverbot"

    let ldflags = [
        $"-X ($go_import_path)/internal/ldflags.Version=($version)"
    ] | str join ' '

    return $ldflags
}

def "main build" [--all, --version: string = "dev"] {
    let ldflags = get-ldflags $version

    let options: list<string> = ["-ldflags", $ldflags]

    if $all {
        build-all $options
    } else {
        build-local $options
    }
}

def main [] {}
