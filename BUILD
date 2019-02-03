load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
go_binary(
	name = "raytrace",
	srcs = ["main.go", "shapes.go", "vectors.go"]
)
go_test(
	name = "shape_test",
	srcs = ["shapes.go", "shapes_test.go", "vectors.go"],
)
