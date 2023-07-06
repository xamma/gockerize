package dockerizer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenerateDockerfile(dockerfilePath, path, baseImage, runImage, buildOutput, exposePort, cmd string) error {
	if dockerfilePath == "" {
		// Use the current working directory as the default Dockerfile path
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current working directory: %v", err)
		}
		dockerfilePath = filepath.Join(cwd, "Dockerfile")
	}

	// Ensure the specified path exists and is a directory
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	// Create the Dockerfile content
	dockerfileContent := fmt.Sprintf("FROM %s AS builder\n", baseImage)
	dockerfileContent += fmt.Sprintf("WORKDIR /app\n")
	dockerfileContent += fmt.Sprintf("COPY %s .\n", path)
	dockerfileContent += fmt.Sprintf("RUN go build -o %s\n", buildOutput)
	dockerfileContent += fmt.Sprintf("\n")
	dockerfileContent += fmt.Sprintf("FROM %s\n", runImage)
	dockerfileContent += fmt.Sprintf("WORKDIR /app\n")
	dockerfileContent += fmt.Sprintf("COPY --from=builder /app/%s .\n", buildOutput)
	dockerfileContent += fmt.Sprintf("\n")
	dockerfileContent += fmt.Sprintf("EXPOSE %s\n", exposePort)
	dockerfileContent += fmt.Sprintf("\n")
	dockerfileContent += fmt.Sprintf("CMD [\"%s\"]\n", cmd)

	// Write the Dockerfile content to the specified path
	err = ioutil.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write Dockerfile: %v", err)
	}

	return nil
}
