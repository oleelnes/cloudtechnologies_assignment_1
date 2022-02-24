package handler

const VERSION string = "v1"

const ENV_PORT = "PORT"
const DEFAULT_PORT = "8080"

//url paths

const UNI_INFO = "uniinfo"
const UNISEARCHER = "unisearcher"
const COUNTRIES = "neighbourunis"
const DIAGNOSTICS = "diag"

const DEFAULT_PATH = "/"
const CORE_PATH = "/" + UNISEARCHER + "/"
const VERSION_PATH = CORE_PATH + VERSION + "/"
const UNI_INFO_PATH = VERSION_PATH + UNI_INFO + "/"
const COUNTRIES_PATH = VERSION_PATH + COUNTRIES + "/"
const DIAGNOSTICS_PATH = VERSION_PATH + DIAGNOSTICS + "/"
