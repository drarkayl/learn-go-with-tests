// Jenkinsfile for a basic Go project CI pipeline
pipeline {
    agent any // This means the pipeline can run on any available Jenkins agent.

    // This section defines the tools needed for the pipeline.
    // Jenkins will automatically install the specified Go version.
    tools {
        go 'go1.24.0' // Use the Go version you want. Check Jenkins' "Global Tool Configuration" to see what's available.
    }

    // Environment variables for our Go project.
    // GO111MODULE is set to on to enable Go Modules.
    environment {
        GO111MODULE = 'on'
    }

    // The stages of our CI pipeline.
    stages {
        stage('Checkout') {
            steps {
                echo 'Checking out source code...'
                // The `git` step checks out the code from the repository.
                // We'll use the scm shortcut for Multibranch Pipelines.
                // For a simple pipeline, you can specify the git URL here.
                git url: 'https://github.com/drarkayl/learn-go-with-tests.git'
            }
        }

        stage('Build') {
            steps {
                echo 'Building Go application...'
                // Build the main package.
                // `sh` step executes a shell command.
                sh 'go build -o app ./...'
            }
        }

        stage('Test') {
            steps {
                echo 'Running unit tests...'
                // Run all tests in the project.
                // The `-v` flag provides verbose output.
                sh 'go test -v ./...'
            }
        }
    }

    // The post section defines actions to take after the pipeline has finished.
    // It's useful for cleanup, sending notifications, etc.
    post {
        always {
            echo 'Pipeline finished.'
        }
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}