FROM amazonlinux:1

RUN yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
RUN yum --enablerepo=epel install -y ocl-icd ocl-icd-devel opencl-headers golang zip unzip

RUN groupadd --gid 1000 commp && \
  useradd --gid 1000 --uid 1000 --create-home --shell /bin/bash commp

USER commp

RUN curl https://sh.rustup.rs -sSf | sh -s -- --default-toolchain nightly-2019-11-06 -y

ENV PATH /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/home/commp/.cargo/bin

VOLUME /home/commp/build

WORKDIR /home/commp

USER commp:commp
