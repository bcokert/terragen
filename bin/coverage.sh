#!/usr/bin/env bash

COVERAGE_DIR=.coverage-reports
profile="$COVERAGE_DIR/cover.out"
mode=count

generate_cover_data() {
    rm -rf "$COVERAGE_DIR"
    mkdir "$COVERAGE_DIR"

    for pkg in "$@"; do
        f="$COVERAGE_DIR/$(echo $pkg | tr / -).cover"
        go test -covermode="$mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $mode" >"$profile"
    if [[ "$(ls ${COVERAGE_DIR} | grep -c .cover)" > 0 ]]; then
        grep -h -v "^mode:" "$COVERAGE_DIR"/*.cover >>"$profile"
    else
        echo "No tests to run!"
    fi
}

generate_cover_data $(go list ./...)
echo "Run make view-coverage to see the coverage report"
