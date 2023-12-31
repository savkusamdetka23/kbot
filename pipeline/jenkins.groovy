pipeline {
    agent any
        environment {
            REPO = 'https://github.com/savkusamdetka23/kbot.git'
            BRANCH = 'main'
            GHCR_REPOSITORY = 'savkusamdetka23/kbot'

        }
    parameters {

        choice(name: 'OS', choices: ['linux', 'darwin', 'macos', 'arm', 'windows', 'all'], description: 'Choose OS')
        choice(name: 'ARCH', choices: ['arm64', 'amd64', 'all'], description: 'Choose architecture')
        string(name: 'ContainerRegistry', defaultValue: 'ghcr.io/savkusamdetka23', description: 'Container Registry')

    }
    stages {
        stage('SelectedParams') {
            steps {
                echo "Build for platform ${params.OS}"

                echo "Build for arch: ${params.ARCH}"

                echo "Build for ContainerRegistry: ${params.ContainerRegistry}"

            }
        }
        stage('clone') {
            steps {
                echo 'CLONE REPOSITORY'
                    git branch: "${BRANCH}", url: "${REPO}"
            }
        }
        stage('build') {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'make build TARGETOS=${OS} TARGETARCH=${ARCH}'
            }
        }
        stage('image') {
            steps {
                echo 'IMAGE CREATION STARTED'
                sh 'make image'
            }
        }
        stage('push') {
            steps {
                script {
                    
                    // Authenticate with GitHub Container Registry
                    withCredentials([string(credentialsId: 'ghcr-token', variable: 'GHCR_TOKEN')]) {
                        sh "echo $GHCR_TOKEN | docker login ghcr.io -u ${params.ContainerRegistry} --password-stdin"
                    }

                    // Tag and push the Docker image to GHCR
                    sh "make push"
                }
            }
        }
    }
}