package repository

type BuildSystem string

const (
	UnknownBuild BuildSystem = "unknown"

	GoModules BuildSystem = "go-mod"

	NPM BuildSystem = "npm"

	Yarn BuildSystem = "yarn"

	PNPM BuildSystem = "pnpm"

	Maven BuildSystem = "maven"

	Gradle BuildSystem = "gradle"

	CMake BuildSystem = "cmake"

	Cargo BuildSystem = "cargo"
)
