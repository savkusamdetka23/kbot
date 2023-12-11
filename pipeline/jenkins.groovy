pipeline {
    agent any
        environment {
            REPO = 'https://github.com/savkusamdetka23/kbot.git'
            BRANCH = 'main'
        }
    parameters {

        choice(name: 'OS', choices: ['linux', 'darwin', 'macos', 'arm', 'windows', 'all'], description: 'Choose OS')
        choice(name: 'ARCH', choices: ['arm64', 'amd64', 'all'], description: 'Choose architecture')
        choice(name: 'ContainerRegistry', defaultValue: 'ghcr.io/savkusamdetka23', description: 'Container Registry')

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
        stage('test') {
            steps {
                echo 'TEST EXECUTION STARTED'
                sh 'make test'
            }
        }
        stage('build') {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'make build TARGETOS=${params.OS} TARGETARCH=${params.ARCH}'
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
                    docker.withRegistry( '', 'ghcr') {
                        sh 'make push'
                    }
                }
            }
        }
    }
}