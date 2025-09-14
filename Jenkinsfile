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
                checkout scm
            }
        }

        stage('Build Networking App') {
        steps {
            echo 'Building Go networking application...'
            // Specify the exact directory of the main package.
            sh 'go build -o app-networking ./networking-exercise'
        }
    }

        stage('Test Networking App') {
            steps {
                echo 'Running integration tests for networking exercise...'
                // Go's test runner will automatically find and run the tests in this directory.
                sh 'go test -v ./networking-exercise'
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