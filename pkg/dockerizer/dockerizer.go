package dockerizer

import (
	"fmt"
	"os"
	"path/filepath"
)

func findRequirementsFile(directory string) (string, error) {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Base(path) == "requirements.txt" {
			return fmt.Errorf(path) // Stop walking and return the path
		}
		return nil
	})
	if err == nil {
		return "", fmt.Errorf("requirements.txt not found in %s", directory)
	}
	return err.Error(), nil
}

func generateBaseImage(baseImage, projectType string) string {
	if projectType == "python" {
		return "python:3.11-slim"
	}
	if baseImage == "" {
		// we want this primarily for Go projects, so take this as default
		return "golang:1.20-alpine"
	}
	return baseImage
}

func GenerateDockerfile(dockerfilePath, path, baseImage, runImage, buildOutput, exposePort, cmd, projectType string) error {
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

	var dockerfileContent string

	baseImage = generateBaseImage(baseImage, projectType)

	switch projectType {
	case "go":
		dockerfileContent = fmt.Sprintf("FROM %s AS builder\n", baseImage)
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("WORKDIR /app\n")
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("COPY %s .\n", path)
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("RUN go build -o %s\n", buildOutput)
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("FROM %s\n", runImage)
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("WORKDIR /app\n")
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("COPY --from=builder /app/%s .\n", buildOutput)
	case "python":
		dockerfileContent = fmt.Sprintf("FROM %s\n", baseImage)
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("RUN mkdir /app\n")
		dockerfileContent += fmt.Sprintf("\n")

		// Look for requirements.txt in the project directory and its subdirectories recursively
		requirementsPath, err := findRequirementsFile(path)
		if err != nil {
			return err
		}

		// Include requirements.txt and install dependencies if file was found
		if requirementsPath != "" {
			dockerfileContent += fmt.Sprintf("COPY %s /opt/requirements.txt\n", requirementsPath)
			dockerfileContent += fmt.Sprintf("\n")
			dockerfileContent += fmt.Sprintf("RUN pip install --no-cache-dir -r /opt/requirements.txt\n")
		}

		// Copy Python project files (hardcoded lol)
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("COPY /src/main /app\n")
		dockerfileContent += fmt.Sprintf("\n")
		dockerfileContent += fmt.Sprintf("WORKDIR /app\n")
	}

	// Common parts for both Go and Python projects (kinda, CMD not really working)
	dockerfileContent += fmt.Sprintf("\n")
	dockerfileContent += fmt.Sprintf("EXPOSE %s\n", exposePort)
	dockerfileContent += fmt.Sprintf("\n")
	dockerfileContent += fmt.Sprintf("CMD [\"%s\"]\n", cmd)

	// Create the Dockerfile and write the content to the file in the specified path
	file, err := os.Create(dockerfilePath)
	if err != nil {
		return fmt.Errorf("failed to create Dockerfile: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(dockerfileContent)
	if err != nil {
		return fmt.Errorf("failed to write Dockerfile content: %v", err)
	}

	return nil
}