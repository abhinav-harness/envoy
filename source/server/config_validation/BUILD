load("//bazel:envoy_build_system.bzl", "envoy_cc_library", "envoy_package")

licenses(["notice"])  # Apache 2

envoy_package()

envoy_cc_library(
    name = "admin_lib",
    srcs = ["admin.cc"],
    hdrs = ["admin.h"],
    deps = [
        "//envoy/network:listen_socket_interface",
        "//envoy/server:admin_interface",
        "//source/common/common:assert_lib",
        "//source/common/network:listen_socket_lib",
        "//source/server/admin:config_tracker_lib",
    ],
)

envoy_cc_library(
    name = "api_lib",
    srcs = ["api.cc"],
    hdrs = ["api.h"],
    deps = [
        ":dispatcher_lib",
        "//envoy/api:api_interface",
        "//envoy/filesystem:filesystem_interface",
        "//source/common/api:api_lib",
        "@envoy_api//envoy/config/bootstrap/v3:pkg_cc_proto",
    ],
)

envoy_cc_library(
    name = "cluster_manager_lib",
    srcs = ["cluster_manager.cc"],
    hdrs = ["cluster_manager.h"],
    deps = [
        "//envoy/upstream:cluster_manager_interface",
        "//source/common/common:utility_lib",
        "//source/common/http:context_lib",
        "//source/common/upstream:cluster_manager_lib",
        "@envoy_api//envoy/config/bootstrap/v3:pkg_cc_proto",
        "@envoy_api//envoy/config/core/v3:pkg_cc_proto",
    ],
)

envoy_cc_library(
    name = "dispatcher_lib",
    srcs = ["dispatcher.cc"],
    hdrs = [
        "connection.h",
        "dispatcher.h",
        "dns.h",
    ],
    deps = [
        "//envoy/event:dispatcher_interface",
        "//envoy/network:connection_interface",
        "//source/common/common:assert_lib",
        "//source/common/event:dispatcher_lib",
    ],
)

envoy_cc_library(
    name = "dns_lib",
    srcs = ["dns.cc"],
    hdrs = ["dns.h"],
    deps = [
        "//envoy/event:dispatcher_interface",
        "//envoy/network:dns_interface",
    ],
)

envoy_cc_library(
    name = "server_lib",
    srcs = ["server.cc"],
    hdrs = ["server.h"],
    external_deps = ["abseil_optional"],
    deps = [
        ":admin_lib",
        ":api_lib",
        ":cluster_manager_lib",
        ":dns_lib",
        "//envoy/server:drain_manager_interface",
        "//envoy/server:instance_interface",
        "//envoy/ssl:context_manager_interface",
        "//envoy/tracing:tracer_interface",
        "//source/common/access_log:access_log_manager_lib",
        "//source/common/common:assert_lib",
        "//source/common/common:random_generator_lib",
        "//source/common/common:utility_lib",
        "//source/common/config:utility_lib",
        "//source/common/grpc:common_lib",
        "//source/common/local_info:local_info_lib",
        "//source/common/protobuf:utility_lib",
        "//source/common/quic:quic_stat_names_lib",
        "//source/common/router:context_lib",
        "//source/common/router:rds_lib",
        "//source/common/runtime:runtime_lib",
        "//source/common/stats:stats_lib",
        "//source/common/thread_local:thread_local_lib",
        "//source/common/version:version_lib",
        "//source/server:configuration_lib",
        "//source/server:server_lib",
        "//source/server/admin:admin_lib",
        "@envoy_api//envoy/config/bootstrap/v3:pkg_cc_proto",
        "@envoy_api//envoy/config/core/v3:pkg_cc_proto",
        "@envoy_api//envoy/config/listener/v3:pkg_cc_proto",
    ],
)
