go_library(
    name = "parse",
    srcs = [
        "init.go",
        "parse_step.go",
        "rules.go",
        "suggest.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//src/cli",
        '//src/core',
        "//src/fs",
        '//src/parse/asp',
        '//src/parse/asp/builtins',
        "//src/parse/rules",
        '//src/parse/cffi:please_parser_static',
        '//src/parse/skylark',
        '//src/utils',
        "//src/worker",
        '//third_party/go:gcfg',
        '//third_party/go:logging',
    ],
    visibility = ['PUBLIC'],
)

filegroup(
    name = 'builtin_rules',
    srcs = glob(['rules/*.build_defs']),
    visibility = [
        '//src/parse/asp/builtins:all',
        '//src/parse/skylark:all',
    ],
)

go_test(
    name = "parse_step_test",
    srcs = ["parse_step_test.go"],
    deps = [
        ":parse",
        "//src/core",
        "//third_party/go:testify",
    ],
)

go_test(
    name = "suggest_test",
    srcs = ["suggest_test.go"],
    deps = [
        ":parse",
        "//third_party/go:testify",
    ],
)

go_test(
    name = "rules_test",
    srcs = ["rules_test.go"],
    deps = [
        ":parse",
        "//third_party/go:testify",
    ],
)
