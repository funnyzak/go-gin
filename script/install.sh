#!/bin/bash

#========================================================
#   System Required: Debian 8+ / Ubuntu 16.04+ / Centos 7+ /
#   Description: GO-GIN Management script
#   Github: https://github.com/funnyzak/go-gin
#========================================================

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'
export PATH=$PATH:/usr/local/bin

SCRIPT_VERSION="0.0.1"
SCRIPT_NAME="GO-GIN Management Script"
GG_DESCRIPTION="Go-Gin is a web service based on Golang and Gin framework."
GG_NAME="go-gin"
GG_GITHUB_REPO_NAME="funnyzak/${GG_NAME}"

GG_SERVICE_NAME="${GG_NAME}"
CGS_BASE_PATH="/opt/${GG_SERVICE_NAME}"
GG_SERVICE_PATH="${CGS_BASE_PATH}/${GG_SERVICE_NAME}"
CGS_CONFIG_PATH="${CGS_BASE_PATH}/config.yaml"
GG_LOG_PATH="${CGS_BASE_PATH}/logs/log.log"
CGS_SERVICE_PATH="/etc/systemd/system/${GG_SERVICE_NAME}.service"
GG_RELEASES_DATA_URL="https://api.github.com/repos/${GG_GITHUB_REPO_NAME}/releases"

GG_LATEST_VERSION=""
GG_LATEST_VERSION_ZIP_NAME=""
GG_LATEST_VERSION_DOWNLOAD_URL=""

GG_RAW_URL="https://raw.githubusercontent.com/${GG_GITHUB_REPO_NAME}/main"
GG_CONFIG_SAMPLE_URL="${GG_RAW_URL}/config.yaml.example"

os_arch=""
[ -e /etc/os-release ] && cat /etc/os-release | grep -i "PRETTY_NAME" | grep -qi "alpine" && os_alpine='1'

start_check() {
    [ "$os_alpine" != 1 ] && ! command -v systemctl >/dev/null 2>&1 && echo "This system is not supported: systemctl not found" && exit 1

    # check root
    [[ $EUID -ne 0 ]] && echo -e "${red}ERROR: ${plain} This script must be run with the root user!\n" && exit 1

    # check arch, only support amd64, arm64, arm
    if [[ $(uname -m | grep 'x86_64') != "" ]]; then
        os_arch="amd64"
    elif [[ $(uname -m | grep 'aarch64\|armv8b\|armv8l') != "" ]]; then
        os_arch="arm64"
    elif [[ $(uname -m | grep 'arm') != "" ]]; then
        os_arch="arm"
    fi

    if [[ -z ${os_arch} ]]; then
        echo -e "${red}ERROR: ${plain} This system is not supported: ${os_arch}\n" && exit 1
    fi

    ping_check

    GG_LATEST_VERSION=$(get_service_version)
    if [ -z "${GG_LATEST_VERSION}" ]; then
        echo -e "${red}ERROR${plain}: Get ${GG_SERVICE_NAME} latest version failed."
        exit 1
    fi

    GG_LATEST_VERSION_ZIP_NAME="${GG_SERVICE_NAME}-linux-${os_arch}-${GG_LATEST_VERSION}.zip"
    GG_LATEST_VERSION_DOWNLOAD_URL="https://github.com/${GG_GITHUB_REPO_NAME}/releases/download/${GG_LATEST_VERSION}/${GG_LATEST_VERSION_ZIP_NAME}"

}

ping_check() {
    ping -c 1 github.com > /dev/null 2>&1
    if [ $? -ne 0 ]; then
        echo -e "${red}ERROR${plain}: ping github.com failed. Please check your network."
        exit 1
    fi
}

confirm() {
  if [[ $# > 1 ]]; then
    echo && read -e -p "$1 [y/n]: " temp
    if [[ x"${temp}" == x"" ]]; then
        temp=$2
    fi
  else
    read -e -p "$1 [y/n]: " temp
  fi
  if [[ x"${temp}" == x"y" || x"${temp}" == x"Y" ]]; then
    return 0
  else
    return 1
  fi
}

install_service() {
  echo -e "Install ${green}${GG_SERVICE_NAME}${plain} service..."
  if [ -f "${CGS_SERVICE_PATH}" ]; then
    echo -e "${red}${GG_SERVICE_NAME}${plain} service already exists."
  fi
  if [ ! -d "${CGS_BASE_PATH}" ]; then
    mkdir -p ${CGS_BASE_PATH}
  fi
  download_service_app
  download_service_config
  # download_service_template
}

service_action() {
  echo -e "${action} ${green}${GG_SERVICE_NAME}${plain} service..."

  local action=$1
  local zero_flag=$2

  systemctl ${action} ${GG_SERVICE_NAME}

  if [[ $? -ne 0 ]]; then
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} ${action} failed."
    return 0
  fi

  echo -e "${green}${GG_SERVICE_NAME}${plain} service ${action} success."
  if [[ -n "${zero_flag}" ]]; then
    before_show_menu
  fi
}

start_service() {
  service_action start $1
}

stop_service() {
  service_action stop $1
}

restart_service() {
  service_action restart $1
}

show_service_status() {
  service_action status
}

upgrade_service() {
  echo -e "Upgrade ${green}${GG_SERVICE_NAME}${plain} service..."

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

uninstall_service() {
  echo -e "Uninstall ${green}${GG_SERVICE_NAME}${plain} service..."

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

show_service_log() {
  echo -e "Show ${green}${GG_SERVICE_NAME}${plain} service log..."

  watch -n 1 tail -n 20 ${GG_LOG_PATH}
  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

edit_service_config() {
  echo -e "Edit ${green}${GG_SERVICE_NAME}${plain} service config..."

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

download_service_config() {
  download_file "${GG_CONFIG_SAMPLE_URL}" "${CGS_CONFIG_PATH}"
  if [ $? -ne 0 ]; then
    return 0
  fi
  if [ $? -ne 0 ]; then
    return 0
  fi
}

download_service_app() {
  download_file "${GG_LATEST_VERSION_DOWNLOAD_URL}" "/tmp/${GG_LATEST_VERSION_ZIP_NAME}"
  if [ $? -ne 0 ]; then
    return 0
  fi
  if [ -f "${GG_SERVICE_PATH}" ]; then
    rm -f ${GG_SERVICE_PATH}
  fi
  unzip -o /tmp/${GG_LATEST_VERSION_ZIP_NAME} -d ${CGS_BASE_PATH} > /dev/null 2>&1
  if [ $? -ne 0 ]; then
    echo -e "${red}ERROR${plain}: Unzip ${GG_LATEST_VERSION_ZIP_NAME} failed."
    return 0
  fi
}

get_service_version() {
  local latest_version=$(curl -s ${GG_RELEASES_DATA_URL} | grep "tag_name" | head -n 1 | awk -F '"' '{print $4}')
  if [ -n "${latest_version}" ]; then
    echo "${latest_version}"
  else
    echo ""
  fi
}

download_file() {
  local url=$1
  local file=$2
  wget -t 3 -T 15 -O ${file} ${url} > /dev/null 2>&1
  if [ $? -ne 0 ]; then
    echo -e "${red}ERROR${plain}: Download ${url} failed."
    return 0
  fi
}

show_usage() {
  echo -e "\nUsage: ${green}go-gin${plain} [option]"
  echo "Options:"
  echo -e "  ${green}install${plain}          Install service"
  echo -e "  ${green}start${plain}            Start service"
  echo -e "  ${green}stop${plain}             Stop service"
  echo -e "  ${green}restart${plain}          Restart service"
  echo -e "  ${green}upgrade${plain}          Upgrade service"
  echo -e "  ${green}uninstall${plain}        Uninstall service"
  echo -e "  ${green}status${plain}           Show service status"
  echo -o "  ${green}log${plain}              Show service log"
  echo -e "  ${green}edit${plain}             Edit service config"
}

show_menu() {
  echo -e ">
    ${green}${SCRIPT_NAME} ${plain}${red}v${SCRIPT_VERSION}${plain}
    ————————————————
    ${green}1.${plain} Install service
    ${green}2.${plain} Start service
    ${green}3.${plain} Stop service
    ${green}4.${plain} Restart service
    ${green}5.${plain} Upgrade service
    ${green}6.${plain} Uninstall service
    ${green}7.${plain} Show service status
    ${green}8.${plain} Show service log
    ${green}9.${plain} Edit service config
    ————————————————
    ${green}0.${plain} Exit
    "
  echo && read -ep "Please enter a number [0-9]: " num

  case "${num}" in
  0)
    exit 0
    ;;
  1)
    install_service
    ;;
  2)
    start_service
    ;;
  3)
    stop_service
    ;;
  4)
    restart_service
    ;;
  5)
    upgrade_service
    ;;
  6)
    uninstall_service
    ;;
  7)
    show_service_status
    ;;
  8)
    show_service_log
    ;;
  9)
    edit_service_config
    ;;
  *)
    echo -e "${red}Please enter the correct number [0-9]${plain}"
    ;;
  esac
}

before_show_menu() {
  echo && echo -n -e "${yellow}* Press Enter to return to the main menu. *${plain}" && read temp
  show_menu
}

if [ $# -gt 0 ] && [ -n "$1" ]; then
  case $1 in
  "install")
    install_service
    ;;
  "start")
    start_service
    ;;
  "stop")
    stop_service
    ;;
  "restart")
    restart_service
    ;;
  "upgrade")
    upgrade_service
    ;;
  "uninstall")
    uninstall_service
    ;;
  "status")
    show_service_status
    ;;
  "log")
    show_service_log
    ;;
  "edit")
    edit_service_config
    ;;
  *)
    $@ || show_usage
    ;;
  esac
else
  show_menu
fi
