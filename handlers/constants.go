package handler

const VERSION string = "v1"

const ENV_PORT = "PORT"
const DEFAULT_PORT = "8080"

//url paths
const DEFAULT_PATH = "/"
const CORE_PATH = "/unisearcher/" + VERSION + "/"
const UNI_INFO_PATH = CORE_PATH + "uniinfo/"
const COUNTRIES_PATH = CORE_PATH + "neighbourunis/"
const DIAGNOSTICS_PATH = CORE_PATH + "diag/"
