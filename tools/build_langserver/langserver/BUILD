go_library(
    name = "langserver",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = ["//tools/build_langserver/..."],
    deps = [
        "//src/core",
        "//src/fs",
        "//src/parse",
        "//src/parse/asp",
        "//src/parse/rules",
        "//src/query",
        "//third_party/go:jsonrpc2",
        "//third_party/go:logging",
        "//tools/build_langserver/lsp",
    ],
)

go_test(
    name = "langserver_test",
    srcs = glob(
        [
            "*_test.go",
            "setup_test.go",
        ],
        exclude = ["utils_test.go"],
    ),
    data = ["test_data"],
    deps = [
        ":langserver",
        "//src/parse/rules",
        "//third_party/go:testify",
    ],
)

go_test(
    name = "utils_test",
    srcs = ["utils_test.go"],
    deps = [
        ":langserver",
        "//third_party/go:testify",
    ],
)
