def mainDir="."
def dockerImgName="msago"
def dockerRegistory="yuhojin"
def ecrLoginHelper="docker-credential-ecr-login"
def region="ap-northeast-2"
def ecrUrl="297064282309.dkr.ecr.ap-northeast-2.amazonaws.com"
def repository="go_msa"
def deployHost=""

pipeline {
    agent any

    stages {

        stage("Pull codes from Github"){
            steps {
                checkout scm
            }
        }

        stage("Test Cli") {
            steps {
               sh '''
                    export GOPATH=$WORKSPACE
                    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
                    go version
                '''
            }
        }

        stage("Build Docker Image"){
            steps {
                sh "/usr/local/bin/docker build -t ${dockerRegistory}/${dockerImgName}:${currentBuild.number} ."
                sh "/usr/local/bin/docker tag ${dockerRegistory}/${dockerImgName}:${currentBuild.number} ${dockerImgName}:latest"
            }

            post {
                success {
                    echo 'succes docker making docker image'
                }
                failure {
                    error 'fail dockerizing project' // exit pipeline
                }
            }
        }

        stage("Push To AWS ECR") {
            steps {
                script {
                    // cleanup current user docker credentials
                    sh 'rm -f ~/.dockercfg ~/.docker/config.json || true'

                    docker.withRegistry("https://${ecrUrl}", "ecr:${region}:aws-key") {
                        docker.image("${dockerRegistory}/${dockerImgName}:${currentBuild.number}").push()
                        docker.image("${dockerImgName}:latest").push()
                    }
                }
            }
        }
    }
}