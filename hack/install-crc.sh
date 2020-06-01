#!/usr/bin/env bash

ROOTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
BINDIR="$ROOTDIR/bin"
CRCDIR="$ROOTDIR/crc"
​
if [ "`which expect`" = "" ];then
   echo "ERROR: EXPECT is not installed, execute 'yum install -y expect' and then re-run this script"
   exit 1
fi
​
export PATH="$BINDIR:$PATH"
​
if [ ! -f $ROOTDIR/install.properties ];then
    echo "install.properties file is missing, exiting..."
    exit 1
fi
​
# stop and delete any previous CRC environment that is running
pids=`ps -ef | grep crc | grep -v grep | grep -v install-crc | awk '{print $2}' || true`
if [ "$pids" != "" ];then
    crc stop || true
    crc delete --clear-cache -f || true
fi
​
pids=`ps -ef | grep crc | grep -v grep | grep -v install-crc | awk '{print $2}' || true`
for i in $pids; do
  echo "Terminating pid:$i"
  ps -ef | grep $i
  kill -9 $i
done
​
if [ -d ~/.crc ];then
    rm -rf ~/.crc
fi
​
if [ -d $BINDIR ];then
   rm -rf $BINDIR
fi
mkdir -p $BINDIR
​
if [ -d $CRCDIR ];then
   rm -rf $CRCDIR
fi
mkdir -p $CRCDIR
​
cd $CRCDIR
# get the CRC installation
if [ "`grep crc.server.url $ROOTDIR/install.properties`" != "" ];then
   url=`grep crc.server.url $ROOTDIR/install.properties | awk -F= '{print $2}'`
   if [ "$url" = "" ];then
       echo "WARNING: crc.server.url property is empty, using default"
       url="https://mirror.openshift.com/pub/openshift-v4/clients/crc/latest/crc-macos-amd64.tar.xz"
   fi
else
   echo "WARNING: crc.server.url property is missing, using default"
   url="https://mirror.openshift.com/pub/openshift-v4/clients/crc/latest/crc-macos-amd64.tar.xz"
fi
wget $url
name=`basename $url`
tar -xvf $name
cd crc*/
cp crc $BINDIR
​
cd $BINDIR
​
# Get the pull request needed to initialize CRC
if [ "`grep crc.token.file $ROOTDIR/install.properties`" != "" ];then
   token_file=`grep crc.token.file $ROOTDIR/install.properties | awk -F= '{print $2}'`
   if [ "$token_file" != "" ];then
       token_file="$ROOTDIR/$token_file"
   else
       echo "WARNING: crc.token.file property is empty, using default"
       token_file="$ROOTDIR/token"
   fi
else
   echo "WARNING: crc.token.file property is missing, using default"
   token_file="$ROOTDIR/token"
fi
export OFFLINE_ACCESS_TOKEN=`cat $token_file`
export BEARER=$(curl \
--silent \
--data-urlencode "grant_type=refresh_token" \
--data-urlencode "client_id=cloud-services" \
--data-urlencode "refresh_token=${OFFLINE_ACCESS_TOKEN}" \
https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token | \
jq -r .access_token)
curl -X POST https://api.openshift.com/api/accounts_mgmt/v1/access_token --header "Content-Type:application/json" --header "Authorization: Bearer $BEARER" > $BINDIR/pullsecret.json
rm -rf $token_file
​
crc version
# setting cpus and memory values based on properties
if [ "`grep crc.cpus $ROOTDIR/install.properties`" != "" ];then
    cpus=`grep crc.cpus $ROOTDIR/install.properties | awk -F= '{print $2}'`
    if [ "$cpus" = "" ];then
        echo "WARNING: crc.cpus property is empty, using default"
        cpus="10"
    fi
else
   echo "WARNING: crc.cpus property is missing, using default"
   cpus="10"
fi
crc config set cpus $cpus
​
if [ "`grep crc.memory $ROOTDIR/install.properties`" != "" ];then
    memory=`grep crc.memory $ROOTDIR/install.properties | awk -F= '{print $2}'`
    if [ "$memory" = "" ];then
        echo "WARNING: crc.memory property is empty, using default"
        memory="10240"
    fi
else
   echo "WARNING: crc.memory property is missing, using default"
   memory="10240"
fi
crc config set memory $memory
​
if [ -f $ROOTDIR/sudo.password ];then
    sudopasswd="`cat $ROOTDIR/sudo.password`"
    if [ "$sudopasswd" != "" ];then
    	expect -c "set timeout -1
spawn crc setup
expect \"Password:*\"
send \"$sudopasswd\r\"
expect eof"
    else
        crc setup
    fi
else
   crc setup
fi
rm -rf $ROOTDIR/sudo.password
​
crc start --pull-secret-file $BINDIR/pullsecret.json
​
crc status
​
CACHE_DIR="$(crc status | grep "Cache Directory" | awk '{ print $3 }')"
KUBECONFIG="$(find ${CACHE_DIR} -type f -name kubeconfig | xargs ls -t | head -1)"
echo "KUBECONFIG: $KUBECONFIG"
cp $KUBECONFIG ~/.crc/config
KUBE_PASS="$(cat $(find ${CACHE_DIR} -type f -name kubeadmin-password | xargs ls -t | head -1))"
OCP_HOST=$(cat ${KUBECONFIG} | grep server | awk -F'[:]' '{print $2":"$3":"$4}')
​
# CRC automatically installs the oc command in ~/.crc/bin
# all we need to do is copy oc to kubectl and copy them both to the bin
# directory for use. The bin directory was added to the PATH variable at the beginning
# of this script
cp ~/.crc/bin/oc $BINDIR
cp ~/.crc/bin/oc $BINDIR/kubectl
​
eval $(crc oc-env)
​
echo "Executing OC Login..."
oc login ${OCP_HOST} -u kubeadmin -p ${KUBE_PASS}
​
oc projects
kubectl get nodes
​
