# install dependencies
sudo apt-get -y update
sudo apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# install docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get -y update
sudo apt-get -y install docker-ce docker-ce-cli containerd.io


# install more dependencies
sudo apt-get -y update && \
  apt-get install --no-install-recommends bash curl wget \
  && rm -rf /var/lib/apt/lists/*

# manually installing yq
sudo wget -O /usr/local/bin/yq https://github.com/mikefarah/yq/releases/download/3.3.0/yq_linux_amd64
sudo chmod +x /usr/local/bin/yq
