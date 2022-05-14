/**
 * Sets a cookie
 * @param {string} key cookie key
 * @param {string} value cookie value
 * @param {float} expiry time to the cookie expires (in minutes)
 */
function setCookie(key, value, expiry) {
    // expiry is in minutes
    var expires = new Date();
    expires.setTime(expires.getTime() + (expiry * 60 * 1000));
    document.cookie = key + '=' + value + ';expires=' + expires.toUTCString();
}

/**
 * Gets a cookie
 * @param {string} key cookie key
 * @returns keyValue[2] if keyValue is not null
 */
function getCookie(key) {
    // This regex is from stackoverflow, don't ask me how tf this works 
    var keyValue = document.cookie.match('(^|;) ?' + key + '=([^;]*)(;|$)');
    return keyValue ? keyValue[2] : null;
}

/**
 * Deletes a cookie
 * @param {string} key cookie key
 */
function eraseCookie(key) {
    var keyValue = getCookie(key);
    setCookie(key, keyValue, '-1');
}