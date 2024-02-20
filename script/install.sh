#!/bin/bash

#========================================================
#   System Required: Debian 8+ / Ubuntu 16.04+ / Centos 7+ /
#   Description: GO-GIN script
#   Github: https://github.com/funnyzak/go-gin
#========================================================

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'
export PATH=$PATH:/usr/local/bin

SCRIPT_VERSION="0.0.1"                                                      # script version
SCRIPT_NAME="GO-GIN Script"                                                 # script name
GG_DESCRIPTION="Go-Gin is a web service based on Golang and Gin framework." # service description

GG_NAME="go-gin"                   # service name
GG_REPO_NAME="funnyzak/${GG_NAME}" # service repo name
GG_REPO_BRANCH="main"              # service repo branch

GG_SERVICE_NAME="${GG_NAME}"                                                 # service system name, as a systemd service name, shell name, tmp file name, etc.
GG_WORK_PATH="/opt/${GG_SERVICE_NAME}"                                       # service workdir path
GG_SERVICE_PATH="${GG_WORK_PATH}/${GG_SERVICE_NAME}"                         # service exec app path
GG_CONFIG_PATH="${GG_WORK_PATH}/${GG_SERVICE_NAME}.yaml"                     # service config file path
GG_SYSTEMD_PATH="/etc/systemd/system/${GG_SERVICE_NAME}.service"             # systemd service config file path
GG_RELEASES_DATA_URL="https://api.github.com/repos/${GG_REPO_NAME}/releases" # service releases data url for get latest version

GG_LATEST_VERSION=""              # service latest version, init by
GG_LATEST_VERSION_ZIP_NAME=""     # service latest version zip name
GG_LATEST_VERSION_DOWNLOAD_URL="" # service latest version download url

GG_RAW_URL="https://raw.githubusercontent.com/${GG_REPO_NAME}/${GG_REPO_BRANCH}" # service attachment prefix url
GG_CONFIG_SAMPLE_URL="${GG_RAW_URL}/config.example.yaml"                         # service sample config download url
SCRIPT_DOWNLINK="${GG_RAW_URL}/script/install.sh"                                # script download url

system_architecture="" # system arch

[ -e /etc/os-release ] && cat /etc/os-release | grep -i "PRETTY_NAME" | grep -qi "alpine" && os_alpine='1'

check_system_requirements() {
  [ "$os_alpine" != 1 ] && ! command -v systemctl >/dev/null 2>&1 && echo "This system is not supported: systemctl not found" && exit 1

  # check root
  [[ $EUID -ne 0 ]] && echo -e "${red}ERROR: ${plain} This script must be run with the root user!\n" && exit 1

  # check arch, only support amd64, arm64, arm
  if [[ $(uname -m | grep 'x86_64') != "" ]]; then
    system_architecture="amd64"
  elif [[ $(uname -m | grep 'aarch64\|armv8b\|armv8l') != "" ]]; then
    system_architecture="arm64"
  elif [[ $(uname -m | grep 'arm') != "" ]]; then
    system_architecture="arm"
  fi

  if [[ -z ${system_architecture} ]]; then
    echo -e "${red}ERROR: ${plain} This system is not supported: ${system_architecture}\n" && exit 1
  fi
}

ping_check() {
  ping -c 1 github.com >/dev/null 2>&1
  handle_error $? "ping github.com failed. Please check your network."
}

handle_error() {
  if [[ $1 -ne 0 ]]; then
    echo -e "${red}ERROR${plain}: $2"
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

execute_funcs() {
  for one_func in "$@"; do
    echo -e "Execute function: ${green}${one_func}${plain}"
    if [[ $one_func == *' '* ]]; then
      eval $one_func
    else
      $one_func
    fi
    if [ $? -ne 0 ]; then
      echo -e "${red}ERROR${plain}: execute function list ${@} failed, failed function is ${yellow}${one_func}${plain}."
      return 1
    fi
  done
}

install_service() {
  echo -e "Install ${green}${GG_SERVICE_NAME}${plain} service..."
  if service_exists; then
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} service is already installed."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi

  execute_funcs "get_service_latest_version" "download_service_app" "download_service_template" "download_service_config" "enable_service 0" "start_service 0"
  if [ $? -ne 0 ]; then
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} service install failed."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi

  echo -e "${green}${GG_SERVICE_NAME}${plain} service for ${system_architecture} install success. the latest version is ${GG_LATEST_VERSION}. Enjoy it!"

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

get_service_log_path() {
  local log_path=""
  local log_paths=$(grep -i "path" ${GG_CONFIG_PATH} | awk '{print $2}')
  if [ -n "${log_paths}" ]; then
    for i in ${log_paths}; do
      if [[ $i == *log* ]]; then
        log_path=$i
      fi
    done
  else
    echo ""
  fi
  if [ -n "${log_path}" ]; then
    if [[ ${log_path} != /* ]]; then
      log_path=$(dirname ${GG_SERVICE_PATH})/${log_path}
      echo ${log_path}
    fi
  else
    echo ""
  fi

}

service_exists() {
  if systemctl --all --type=service | grep -Fq "${GG_SERVICE_NAME}.service"; then
    return 0
  else
    return 1
  fi
}

service_active() {
  if systemctl is-active --quiet ${GG_SERVICE_NAME}; then
    return 0
  else
    return 1
  fi
}

handle_service_action() {
  local action=$1
  shift

  echo -e "${action} ${green}${GG_SERVICE_NAME}${plain} service..."

  systemctl ${action} ${GG_SERVICE_NAME}

  if [[ $? -ne 0 ]]; then
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} ${action} failed."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi

  echo -e "${green}${GG_SERVICE_NAME}${plain} service ${action} success."

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

create_service_workdir() {
  if [ ! -d "${GG_WORK_PATH}" ]; then
    mkdir -p ${GG_WORK_PATH}
  fi
}

enable_service() {
  echo -e "Enable ${green}${GG_SERVICE_NAME}${plain} service..."
  systemctl enable ${GG_SERVICE_NAME}
  if [ $? -ne 0 ]; then
    echo -e "${red}ERROR${plain}: Enable ${GG_SERVICE_NAME} service failed."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi
  echo -e "${green}${GG_SERVICE_NAME}${plain} service enable success."
  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

start_service() {
  handle_service_action start $1
}

stop_service() {
  handle_service_action stop $1
}

disable_service() {
  handle_service_action disable $1
}

restart_service() {
  handle_service_action restart $1
}

show_service_status() {
  handle_service_action status $1
}

upgrade_service() {
  echo -e "Upgrade ${green}${GG_SERVICE_NAME}${plain} service..."

  if service_exists; then
    execute_funcs "get_service_latest_version" "download_service_app" "handle_service_action restart"
    if [ $? -ne 0 ]; then
      echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} service upgrade failed."
      if [[ $# == 0 ]]; then
        before_show_menu
      else
        return 1
      fi
    fi
    echo -e "${green}${GG_SERVICE_NAME}${plain} service for ${system_architecture} upgrade success. the latest version is ${GG_LATEST_VERSION}. Enjoy it!"
  else
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} service not installed or upgrade failed."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

uninstall_service() {
  echo -e "Uninstall ${green}${GG_SERVICE_NAME}${plain} service..."

  if service_exists; then
    execute_funcs "stop_service 0" "disable_service 0"
    systemctl daemon-reload
    rm -f ${GG_SYSTEMD_PATH}
    rm -f ${GG_SERVICE_PATH}
    rm -f ${GG_CONFIG_PATH}
    rm -rf ${GG_WORK_PATH}
    systemctl daemon-reload
    systemctl reset-failed
    echo -e "${green}${GG_SERVICE_NAME}${plain} service uninstall success. Goodbye!"
  else
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} service uninstall failed. Please check ${GG_SERVICE_NAME} service is installed."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

show_service_log() {
  echo -e "Show ${green}${GG_SERVICE_NAME}${plain} service log..."

  if ! service_exists || [ -z "$(get_service_log_path)" ] || [ ! -f "$(get_service_log_path)" ]; then
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} service log not found."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi

  echo -e "Press ${red}Ctrl+C${plain} to exit."
  watch -n 1 tail -n 20 "$(get_service_log_path)"

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

edit_service_config() {
  echo -e "Edit ${green}${GG_SERVICE_NAME}${plain} service config..."

  if ! service_exists; then
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} service not installed. Please install it first."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi

  if [ ! -f "${GG_CONFIG_PATH}" ]; then
    echo -e "${red}ERROR${plain}: ${GG_SERVICE_NAME} config file not found."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi

  if command -v vim >/dev/null 2>&1; then
    vim ${GG_CONFIG_PATH}
  elif command -v nano >/dev/null 2>&1; then
    nano ${GG_CONFIG_PATH}
  elif command -v vi >/dev/null 2>&1; then
    vi ${GG_CONFIG_PATH}
  else
    echo -e "${red}ERROR${plain}: No editor found."
    if [[ $# == 0 ]]; then
      before_show_menu
    else
      return 1
    fi
  fi
  echo -e "${green}${GG_SERVICE_NAME}${plain} service config edit success."
  if confirm "Do you want to restart ${GG_SERVICE_NAME} service?"; then
    execute_funcs "handle_service_action restart"
  fi

  if [[ $# == 0 ]]; then
    before_show_menu
  fi
}

download_service_config() {
  create_service_workdir
  download_file "${GG_CONFIG_SAMPLE_URL}" "${GG_CONFIG_PATH}"
  if [ $? -ne 0 ]; then
    return 1
  fi
}

download_service_template() {
  create_service_workdir
  download_file "${GG_RAW_URL}/script/${GG_SERVICE_NAME}.service" "${GG_SYSTEMD_PATH}"
  if [ $? -ne 0 ]; then
    return 1
  fi
}

download_service_app() {
  create_service_workdir
  download_file "${GG_LATEST_VERSION_DOWNLOAD_URL}" "/tmp/${GG_LATEST_VERSION_ZIP_NAME}"
  if [ $? -ne 0 ]; then
    return 1
  fi
  if [ -f "${GG_SERVICE_PATH}" ]; then
    rm -f ${GG_SERVICE_PATH}
  fi
  echo -e "Unzip ${GG_LATEST_VERSION_ZIP_NAME} to ${GG_WORK_PATH}..."
  unzip -o /tmp/${GG_LATEST_VERSION_ZIP_NAME} -d ${GG_WORK_PATH} >/dev/null 2>&1
  if [ $? -ne 0 ]; then
    echo -e "${red}ERROR${plain}: Unzip ${GG_LATEST_VERSION_ZIP_NAME} failed."
    return 1
  fi
  chmod +x ${GG_SERVICE_PATH}
}

get_service_latest_version() {
  if [ -n "${GG_LATEST_VERSION}" ]; then
    return 0
  fi
  GG_LATEST_VERSION=$(curl -s ${GG_RELEASES_DATA_URL} | grep "tag_name" | head -n 1 | awk -F '"' '{print $4}')
  if [ -z "${GG_LATEST_VERSION}" ]; then
    echo -e "${red}ERROR${plain}: Get ${GG_SERVICE_NAME} latest version failed."
    exit 1
  fi
  GG_LATEST_VERSION_ZIP_NAME="${GG_SERVICE_NAME}-linux-${system_architecture}-${GG_LATEST_VERSION}.zip"
  GG_LATEST_VERSION_DOWNLOAD_URL="https://github.com/${GG_REPO_NAME}/releases/download/${GG_LATEST_VERSION}/${GG_LATEST_VERSION_ZIP_NAME}"
}

download_file() {
  local url=$1
  local file=$2
  echo -e "Download ${url} to ${file}..."
  wget -t 3 -T 15 -O ${file} ${url} >/dev/null 2>&1
  if [ $? -ne 0 ]; then
    echo -e "${red}ERROR${plain}: Download ${url} failed."
    return 1
  fi
}

update_script() {
  echo -e "Updating script..."

  curl -sL ${SCRIPT_DOWNLINK} -o /tmp/${GG_SERVICE_NAME}.sh
  new_version=$(cat /tmp/${GG_SERVICE_NAME}.sh | grep "SCRIPT_VERSION" | head -n 1 | awk -F "=" '{print $2}' | sed 's/\"//g;s/,//g;s/ //g')
  if [ ! -n "$new_version" ]; then
    echo -e "${red}ERROR${plain}: Script download failed, please check your network. Download URL: ${SCRIPT_DOWNLINK}"
    return 1
  fi
  echo -e "The latest version is ${new_version}."
  mv -f /tmp/${GG_SERVICE_NAME}.sh ./${GG_SERVICE_NAME}.sh && chmod a+x ./${GG_SERVICE_NAME}.sh

  echo -e "3s later, restart script..."
  sleep 3s
  clear
  exec ./${GG_SERVICE_NAME}.sh
  exit 0
}

show_usage() {
  echo -e "\nUsage: $0 [action] [arguments]"
  echo -e ""
  echo "Available actions:"
  echo -e "  ${green}install${plain}          Install service"
  echo -e "  ${green}start${plain}            Start service"
  echo -e "  ${green}stop${plain}             Stop service"
  echo -e "  ${green}restart${plain}          Restart service"
  echo -e "  ${green}upgrade${plain}          Upgrade service"
  echo -e "  ${green}uninstall${plain}        Uninstall service"
  echo -e "  ${green}status${plain}           Show service status"
  echo -e "  ${green}log${plain}              Show service log"
  echo -e "  ${green}edit${plain}             Edit service config"
  echo -e "  ${green}update${plain}           Update script"
  echo -e ""
  echo -e "Example:"
  echo -e "  $0 ${green}install${plain}"
  echo -e "  $0 ${green}start${plain}"
  echo -e "  $0 ${green}stop${plain}"
  echo -e ""
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
    ${green}10.${plain} Update script
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
  10)
    update_script
    ;;
  *)
    echo -e "${red}Please enter the correct number [0-10]${plain}"
    ;;
  esac
}

before_show_menu() {
  echo && echo -n -e "${yellow}* Press Enter to return to the main menu. *${plain}" && read temp
  show_menu
}

check_system_requirements

if [ $# -gt 0 ] && [ -n "$1" ]; then
  case $1 in
  "install")
    install_service 0
    ;;
  "start")
    start_service 0
    ;;
  "stop")
    stop_service 0
    ;;
  "restart")
    restart_service 0
    ;;
  "upgrade")
    upgrade_service 0
    ;;
  "uninstall")
    uninstall_service 0
    ;;
  "status")
    show_service_status 0
    ;;
  "log")
    show_service_log 0
    ;;
  "edit")
    edit_service_config 0
    ;;
  "update")
    update_script 0
    ;;
  *)
    $@ || show_usage 0
    ;;
  esac
else
  show_menu
fi
