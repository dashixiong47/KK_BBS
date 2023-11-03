// composables/useCookies.js

import { useCookie } from "#app"
export default function useCookies() {
    let setCookie = (key, value = null) => {
        let cookie = useCookie(key);
        cookie.value = value
    }
    let getCookie = (key) => {
        let cookie = useCookie(key);
        return cookie.value
    }
    return { setCookie,getCookie };
}
