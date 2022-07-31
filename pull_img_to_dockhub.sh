
###Build docker
#docker build -t httpserver:v1.0 -f dockerfile  .


###start project
#docker run -d httpserver:v1.0 -p 8080:8080


###Pull image to dockerhub
#sh pull_img_to_dockhub.sh httpserver v1.0
REPOSITORY = "registry.cn-beijing.aliyuncs.com"
PROJECT = "$1"
IMAGE_NAME = "$1:$2"

sudo docker login --username=fernandoander -p [password] ${REPOSITORY}
sudo docker tag ${IMAGE_NAME} ${REPOSITORY}/${PROJECT}/${IMAGE_NAME}
sudo docker push  ${REPOSITORY}/${PROJECT}/${IMAGE_NAME}
