pipeline {
    agent none

    stages {
        stage('Build app') {
            agent {
                kubernetes {
                    yamlFile "ci/go-container.yaml"
                }
            }
            steps {
                container('kubectl-go') {
                    sh 'go get -d'
                    echo 'make test...'
                    sh 'go build -o pirates *.go'
                    stash name: "app", includes: "pirates"
                } 
            }
        }
        stage('Build image') {
            agent {
                kubernetes {
                    yamlFile "ci/kaniko.yaml"
                }
            }
            steps {
                container('kaniko') {
                    unstash 'app'
                    sh '/kaniko/executor --force -f `pwd`/Dockerfile.embed -c `pwd` --insecure --skip-tls-verify --cache=true --destination mich4n/pirates:latest'
                }
            }
        }
        stage('Deploy dev') {
            agent {
                kubernetes {
                    serviceAccount 'deployer'
                    containerTemplate {
                        name 'kubectl'
                        image 'mich4n/ci'
                        ttyEnabled true
                        command 'cat'
                    }
                }
            }
            steps {
                sh 'cd pirates-kubernetes/overlays/dev/ && kustomize build | kubectl apply -f-'
            }
        }
    }
}
