def mainDir="."
def dockerImgName="msago"
def docker="/usr/local/bin/docker"
def dockerRegistory="yuhojin"
def ecrLoginHelper="docker-credential-ecr-login"
def region="ap-northeast-2"
def ecrUrl="297064282309.dkr.ecr.ap-northeast-2.amazonaws.com"
def repository="go_msa"
def deployHost=""


// 개인적으로 cli 명렁어들에 대해서 경로를 직접 지정하는 것이 굉장히 마음에 들지 않지만...
// 해당 Jenkins는 Local에서 돌아가고 있고, 저의 Local인 맥북은 제가 막 테스트 하느라
// 가끔씩 Path가 엉망으로 설정이 되어 있는 경우가 있습니다.. 어쩔경우에는 그냥 없는 경우도;;
// 그래서 일단은 직접 경로를 설정해 주는 것으로 사용 중입니다.
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
                sh "${docker} build -t ${dockerRegistory}/${dockerImgName}:${currentBuild.number} ."
                sh "${docker} tag ${dockerRegistory}/${dockerImgName}:${currentBuild.number} ${ecrUrl}/${dockerImgName}:latest"
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
                    // 어차피 aws configure를 해두었다.
                    // cleanup current user docker credentials
                    sh 'rm -f ~/.dockercfg ~/.docker/config.json || true'
                    sh "/usr/local/bin/aws ecr get-login-password --region ${region} | ${docker} login --username AWS --password-stdin ${ecrUrl}"
                    // 기존에 이미 agent에서 login해 두었지만 혹시 모르니 테스트 용도로 재 로그인 시도

                    sh "${docker} push ${ecrUrl}/${dockerImgName}:latest"
                }
            }

            post {
                success {
                    echo 'succes push docker image to ECR'
                }
                failure {
                    error 'fail push docker image to ECR' // exit pipeline
                }
            }
        }
    }
}